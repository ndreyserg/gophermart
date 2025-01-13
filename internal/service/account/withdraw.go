package account

import (
	"context"
	"fmt"
)

func (s *service) Withdraw(ctx context.Context, userID int, orderNumber string, sum float64) error {
	err := s.checker.Check(orderNumber)

	if err != nil {
		return fmt.Errorf("witdraw error: %w", err)
	}

	account, err := s.GetOrCreate(ctx, userID)

	if err != nil {
		return fmt.Errorf("witdraw error: %w", err)
	}
	err = s.accountReposity.Withdraw(ctx, account.ID, sum, orderNumber)
	if err != nil {
		return fmt.Errorf("witdraw error: %w", err)
	}
	return nil
}
