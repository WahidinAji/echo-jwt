#!/bin/bash
export APP_NAME=SimpleApp
export SERVER_PORT=8080
export PG_URL="YourPgSqlUrl"
export JWT_SECRET="YourJWTSecret"
go run  ./main.go