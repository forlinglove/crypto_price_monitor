package main

import (
	"encoding/json"
	"os"
)

// Alert represents a price alert configuration
type Alert struct {
	Symbol    string  `json:"symbol"`
	Condition string  `json:"condition"` // "above" or "below"
	Price     float64 `json:"price"`
	Message   string  `json:"message"`
}

// Config holds the application configuration
type Config struct {
	APIKey        string   `json:"api_key"`
	UpdateInterval int     `json:"update_interval"` // in seconds
	Currencies    []string `json:"currencies"`
	Alerts        []Alert  `json:"alerts"`
}

// LoadConfig loads configuration from JSON file
func LoadConfig(filename string) (*Config, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}