package services

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/errors"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/utils"
	"github.com/google/uuid"
)

type OTPService struct {
	db *sql.DB
}

func NewOTPService(db *sql.DB) *OTPService {
	return &OTPService{
		db: db,
	}
}

func (s *OTPService) GenerateOTP(email string) (string, error) {
	// Generate a 6-digit OTP
	otp, err := utils.GenerateNumericOTP(6)
	if err != nil {
		return "", fmt.Errorf("generate otp: %w", err)
	}

	// Create OTP token record
	id := uuid.New()
	now := time.Now()
	expiresAt := now.Add(10 * time.Minute) // OTP expires in 10 minutes

	_, err = s.db.Exec(queries.InsertOTPToken,
		sql.Named("id", id),
		sql.Named("email", email),
		sql.Named("token_hash", utils.Hash(otp)),
		sql.Named("expires_at", expiresAt),
		sql.Named("created_at", now),
	)
	if err != nil {
		return "", fmt.Errorf("save otp token: %w", err)
	}

	return otp, nil
}

func (s *OTPService) VerifyOTP(email, otp string) (bool, error) {
	var token models.OTPToken

	// Then verify the OTP
	err := s.db.QueryRow(queries.SelectOTPToken,
		sql.Named("email", email),
		sql.Named("token_hash", utils.Hash(otp)),
	).Scan(
		&token.Id,
		&token.Email,
		&token.TokenHash,
		&token.ExpiresAt,
		&token.CreatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.New("invalid otp")
		}
		return false, fmt.Errorf("verify otp: %w", err)
	}

	// Check if OTP is expired
	if time.Now().After(token.ExpiresAt) {
		return false, errors.New("otp expired")
	}

	// Delete the OTP after successful verification
	_, err = s.db.Exec(queries.DeleteOTPToken, sql.Named("token_hash", token.TokenHash))
	if err != nil {
		return false, fmt.Errorf("delete otp: %w", err)
	}

	return true, nil
}
