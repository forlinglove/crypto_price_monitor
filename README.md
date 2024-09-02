## Cryptocurrency Price Monitor

A simple Go application that monitors cryptocurrency prices and displays alerts based on configured rules.

### Features

- Monitors multiple cryptocurrencies simultaneously
- Configurable check intervals
- Price-based alert system
- Support for "above" and "below" price conditions
- Real-time price display in console

### Prerequisites

- Go 1.21 or later
- Internet connection (for API calls)

### Installation

1. **Create project directory:**
   ```bash
   mkdir crypto-monitor
   cd crypto-monitor
   ```

2. **Save all the provided files in the directory**

3. **Initialize Go module:**
   ```bash
   go mod init crypto-monitor
   go mod tidy
   ```

### Configuration

Edit `config.json` to customize:

- **cryptocurrencies**: List of crypto IDs (use CoinGecko API IDs)
- **check_interval_seconds**: How often to check prices (in seconds)
- **alerts**: Array of alert rules with conditions and messages

**Common cryptocurrency IDs:**
- `bitcoin` (BTC)
- `ethereum` (ETH)
- `cardano` (ADA)
- `solana` (SOL)
- `ripple` (XRP)
- `dogecoin` (DOGE)

### Usage

1. **Run the application:**
   ```bash
   go run .
   ```

2. **The application will:**
   - Display current prices for configured cryptocurrencies
   - Show alerts when price conditions are met
   - Continue monitoring until stopped (Ctrl+C)

### Example Output
```
🚀 Cryptocurrency Price Monitor Started...
[14:30:15] bitcoin: $48523.45
[14:30:15] ethereum: $3250.67
[14:30:15] cardano: $0.52
[14:30:15] solana: $98.75
---
[14:30:45] bitcoin: $48789.12
[14:30:45] ethereum: $3245.89
🔔 🚨 ALERT: ethereum is $3245.89 (below $3000.00) - Ethereum dropped below threshold
[14:30:45] cardano: $0.53
[14:30:45] solana: $101.25
🔔 🚨 ALERT: solana is $101.25 (above $100.00) - Solana is performing well
---
```

### Customization

**Add more notification methods:**
Modify `alert_manager.go` to add:
- Email notifications
- SMS alerts
- Desktop notifications
- Webhook calls

**Add more exchanges:**
Extend `price_tracker.go` to support multiple price sources like:
- Binance API
- Coinbase API
- Kraken API

### Notes

- Uses CoinGecko free API (rate limits may apply)
- For production use, consider adding error handling and retry logic
- Prices are in USD
- Alert messages are displayed in console only

### Stopping the Application

Press `Ctrl+C` to gracefully stop the monitoring process.

This provides a solid foundation that you can extend with additional features like database storage, web interface, or more sophisticated alerting mechanisms.