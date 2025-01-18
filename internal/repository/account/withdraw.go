package account

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) Withdraw(ctx context.Context, accountID int, sum float64, orderNumber string) error {
	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("withdraw start trans error: %w", err)
	}

	row := tx.QueryRowContext(
		ctx,
		"select balance from accounts where id = $1 for update",
		accountID,
	)

	if row.Err() != nil {
		_ = tx.Rollback()
		return fmt.Errorf("withdraw exec get balance error: %w", err)
	}

	var balance float64
	err = row.Scan(&balance)

	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("withdraw scan get balance error: %w", err)
	}

	balance -= sum
	if balance < 0 {
		return model.ErrAccountNegativeBalance
	}

	_, err = tx.ExecContext(ctx, "update accounts set balance = balance - $1 where id = $2", sum, accountID)

	if err != nil {
		_ = tx.Rollback()
		return fmt.Errorf("withdraw update balance error: %w", err)
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
		_ = tx.Rollback()
		return fmt.Errorf("withdraw insert withdrawals error: %w", err)
	}

	err = tx.Commit()

	if err != nil {
		return fmt.Errorf("withdraw commit error: %w", err)
	}
	return nil
}
