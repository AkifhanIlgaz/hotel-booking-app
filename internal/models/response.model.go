package models

type Response struct {
	Status  string
	Message string
	Error   string
	Payload any
}
