.PHONY: start-dev stop-dev restart-dev setup migration rollback

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
