package main

import (
	"fmt"
	"log"
)

type AlertManager struct {
	alerts []AlertRule
}

func NewAlertManager(alerts []AlertRule) *AlertManager {
	return &AlertManager{
		alerts: alerts,
	}
}

func (am *AlertManager) CheckAlerts(symbol string, currentPrice float64) {
	for _, alert := range am.alerts {
		if alert.Symbol == symbol {
			triggered := false
			
			switch alert.Condition {
			case "above":
				if currentPrice > alert.Price {
					triggered = true
				}
			case "below":
				if currentPrice < alert.Price {
					triggered = true
				}
			default:
				log.Printf("Unknown alert condition: %s", alert.Condition)
				continue
			}
			
			if triggered {
				am.triggerAlert(alert, currentPrice)
			}
		}
	}
}

func (am *AlertManager) triggerAlert(alert AlertRule, currentPrice float64) {
	message := fmt.Sprintf("🚨 ALERT: %s is $%.2f (%s $%.2f) - %s", 
		alert.Symbol, currentPrice, alert.Condition, alert.Price, alert.Message)
	
	fmt.Println("🔔 " + message)
	
	// Here you could add additional notification methods:
	// - Send email
	// - Send SMS
	// - Desktop notification
	// - Webhook
}