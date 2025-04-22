package models

import "github.com/google/uuid"

type Hotel struct {
	Id            uuid.UUID `json:"id" `
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Location      Location  `json:"location"`
	ImageUrl      string    `json:"image_url"`
	PricePerNight string    `json:"price_per_night"`
	PhoneNumber   string    `json:"phone_number"`
	Rating        float64   `json:"rating"`
	Features      []string  `json:"features"`
	CreatedAt     string    `json:"created_at"`
}

type Location struct {
	City    string `json:"city"`
	Country string `json:"country"`
}
