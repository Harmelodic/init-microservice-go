#!/usr/bin/env bash

# ==== CI PIPELINES ====
build: generate
	go build -o app -v ./...

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

# ==== DEV SCRIPTS ====
run: install
	go run ./src
