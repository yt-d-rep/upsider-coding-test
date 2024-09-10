.PHONY: start-dev stop-dev restart-dev migration rollback setup serve gen-wire gen-mock test

start-dev:
	docker compose up -d --build

stop-dev:
	docker compose down

restart-dev:
	make stop-dev
	make start-dev

setup:
	go install github.com/google/wire/cmd/wire@latest
	go install go.uber.org/mock/mockgen@latest

migrate:
	migrate -source file://./db/migration -database postgres://main:main@db:5432/main?sslmode=disable up

rollback:
	migrate -source file://./db/migration -database postgres://main:main@db:5432/main?sslmode=disable down

serve:
	go run cmd/main.go

gen-wire:
	wire ./...

gen-mock:
	bash ./.script/gen-mock.sh

test:
	go test -v ./test/...