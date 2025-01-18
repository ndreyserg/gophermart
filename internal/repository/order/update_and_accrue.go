package order

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) UpdateAndAccrue(ctx context.Context, order *model.Order, accountID int) error {
	tx, err := r.db.Begin()

	if err != nil {
		return fmt.Errorf("update and accrue order %s start error: %w", order.Number, err)
	}

	_, err = tx.ExecContext(
		ctx,
		`update orders set status = $1, accrual = $2 where number = $3`,
		order.Status,
		order.Accrual,
		order.Number,
	)

	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update and accrue order %s update error: %w", order.Number, err)
	}

	_, err = tx.Exec(
		"update accounts set balance = balance + $1 where id = $2",
		order.Accrual,
		accountID,
	)

	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("update and accrue order %s accrue error: %w", order.Number, err)
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("update and accrue order %s commit error: %w", order.Number, err)
	}
	return nil
}
