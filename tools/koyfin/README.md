# Koyfin CLI Tools

Command-line tools for accessing Koyfin financial data. Search stocks, get snapshots, fetch time series data, retrieve earnings call transcripts, analyze ETF holdings, and run stock screeners.

## Overview

```bash
# Search for a stock
./scripts/koyfin search -q "Apple"

# Get snapshot data for multiple tickers
./scripts/koyfin snapshot -kids <list_of_koyfin_ids> -category Equity

# Get 1 year of price data
./scripts/koyfin ticker-data -id <koyfin_id> -key "p_candle_range" -date-from "2024-01-01"

# List recent earnings call transcripts
./scripts/koyfin transcript -action list -kid "AAPL:US" -limit 5

# Run a stock screener (large cap tech)
./scripts/koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":10000}]' -page-size 50
```

## Commands

| Command | Description |
|---------|-------------|
| `search` | Search for stocks/ETFs by name |
| `snapshot` | Get current financial metrics |
| `ticker-data` | Get time series data |
| `transcript` | Earnings call transcripts |
| `schema` | Get indicator schema reference |
| `etf-holdings` | Get ETF holdings data |
| `screener-schema` | Get screener filter schema |
| `screener` | Run stock screener with filters |

### search

Search for stocks or ETFs by name.

```bash
./scripts/koyfin search -q "Apple"
./scripts/koyfin search -q "SPY:US" -categories "ETF"
```

| Flag | Description | Default |
|------|-------------|---------|
| `-q` | Search query (required) | - |
| `-categories` | Categories (comma-separated) | Equity,ETF |
| `-primary-only` | Primary exchange only | false |

### snapshot

Get current snapshot data with financial metrics.

```bash
./scripts/koyfin snapshot -kids <list_of_koyfin_ids> -category Equity
```

| Flag | Description | Default |
|------|-------------|---------|
| `-kids` | Comma-separated Koyfin IDs (required) | - |
| `-category` | Equity or ETF | Equity |

**Limits:** Max 32 KIDs (Equity), 2 KIDs (ETF)

### ticker-data

Get time series data for a ticker.

```bash
./scripts/koyfin ticker-data -id "AAPL:US" -key "p_candle_range" -date-from "2024-01-01"
./scripts/koyfin ticker-data -id "AAPL:US" -key "f_r" -date-from "2020-01-01" -fin-period "quarterly"
```

| Flag | Description | Default |
|------|-------------|---------|
| `-id` | Koyfin ID (required) | - |
| `-key` | Indicator key (required) | - |
| `-date-from` | Start date YYYY-MM-DD (required) | - |
| `-date-to` | End date YYYY-MM-DD | Today |
| `-currency` | Data currency | USD |
| `-agg-period` | day, monthly, quarterly, annually | day |
| `-fin-period` | quarterly, annual, LTM | - |

### transcript

Manage earnings call transcripts.

```bash
# List transcripts for a ticker
./scripts/koyfin transcript -action list -kid "AAPL:US" -limit 5

# Get a specific transcript
./scripts/koyfin transcript -action get -transcript-id 123456

# Get transcript summary
./scripts/koyfin transcript -action summary -transcript-id 123456
```

| Flag | Description | Default |
|------|-------------|---------|
| `-action` | list, get, summary (required) | - |
| `-kid` | Koyfin ID (required for list) | - |
| `-transcript-id` | Key dev ID (required for get/summary) | - |
| `-limit` | Max results for list (1-64) | 10 |

### schema

Get indicator schema reference.

```bash
./scripts/koyfin schema -asset-type Equity -indicator-type financials
./scripts/koyfin schema -asset-type Equity -indicator-type ratios
```

| Flag | Description | Default |
|------|-------------|---------|
| `-asset-type` | Asset type | Equity |
| `-indicator-type` | financials, ratios, forward_estimates, market_data (required) | - |

### etf-holdings

Get ETF holdings data.

