package service

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

type UserService interface {
	Register(ctx context.Context, login string, pass string) (*model.User, error)
	Login(ctx context.Context, login string, pass string) (*model.User, error)
}
