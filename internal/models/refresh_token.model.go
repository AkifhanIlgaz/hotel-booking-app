package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Id        uuid.UUID `json:"id"`
	UserId    uuid.UUID `json:"user_id"`
	TokenHash string    `json:"-"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
	// DeviceInfo DeviceInfo
}

// Kullanıcının farklı cihazlardan açtığı oturumlar takip edilmek istenirse implement edilebilir
type DeviceInfo struct {
	IP        string
	UserAgent string
	OS        string
	Browser   string
	Device    string
}
