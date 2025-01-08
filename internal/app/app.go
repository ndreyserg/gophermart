package app

import (
	"context"
	"database/sql"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ndreyserg/gophermart/internal/api"
)

type App struct {
	serviceProvider *serviceProvider
	db              *sql.DB
}

func (a *App) Run() error {
	return http.ListenAndServe(
		":8080",
		api.NewRouter(
			a.serviceProvider.UserService(),
			a.serviceProvider.OrderService(),
			a.serviceProvider.AccountService(),
			"eeee",
		),
	)
}

func (a *App) init(ctx context.Context) error {
	err := a.initDB(ctx, "postgres://gophermart:gophermart@localhost:5432/gophermart")
	if err != nil {
		return err
	}
	a.serviceProvider = newServiceProvider(a.db)
	return nil
}

func (a *App) initDB(ctx context.Context, dsn string) error {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return err
	}

	if err = db.PingContext(ctx); err != nil {
		return err
	}

	a.db = db
	return nil
}

func NewApp(ctx context.Context) (*App, error) {
	app := App{}
	err := app.init(ctx)

	if err != nil {
		return nil, err
	}
	return &app, nil
}
