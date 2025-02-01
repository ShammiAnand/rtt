PROJECT_ROOT := $(shell pwd)

.PHONY: build test run lint

build:
	@echo "Building rtt binary..."
	@go build -o bin/rtt main.go

run: build
	@./bin/rtt

test:
	@go test -v ./...

lint:
	@echo "Running linters..."
	@golangci-lint run

