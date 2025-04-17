package migrations

import (
	"database/sql"
	"fmt"

	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/schemas"
)

func Init(db *sql.DB) error {
	err := createTables(db, schemas.All()...)
	if err != nil {
		return fmt.Errorf("failed to migrate: %w", err)
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
