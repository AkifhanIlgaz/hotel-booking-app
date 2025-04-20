package models

import (
	"time"

	"github.com/google/uuid"
)

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

func (r Role) IsValid() bool {
	switch r {
	case RoleAdmin, RoleUser:
		return true
	default:
		return false
	}
}

// TODO: Should I use custom IsValid function or validate struct tag

type User struct {
	Id           uuid.UUID `json:"id" validate:"required,uuid"`
	Name         string    `json:"name" validate:"required,min=3,max=50"`
	Email        string    `json:"email" validate:"required,email"`
	PasswordHash string    `json:"-" validate:"required,min=8"`
	Role         Role      `json:"role" validate:"required,oneof=admin user"`
	CreatedAt    time.Time `json:"created_at" validate:"required"`
}

// TODO: request'i düzenlemek için method yaz Trimspace lower
type RegistrationRequest struct {
	Name     string `json:"name" binding:"required,min=3,max=50"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type ChangePasswordRequest struct {
	ResetToken string `json:"reset_token" binding:"required"`
	Password   string `json:"password" binding:"required,min=8"`
}
