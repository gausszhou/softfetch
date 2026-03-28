# Makefile for SoftFetch project

# Variables
APP_NAME := softfetch
BUILD_DIR := bin
MAIN_PACKAGE := ./cmd/softfetch
VERSION := $(shell git describe --tags --always --dirty 2>/dev/null || echo "dev")
COMMIT := $(shell git rev-parse --short HEAD 2>/dev/null || echo "unknown")
BUILD_DATE := $(shell date -u '+%Y-%m-%d %H:%M:%S')
LDFLAGS := -ldflags "-X main.version=$(VERSION) -X main.commit=$(COMMIT) -X 'main.buildDate=$(BUILD_DATE)'"

# Default target
.PHONY: all
all: build

# Detect OS
ifeq ($(OS),Windows_NT)
    DETECTED_OS := Windows
else
    DETECTED_OS := $(shell uname -s)
endif

# Build the application
.PHONY: build
build:
	@echo "Building $(APP_NAME)..."
	@go build $(LDFLAGS) -o $(BUILD_DIR)/$(APP_NAME) $(MAIN_PACKAGE)
	@echo "Build complete: $(BUILD_DIR)/$(APP_NAME)"

# Run tests
.PHONY: test
test:
	@echo "Running tests..."
	@go test -v ./...

# Run tests with coverage
.PHONY: test-coverage
test-coverage:
	@echo "Running tests with coverage..."
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "Coverage report generated: coverage.html"

# Lint the code
.PHONY: lint
lint:
	@echo "Running linter..."
	@if command -v golangci-lint >/dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Install it with:"; \
		echo "go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest"; \
	fi

# Format the code
.PHONY: fmt
fmt:
	@echo "Formatting code..."
	@go fmt ./...

# Clean build artifacts
.PHONY: clean
clean:
	@echo "Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f coverage.out coverage.html
	@echo "Clean complete"

# Install dependencies
.PHONY: deps
deps:
	@echo "Installing dependencies..."
	@go mod download
	@go mod tidy

# Install the application
.PHONY: install
install: build
	@echo "Installing $(APP_NAME)..."
	@mkdir -p $$HOME/bin
	@cp -f $(BUILD_DIR)/$(APP_NAME) $$HOME/bin/
	@echo "Installed to $$HOME/bin/$(APP_NAME)"
	@echo "Add $$HOME/bin to your PATH to use $(APP_NAME) from anywhere"

# Uninstall the application
.PHONY: uninstall
uninstall:
	@echo "Uninstalling $(APP_NAME)..."
	@rm -f $$HOME/bin/$(APP_NAME)
	@echo "Removed from $$HOME/bin/"

# Run the application
.PHONY: run
run: build
	@echo "Running $(APP_NAME)..."
	@./$(BUILD_DIR)/$(APP_NAME)

# Docker build
.PHONY: docker-build
docker-build:
	@echo "Building Docker image..."
	@docker build -t $(APP_NAME):$(VERSION) .

# Help
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  build          - Build the application"
	@echo "  test           - Run tests"
	@echo "  test-coverage  - Run tests with coverage"
	@echo "  lint           - Lint the code"
	@echo "  fmt            - Format the code"
	@echo "  clean          - Clean build artifacts"
	@echo "  deps           - Install dependencies"
	@echo "  install        - Install the application"
	@echo "  uninstall      - Uninstall the application"
	@echo "  run            - Run the application"
	@echo "  docker-build   - Build Docker image"
	@echo "  help           - Show this help message"