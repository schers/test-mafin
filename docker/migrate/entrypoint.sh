#!/usr/bin/env bash

migrate -source file://migrations -database postgres://"$POSTGRES_USER":"$POSTGRES_PASSWORD"@postgres:5432/"$POSTGRES_DB"?sslmode=disable up