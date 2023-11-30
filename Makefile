OS := $(shell uname)
.PHONY: setup_db

all: setup_db nats go 

stop: rm_db stan_stop 

setup_db:
ifeq ($(OS), Darwin)  
	sh ./scripts/db_start_mac.sh
else ifeq ($(OS), Linux) 
	sh ./scripts/db_start.sh
endif

rm_db:
	docker stop orders
	docker rm orders

# # # # # # # # # # # # # # # # # # # # # # # # # # # # # # # 

nats:
	nats-streaming-server -sc config/nats-streaming-config.yaml &

stan_stop:
	sh ./scripts/stan_stop.sh

go:
	go run cmd/main.go

vegeta:
	vegeta attack -targets=./scripts/attack-targets.txt -rate=50 -duration=40s | tee results.bin | vegeta report

test: 
	go test a_test.go