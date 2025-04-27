package services

import (
	"fmt"
)

// This function simulates notifying the analytics service
func NotifyAnalytics(vin string) {
	fmt.Printf("[Service] Analytics notified for VIN: %s\n", vin)
}
// In a real-world scenario, this would involve making an HTTP request to the analytics service.