```bash
./scripts/koyfin etf-holdings -kids <list_of_koyfin_ids> -category ETF
```

| Flag | Description | Default |
|------|-------------|---------|
| `-kids` | Koyfin IDs, max 2 (required) | - |
| `-category` | Must be ETF | ETF |

### screener-schema

Get available screener filters.

```bash
./scripts/koyfin screener-schema -asset-type Equity
```

### screener

Run stock screener with custom filters.

```bash
# Large cap (>10B)
./scripts/koyfin screener -filters '[{"key":"mkt","min":10000}]'

# Tech sector, 1B-10B market cap
./scripts/koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":1000,"max":10000}]' -page-size 50

# EV/EBITDA < 10
./scripts/koyfin screener -filters '[{"key":"evebitdaltm","max":10}]'
```

| Flag | Description | Default |
|------|-------------|---------|
| `-filters` | JSON filter array (required) | - |
| `-page-size` | Results per page (max 300) | 100 |

**Common Filter Keys:**

| Key | Type | Description |
|-----|------|-------------|
| `mkt` | Numeric | Market cap (in millions) |
| `t_sec` | Enum | Sector |
| `iso2` | Enum | Country code |
| `evebitdaltm` | Numeric | EV/EBITDA LTM |
| `pf_fcf_margin-LTM` | Numeric | FCF margin LTM |
| `chgYTDPct_L` | Numeric | YTD change % |

## Authentication

```bash
./scripts/koyfin auth -email <Koyfin account email> -password <Koyfin account password>
```

### Excel Export

Export snapshot data to Excel with formatted sheets:

```bash
./scripts/koyfin snapshot -kids <list_of_koyfin_ids> | \
    python3 ./scripts/utils/excel_export.py -o <output.xlsx>
```

**Excel Sheets Created:**

| Sheet | Description |
|-------|-------------|
| **Summary** | Key metrics: Price, Market Cap, P/E, EV/EBITDA, Margins |
| **Raw Data** | All snapshot metrics |
| **Ratios** | P/E, EV/EBITDA, EV/Sales, Margins, P/FCF |
| **Growth** | 1Y/3Y/5Y CAGR, YTD, estimate vs LTM |

### Install Python Dependencies

```bash
pip3 install -r ./scripts/utils/requirements.txt
```

### Format Output as Tables

```bash
# Format search results
./scripts/koyfin search -q "Apple" | python3 ./scripts/utils/process.py search

# Format snapshot
./scripts/koyfin snapshot -kids <list_of_koyfin_ids> | python3 ./scripts/utils/process.py snapshot
```

## Examples

```bash
# Search for Apple
./scripts/koyfin search -q "Apple"

# Get snapshot for multiple tickers
./scripts/koyfin snapshot -kids <list_of_koyfin_ids> -category Equity

# Get 1 year of daily price data
./scripts/koyfin ticker-data -id "AAPL:US" -key "p_candle_range" -date-from "2024-01-01"

# Get quarterly revenue
./scripts/koyfin ticker-data -id "AAPL:US" -key "f_r" -date-from "2020-01-01" -fin-period quarterly

# List recent transcripts
./scripts/koyfin transcript -action list -kid "AAPL:US" -limit 5

# Get transcript by ID
./scripts/koyfin transcript -action get -transcript-id 123456

# Screen for large cap tech stocks
./scripts/koyfin screener -filters '[{"key":"t_sec","values":["Information Technology"]},{"key":"mkt","min":10000}]' -page-size 50

# Export to Excel (Linux example)
./scripts/koyfin snapshot -kids "AAPL:US,MSFT:US" | python3 ~/.local/bin/koyfin-utils/excel_export.py -o output.xlsx
```

## Troubleshooting

### "command not found: koyfin"

Ensure `/scripts` is in your AI tool skill directory

### "Error loading session"

Re-run the setup script to re-authenticate:

```bash
./tool_build.sh koyfin
```

### "No data" or API errors

Verify your Koyfin credentials are correct.