#!/usr/bin/env bash

# This Makefile doesn't track file changes and built executables like a normal Makefile.
# Instead, it defines "CI Pipelines" by defining targets that don't exist that depend on each other.
# This results in very simply-defined pipelines, at the cost of some `make` efficiencies.


# ==== CI PIPELINES ====
ci-build: ensure-no-changes test build

ensure-no-changes:
	git diff --exit-code

build: clean generate
	go build -o ./bin/app -v ./cmd/app
	cp -r migrations ./bin/migrations

test: generate
	go test -v ./...

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
run: install
	docker run -d --rm --name make_postgres -it -p 5432:5432 \
		-e POSTGRES_USER=init-microservice-go \
 		-e POSTGRES_PASSWORD=password \
 		-e POSTGRES_DB=service_db \
 		postgres:latest
	sleep 3 # postgres takes a hot second to be ready
	bash -c "trap 'trap - SIGINT SIGTERM ERR; docker stop make_postgres; exit 1; exit 1' SIGINT SIGTERM ERR; ${MAKE} run_internal"

run-podman: install
	podman run -d --rm --name make_postgres -it -p 5432:5432 \
		-e POSTGRES_USER=init-microservice-go \
 		-e POSTGRES_PASSWORD=password \
 		-e POSTGRES_DB=service_db \
 		postgres:latest
	sleep 3 # postgres takes a hot second to be ready
	bash -c "trap 'trap - SIGINT SIGTERM ERR; podman stop make_postgres; exit 1; exit 1' SIGINT SIGTERM ERR; ${MAKE} run_internal"

PROJECT_DIR := $(shell pwd)
run_internal:
	go run ./cmd/app \
		--migrations-directory $(PROJECT_DIR)/migrations \
		--db-connection-string "postgres://init-microservice-go:password@localhost:5432/service_db?sslmode=disable"
