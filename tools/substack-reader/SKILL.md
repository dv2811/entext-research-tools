---
name: substack-cli
description: Substack CLI tools for reading Substack content including inbox posts, article retrieval, and post search. Use when you need to access Substack reader data from the command line.
---

# Substack Reader CLI Tools

## Authentication

Use the `auth` command to authenticate:

```bash
substack auth
```

Enter your email, then paste the authentication URL from your email. Session is saved automatically.

## CLI Commands

### auth

Authenticate with Substack via email link.

```bash
substack auth
```

### inbox

Get chronological inbox posts.

```bash
substack inbox
substack inbox -after "2024-01-01T00:00:00.000Z"
```

| Flag | Description | Default |
|------|-------------|---------|
| `-after` | Timestamp cursor (RFC3339) | - |

### article

Get article content by post ID.

```bash
substack article -post-id 123456
substack article -post-id 123456 -base-url "substack.com"
```

| Flag | Description | Default |
|------|-------------|---------|
| `-post-id` | Post ID (required) | - |
| `-base-url` | Custom domain | - |

### search

Search Substack posts.

```bash
substack search -query "AI" -mode top
substack search -query "tech" -mode all -page 1
substack search -query "news" -mode subscribed -language en
```

| Flag | Description | Default |
|------|-------------|---------|
| `-query` | Search query (required) | - |
| `-mode` | top, all, subscribed | all |
| `-page` | Page number (0-10) | 0 |
| `-language` | 2-letter language code | - |

**Search Modes:**
- **top**: Top-ranked results (cursor pagination)
- **all**: All posts (page pagination)
- **subscribed**: Subscribed publications only

## Usage Examples

### Authenticate and get inbox

```bash
substack auth -email "user@example.com"
substack inbox | jq '.data.posts[] | {title, author: .publishedBylines[0].name}'
```

### Search and retrieve article

```bash
POST_ID=$(substack search -query "AI" -mode top | jq '.data.results[0].id')
substack article -post-id $POST_ID
```

## Output

All commands output JSON to stdout:

```bash
substack inbox | jq '.data.posts | length'
substack search -query "AI" | jq '.data.results[].title'
```

## Session Location

- Default: `~/.config/substack-reader/session.json` (Linux)
- macOS: `~/Library/Application Support/substack-reader/session.json`
- Windows: `%APPDATA%\substack-reader\session.json`
- Custom: Set `SUBSTACK_SESSION_FILE`
