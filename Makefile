#!/usr/bin/env bash

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
	${GOPATH}/bin/golangci-lint run

install:
	go mod download
	go install github.com/golangci/golangci-lint/v2/cmd/golangci-lint@latest
