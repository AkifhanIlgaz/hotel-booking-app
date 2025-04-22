package services

import (
	"database/sql"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
)

type HotelService struct {
	db *sql.DB
}

func NewHotelService(db *sql.DB) *HotelService {
	return &HotelService{
		db: db,
	}
}

// Todo: Implement
func (hs *HotelService) GetHotels() ([]*models.Hotel, error) {
	return nil, nil
}
