package model

import (
	"errors"
)

var (
	ErrorUserNotFound             = errors.New("user not found")
	ErrorUserAllreadyExist        = errors.New("user allready exists")
	ErrorOrderAllreadyExist       = errors.New("order allready exists")
	ErrorOrderAllreadyExistOnUser = errors.New("order allready exists on user")
	ErrorOrderNotFound            = errors.New("order not found")
	ErrorAccountNotFound          = errors.New("account not found")
	ErrorAccountNegativeBalance   = errors.New("negative balance")
)
