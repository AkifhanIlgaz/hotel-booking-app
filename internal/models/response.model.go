package models

type Response struct {
	Status  string
	Message string // user-friendly message
	Error   string // developer-friendly message
	Payload any
}
