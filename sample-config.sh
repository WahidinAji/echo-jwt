#!/bin/bash
export APP_NAME=SimpleApp
export SERVER_PORT=9000
export SERVER_READ_TIMEOUT_IN_MINUTE=2
export SERVER_WRITE_TIMEOUT_IN_MINUTE=2
export DB_USER=postgres
export DB_PASS=postgres
export DB_NAME=echo_jwt
export PG_URL="YourPgSqlUrl"
export JWT_SECRET="YourJWTSecret"
go run  ./main.go