PROJECT_ROOT := $(shell pwd)

.PHONY: build test run

build:
	@echo "Building `rtt` binary..."
	@mkdir -p bin
	@go build -o bin/rtt cmd/main.go

run: build
	@./bin/rtt

test:
	@go test -v ./...

clean:
	@echo "Cleaning up..."
	@rm -rf bin/
	@rm -f coverage.out coverage.html

lint:
	@echo "Running linters..."
	@golangci-lint run

