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

func (s *OTPService) GenerateOTP(userId uuid.UUID) (string, error) {
	// Generate a 6-digit OTP
	otp, err := utils.GenerateNumericOTP(6)
	if err != nil {
		return "", fmt.Errorf("generate otp: %w", err)
	}

	// Create OTP token record
	id := uuid.New()
	now := time.Now()
	expiresAt := now.Add(2 * time.Minute) // OTP expires in 10 minutes

	_, err = s.db.Exec(queries.InsertOTPToken,
		id,
		userId,
		otp,
		expiresAt,
		now,
		false,
	)
	if err != nil {
		return "", fmt.Errorf("save otp token: %w", err)
	}

	return otp, nil
}

func (s *OTPService) VerifyOTP(email, otp string) (bool, error) {
	var token models.OTPToken
	var userId uuid.UUID

	// First get the user ID from email
	err := s.db.QueryRow(queries.SelectUserIdByEmail, email).Scan(&userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errors.ErrUserNotFound
		}
		return false, fmt.Errorf("get user id: %w", err)
	}

	// Then verify the OTP
	err = s.db.QueryRow(queries.SelectOTPToken, userId, otp).Scan(
		&token.Id,
		&token.UserId,
		&token.Token,
		&token.ExpiresAt,
		&token.CreatedAt,
		&token.Used,
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

	// Check if OTP is already used
	if token.Used {
		return false, errors.New("otp already used")
	}

	// Mark OTP as used
	_, err = s.db.Exec(queries.MarkOTPAsUsed, token.Id)
	if err != nil {
		return false, fmt.Errorf("mark otp as used: %w", err)
	}

	return true, nil
}
