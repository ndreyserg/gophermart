package account

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetWithdrawals(ctx context.Context, userID int) ([]*model.AccountWithdrawals, error) {
	acc, err := s.GetOrCreate(ctx, userID)

	if err != nil {
		return nil, fmt.Errorf("get withdrawal error: %w,", err)
	}

	res, err := s.accountReposity.GetWithdrawals(ctx, acc.ID)
	if err != nil {
		return nil, fmt.Errorf("get withdrawal error: %w,", err)
	}
	return res, nil
}
