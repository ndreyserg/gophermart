package order

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) GetByUserID(ctx context.Context, userID int) ([]*model.Order, error) {
	res := []*model.Order{}
	rows, err := r.db.QueryContext(
		ctx,
		`select id, number, status, accrual, user_id, to_char(uploaded_at, 'YYYY-MM-DD"T"HH24:MI:SSTZH:TZM') 
		from orders where user_id = $1 
		order by uploaded_at desc`,
		userID,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		order := new(model.Order)
		err := rows.Scan(&order.ID, &order.Number, &order.Status, &order.Accrual, &order.UserID, &order.UploadedAt)
		if err != nil {
			return nil, err
		}
		res = append(res, order)
	}
	return res, nil
}
