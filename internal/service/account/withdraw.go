package account

import (
	"context"
)

func (s *service) Withdraw(ctx context.Context, userID int, orderNumber string, sum float64) error {
	account, err := s.getOrCreate(ctx, userID)

	if err != nil {
		return err
	}

	return s.accountReposity.Withdraw(ctx, account.ID, sum, orderNumber)
}
