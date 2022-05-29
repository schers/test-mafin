#!make
SHELL = /bin/bash

DOCKER_COMPOSE = docker-compose -f docker-compose.yml

.PHONY: all install-dev build up start down destroy stop restart migrate-up migrate-down-all

include .env

all: install-dev

install-dev:
	$(DOCKER_COMPOSE) up -d --build $(c)
	$(MAKE) migrate-up
build:
	$(DOCKER_COMPOSE) build $(c)
up:
	$(DOCKER_COMPOSE) up -d $(c)
start:
	$(DOCKER_COMPOSE) start $(c)
down:
	$(DOCKER_COMPOSE) down --remove-orphans $(c)
destroy:
	$(DOCKER_COMPOSE) down -v $(c)
stop:
	$(DOCKER_COMPOSE) stop $(c)
restart:
	$(DOCKER_COMPOSE) stop $(c)
	$(DOCKER_COMPOSE) up -d $(c)
migrate-up:
	docker run -v `pwd`/migrations:/migrations --network host migrate/migrate -path=/migrations/ \
        -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_LOCAL_PORT)/$(POSTGRES_DB)?sslmode=disable up
migrate-down-all:
	echo "y" | docker run -v `pwd`/migrations:/migrations --network host migrate/migrate -path=/migrations/ \
        -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_LOCAL_PORT)/$(POSTGRES_DB)?sslmode=disable down --all
