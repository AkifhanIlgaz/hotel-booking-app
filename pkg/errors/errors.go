package errors

import "errors"

var (
	As = errors.As
	Is = errors.Is
)

var (
	ErrEmailTaken = errors.New("email address is already taken")
)
