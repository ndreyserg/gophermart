package account

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) GetWithdrawals(ctx context.Context, userID int) ([]*model.AccountWithdrawals, error) {

	acc, err := s.getOrCreate(ctx, userID)

	if err != nil {
		return nil, err
	}

	return s.accountReposity.GetWithdrawals(ctx, acc.ID)
}
