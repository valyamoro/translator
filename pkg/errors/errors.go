package errors 

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrTokenExpired = errors.New("token has expired")
	ErrTokenInvalid = errors.New("token is invalid")
)
