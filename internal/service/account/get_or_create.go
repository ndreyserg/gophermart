package account

import (
	"context"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) getOrCreate(ctx context.Context, userID int) (*model.Account, error) {
	account, err := s.accountReposity.FindByUserID(ctx, userID)

	if err != nil && !errors.Is(err, model.ErrorAccountNotFound) {
		return nil, err
	}

	if account == nil {
		account, err = s.accountReposity.Create(ctx, userID)
		if err != nil {
			return nil, err
		}
	}
	return account, nil
}
