package repository

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, name string, passHash string) (*model.User, error)
	FindByLoginAndPassHash(ctx context.Context, login string, passHash string) (*model.User, error)
}

type OrderRepository interface {
	Create(ctx context.Context, number string, userID int) (*model.Order, error)
	GetByUserID(ctx context.Context, userID int) ([]*model.Order, error)
	FindByNumber(ctx context.Context, number string) (*model.Order, error)
}

type AccountReposity interface {
	Create(ctx context.Context, userID int) (*model.Account, error)
	FindByUserID(ctx context.Context, userID int) (*model.Account, error)
	Withdraw(ctx context.Context, accountID int, sum float64, orderNumber string) error
}
