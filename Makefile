#!/usr/bin/env bash

build: test
	go build -o app -v ./...

test: generate
	go test -v -race ./...

generate: lint
	go generate ./...

lint:
	go mod verify
	go mod tidy
	go fmt ./...
	go vet ./...
	golangci-lint run
