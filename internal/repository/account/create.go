package account

import (
	"context"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) Create(ctx context.Context, userID int) (*model.Account, error) {
	row := r.db.QueryRowContext(
		ctx,
		`insert into accounts (user_id, balance) values ($1, $2) 
		returning id,user_id,balance`,
		userID,
		0,
	)

	if row.Err() != nil {
		return nil, fmt.Errorf("query exec error: %w", row.Err())
	}

	account := model.Account{}

	err := row.Scan(&account.ID, &account.UserID, &account.Balance)

	if err != nil {
		return nil, fmt.Errorf("row scan error: %w", err)
	}

	return &account, nil
}
