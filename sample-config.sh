#!/bin/bash
export APP_NAME=SimpleApp
export SERVER_PORT=8080
export PG_URL="YourPgSqlUrl"
export JWT_SECRET="YourJWTSecret"
export DB_TESTING="postgres://postgres:postgres@localhost:5432/postgres"
go run  ./main.go