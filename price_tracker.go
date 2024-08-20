package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

// PriceData represents cryptocurrency price information
type PriceData struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

// PriceTracker manages cryptocurrency price monitoring
type PriceTracker struct {
	config     *Config
	httpClient *http.Client
}

// NewPriceTracker creates a new price tracker instance
func NewPriceTracker(config *Config) *PriceTracker {
	return &PriceTracker{
		config: config,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}
}

// GetPrices fetches current prices for specified cryptocurrencies
func (pt *PriceTracker) GetPrices(symbols []string) (map[string]float64, error) {
	// Using CoinGecko API (free tier, no API key required for basic price data)
	url := "https://api.coingecko.com/api/v3/simple/price?ids="
	
	// Convert symbols to CoinGecko IDs (simplified mapping)
	for i, symbol := range symbols {
		if i > 0 {
			url += ","
		}
		// Simple mapping - you can extend this
		switch symbol {
		case "BTC":
			url += "bitcoin"
		case "ETH":
			url += "ethereum"
		case "ADA":
			url += "cardano"
		case "DOT":
			url += "polkadot"
		default:
			url += symbol
		}
	}
	url += "&vs_currencies=usd"

	resp, err := pt.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]map[string]float64
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	prices := make(map[string]float64)
	for coinID, data := range result {
		if usdPrice, exists := data["usd"]; exists {
			// Map back to our symbols
			switch coinID {
			case "bitcoin":
				prices["BTC"] = usdPrice
			case "ethereum":
				prices["ETH"] = usdPrice
			case "cardano":
				prices["ADA"] = usdPrice
			case "polkadot":
				prices["DOT"] = usdPrice
			default:
				prices[coinID] = usdPrice
			}
		}
	}

	return prices, nil
}

// CheckAlerts verifies if any price alerts should be triggered
func (pt *PriceTracker) CheckAlerts(prices map[string]float64) {
	for _, alert := range pt.config.Alerts {
		currentPrice, exists := prices[alert.Symbol]
		if !exists {
			continue
		}

		triggered := false
		message := ""

		switch alert.Condition {
		case "above":
			if currentPrice > alert.Price {
				triggered = true
				message = fmt.Sprintf("🚨 ALERT: %s is above $%.2f (Current: $%.2f) - %s", 
					alert.Symbol, alert.Price, currentPrice, alert.Message)
			}
		case "below":
			if currentPrice < alert.Price {
				triggered = true
				message = fmt.Sprintf("🚨 ALERT: %s is below $%.2f (Current: $%.2f) - %s", 
					alert.Symbol, alert.Price, currentPrice, alert.Message)
			}
		}

		if triggered {
			fmt.Printf("\n%s\n", message)
		}
	}
}

// StartMonitoring begins the continuous price monitoring
func (pt *PriceTracker) StartMonitoring() {
	ticker := time.NewTicker(time.Duration(pt.config.UpdateInterval) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			prices, err := pt.GetPrices(pt.config.Currencies)
			if err != nil {
				log.Printf("Error fetching prices: %v", err)
				continue
			}

			// Display current prices
			fmt.Printf("\n📊 Current Prices (%s):\n", time.Now().Format("15:04:05"))
			for symbol, price := range prices {
				fmt.Printf("  %s: $%.2f\n", symbol, price)
			}

			// Check alerts
			pt.CheckAlerts(prices)
		}
	}
}