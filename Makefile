#!/usr/bin/env bash

# This Makefile doesn't track file changes and built executables like a normal Makefile.
# Instead, it defines "CI Pipelines" by defining targets that don't exist that depend on each other.
# This results in very simply-defined pipelines, at the cost of some `make` efficiencies.


# ==== CI PIPELINES ====
build: clean generate
	go build -o ./bin/app -v ./cmd/app
	cp -r migrations ./bin/migrations

test: generate
	go test -v -race ./...

generate: lint
	go generate ./...

lint: install
	go mod verify
	go mod tidy
	go fmt ./...
	go vet ./...
	golangci-lint run

install:
	go mod download
	which golangci-lint # Check golangci-lint is installed

clean:
	rm -rf ./bin

# ==== DEV SCRIPTS ====
PROJECT_DIR := $(shell pwd)

run: install
	docker run -d --rm --name make_postgres -it -p 5432:5432 \
		-e POSTGRES_USER=init-microservice-go \
 		-e POSTGRES_PASSWORD=password \
 		-e POSTGRES_DB=service_db \
 		postgres:latest
	bash -c "trap 'trap - SIGINT SIGTERM ERR; docker stop make_postgres; exit 1; exit 1' SIGINT SIGTERM ERR; ${MAKE} run_internal"

# TODO: This is broken with new migration code - FIX!
run_internal:
	go run ./cmd/app $(PROJECT_DIR)/migrations
