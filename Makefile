.PHONY: help build clean test vet run-auth run-inbox run-search run-article release-all release-linux release-darwin release-windows

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOVET=$(GOCMD) vet
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod

# Binary names
BINARY_NAME=substack
BINARY_DIR=bin

# Source directories
SRC_DIR=tools/substack-reader/src
INTERNAL_DIR=internal

# Build output paths
BUILD_LINUX=$(BINARY_DIR)/linux/$(BINARY_NAME)
BUILD_DARWIN=$(BINARY_DIR)/darwin/$(BINARY_NAME)
BUILD_WINDOWS=$(BINARY_DIR)/windows/$(BINARY_NAME).exe
BUILD_LOCAL=$(BINARY_DIR)/$(BINARY_NAME)

# Default target
help:
	@echo "Substack CLI - Makefile Commands"
	@echo "================================"
	@echo ""
	@echo "Build:"
	@echo "  make build          - Build for current platform"
	@echo "  make release-all    - Build for all platforms (linux, darwin, windows)"
	@echo "  make release-linux  - Build for Linux"
	@echo "  make release-darwin - Build for macOS"
	@echo "  make release-windows- Build for Windows"
	@echo ""
	@echo "Test & Lint:"
	@echo "  make test           - Run tests"
	@echo "  make vet            - Run go vet"
	@echo "  make clean          - Clean build artifacts"
	@echo ""
	@echo "Run (requires session):"
	@echo "  make run-auth       - Run auth command"
	@echo "  make run-inbox      - Run inbox command"
	@echo "  make run-search     - Run search command"
	@echo "  make run-article    - Run article command"
	@echo ""
	@echo "Dependencies:"
	@echo "  make deps           - Download dependencies"
	@echo "  make tidy           - Tidy go.mod"

# Build for current platform
build: $(BINARY_DIR)
	$(GOBUILD) -o $(BUILD_LOCAL) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_LOCAL)"

# Create bin directory
$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)
	mkdir -p $(BINARY_DIR)/linux
	mkdir -p $(BINARY_DIR)/darwin
	mkdir -p $(BINARY_DIR)/windows

# Build for all platforms
release-all: release-linux release-darwin release-windows
	@echo "✓ Built all platforms"

# Build for Linux
release-linux: $(BINARY_DIR)
	GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BUILD_LINUX) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_LINUX)"

# Build for macOS (Intel & Apple Silicon)
release-darwin: $(BINARY_DIR)
	GOOS=darwin GOARCH=amd64 $(GOBUILD) -o $(BUILD_DARWIN)_amd64 ./$(SRC_DIR)/
	GOOS=darwin GOARCH=arm64 $(GOBUILD) -o $(BUILD_DARWIN)_arm64 ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_DARWIN)_amd64 and $(BUILD_DARWIN)_arm64"

# Build for Windows
release-windows: $(BINARY_DIR)
	GOOS=windows GOARCH=amd64 $(GOBUILD) -o $(BUILD_WINDOWS) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_WINDOWS)"

# Run tests
test:
	$(GOTEST) -v ./$(INTERNAL_DIR)/...

# Run go vet
vet:
	$(GOVET) ./...

# Clean build artifacts
clean:
	rm -rf $(BINARY_DIR)
	@echo "✓ Cleaned build artifacts"

# Run auth command
run-auth: build
	./$(BUILD_LOCAL) auth

# Run inbox command
run-inbox: build
	./$(BUILD_LOCAL) inbox

# Run search command
run-search: build
	./$(BUILD_LOCAL) search -query "test"

# Run article command
run-article: build
	./$(BUILD_LOCAL) article -post-id 123456

# Download dependencies
deps:
	$(GOGET) ./...

# Tidy go.mod
tidy:
	$(GOMOD) tidy

# Install to local bin
install: build
	cp $(BUILD_LOCAL) ~/bin/$(BINARY_NAME)
	@echo "✓ Installed to ~/bin/$(BINARY_NAME)"

# Uninstall from local bin
uninstall:
	rm -f ~/bin/$(BINARY_NAME)
	@echo "✓ Uninstalled from ~/bin/$(BINARY_NAME)"
