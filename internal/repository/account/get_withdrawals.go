package account

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) GetWithdrawals(ctx context.Context, accID int) ([]*model.AccountWithdrawals, error) {
	rows, err := r.db.QueryContext(
		ctx,
		`select id, sum, account_id, order_number,  to_char(processed_at, 'YYYY-MM-DD"T"HH24:MI:SSTZH:TZM')  from withdrawals 
		where account_id = $1
		order by processed_at desc`,
		accID,
	)

	if err != nil {
		return nil, fmt.Errorf("get withdrawals exec error: %w", err)
	}
	defer func() {
		_ = rows.Close()
	}()

	res := []*model.AccountWithdrawals{}

	for rows.Next() {
		w := new(model.AccountWithdrawals)
		err := rows.Scan(&w.ID, &w.Sum, &w.AccountID, &w.OrderNumber, &w.ProcessedAt)
		if err != nil {
			return nil, fmt.Errorf("get withdrawals scan error: %w", err)
		}
		res = append(res, w)
	}

	err = rows.Err()

	if err != nil {
		return nil, fmt.Errorf("get withdrawals error: %w", err)
	}

	return res, nil
}
