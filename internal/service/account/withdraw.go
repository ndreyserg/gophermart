package account

import (
	"context"
)

func (s *service) Withdraw(ctx context.Context, userID int, orderNumber string, sum float64) error {

	err := s.orders.CheckNumber(orderNumber)

	if err != nil {
		return err
	}

	account, err := s.getOrCreate(ctx, userID)

	if err != nil {
		return err
	}

	return s.accountReposity.Withdraw(ctx, account.ID, sum, orderNumber)
}
