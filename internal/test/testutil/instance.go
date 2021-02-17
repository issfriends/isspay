package testutil

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"gorm.io/gorm/logger"

	"github.com/issfriends/isspay/internal/app"
	"github.com/issfriends/isspay/internal/delivery/bot"
	"github.com/issfriends/isspay/internal/delivery/restful"
	"github.com/issfriends/isspay/internal/repository/database"
	"github.com/labstack/echo/v4"
	"github.com/pressly/goose"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gox/container"
	"github.com/vx416/gox/dbprovider"
	"gorm.io/gorm"

	"go.uber.org/fx"
)

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

type TestInstance struct {
	*container.Builder
	Ctx      context.Context
	DB       dbprovider.GormProvider
	Database *database.Database
	Svc      *app.Service
	Serv     *echo.Echo
	App      *fx.App
}

// func (ti *TestInstance) BuildRedis() (*redis.Client, error) {
// 	return ti.RunRedis("isspay_redis")
// }

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

func (ti *TestInstance) KillAll() error {
	return ti.PruneAll()
}

func (ti *TestInstance) BuildGorm() (dbprovider.GormProvider, error) {
	pgCfg, err := ti.RunPg("isspay_test", "isspay_test", "15432")
	if err != nil {
		return nil, err
	}
	dbCfg := &dbprovider.DBConfig{
		Host:     pgCfg.Host,
		Port:     pgCfg.Port,
		User:     pgCfg.Username,
		Password: pgCfg.Password,
		DBName:   pgCfg.DBName,
		Type:     dbprovider.Pg,
	}

	gormConfig := &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: 200 * time.Millisecond,
			LogLevel:      logger.Info,
			Colorful:      true,
		}),
	}

	db, err := dbprovider.NewGorm(dbCfg, gormConfig)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (ti *TestInstance) RunMigration(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	migrationPath := filepath.Join(dir, "../../../deployments/migrations")

	return goose.Up(sqlDB, migrationPath)
}

func (ti *TestInstance) ProvideDB() fx.Option {
	return fx.Options(
		fx.Provide(
			ti.BuildGorm,
			database.New,
		),
		fx.Invoke(ti.RunMigration, ti.SetupFactory),
		fx.Populate(&ti.DB, &ti.Database),
	)
}

func (ti *TestInstance) ProvideSvc() fx.Option {
	return fx.Options(
		ti.ProvideDB(),
		fx.Provide(
			app.New,
		),
		fx.Populate(&ti.Svc),
	)
}

func (ti *TestInstance) ProvideRestfulHandler() fx.Option {
	return fx.Options(
		ti.ProvideSvc(),
		fx.Provide(
			echo.New,
			restful.New,
		),
		fx.Invoke(restful.Routes),
		fx.Populate(&ti.Serv),
	)
}

func (ti *TestInstance) ProvideBotHandler() fx.Option {
	return fx.Options(
		ti.ProvideSvc(),
		fx.Provide(
			bot.New,
		),
	)
}

func (ti *TestInstance) SetupFactory(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	gofactory.Opt().SetDB(sqlDB, "postgres")
	gofactory.Opt().SetTagProcess(gofactory.GormTagProcess)
	return nil
}

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

func (ti *TestInstance) Finish(kill bool) error {
	err := ti.App.Stop(ti.Ctx)
	if err != nil {
		return err
	}
	if kill {
		if err := ti.KillAll(); err != nil {
			return err
		}
	}
	return nil
}
