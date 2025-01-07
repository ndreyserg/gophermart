package user

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s service) Login(ctx context.Context, login string, pass string) (*model.User, error) {
	passHash := getPassHash(pass)
	user, err := s.userRepository.FindByLoginAndPassHash(ctx, login, passHash)

	if err != nil {
		return nil, err
	}

	return user, nil
}
