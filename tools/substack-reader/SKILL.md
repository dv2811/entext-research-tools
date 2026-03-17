---
name: substack-reader
description: Substack CLI tool - part of entext-research-tool skill suite. Access Substack inbox, articles, and search from the command line.
---

# Substack Reader CLI

Part of the **entext-research-tool** skill suite.

## Installation

```bash
./tool_build.sh substack-reader
```

## Authentication

```bash
substack auth
```

Enter your email, then paste the authentication URL from your email. Session is saved automatically.

## Commands

### inbox

Get chronological inbox posts.

```bash
substack inbox
substack inbox -after "2024-01-01T00:00:00.000Z"
```

### article

Get article content by post ID.

```bash
substack article -post-id 123456
```

### search

Search Substack posts.

```bash
substack search -query "AI" -mode top
substack search -query "tech" -mode all
substack search -query "news" -mode subscribed
```

**Search Modes:**
- `top` - Top-ranked results
- `all` - All posts (default)
- `subscribed` - Subscribed publications only

## Examples

```bash
# Get inbox and extract titles
substack inbox | jq '.data.posts[] | .title'

# Search and get article
POST_ID=$(substack search -query "AI" -mode top | jq '.data.results[0].id')
substack article -post-id $POST_ID
```

## Output

All commands output JSON to stdout for piping to `jq`.

## Session Location

- **Linux:** `~/.config/substack-reader/session.json`
- **macOS:** `~/Library/Application Support/substack-reader/session.json`
- **Windows:** `%APPDATA%\substack-reader\session.json`
