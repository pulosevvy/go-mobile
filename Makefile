include .env
export $(shell sed 's/=.*//' .env)

DOCKER_COMPOSE_FILE := docker-compose.yml

go-run:
	go run cmd/app/main.go

migrate-up:
	goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=$(PG_SSLMODE)" up

docker-build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

docker-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

docker-migrate-up:
	docker-compose run --rm app goose -dir db/migrations postgres "postgresql://$(PG_USERNAME):$(PG_PASSWORD)@$(PG_HOST):$(PG_PORT)/$(PG_DATABASE)?sslmode=$(PG_SSLMODE)" up

docker-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down