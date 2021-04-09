package testutil

import (
	"context"
	"errors"
	"fmt"

	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/delivery/restful"
	"github.com/issfriends/isspay/internal/repository/database"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/issfriends/isspay/pkg/i18n"
	"github.com/labstack/echo/v4"
	"github.com/vx416/gox/container"
	"github.com/vx416/gox/dbprovider"

	"go.uber.org/fx"
)

// New new test application instance
func New() (*TestInstance, error) {
	c, err := container.NewConBuilder()
	if err != nil {
		return nil, err
	}
	return &TestInstance{
		Builder: c,
		Ctx:     context.Background(),
	}, nil
}

// TestInstance test application instance
type TestInstance struct {
	*container.Builder
	Ctx      context.Context
	DB       dbprovider.GormProvider
	Database *database.Database
	Svc      *app.App
	Serv     *echo.Echo
	Bot      chatbot.ChatBot
	App      *fx.App
}

// ProvideDB provide db layer fx options
func (ti *TestInstance) ProvideDB() fx.Option {
	return fx.Options(
		fx.Provide(
			ti.buildGorm,
			database.New,
		),
		fx.Invoke(ti.setupMigration, ti.setupFactory, ti.setupLog, i18n.Initi18n),
		fx.Populate(&ti.DB, &ti.Database),
	)
}

// ProvideSvc provide service layer fx options
func (ti *TestInstance) ProvideSvc() fx.Option {
	return fx.Options(
		ti.ProvideDB(),
		fx.Provide(
			app.New,
		),
		fx.Populate(&ti.Svc),
	)
}

// ProvideRestfulHandler provide restful layer fx options
func (ti *TestInstance) ProvideRestfulHandler() fx.Option {
	return fx.Options(
		ti.ProvideSvc(),
		fx.Provide(
			echo.New,
			restful.New,
		),
		fx.Invoke(func(h *restful.Handler, e *echo.Echo) {
			h.Routes(e)
		}),
		fx.Populate(&ti.Serv),
	)
}

// ProvideBotHandler provide bothandler layer fx options
func (ti *TestInstance) ProvideBotHandler() fx.Option {
	return fx.Options(
		ti.ProvideSvc(),
		fx.Provide(
			chatbot.TestBot,
			// bot.New,
		),
		// fx.Invoke(func(h *bot.Handler, chatbot chatbot.ChatBot) {
		// 	h.Routes(chatbot)
		// }),
		fx.Populate(&ti.Bot),
	)
}

// Start start test application
func (ti *TestInstance) Start(option fx.Option) error {
	app := fx.New(
		option,
		fx.NopLogger,
	)

	err := app.Start(ti.Ctx)
	if err != nil {
		return err
	}
	ti.App = app
	return nil
}

// Shutdown shutdown test application
func (ti *TestInstance) Shutdown(kill bool) error {
	err := ti.App.Stop(ti.Ctx)
	if err != nil {
		return err
	}
	if kill {
		if err := ti.killAll(); err != nil {
			return err
		}
	}
	return nil
}

// TruncateTables delete tables
func (ti *TestInstance) TruncateTables(tables ...string) error {
	if ti.DB == nil {
		return errors.New("db is nil")
	}

	sqlDB, err := ti.DB.DB()
	if err != nil {
		return err
	}

	for _, table := range tables {
		_, err := sqlDB.Exec(fmt.Sprintf("TRUNCATE %s CASCADE;", table))
		if err != nil {
			return err
		}
	}
	return nil
}
