package order

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) Create(ctx context.Context, number string, userID int) (*model.Order, error) {
	row := r.db.QueryRowContext(
		ctx,
		`insert into orders (number, status, accrual, user_id, uploaded_at) values ($1, $2, $3, $4, NOW()) 
		on conflict(number) do nothing
		returning id, number, status, accrual, user_id, uploaded_at`,
		number,
		model.OrderStatusNew,
		0,
		userID,
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
