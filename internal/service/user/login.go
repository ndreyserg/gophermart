package user

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s service) Login(ctx context.Context, login string, pass string) (*model.User, error) {
	passHash := getPassHash(pass)
	user, err := s.userRepository.FindByLoginAndPassHash(ctx, login, passHash)

	if err != nil {
		return nil, fmt.Errorf("login serivice error: %w", err)
	}

	return user, nil
}
