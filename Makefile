#! /usr/bin/make
#
# Makefile for swagger-service
#
# Targets:
# - "coverage" runs the tests
# - "build" build the platform
# - "swagger" generate swagger documentation
# - "run" run service
#

TMP="./cmd/api"

build:
	@go build "${TMP}"

coverage:
	@go test -coverprofile=coverage.out "${TMP}"
	@go tool cover -html=coverage.out

swagger:
	@swag init -d "${TMP}" -g openapi.go 

run:
	@go run "${TMP}"