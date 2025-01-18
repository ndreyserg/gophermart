package account

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetBalance(ctx context.Context, userID int) (*model.AccountBalance, error) {
	acc, err := s.GetOrCreate(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("get balance get account error: %w,", err)
	}

	w, err := s.accountReposity.GetWithdrawn(ctx, acc.ID)

	if err != nil {
		return nil, fmt.Errorf("get balance get withdrawn error: %w,", err)
	}

	balance := model.AccountBalance{
		Current:   acc.Balance,
		Withdrawn: w,
	}

	return &balance, nil
}
