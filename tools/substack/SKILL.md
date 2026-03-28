---
name: substack
description: Substack CLI tool for accessing inbox posts, articles, and search. Use when you need to retrieve Substack reader data via command line.
---

# Substack Reader CLI

Part of the **entext-research-tool** skill suite.

## Installation

The tool is built from source to the AI skills directory's `scripts/` folder.

## Authentication

**First-time setup (2-step process):**

```bash
# Step 1: Set email address
substack profile -email "user@example.com"

# Step 2: Complete with authentication link (obtained from email)
substack auth -auth_string "https://substack.com/auth?token=..."
```

Session persists across CLI invocations and auto-refreshes.

## Commands

### profile

Set Substack email address for authentication.

```bash
substack profile -email "user@example.com"
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-email` | string | Substack email address | Yes |

**Examples:**
```bash
# Set email address
substack profile -email "user@example.com"
```

---

### auth

Complete Substack authentication with link or code from email.

```bash
substack auth -auth_string "https://substack.com/auth?token=..."
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-auth_string` | string | Authentication link or code from Substack email | Yes |

**Flow:**
1. Run `substack profile -email <email>`
2. Check your email for the authentication link
3. Run `substack auth -auth_string <link>`

**Examples:**
```bash
# Complete auth with link (for automation)
substack auth -auth_string "https://substack.com/auth?token=..."
```

---

### inbox

Get chronological inbox posts.

```bash
substack inbox
substack inbox -after "2024-01-01T00:00:00.000Z"
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-after` | string | Timestamp cursor for pagination (RFC3339 format) | No |

**Examples:**
```bash
# Get all inbox posts
substack inbox

# Paginate with timestamp cursor
substack inbox -after "2024-01-01T00:00:00.000Z"

# Extract titles with jq
substack inbox | jq '.data.posts[] | .title'

# Count posts
substack inbox | jq '.data.posts | length'
```

---

### article

Get article content by post ID.

```bash
substack article -post-id 123456
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-post-id` | uint | Post ID to retrieve content for | Yes |
| `-base-url` | string | Custom Substack domain (e.g., `substack.com`) | No |
| `-compact-mode` | bool | Compact content mode (default: false) | No |

**Examples:**
```bash
# Get article by ID
substack article -post-id 123456

# Use custom domain
substack article -post-id 123456 -base-url "substack.com"

# Compact mode
substack article -post-id 123456 -compact-mode
```

---

### search

Search Substack posts with different modes.

```bash
substack search -query "AI" -mode top
substack search -query "tech" -mode all
substack search -query "news" -mode subscribed
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-query` | string | Search query | Yes |
| `-mode` | string | Search mode: `top`, `all`, `subscribed` (default: `all`) | No |
| `-page` | int | Page number for pagination (0-10, not used for `top` mode) | No |
| `-language` | string | Language code (2-letter, e.g., `en`) | No |

**Search Modes:**
| Mode | Description |
|------|-------------|
| `top` | Top-ranked results |
| `all` | All posts (default) |
| `subscribed` | Subscribed publications only |

**Examples:**
```bash
# Top results
substack search -query "AI" -mode top

# All posts with pagination
substack search -query "technology" -mode all -page 1

# Subscribed publications only
substack search -query "newsletter" -mode subscribed

# Language filter
substack search -query "openclaw" -language en
```

---

## Examples

### Authentication

```bash
# Non-interactive authentication (automation)
substack profile -email "user@example.com"
substack auth -auth_string "https://substack.com/auth?token=..."
```

### Working with Posts

```bash
# Get inbox and extract titles
substack inbox | jq '.data.posts[] | .title'

# Search and get article
POST_ID=$(substack search -query "AI" -mode top | jq '.data.results[0].id')
substack article -post-id $POST_ID

# Count posts in inbox
substack inbox | jq '.data.posts | length'
```

---

## Managing Large Outputs

Article content and search results can be large. Pipe to files for processing:

```bash
# Article by ID
substack article -post-id 191095022 > article_191095022.json

# Search with query terms + timestamp
substack search -query "openclaw china" > search_openclaw_china_$(date +%Y%m%d).json

# Inbox with date
substack inbox > inbox_$(date +%Y%m%d).json

# Process file later
jq '.data.post.title' article_191095022.json
jq '.data.results[] | {title, id}' search_openclaw_china_$(date +%Y%m%d).json
```

**Filename conventions:**
- Include unique ID when available: `article_191095022.json`
- Include query terms for searches: `search_<query>_<date>.json`
- Use dated filenames for recurring commands: `inbox_YYYYMMDD.json`

---

## Session Location

Session file is stored in the scripts directory:

```
<skills-dir>/substack/scripts/session.json
```

**Note:** This directory is excluded from version control (`.gitignore`) to protect authentication tokens.

---

## Troubleshooting

**Command not found:**

Ensure the binary is in your PATH or use the full path:
```bash
<skills-dir>/substack/scripts/substack <command>
```

**Permission denied:**
```bash
chmod +x <skills-dir>/substack/scripts/substack
```

**Session expired:**
```bash
substack profile -email "user@example.com"
substack auth -auth_string "<link>"
```

**Email not set:**
```bash
substack profile -email "user@example.com"
```

---

## Requirements

- Go 1.21+ (for building from source)
- Substack account
