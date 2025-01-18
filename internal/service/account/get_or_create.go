package account

import (
	"context"
	"errors"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetOrCreate(ctx context.Context, userID int) (*model.Account, error) {
	account, err := s.accountReposity.FindByUserID(ctx, userID)

	if err != nil && !errors.Is(err, model.ErrAccountNotFound) {
		return nil, fmt.Errorf("get or create acc error: %w,", err)
	}

	if account == nil {
		account, err = s.accountReposity.Create(ctx, userID)
		if err != nil {
			return nil, fmt.Errorf("get or create acc error: %w,", err)
		}
	}
	return account, nil
}
