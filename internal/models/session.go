package models

import "time"

type Session struct {
	IDSession string    `json:"id_session"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expired_at"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserAgent string    `json:"user_agent"`
	IPAddress string    `json:"ip_address"`
	UserID    string    `json:"user_id"`
}
