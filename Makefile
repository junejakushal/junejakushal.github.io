.PHONY: help build run test clean docker-build docker-run fmt vet lint deploy

help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Available targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  %-15s %s\n", $$1, $$2}'

build: ## Build the application
	go build -o audio-transcription main.go

run: ## Run the application locally
	go run main.go

test: ## Run tests
	go test -v -race -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

clean: ## Clean build artifacts
	rm -f audio-transcription
	rm -f coverage.out coverage.html
	rm -rf uploads/ outputs/

fmt: ## Format code
	gofmt -w .
	go mod tidy

vet: ## Run go vet
	go vet ./...

lint: fmt vet ## Run linters
	@echo "Linting complete"

docker-build: ## Build Docker image
	docker build -t audio-transcription:latest .

docker-run: ## Run Docker container locally
	docker run -p 8080:8080 --rm audio-transcription:latest

deploy: ## Deploy to Fly.io
	fly deploy

dev: ## Run in development mode (with auto-reload)
	@which air > /dev/null || (echo "Installing air..." && go install github.com/air-verse/air@latest)
	air

all: lint test build ## Run lint, test, and build
