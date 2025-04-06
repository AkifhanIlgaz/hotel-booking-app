package models

import "github.com/google/uuid"

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
	Id           uuid.UUID `json:"id" `
	Name         string    `json:"name"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"-"`
	Role         Role      `json:"role"`
	CreatedAt    string    `json:"created_at"`
}
