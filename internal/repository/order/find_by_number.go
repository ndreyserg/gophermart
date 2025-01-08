package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) FindByNumber(ctx context.Context, number string) (*model.Order, error) {
	row := r.db.QueryRowContext(
		ctx,
		`select id, number, status, accrual, user_id, uploaded_at from orders where number = $1`,
		number,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}

	order := model.Order{}

	err := row.Scan(&order.ID, &order.Number, &order.Status, &order.Accrual, &order.UserID, &order.UploadedAt)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorOrderAllreadyExist
		}
		return nil, err
	}

	return &order, nil
}
