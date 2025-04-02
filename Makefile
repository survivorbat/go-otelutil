MAKEFLAGS := --no-print-directory --silent

default: help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z\._-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

fmt: ## Format go code, tidy the go.mod file and run the linter
	@go mod tidy
	@gofumpt -l -w .
	@golangci-lint run --fix

tools: ## Install extra tools for development
	go install mvdan.cc/gofumpt@latest
	@echo "Make sure to have golangci-lint v2 installed" # v2 is hard to install platform independently

lint: ## Lint the code locally
	golangci-lint run

test: ## Run unit-tests
	go test ./... -timeout=40s -parallel=20

