#!/bin/bash
# Build script for chart-maker tools
# Usage: ./tool_build.sh [tool-name]
#
# Available tools:
#   koyfin    - Koyfin CLI tools
#
# If no tool specified, shows available tools.

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
TOOLS_DIR="$SCRIPT_DIR/tools"

# Show usage if no argument provided
if [ -z "$1" ]; then
    echo "Usage: $0 <tool-name>"
    echo ""
    echo "Available tools:"
    for dir in "$TOOLS_DIR"/*/; do
        if [ -d "$dir" ] && [ -f "$dir/setup.sh" ]; then
            echo "  $(basename "$dir")"
        fi
    done
    echo ""
    echo "Examples:"
    echo "  $0 substack-reader"
    exit 0
fi

TOOL_NAME="$1"
TOOL_DIR="$TOOLS_DIR/$TOOL_NAME"

# Check if tool directory exists
if [ ! -d "$TOOL_DIR" ]; then
    echo "Error: Tool '$TOOL_NAME' not found"
    echo ""
    echo "Available tools:"
    for dir in "$SCRIPT_DIR"/*/; do
        if [ -d "$dir" ] && [ -f "$dir/setup.sh" ]; then
            echo "  $(basename "$dir")"
        fi
    done
    exit 1
fi

# Check if tool has setup.sh
if [ -f "$TOOL_DIR/setup.sh" ]; then
    bash "$TOOL_DIR/setup.sh"
else
    echo "Error: No setup.sh found for '$TOOL_NAME'"
    exit 1
fi
