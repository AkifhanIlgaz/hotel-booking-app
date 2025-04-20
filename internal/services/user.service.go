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
func (us *UserService) RegisterUser(registrationReq models.RegistrationRequest) (uuid.UUID, error) {
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
				return uuid.Nil, errors.ErrEmailTaken
			}
		}
		return uuid.Nil, fmt.Errorf("register user: %w", err)
	}

	return id, nil
}

func (us *UserService) AuthenticateUser(loginReq models.LoginRequest) (*models.User, error) {
	var user models.User

	if err := us.db.QueryRow(queries.SelectUserByEmail, loginReq.Email).Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrUserNotFound
		}
		return nil, fmt.Errorf("register user: %w", err)
	}

	isPasswordTrue := utils.VerifyPassword(loginReq.Password, user.PasswordHash)
	if !isPasswordTrue {
		return nil, errors.ErrWrongPassword
	}

	return &user, nil
}

func (us *UserService) GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	if err := us.db.QueryRow(queries.SelectUserByEmail, email).Scan(&user.Id, &user.Name, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.ErrUserNotFound
		}
		return nil, fmt.Errorf("check is user exists: %w", err)
	}

	return &user, nil
}
