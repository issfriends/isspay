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
		Name string `yaml:"name"`
		Env  Env    `yaml:"env"`
	} `yaml:"app"`

	DB    *dbprovider.DBConfig `yaml:"db"`
	Cache *cache.RedisCfg      `yaml:"cache"`

	Log        *log.Config    `yaml:"log"`
	HTTPServer *server.Config `yaml:"http_server"`

	Secrets *Secrets
}

// // ProvideDB provide fx db options
// func (cfg *Config) ProvideDB() fx.Option {
// 	gormCfg := newGormConfig(cfg.App.Env)

// 	return fx.Options(
// 		fx.Supply(gormCfg),
// 		fx.Provide(dbprovider.NewGorm),
// 	)
// }

// // InvokeServer run server invoker
// func (cfg *Config) InvokeServer() fx.Option {
// 	return fx.Invoke(
// 		cfg.runServer,
// 	)
// }

// func (cfg *Config) runServer(lc fx.Lifecycle, handler *restful.Handler, botHandler *bot.Handler) error {
// 	linebot, err := chatbot.NewLineBot(cfg.Secrets.Linebot)
// 	if err != nil {
// 		return err
// 	}
// 	if err := botHandler.Routes(linebot); err != nil {
// 		return err
// 	}

// 	engine, err := server.NewEcho(cfg.HTTPServer, handler.Routes, linebot.HookOnEcho)
// 	if err != nil {
// 		return err
// 	}

// 	lc.Append(
// 		fx.Hook{
// 			OnStart: func(ctx context.Context) error {
// 				go func() {
// 					if err := engine.Run(); err != nil {
// 						log.Get().Errorf("server: run server failed, err:%+v", err)
// 					}
// 				}()
// 				return nil
// 			},
// 			OnStop: func(ctx context.Context) error {
// 				return engine.Shutdown(ctx)
// 			},
// 		},
// 	)

// 	return nil
// }
