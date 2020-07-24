#!/usr/bin/env bash

wait-for-it postgres:5432 -s -- migrate -source file://migrations -database postgres://"$POSTGRES_USER":"$POSTGRES_PASSWORD"@postgres:5432/"$POSTGRES_DB"?sslmode=disable up