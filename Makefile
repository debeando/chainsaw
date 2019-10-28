SHELL := /bin/bash

.PHONY: help

help: ## Show this help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-40s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

deps: ## Install dependencies
	go get -u github.com/golang/dep/cmd/dep
	go get -u github.com/go-sql-driver/mysql

tests: ## Run tests
	go test -v ./...

build: ## Build binary for local operating system
	go build -o chainsaw main.go
