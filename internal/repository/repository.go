package repository

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, name string, passHash string) (*model.User, error)
	FindByLoginAndPassHash(ctx context.Context, login string, passHash string) (*model.User, error)
}
