---
name: koyfin
description: Koyfin CLI tools for financial data access including stock search, snapshot data, time series, earnings transcripts, ETF holdings, and stock screener. Use when you need to query Koyfin financial data from the command line.
---

# Koyfin CLI Tools

## Installation

The tool is built from source to the AI skills directory's `scripts/` folder. Python utilities are copied to the same directory.

## Commands

### auth

Authenticate with Koyfin using email and password.

```bash
koyfin auth -email "user@example.com" -password "secret"
```

| Flag | Type | Description | Required |
|------|------|-------------|----------|
| `-email` | string | Koyfin email address | Yes |
| `-password` | string | Koyfin password | Yes |

**Examples:**
```bash
# Non-interactive mode (for automation)
koyfin auth -email "user@example.com" -password "secret"
```

Session persists across CLI invocations and tokens are refreshed automatically.

---

### search

Search for stocks/tickers by name.

```bash
koyfin search -q "Apple"
koyfin search -q "SPY" -categories "ETF"
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-q` | string | Ticker or ETF name to search for | - |
| `-categories` | string | Search categories (comma-separated) | `Equity,ETF` |
| `-primary-only` | bool | Use primary exchange only | `false` |

**Examples:**
```bash
# Search for Apple stock
koyfin search -q "Apple"

# Search for ETF
koyfin search -q "SPY" -categories "ETF"

# Primary exchange only
koyfin search -q "Microsoft" -primary-only
```

---

### snapshot

Get snapshot data with financial metrics.

```bash
koyfin snapshot -kids "KID1,KID2,KID3" -category Equity
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-kids` | string | Comma-separated list of Koyfin IDs | - |
| `-category` | string | Category: `Equity` or `ETF` | `Equity` |

**Limits:**
| Category | Max KIDs |
|----------|----------|
| Equity | 32 |
| ETF | 2 |

**Examples:**
```bash
# Get snapshot for multiple stocks
koyfin snapshot -kids "KID1,KID2,KID3" -category Equity

# Get snapshot for ETFs
koyfin snapshot -kids "KID1,KID2" -category ETF
```

---

### ticker-data

Get time series data for a ticker.

```bash
koyfin ticker-data -kid "KID" -key "p_candle_range" -date-from "2024-01-01"
koyfin ticker-data -kid "KID" -key "f_r" -date-from "2020-01-01" -fin-period "quarterly"
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-kid` | string | Koyfin ID for the ticker | - |
| `-key` | string | Indicator key to search for | - |
| `-date-from` | string | Start date in YYYY-MM-DD format | - |
| `-date-to` | string | End date in YYYY-MM-DD format | Today |
| `-currency` | string | Data currency | `USD` |
| `-agg-period` | string | Series granularity: `day`, `monthly`, `quarterly`, `annually` | `day` |
| `-price-format` | string | Price format: `both`, `standard`, `adj` (auto-set based on key) | - |
| `-fin-period` | string | Financial period: `quarterly`, `annual`, `LTM` | - |

**Examples:**
```bash
# Get 1 year price data
koyfin ticker-data -kid "KID" -key "p_candle_range" -date-from "2024-01-01"

# Get quarterly revenue
koyfin ticker-data -kid "KID" -key "f_r" -date-from "2020-01-01" -fin-period "quarterly"

# Get data in different currency
koyfin ticker-data -kid "KID" -key "p_candle_range" -date-from "2024-01-01" -currency "EUR"
```

---

### transcript

Earnings call transcripts (list, get, summary).

```bash
# List transcripts
koyfin transcript -action list -kid "KID" -limit 5

# Get specific transcript
koyfin transcript -action get -transcript-id 12345

# Get transcript summary
koyfin transcript -action summary -transcript-id 12345
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-action` | string | Action: `list`, `get`, `summary` | `list` |
| `-kid` | string | Koyfin ID (required for `list`) | - |
| `-transcript-id` | int | Key development identifier (required for `get`/`summary`) | - |
| `-limit` | int | Maximum results for `list` action (1-64) | `10` |

**Examples:**
```bash
# List recent transcripts
koyfin transcript -action list -kid "KID" -limit 5

# Get transcript by ID
koyfin transcript -action get -transcript-id 12345

# Get AI-generated summary
koyfin transcript -action summary -transcript-id 12345
```

---

### schema

Get indicator schema.

```bash
koyfin schema -asset-type Equity -indicator-type financials
koyfin schema -asset-type Equity -indicator-type ratios
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-asset-type` | string | Asset type | `Equity` |
| `-indicator-type` | string | Indicator type: `financials`, `ratios`, `forward_estimates`, `market_data` | - |

**Examples:**
```bash
# Get financials schema
koyfin schema -asset-type Equity -indicator-type financials

# Get ratios schema
koyfin schema -asset-type Equity -indicator-type ratios

# Get market data schema
koyfin schema -asset-type Equity -indicator-type market_data
```

---

### etf-holdings

Get ETF holdings.

```bash
koyfin etf-holdings -kids "KID1,KID2" -category ETF
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-kids` | string | Comma-separated list of Koyfin IDs for ETFs (max 2) | - |
| `-category` | string | Category (must be `ETF`) | `ETF` |

