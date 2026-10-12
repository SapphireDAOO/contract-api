package server

import "errors"

var (
	ErrLoadEnv = errors.New("error loading .env file")
	ErrSignIn  = errors.New("configuring sign-in")
)
