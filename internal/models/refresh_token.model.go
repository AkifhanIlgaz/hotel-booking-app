package models

import (
	"time"

	"github.com/google/uuid"
)

type RefreshToken struct {
	Id          uuid.UUID
	UserId      uuid.UUID
	HashedToken string
	ExpiresAt   time.Time
	CreatedAt   time.Time
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
