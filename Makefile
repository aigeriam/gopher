SHELL := /bin/zsh

DB_ADDR ?= postgres://admin:adminpassword@localhost:5432/gopher?sslmode=disable
MIGRATIONS_DIR := cmd/migrate/migrations

.PHONY: help dev test db-up db-down db-logs migrate-up migrate-down migrate-version migrate-create

help:
	@printf "Available commands:\n"
	@printf "  make dev                    Start the API with Air\n"
	@printf "  make test                   Run Go tests\n"
	@printf "  make db-up                 Start PostgreSQL\n"
	@printf "  make db-down               Stop PostgreSQL\n"
	@printf "  make db-logs               Show PostgreSQL logs\n"
	@printf "  make migrate-up            Apply pending migrations\n"
	@printf "  make migrate-down          Undo the latest migration\n"
	@printf "  make migrate-version       Show migration version\n"
	@printf "  make migrate-create name=x Create a migration\n"

dev:
	air

test:
	go test ./...

db-up:
	docker compose up -d

db-down:
	docker compose down

db-logs:
	docker compose logs -f db

migrate-up:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_ADDR)" up

migrate-down:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_ADDR)" down 1

migrate-version:
	migrate -path $(MIGRATIONS_DIR) -database "$(DB_ADDR)" version

migrate-create:
	@test -n "$(name)" || (printf "Usage: make migrate-create name=create_posts\n" && exit 1)
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq $(name)
