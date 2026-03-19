#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Tool-specific configuration
TOOL_NAME="substack-reader"
BINARY_NAME="substack"

# Binary and session file directory: <tool_name>/scripts/
BINARY_DIR="$SCRIPT_DIR/scripts"
BIN_FILE="$BINARY_DIR/$BINARY_NAME"

echo "$TOOL_NAME CLI Tools Setup"
echo "================================"
echo ""

# Create directories
mkdir -p "$BINARY_DIR"

# Determine target OS
if [[ "$OSTYPE" == "darwin"* ]]; then
    TARGET_OS="darwin"
elif [[ "$OSTYPE" == "msys" || "$OSTYPE" == "win32" ]]; then
    TARGET_OS="windows"
else
    TARGET_OS="linux"
fi

# Build from source
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed"
    echo ""
    echo "Please install Go from https://go.dev/dl/"
    exit 1
fi

echo "Building from source..."
echo "✓ Go: $(go version)"
echo "✓ Target OS: $TARGET_OS"
GOOS="$TARGET_OS" go build -ldflags='-w -s' -o "$BIN_FILE" "$SCRIPT_DIR/src/"
chmod +x "$BIN_FILE"
echo "✓ Built: $BIN_FILE"

# Done
echo ""
echo "================================"
echo "Setup complete!"
echo ""
echo "Binary and session file location:"
echo "  $BINARY_DIR"
echo ""
echo "Next step: Authenticate"
echo "  $BIN_FILE auth"
echo ""
echo "Usage:"
echo "  $BIN_FILE <command> -h"
