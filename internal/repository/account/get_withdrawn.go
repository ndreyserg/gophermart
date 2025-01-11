package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

func (r *repository) GetWithdrawn(ctx context.Context, accID int) (float64, error) {
	row := r.db.QueryRowContext(
		ctx,
		"select coalesce(sum(w.sum), 0) from withdrawals w where w.account_id = $1",
		accID,
	)

	if row.Err() != nil {
		return 0, fmt.Errorf("get withdrawn exec error %w", row.Err())
	}

	var sum float64

	err := row.Scan(&sum)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, fmt.Errorf("get withdrawn scan error: %w", err)
	}

	return sum, nil
}
