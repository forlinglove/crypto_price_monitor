# Cryptocurrency Price Monitor

A simple Go application that monitors cryptocurrency prices and displays alerts based on configured thresholds.

## Features

- Real-time cryptocurrency price monitoring
- Configurable price alerts (above/below thresholds)
- Support for multiple cryptocurrencies
- Customizable update intervals
- Graceful shutdown handling

## Setup Instructions

1. **Install Go** (version 1.21 or higher)

2. **Create project directory**:
   ```bash
   mkdir crypto-monitor
   cd crypto-monitor
   ```

3. **Save all files** in the project directory

4. **Initialize Go module**:
   ```bash
   go mod init crypto-monitor
   ```

5. **Configure the application**:
   - Edit `config.json` with your preferred settings
   - Add/remove cryptocurrencies from the `currencies` array
   - Configure alerts in the `alerts` section

6. **Run the application**:
   ```bash
   go run .
   ```

## Configuration

The `config.json` file contains:

- `api_key`: API key for price data (currently using free CoinGecko API)
- `update_interval`: How often to check prices (in seconds)
- `currencies`: Array of cryptocurrency symbols to monitor
- `alerts`: Array of alert configurations with:
  - `symbol`: Cryptocurrency symbol
  - `condition`: "above" or "below"
  - `price`: Threshold price
  - `message`: Custom alert message

## Usage

1. The application will start monitoring prices immediately
2. Current prices are displayed every update interval
3. Alerts are triggered when price conditions are met
4. Press `Ctrl+C` to stop the application gracefully

## Supported Cryptocurrencies

Currently supports:
- BTC (Bitcoin)
- ETH (Ethereum) 
- ADA (Cardano)
- DOT (Polkadot)

You can extend the symbol mapping in `price_tracker.go` to add more cryptocurrencies.

## Dependencies

- Standard Go libraries only
- Uses CoinGecko free API for price data
- No external dependencies required

## Notes

- The application uses the free CoinGecko API which has rate limits
- For production use, consider adding error handling and logging
- You can extend functionality by adding more alert types or data sources