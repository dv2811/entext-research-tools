# Load environment variables from .env file
-include .env

# Binary names
BINARY_NAME=substack
BINARY_DIR=bin

# Source directories
SRC_DIR=tools/substack-reader/src
INTERNAL_DIR=internal
SETUP_SCRIPT=tools/substack-reader/setup.sh

# Build output paths
BUILD_LINUX=$(BINARY_DIR)/linux/$(BINARY_NAME)
BUILD_DARWIN=$(BINARY_DIR)/darwin/$(BINARY_NAME)
BUILD_WINDOWS=$(BINARY_DIR)/windows/$(BINARY_NAME).exe
BUILD_LOCAL=$(BINARY_DIR)/$(BINARY_NAME)

.PHONY: help
help: ## Print this help message
	@echo "Substack CLI - Makefile Commands"
	@echo "================================"
	@echo ""
	@awk -F ':.*## ' '/^[a-zA-Z_-]+:.*## / {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""

.PHONY: build release-all release-linux release-darwin release-windows
build: $(BINARY_DIR) ## Build for current platform
	go build -o $(BUILD_LOCAL) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_LOCAL)"

$(BINARY_DIR):
	mkdir -p $(BINARY_DIR)
	mkdir -p $(BINARY_DIR)/linux
	mkdir -p $(BINARY_DIR)/darwin
	mkdir -p $(BINARY_DIR)/windows

release-all: release-linux release-darwin release-windows ## Build for all platforms
	@echo "✓ Built all platforms"

release-linux: $(BINARY_DIR) ## Build for Linux (amd64)
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_LINUX) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_LINUX)"

release-darwin: $(BINARY_DIR) ## Build for macOS (Intel + Apple Silicon)
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DARWIN)_amd64 ./$(SRC_DIR)/
	GOOS=darwin GOARCH=arm64 go build -o $(BUILD_DARWIN)_arm64 ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_DARWIN)_amd64 and $(BUILD_DARWIN)_arm64"

release-windows: $(BINARY_DIR) ## Build for Windows (amd64)
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_WINDOWS) ./$(SRC_DIR)/
	@echo "✓ Built: $(BUILD_WINDOWS)"

.PHONY: setup install uninstall uninstall-all clean-session
setup: ## Run full setup (build + install + auth session)
	@echo "Running full setup..."
	@bash $(SETUP_SCRIPT)

install: build ## Build and install binary to platform-specific path
	@echo "Installing to platform-specific location..."
	@if [ "$$(uname)" = "Darwin" ]; then \
		mkdir -p $$HOME/bin && cp $(BUILD_LOCAL) $$HOME/bin/$(BINARY_NAME) && chmod +x $$HOME/bin/$(BINARY_NAME); \
		echo "✓ Installed to $$HOME/bin/$(BINARY_NAME)"; \
	elif [ "$$(uname)" = "Linux" ]; then \
		mkdir -p $$HOME/.local/bin && cp $(BUILD_LOCAL) $$HOME/.local/bin/$(BINARY_NAME) && chmod +x $$HOME/.local/bin/$(BINARY_NAME); \
		echo "✓ Installed to $$HOME/.local/bin/$(BINARY_NAME)"; \
	else \
		echo "Manual installation required for this platform"; \
		echo "Copy $(BUILD_LOCAL) to your preferred location"; \
	fi

uninstall: ## Remove binary from platform-specific path
	@echo "Removing binary from platform-specific location..."
	@if [ "$$(uname)" = "Darwin" ]; then \
		if [ -f $$HOME/bin/$(BINARY_NAME) ]; then \
			rm -f $$HOME/bin/$(BINARY_NAME) && echo "✓ Removed from $$HOME/bin/$(BINARY_NAME)"; \
		else \
			echo "Binary not found at $$HOME/bin/$(BINARY_NAME)"; \
		fi; \
	elif [ "$$(uname)" = "Linux" ]; then \
		if [ -f $$HOME/.local/bin/$(BINARY_NAME) ]; then \
			rm -f $$HOME/.local/bin/$(BINARY_NAME) && echo "✓ Removed from $$HOME/.local/bin/$(BINARY_NAME)"; \
		else \
			echo "Binary not found at $$HOME/.local/bin/$(BINARY_NAME)"; \
		fi; \
	else \
		echo "Manual uninstallation required for this platform"; \
		echo "Remove the binary from your installation directory"; \
	fi

uninstall-all: uninstall clean-session ## Remove binary + session file + config
	@echo ""
	@echo "Full uninstall complete."
	@echo "To remove SKILL.md from config directory, run:"
	@echo "  rm -rf ~/.config/substack-reader  # Linux"
	@echo "  rm -rf ~/Library/Application\\ Support/substack-reader  # macOS"

clean-session: ## Remove only the session file (for security)
	@echo "Removing session file..."
	@if [ -n "$$SUBSTACK_SESSION_FILE" ]; then \
		rm -f "$$SUBSTACK_SESSION_FILE" && echo "✓ Removed: $$SUBSTACK_SESSION_FILE"; \
	elif [ "$$(uname)" = "Darwin" ]; then \
		rm -f $$HOME/Library/Application\\ Support/substack-reader/session.json && echo "✓ Removed session file (macOS)"; \
	elif [ "$$(uname)" = "Linux" ]; then \
		rm -f $$HOME/.config/substack-reader/session.json && echo "✓ Removed session file (Linux)"; \
	else \
		echo "Session file location unknown - remove manually"; \
	fi

.PHONY: test vet clean
test: ## Run tests
	go test -v ./$(INTERNAL_DIR)/...

vet: ## Run go vet
	go vet ./...

clean: ## Clean build artifacts
	rm -rf $(BINARY_DIR)
	@echo "✓ Cleaned build artifacts"

.PHONY: run-auth run-inbox run-search run-article
run-auth: build ## Run auth command
	./$(BUILD_LOCAL) auth

run-inbox: build ## Run inbox command
	./$(BUILD_LOCAL) inbox

run-search: build ## Run search command with test query
	./$(BUILD_LOCAL) search -query "test"

run-article: build ## Run article command (requires valid post-id)
	./$(BUILD_LOCAL) article -post-id 123456

.PHONY: deps tidy
deps: ## Download dependencies
	go get ./...

tidy: ## Tidy go.mod
	go mod tidy
