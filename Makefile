.PHONY: setup_db

all: run-nats run-go
setup_db:
	sh ./scripts/db_start.sh

rm_db:
	docker stop orders
	docker rm orders

run-nats:
	nats-streaming-server -sc config/nats-streaming-config.yaml &

run-go:
	go run cmd/main.go