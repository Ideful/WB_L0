OS := $(shell uname)
.PHONY: setup_db

all: nats go

rm_db:
	docker stop orders
	docker rm orders

nats:
	nats-streaming-server -sc config/nats-streaming-config.yaml &

stan_stop:
	sh ./scripts/stan_stop.sh

go:
	go run cmd/main.go

setup_db:
ifeq ($(OS), Darwin)  
	sh ./scripts/db_start_mac.sh
else ifeq ($(OS), Linux) 
	sh ./scripts/db_start.sh
endif