**Examples:**
```bash
# Get holdings for single ETF
koyfin etf-holdings -kids "KID1" -category ETF

# Get holdings for multiple ETFs
koyfin etf-holdings -kids "KID1,KID2" -category ETF
```

---

### screener-schema

Get screener filter schema.

```bash
koyfin screener-schema -asset-type Equity
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-asset-type` | string | Asset type | `Equity` |

**Examples:**
```bash
# Get screener schema for Equity
koyfin screener-schema -asset-type Equity
```

---

### screener

Run stock screener.

```bash
koyfin screener -filters '[{"key":"mkt","min":10000}]'
koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":1000,"max":10000}]'
```

| Flag | Type | Description | Default |
|------|------|-------------|---------|
| `-filters` | string | JSON array of filter conditions | - |
| `-page-size` | int | Page size (max 300) | `100` |

**Filter Types & Units:**

| Filter Key | Unit | Example |
|------------|------|---------|
| `mkt` | Millions USD | `{"min":1000}` = >$1B |
| `evebitdaltm` | Ratio (not %) | `{"max":10}` = <10x |
| `pf_fcf_margin-LTM` | Decimal (not %) | `{"min":0.10}` = >10% |
| `f_bet1yr` | Decimal (not %) | `{"min":0.05}` = >5% |
| `chg1mPct_L`, `chg3mPct_L`, `chgYTDPct_L` | Decimal (not %) | `{"min":0.10}` = >10% |
| `pe_ratio` | Ratio | `{"max":20}` = <20x |
| `pb_ratio` | Ratio | `{"max":3}` = <3x |
| `div_yield` | Decimal (not %) | `{"min":0.03}` = >3% |
| `iso2` | Country code | `{"values":["JP","US"]}` |
| `t_sec` | Sector name | `{"values":["Information Technology"]}` |

**Examples:**
```bash
# Large cap (>10B)
koyfin screener -filters '[{"key":"mkt","min":10000}]'

# Tech sector, 1B-10B market cap
koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":1000,"max":10000}]'

# EV/EBITDA < 10
koyfin screener -filters '[{"key":"evebitdaltm","max":10}]'

# Custom page size
koyfin screener -filters '[{"key":"mkt","min":10000}]' -page-size 50
```

---

## Python Utilities

Location: `scripts/` (same directory as binary)

| File | Description |
|------|-------------|
| `excel_export.py` | Export snapshot data to Excel with multiple sheets |
| `process.py` | Format output as tables |
| `requirements.txt` | Python dependencies for Excel export |

### Excel Export (Snapshot to XLSX)

Export snapshot data to Excel with multiple sheets:

```bash
koyfin snapshot -kids "KID1,KID2,KID3" | python3 excel_export.py -o snapshot.xlsx
```

**Excel Sheets Created:**

| Sheet | Description |
|-------|-------------|
| **Summary** | Key metrics: Price, Market Cap, P/E, EV/EBITDA, Margins, Growth |
| **Raw Data** | All snapshot metrics from API |
| **Ratios** | Calculated ratios: P/E, EV/EBITDA, EV/Sales, Margins, P/FCF |
| **Growth** | Growth rates: 1Y/3Y/5Y price CAGR, YTD, estimate vs LTM |

### Install Dependencies (for Excel export)

```bash
pip3 install -r requirements.txt
```

### Format Output

```bash
# Format search results as table
koyfin search -q "Apple" | python3 process.py search

# Format snapshot
koyfin snapshot -kids "KID1,KID2" | python3 process.py snapshot

# Format time series
koyfin ticker-data -kid "KID" -key "p_candle_range" -date-from "2024-01-01" | python3 process.py ticker
```

---

## Examples

```bash
# Non-interactive authentication
koyfin auth -email "user@example.com" -password "secret"

# Search for Apple
koyfin search -q "Apple"

# Get snapshot
koyfin snapshot -kids "KID1,KID2" -category Equity

# Get 1 year price data
koyfin ticker-data -kid "KID" -key "p_candle_range" -date-from "2024-01-01"

# List recent transcripts
koyfin transcript -action list -kid "KID" -limit 5

# Get transcript
koyfin transcript -action get -transcript-id 12345

# Get transcript summary
koyfin transcript -action summary -transcript-id 12345

# Screen for large cap tech
koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":10000}]' -page-size 50

# Export snapshot to Excel
koyfin snapshot -kids "KID1,KID2" | python3 excel_export.py -o snapshot.xlsx
```

---

## Session Location

Session file is stored in the scripts directory:

```
<skills-dir>/koyfin/scripts/session.json
```

**Note:** This directory is excluded from version control (`.gitignore`) to protect authentication tokens.

---

## Troubleshooting

**Command not found:**

Ensure the binary is in your PATH or use the full path:
```bash
<skills-dir>/koyfin/scripts/koyfin <command>
```

**Permission denied:**
```bash
chmod +x <skills-dir>/koyfin/scripts/koyfin
```

**Session expired:**
```bash
koyfin auth -email "user@example.com" -password "secret"
```

---

## Requirements

- Go 1.21+ (for building from source)
- Koyfin account
- Python 3 (optional, for Excel export utilities)
