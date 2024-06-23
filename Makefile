# Go parameters
GO := go
BUILD_DIR := dist

# Determine OS and ARCH
OS ?= $(shell go env GOOS)
ARCH ?= $(shell go env GOARCH)

# Output binary name
OUTPUT := $(BUILD_DIR)/$(OS)/$(ARCH)/flux-graph

# Build the binary
build:
	mkdir -p $(BUILD_DIR)
	$(GO) build -o $(OUTPUT)
	@echo "Binary at: $(OUTPUT)"

lint: 
	@echo "==> Running golangci-lint"
	@golangci-lint run --config .golangci.yml

test:
	$(GO) test -v ./...

# Clean up
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build clean
