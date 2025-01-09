package account

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetBalance(ctx context.Context, userID int) (*model.AccountBalance, error) {
	acc, err := s.getOrCreate(ctx, userID)

	if err != nil {
		return nil, err
	}

	w, err := s.accountReposity.GetWithdrawn(ctx, acc.ID)

	if err != nil {
		return nil, err
	}

	balance := model.AccountBalance{
		Current:   acc.Balance,
		Withdrawn: w,
	}

	return &balance, nil
}
