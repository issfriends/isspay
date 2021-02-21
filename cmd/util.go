package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"runtime"
	"syscall"
	"time"

	"github.com/issfriends/isspay/pkg/factory"
	"github.com/pressly/goose"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gox/dbprovider"
	"github.com/vx416/gox/log"
	"go.uber.org/fx"
)

func runApp(app *fx.App) {
	ctx := context.Background()
	err := app.Start(ctx)
	if err != nil {
		log.Get().Errorf("server: start app failed, err:%+v", err)
		os.Exit(1)
	}

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-sigterm:
		log.Get().Debug("server: shutdown process start")
	}
	stopCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := app.Stop(stopCtx); err != nil {
		log.Get().Errorf("server: shutdown process failed, err:%+v", err)
	}
}

func recoverPanic() {
	if r := recover(); r != nil {
		var msg string
		for i := 2; ; i++ {
			_, file, line, ok := runtime.Caller(i)
			if !ok {
				break
			}
			msg = msg + fmt.Sprintf("%s:%d\n", file, line)
		}
		log.Get().Errorf("%s\n↧↧↧↧↧↧ PANIC ↧↧↧↧↧↧\n%s↥↥↥↥↥↥ PANIC ↥↥↥↥↥↥", r, msg)
	}
}

func runMigrations(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	migrationPath := filepath.Join(dir, "../deployments/migrations")
	return goose.Up(sqlDB, migrationPath)
}

func initTestData(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	gofactory.Opt().SetDB(sqlDB, "postgres")
	gofactory.Opt().SetTagProcess(gofactory.GormTagProcess)

	_, err = factory.Product.InsertN(100)
	return err
}
