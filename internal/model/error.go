package model

import (
	"errors"
)

var (
	ErrorUserNotFound      = errors.New("user not found")
	ErrorUserAllreadyExist = errors.New("user allready exists")
)
