package errors

import "errors"

var (
	As  = errors.As
	Is  = errors.Is
	New = errors.New
)

var (
	ErrEmailTaken           = errors.New("email address is already taken")
	ErrUserNotFound         = errors.New("user not found")
	ErrWrongPassword        = errors.New("wrong password")
	ErrTokenExpired         = errors.New("token expired")
	ErrNotFoundRefreshToken = errors.New("refresh token not found")
	ErrInvalidAuthHeader    = errors.New("invalid auth header")
	ErrInvalidToken         = errors.New("invalid token")
	ErrMissingAuthHeader    = errors.New("missing auth header")
)
