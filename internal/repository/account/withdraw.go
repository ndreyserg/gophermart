package account

import (
	"context"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (s *repository) Withdraw(ctx context.Context, accountID int, sum float64, orderNumber string) error {
	tx, err := s.db.Begin()
	if err != nil {
		return err
	}

	row := tx.QueryRowContext(
		ctx,
		"select balance from accounts where id = $1 for update",
		accountID,
	)

	if row.Err() != nil {
		tx.Rollback()
		return row.Err()
	}

	var balance float64
	err = row.Scan(&balance)

	if err != nil {
		tx.Rollback()
		return err
	}

	balance -= sum
	if balance < 0 {
		return model.ErrorAccountNegativeBalance
	}

	_, err = tx.ExecContext(ctx, "update accounts set balance = balance - $1 where id = $2", sum, accountID)

	if err != nil {
		tx.Rollback()
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		`insert into withdrawals (sum, account_id, order_number, processed_at) 
		values ($1, $2, $3, NOW())`,
		sum,
		accountID,
		orderNumber,
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}
