package main

import (
	"encoding/json"
	"os"
	"time"
)

type Config struct {
	Cryptos       []string      `json:"cryptocurrencies"`
	CheckInterval time.Duration `json:"check_interval_seconds"`
	Alerts        []AlertRule   `json:"alerts"`
}

type AlertRule struct {
	Symbol    string  `json:"symbol"`
	Condition string  `json:"condition"` // "above", "below"
	Price     float64 `json:"price"`
	Message   string  `json:"message"`
}

func LoadConfig(filename string) (*Config, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	
	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}
	
	// Convert seconds to duration
	config.CheckInterval = config.CheckInterval * time.Second
	
	return &config, nil
}