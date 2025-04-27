package migrations

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/schemas"
)

func Init(db *sql.DB) error {
	err := createTables(db, schemas.All()...)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	err = addHotels(db)
	if err != nil {
		return fmt.Errorf("failed to add hotels: %w", err)
	}

	return nil
}

func addHotels(db *sql.DB) error {
	doc, err := os.ReadFile("c:/Users/AKIF/Desktop/workspace/hotel-booking-app/mock/hotels.json")
	if err != nil {
		return fmt.Errorf("failed to read hotels.json: %w", err)
	}

	var hotels []models.Hotel
	err = json.Unmarshal(doc, &hotels)
	if err != nil {
		return fmt.Errorf("failed to unmarshal hotels.json: %w", err)
	}

	for _, hotel := range hotels {
		_, err = db.Exec(queries.InsertHotelQuery,
			sql.Named("id", hotel.Id),
			sql.Named("name", hotel.Name),
			sql.Named("description", hotel.Description),
			sql.Named("city", hotel.Location.City),
			sql.Named("country", hotel.Location.Country),
			sql.Named("image_url", hotel.ImageUrl),
			sql.Named("price_per_night", hotel.PricePerNight),
			sql.Named("rating", hotel.Rating),
			sql.Named("phone_number", hotel.PhoneNumber),
			sql.Named("features", strings.Join(hotel.Features, ",")),
			sql.Named("created_at", hotel.CreatedAt),
		)
		if err != nil {
			return fmt.Errorf("failed to insert hotel: %w", err)
		}
	}

	return nil
}

func createTables(db *sql.DB, tables ...string) error {
	for _, table := range tables {
		_, err := db.Exec(table)
		if err != nil {
			return fmt.Errorf("failed to create table: %w", err)
		}
		// ? Should I check rowsAffected ?
	}

	return nil

}
