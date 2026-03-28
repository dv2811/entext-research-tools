# entext-research-tools

Go-based CLI tools for accessing external content platforms.

## Requirements

- **Go 1.21+** (required for building from source)
- Platform account for each tool (Substack, Koyfin, etc.)

## Installation

### Option 1: Build from source (recommended)

```bash
# Clone the repository
git clone https://github.com/dv2811/entext-research-tools.git
cd entext-research-tools

# Build and deploy a tool
./tool_build.sh <tool-name> <skills-dir>
```

**Arguments:**
- `<tool-name>` - Tool to build (`substack`, `koyfin`)
- `<skills-dir>` - AI skills directory for deployment

**Example:**
```bash
./tool_build.sh substack /path/to/ai-tool-config/skills
./tool_build.sh koyfin /path/to/ai-tool-config/skills
```

### Option 2: Download pre-built binary

Download the release archive for your platform from [Releases](https://github.com/dv2811/entext-research-tools/releases).

**Available platforms:**
| OS | Architectures |
|----|---------------|
| macOS | `amd64`, `arm64` (M1/M2) |
| Linux | `amd64`, `arm64`, `arm` |
| Windows | `amd64`, `386` |

**Archive structure:**
```
<tool>-<os>-<arch>.tar.gz
├── scripts/
│   └── <tool>         # Binary
├── SKILL.md           # Documentation
└── utils/             # Additional utilities (if applicable)
```

Extract and copy contents to your AI skills directory.

## Project Structure

```
entext-research-tools/
├── .github/
│   └── workflows/
│       └── release.yml      # CI/CD for building releases
├── .claude-plugin/
│   └── plugin.json          # Plugin manifest
├── tools/
│   ├── <tool-name>/         # CLI tools
│   │   ├── setup.sh         # Tool-specific messages
│   │   ├── SKILL.md         # Tool documentation
│   │   ├── src/             # Source code
│   │   └── utils/           # Additional utilities (optional)
├── internal/
│   ├── substack/            # Substack API client
│   ├── koyfin/              # Koyfin API client
│   └── ...
├── bin/
│   ├── build.sh             # Script to build all platform binaries
│   ├── darwin/              # Pre-built macOS binaries
│   ├── linux/               # Pre-built Linux binaries
│   └── windows/             # Pre-built Windows binaries
├── go.mod
├── tool_build.sh            # Main build script
└── README.md
```

## Available Tools

| Tool | Description | Documentation |
|------|-------------|---------------|
| `substack` | Substack inbox, articles, and search | [tools/substack/SKILL.md](tools/substack/SKILL.md) |
| `koyfin` | Koyfin financial data (snapshots, time series, transcripts, screener) | [tools/koyfin/SKILL.md](tools/koyfin/SKILL.md) |

## Quick Start

### Substack

```bash
# Build and deploy
./tool_build.sh substack /path/to/skills

# Authenticate (2-step)
/path/to/skills/substack/scripts/substack profile -email "user@example.com"
/path/to/skills/substack/scripts/substack auth -auth_string "<link-from-email>"

# Get inbox
/path/to/skills/substack/scripts/substack inbox
```

### Koyfin

```bash
# Build and deploy
./tool_build.sh koyfin /path/to/skills

# Authenticate
/path/to/skills/koyfin/scripts/koyfin auth -email "user@example.com" -password "<password>"

# Search for a ticker
/path/to/skills/koyfin/scripts/koyfin search -q "Apple"

# Export snapshot to Excel
/path/to/skills/koyfin/scripts/koyfin snapshot -kids "KID1,KID2" | \
  python3 /path/to/skills/koyfin/scripts/excel_export.py -o snapshot.xlsx
```

## Building Releases

To build pre-built binaries for all platforms:

```bash
./bin/build.sh
```

This creates binaries in `bin/<os>/` directories for distribution.

## Tool-Specific Documentation

- **substack:** [tools/substack/SKILL.md](tools/substack/SKILL.md)
- **koyfin:** [tools/koyfin/SKILL.md](tools/koyfin/SKILL.md)
