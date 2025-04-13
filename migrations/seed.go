package migrations

import (
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/schemas"
)

func Init(db *sql.DB) error {
	err := createUsersTable(db)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
	}

	return nil
}

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(schemas.Users)
	if err != nil {
		return fmt.Errorf("failed to create users table: %w", err)
	}

	// ? Should I check rowsAffected ?
	return nil
}
