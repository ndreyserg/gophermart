package order

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) Update(ctx context.Context, order *model.Order) error {
	row := r.db.QueryRowContext(
		ctx,
		`update orders set status = $1, accrual = $2 where number = $3`,
		order.Status,
		order.Accrual,
		order.Number,
	)

	if row.Err() != nil {
		return fmt.Errorf("order create query error: %w", row.Err())
	}
	return nil
}
