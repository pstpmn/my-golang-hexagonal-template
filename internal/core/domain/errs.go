package domain

import (
	"errors"
)

var (
	ErrForbidden     = errors.New("user is forbidden to access the resource")
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user ID")
	ErrUserInactive  = errors.New("user is inactive")
	ErrDuplicateUser = errors.New("user already exists")
	ErrUnauthorized  = errors.New("user is unauthorized to access the resource")
)
