package user

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) Register(ctx context.Context, login string, pass string) (*model.User, error) {
	hash := getPassHash(pass)
	user, err := s.userRepository.Create(ctx, login, hash)
	if err != nil {
		return nil, fmt.Errorf("register user service error: %w", err)
	}
	return user, nil
}
