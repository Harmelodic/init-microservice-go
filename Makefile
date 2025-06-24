#!/usr/bin/env bash

# ==== CI PIPELINES ====
build: clean generate
	go build -o ./bin/app -v ./cmd/app

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
run: install
	go run ./cmd/app
