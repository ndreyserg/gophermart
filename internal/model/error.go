package model

import (
	"errors"
)

var (
	ErrUserNotFound                 = errors.New("user not found")
	ErrUserAllreadyExist            = errors.New("user allready exists")
	ErrOrderAllreadyExist           = errors.New("order allready exists")
	ErrOrderAllreadyExistOnUser     = errors.New("order allready exists on user")
	ErrOrderNotFound                = errors.New("order not found")
	ErrAccountNotFound              = errors.New("account not found")
	ErrAccountNegativeBalance       = errors.New("negative balance")
	ErrUncorrectOrederNumber        = errors.New("uncorrect order number")
	ErrAccrualSystemTooManyRequests = errors.New("accrual system too many requests")
)
