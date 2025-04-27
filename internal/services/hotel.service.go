package services

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
)

type HotelService struct {
	db *sql.DB
}

func NewHotelService(db *sql.DB) *HotelService {
	return &HotelService{
		db: db,
	}
}

func (hs *HotelService) GetHotels(params models.HotelFilterParams) ([]models.Hotel, error) {
	query := queries.BuildHotelsQueryWithParams(params)

	rows, err := hs.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to get hotels: %w", err)
	}

	var hotels []models.Hotel

	defer rows.Close()

	for rows.Next() {
		var hotel models.Hotel
		var features string
		err := rows.Scan(
			&hotel.Id,
			&hotel.Name,
			&hotel.Description,
			&hotel.Location.City,
			&hotel.Location.Country,
			&hotel.ImageUrl,
			&hotel.PricePerNight,
			&hotel.Rating,
			&hotel.PhoneNumber,
			&features,
			&hotel.CreatedAt,
		)
		if err != nil {
			return hotels, fmt.Errorf("failed to scan hotel: %w", err)
		}
		hotel.Features = strings.Split(features, ",")

		hotels = append(hotels, hotel)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("get all hotels with filter: %w", err)
	}

	return hotels, nil
}

func (hs *HotelService) GetHotelById(hotelId string) (models.Hotel, error) {
	var hotel models.Hotel
	var features string

	err := hs.db.QueryRow(queries.SelectHotelById, sql.Named("id", hotelId)).Scan(&hotel.Id,
		&hotel.Name,
		&hotel.Description,
		&hotel.Location.City,
		&hotel.Location.Country,
		&hotel.ImageUrl,
		&hotel.PricePerNight,
		&hotel.Rating,
		&hotel.PhoneNumber,
		&features,
		&hotel.CreatedAt)
	if err != nil {
		return models.Hotel{}, fmt.Errorf("get hotel by id: %w", err)
	}

	hotel.Features = strings.Split(features, ",")

	return hotel, nil
}
