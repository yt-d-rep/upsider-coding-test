.PHONY: start-dev stop-dev restart-dev migration rollback setup serve gen-wire gen-mock

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

migration:
	cd migration && go run migrate.go up

rollback:
	cd migration && go run migrate.go down

serve:
	go run cmd/main.go

gen-wire:
	wire ./...

gen-mock:
	bash ./.script/gen-mock.sh