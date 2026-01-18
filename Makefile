# Makefile for s3heck

BINARY_NAME := s3heck
BUILD_DIR := bin

# Git information
VERSION := $(shell git describe --tags --always --dirty)
COMMIT := $(shell git rev-parse --short HEAD)
DATE := $(shell date -u +%Y-%m-%dT%H:%M:%SZ)

# Linker flags
LDFLAGS := -ldflags "-X github.com/kanywst/s3heck/cmd.Version=$(VERSION) -X github.com/kanywst/s3heck/cmd.Commit=$(COMMIT) -X github.com/kanywst/s3heck/cmd.Date=$(DATE)"

.PHONY: all
all: lint test build

.PHONY: build
build:
	@echo "📦 Building $(BINARY_NAME)..."
	@mkdir -p $(BUILD_DIR)
	go build $(LDFLAGS) -o $(BUILD_DIR)/$(BINARY_NAME) .

.PHONY: run
run: build
	@echo "🚀 Running $(BINARY_NAME)..."
	@$(BUILD_DIR)/$(BINARY_NAME)

.PHONY: clean
clean:
	@echo "🧹 Cleaning..."
	@rm -rf $(BUILD_DIR)
	@rm -f demo.gif

.PHONY: test
test:
	@echo "🧪 Testing..."
	go test -v ./...

.PHONY: lint
lint:
	@echo "🔍 Linting..."
	@if command -v golangci-lint > /dev/null; then \
		golangci-lint run; \
	else \
		echo "golangci-lint not installed. Skipping."; \
	fi

.PHONY: demo
demo:
	@echo "🎥 Generating demo GIF..."
	vhs demo.tape

.PHONY: release-dry
release-dry:
	@echo "📦 Testing release build..."
	goreleaser release --snapshot --clean
