package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/config"
)

func Connect(psqlConfig config.PostgresConfig) (*sql.DB, error) {
	connString := generateConnString(psqlConfig)

	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	db.SetMaxIdleConns(psqlConfig.MaxIdleConns)
	db.SetMaxOpenConns(psqlConfig.MaxOpenConns)
	db.SetConnMaxLifetime(time.Duration(psqlConfig.ConnMaxLifetimeMin) * time.Minute)
	db.SetConnMaxIdleTime(time.Duration(psqlConfig.ConnMaxIdleTimeMin) * time.Minute)

	return db, nil
}

func generateConnString(psqlConfig config.PostgresConfig) string {
	connString := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		psqlConfig.Host, psqlConfig.Port, psqlConfig.User, psqlConfig.Password, psqlConfig.DBName)

	return connString
}
