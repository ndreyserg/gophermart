package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ndreyserg/gophermart/internal/api"
	"github.com/ndreyserg/gophermart/internal/config"
)

type App struct {
	serviceProvider *serviceProvider
	db              *sql.DB
	config          *config.Config
}

func (a *App) Run() error {
	err := http.ListenAndServe(
		a.config.RunAddress,
		api.NewRouter(
			a.serviceProvider.UserService(),
			a.serviceProvider.OrderService(),
			a.serviceProvider.AccountService(),
			a.config.SecretKey,
		),
	)

	if err != nil {
		return fmt.Errorf("app run error: %w", err)
	}

	return nil
}

func (a *App) init(ctx context.Context) error {
	a.config = config.NewConfig()
	err := a.initDB(ctx, a.config.DatabaseURI)
	if err != nil {
		return err
	}
	a.serviceProvider = newServiceProvider(a.db)
	return nil
}

func (a *App) initDB(ctx context.Context, dsn string) error {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return fmt.Errorf("open db error: %w", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return fmt.Errorf("ping db error: %w", err)
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
