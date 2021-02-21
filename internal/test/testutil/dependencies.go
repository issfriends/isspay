package testutil

import (
	"path/filepath"
	"runtime"
	"time"

	"github.com/pressly/goose"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gox/dbprovider"
	"github.com/vx416/gox/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (ti *TestInstance) buildGorm() (dbprovider.GormProvider, error) {
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
		Logger: log.NewGormLogger(logger.Config{
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

func (ti *TestInstance) setupFactory(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	gofactory.Opt().SetDB(sqlDB, "postgres")
	gofactory.Opt().SetTagProcess(gofactory.GormTagProcess)
	return nil
}

func (ti *TestInstance) setupLog() error {
	logCfg := log.Config{
		AppName: "isspay_test",
		Env:     "dev",
		Type:    "zap",
		Level:   "debug",
	}
	if _, err := logCfg.Build(); err != nil {
		return err
	}
	return nil
}

func (ti *TestInstance) setupMigration(db dbprovider.GormProvider) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	migrationPath := filepath.Join(dir, "../../../deployments/migrations")

	return goose.Up(sqlDB, migrationPath)
}

func (ti *TestInstance) killAll() error {
	return ti.PruneAll()
}
