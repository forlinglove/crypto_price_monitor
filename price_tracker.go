package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CryptoPrice struct {
	Symbol string  `json:"symbol"`
	Price  float64 `json:"price,string"`
}

type PriceTracker struct {
	cryptos       []string
	checkInterval time.Duration
}

func NewPriceTracker(cryptos []string, interval time.Duration) *PriceTracker {
	return &PriceTracker{
		cryptos:       cryptos,
		checkInterval: interval,
	}
}

func (pt *PriceTracker) StartMonitoring(alertManager *AlertManager) {
	ticker := time.NewTicker(pt.checkInterval)
	defer ticker.Stop()
	
	for {
		select {
		case <-ticker.C:
			pt.checkPrices(alertManager)
		}
	}
}

func (pt *PriceTracker) checkPrices(alertManager *AlertManager) {
	for _, symbol := range pt.cryptos {
		price, err := pt.fetchPrice(symbol)
		if err != nil {
			log.Printf("Error fetching price for %s: %v", symbol, err)
			continue
		}
		
		fmt.Printf("[%s] %s: $%.2f\n", time.Now().Format("15:04:05"), symbol, price)
		alertManager.CheckAlerts(symbol, price)
	}
	fmt.Println("---")
}

func (pt *PriceTracker) fetchPrice(symbol string) (float64, error) {
	// Using CoinGecko API (free tier)
	url := fmt.Sprintf("https://api.coingecko.com/api/v3/simple/price?ids=%s&vs_currencies=usd", symbol)
	
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}
	
	var result map[string]map[string]float64
	err = json.Unmarshal(body, &result)
	if err != nil {
		return 0, err
	}
	
	if cryptoData, exists := result[symbol]; exists {
		if price, exists := cryptoData["usd"]; exists {
			return price, nil
		}
	}
	
	return 0, fmt.Errorf("price not found for %s", symbol)
}