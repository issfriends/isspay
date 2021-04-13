package config

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"

	"go.uber.org/fx"

	"github.com/issfriends/isspay/pkg/server"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/vx416/gox/cache"
	"github.com/vx416/gox/dbprovider"
	"github.com/vx416/gox/log"
)

// Env application environment
type Env string

// Dev is dev mode
func (env Env) Dev() bool {
	return strings.ToLower(string(env)) == "dev"
}

// Prod is production mode
func (env Env) Prod() bool {
	return strings.ToLower(string(env)) == "production" || strings.ToLower(string(env)) == "prod"
}

var (
	_config *Config
	once    sync.Once
)

// Init initialize global singleton configuration
func Init() (*Config, error) {
	var err error
	once.Do(func() {
		configPath := os.Getenv("CONFIG_PATH")
		if configPath == "" {
			_, f, _, _ := runtime.Caller(0)
			dir := filepath.Dir(f)
			configPath = filepath.Join(dir, "../../configs")
		}
		configName := os.Getenv("CONFIG_NAME")
		if configName == "" {
			configName = "app"
		}

		viper.SetConfigName(configName)
		viper.AddConfigPath(configPath)
		viper.SetConfigType("yaml")
		err = viper.ReadInConfig()
		if err != nil {
			return
		}

		err = viper.Unmarshal(&_config, func(decode *mapstructure.DecoderConfig) {
			decode.TagName = "yaml"
		})
		if err != nil {
			return
		}

		secrets, serr := initSecret()
		if serr != nil {
			err = serr
			return
		}
		_config.Secrets = secrets

	})
	if err != nil {
		return nil, err
	}
	return _config, nil
}

// Get get global singleton configuration
func Get() *Config {
	if _config == nil {
		_, err := Init()
		panic(err)
	}
	return _config
}

// Config application configuration
type Config struct {
	fx.Out

	App struct {
		Name              string `yaml:"name"`
		Env               Env    `yaml:"env"`
		IntervieweeAmount int64  `yaml:"interviewee_amount"`
	} `yaml:"app"`

	DB    *dbprovider.DBConfig `yaml:"db"`
	Cache *cache.RedisCfg      `yaml:"cache"`

	Log        *log.Config    `yaml:"log"`
	HTTPServer *server.Config `yaml:"http_server"`

	Secrets *Secrets
}

// ProvideDB provide fx db options
func (cfg *Config) ProvideInfra() fx.Option {
	gormCfg := newGormConfig(cfg.App.Env)

	return fx.Options(
		fx.Supply(*cfg),
		fx.Supply(gormCfg),
		fx.Provide(dbprovider.NewGorm),
		fx.Provide(cache.NewRedis),
	)
}
