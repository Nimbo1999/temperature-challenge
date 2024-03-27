.PHONY: start-build start stop test

start-build:
	docker compose up -d --build

start:
	docker compose up -d

stop:
	docker compose down

test:
	go test ./...
