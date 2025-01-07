package user

import (
	"context"
	"database/sql"
	"errors"

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
		return nil, row.Err()
	}

	user := model.User{}

	err := row.Scan(&user.ID, &user.Login)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, model.ErrorUserNotFound
		}
		return nil, err
	}
	return &user, nil
}
