package mysql

import "errors"

var (
	ErrorUserExist    = errors.New("user already exist")
	ErrorUserNotExist = errors.New("user does not exist")
)
