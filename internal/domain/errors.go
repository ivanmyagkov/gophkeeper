package domain

import "errors"

//user
var (
	ErrUserNotFound                    = errors.New("user doesn't exists")
	ErrUserNotFoundOrSessionWasExpired = errors.New("user doesn't exists or session was expired")
	ErrUserAlreadyExists               = errors.New("user with such login already exists")
	ErrUserBadPassword                 = errors.New("bad password")

	ErrSessionNotFound      = errors.New("session was not found")
	ErrSessionAlreadyExists = errors.New("session is already exist")
)

//materials
var (
	ErrDataNotFound = errors.New("data was not found")
)

var (
	ErrInternalServerError = errors.New("internal server error")
)
