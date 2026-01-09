APP_NAME=go-flight-search
COMPOSE=docker compose

.PHONY: help up down build rebuild logs ps restart clean dev test lint redis-cli

help:
	@echo ""
	@echo "Available commands:"
	@echo "  make up        - Build & start all services"
	@echo "  make down      - Stop all services"
	@echo "  make restart   - Restart all services"
	@echo "  make build     - Build docker images"
	@echo "  make rebuild   - Rebuild docker images (no cache)"
	@echo "  make logs      - Show logs"
	@echo "  make ps        - Show running containers"
	@echo "  make clean     - Stop and remove volumes"
	@echo "  make dev       - Run API locally (no docker)"
	@echo "  make test      - Run go test"
	@echo "  make lint      - Run golangci-lint (if installed)"
	@echo "  make redis-cli - Enter redis CLI"
	@echo ""

up:
	$(COMPOSE) up -d --build

down:
	$(COMPOSE) down

restart:
	$(COMPOSE) down
	$(COMPOSE) up -d --build

build:
	$(COMPOSE) build

rebuild:
	$(COMPOSE) build --no-cache

logs:
	$(COMPOSE) logs -f

ps:
	$(COMPOSE) ps

clean:
	$(COMPOSE) down -v

dev:
	go run ./cmd/api

test:
	go test ./...

lint:
	golangci-lint run

redis-cli:
	docker exec -it go-flight-search-redis redis-cli -a supersecretpassword
