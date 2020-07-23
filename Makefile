#!make
SHELL = /bin/bash

include .env

migrate-up:
	migrate -source file://migrations -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_LOCAL_PORT)/$(POSTGRES_DB)?sslmode=disable up

migrate-down-all:
	migrate -source file://migrations -database postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:$(POSTGRES_LOCAL_PORT)/$(POSTGRES_DB)?sslmode=disable down
