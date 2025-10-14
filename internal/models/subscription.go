package models

import "time"

type Subscription struct {
	ID            string    `json:"id"`
	ServiceName   string    `json:"service_name"`
	Price         float64   `json:"price"`
	Category      string    `json:"category"`
	PaymentMethod string    `json:"payment_method"`
	BillingCycle  string    `json:"billing_cycle"`
	IsPinned      bool      `json:"is_pinned"`
	Notes         string    `json:"notes"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	UserID        string    `json:"user_id"`
}
