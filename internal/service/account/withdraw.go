package account

import (
	"context"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *service) Withdraw(ctx context.Context, userID int, orderNumber string, sum float64) error {
	account, err := s.accountReposity.FindByUserID(ctx, userID)

	if err != nil && !errors.Is(err, model.ErrorAccountNotFound) {
		return err
	}

	if account == nil {
		account, err = s.accountReposity.Create(ctx, userID)
		if err != nil {
			return err
		}
	}

	return s.accountReposity.Withdraw(ctx, account.ID, sum, orderNumber)
}
