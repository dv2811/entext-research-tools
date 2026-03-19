#!/bin/bash
set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(dirname "$SCRIPT_DIR")"

# Tool-specific configuration
TOOL_NAME="koyfin"
BINARY_NAME="koyfin"

# Binary and session file directory: <tool_name>/scripts/
BINARY_DIR="$SCRIPT_DIR/scripts"
BIN_FILE="$BINARY_DIR/$BINARY_NAME"

echo "$TOOL_NAME CLI Tools Setup"
echo "======================"
echo ""

# Create directories
mkdir -p "$BINARY_DIR"

# Done
echo ""
echo "======================"
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
