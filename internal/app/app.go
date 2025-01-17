package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/ndreyserg/gophermart/internal/api"
	"github.com/ndreyserg/gophermart/internal/config"
	"github.com/ndreyserg/gophermart/internal/db"
	"github.com/ndreyserg/gophermart/internal/logger"
)

type App struct {
	serviceProvider *serviceProvider
	db              *sql.DB
	config          *config.Config
}

func (a *App) Run() error {
	logger.Log.Info("start server ", a.config.RunAddress)
	err := http.ListenAndServe(
		a.config.RunAddress,
		api.NewRouter(
			a.serviceProvider.UserService(),
			a.serviceProvider.OrderService(),
			a.serviceProvider.AccountService(),
			a.serviceProvider.SessionService(),
		),
	)

	if err != nil {
		return fmt.Errorf("app run error: %w", err)
	}

	return nil
}

func (a *App) init(ctx context.Context) error {
	a.config = config.NewConfig()
	err := logger.Initialize(a.config.LogLevel)

	if err != nil {
		return fmt.Errorf("app init logger error %w", err)
	}

	conn, err := db.NewDB(ctx, a.config.DatabaseURI)
	if err != nil {
		return fmt.Errorf("app init db error %w", err)
	}
	a.db = conn
	a.serviceProvider = newServiceProvider(a.db, a.config.AccrualSystemAderess, a.config.SecretKey)
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
