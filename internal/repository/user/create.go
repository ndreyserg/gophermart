package user

import (
	"context"
	"database/sql"
	"errors"

	"github.com/ndreyserg/gophermart/internal/model"
)

func (r *repository) Create(ctx context.Context, login string, passHash string) (*model.User, error) {
	row := r.db.QueryRowContext(
		ctx,
		`insert into users (login, pass_hash) values ($1, $2) 
		on conflict do nothing
		returning id,login`,
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
			return nil, model.ErrorUserAllreadyExist
		}
		return nil, err
	}

	return &user, nil
}
