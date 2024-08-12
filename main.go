package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	fmt.Println("🚀 Cryptocurrency Price Monitor Started...")
	
	// Load configuration
	config, err := LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Initialize price tracker
	tracker := NewPriceTracker(config.Cryptos, config.CheckInterval)
	
	// Initialize alert manager
	alertManager := NewAlertManager(config.Alerts)
	
	// Start monitoring
	go tracker.StartMonitoring(alertManager)
	
	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	<-sigChan
	fmt.Println("\n🛑 Shutting down cryptocurrency monitor...")
}