package models

import (
	"time"

	"github.com/google/uuid"
)

type OTPToken struct {
	Id        uuid.UUID
	UserId    uuid.UUID
	Token     string
	ExpiresAt time.Time
	CreatedAt time.Time
}

type OTPRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type OTPVerificationRequest struct {
	Email string `json:"email" binding:"required,email"`
	OTP   string `json:"otp" binding:"required,min=6,max=6"`
}
