package services

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/AkifhanIlgaz/hotel-booking-app/internal/models"
	"github.com/AkifhanIlgaz/hotel-booking-app/migrations/queries"
	"github.com/AkifhanIlgaz/hotel-booking-app/pkg/utils"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgerrcode"
)

type UserService struct {
	db *sql.DB
}

func NewUserService(db *sql.DB) *UserService {
	return &UserService{
		db: db,
	}
}

// This function will create default user
func (us UserService) RegisterUser(registrationReq models.RegistrationRequest) (uuid.UUID, error) {
	hashedPassword, err := utils.HashPassword(registrationReq.Password)
	if err != nil {
		return uuid.Nil, fmt.Errorf("register user: %w", err)
	}

	id := uuid.New()

	if _, err := us.db.Exec(queries.InsertUser,
		id,
		registrationReq.Name,
		registrationReq.Email,
		hashedPassword,
		models.RoleUser,
		time.Now()); err != nil {
		var pgError *pgconn.PgError
		if errors.As(err, &pgError) {
			if pgError.Code == pgerrcode.UniqueViolation {
				return uuid.Nil, errors.New("This email address is already registered.")
			}
		}
		return uuid.Nil, fmt.Errorf("register user: %w", err)
	}

	return id, nil
}
