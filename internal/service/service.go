package service

import (
	"context"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/model"
)

type UserService interface {
	Register(ctx context.Context, login string, pass string) (*model.User, error)
	Login(ctx context.Context, login string, pass string) (*model.User, error)
}

type OrderService interface {
	Create(ctx context.Context, number string, userID int) (*model.Order, error)
	GetByUserID(ctx context.Context, userID int) ([]*model.Order, error)
}

type AccrualService interface {
	AccrueAsync(*model.Order)
}

type AccountService interface {
	Withdraw(ctx context.Context, userID int, orderNumber string, sum float64) error
	GetBalance(ctx context.Context, userID int) (*model.AccountBalance, error)
	GetWithdrawals(ctx context.Context, userID int) ([]*model.AccountWithdrawals, error)
	GetOrCreate(ctx context.Context, userID int) (*model.Account, error)
}

type CheckerSevice interface {
	Check(num string) error
}

type SessionService interface {
	Open(userID int, w http.ResponseWriter, r *http.Request) error
	GetUserID(r *http.Request) (int, error)
}
