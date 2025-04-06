package models

import "github.com/google/uuid"

type Hotel struct {
	Id            uuid.UUID `json:"id" `
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Location      string    `json:"location"`
	ImageUrl      string    `json:"image_url"`
	PricePerNight string    `json:"price_per_night"`
	Rating        int       `json:"rating"`
	Features      []string  `json:"features"`
}
