package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) FindByLoginAndPassHash(ctx context.Context, login string, passHash string) (*model.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		`select id, login from users where login = $1 and pass_hash = $2`,
		login,
		passHash,
	)

	if row.Err() != nil {
		return nil, fmt.Errorf("find user query error: %w,", row.Err())
	}

	user := model.User{}

	err := row.Scan(&user.ID, &user.Login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrUserNotFound
		}
		return nil, fmt.Errorf("find user scan error: %w,", err)
	}
	return &user, nil
}
