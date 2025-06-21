# Makefile for Fiber Boilerplate

# Variables
APP_NAME=fiber-boilerplate
GO_FILES=$(shell find . -name "*.go" -type f)

# Default target
.PHONY: all
all: help

# Help target
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make dev       - Run the application with Air for hot reloading"
	@echo "  make build     - Build the application"
	@echo "  make run       - Run the application without hot reloading"
	@echo "  make clean     - Clean build artifacts"
	@echo "  make test      - Run tests"
	@echo "  make help      - Show this help message"

# Development with hot reloading
.PHONY: dev
dev:
	@echo "Starting development server with Air..."
	@./dev.sh

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go build -o $(APP_NAME) .

# Run without hot reloading
.PHONY: run
run:
	@echo "Running $(APP_NAME)..."
	@go run main.go

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning build artifacts..."
	@rm -rf $(APP_NAME) tmp/

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...
