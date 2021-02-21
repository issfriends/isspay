package config

import (
	"time"

	"github.com/vx416/gox/log"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func newGormConfig(env Env) *gorm.Config {
	logCfg := logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      logger.Warn,
		Colorful:      false,
	}

	if env.Dev() {
		logCfg.LogLevel = logger.Info
		logCfg.Colorful = true
	}

	gormLogger := log.NewGormLogger(logCfg)

	gormCfg := &gorm.Config{}
	gormCfg.Logger = gormLogger
	gormCfg.NowFunc = func() time.Time { return time.Now().UTC() }

	return gormCfg
}
