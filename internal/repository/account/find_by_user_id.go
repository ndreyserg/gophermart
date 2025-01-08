package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) FindByUserID(ctx context.Context, userID int) (*model.Account, error) {
	row := r.db.QueryRowContext(
		ctx,
		`select id, user_id, balance from accounts where user_id = $1`,
		userID,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}

	account := model.Account{}

	err := row.Scan(&account.ID, &account.UserID, &account.Balance)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorAccountNotFound
		}
		return nil, err
	}
	return &account, nil
}
