package domain

import "errors"

var (
	ErrNotFound        = errors.New("not found")
	ErrAlreadyExists   = errors.New("already exists")
	ErrInvalidInput    = errors.New("invalid input")
	ErrUnauthorized    = errors.New("unauthorized")
	ErrForbidden       = errors.New("forbidden")
	ErrInternalServer  = errors.New("internal server error")
	ErrUserNotFound    = errors.New("user data not found")
	ErrProfileNotFound = errors.New("profile data not found")
)
