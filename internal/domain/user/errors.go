package user

import "errors"

var (
	ErrorUsernameTaken = errors.New("username already taken")
	ErrorUserNotFound  = errors.New("user not found")
)
