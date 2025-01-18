package account

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) FindByUserID(ctx context.Context, userID int) (*model.Account, error) {
	row := r.db.QueryRowContext(
		ctx,
		`select id, user_id, balance from accounts where user_id = $1`,
		userID,
	)

	if row.Err() != nil {
		return nil, fmt.Errorf("repo find exec error: %w", row.Err())
	}

	account := model.Account{}

	err := row.Scan(&account.ID, &account.UserID, &account.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrAccountNotFound
		}
		return nil, fmt.Errorf("repo find scan error: %w", err)
	}
	return &account, nil
}
