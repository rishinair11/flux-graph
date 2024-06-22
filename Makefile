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

# Clean up
clean:
	rm -rf $(BUILD_DIR)

.PHONY: build clean
