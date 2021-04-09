package delivery

import (
	"context"

	"github.com/issfriends/isspay/internal/delivery/bot"
	"github.com/issfriends/isspay/internal/delivery/restful"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/issfriends/isspay/pkg/config"
	"github.com/issfriends/isspay/pkg/server"
	"github.com/vx416/gox/log"
	"go.uber.org/fx"
)

func RunServer(lc fx.Lifecycle, handler *restful.Handler, botHandler *bot.Handler) error {
	cfg := config.Get()
	linebot, err := chatbot.NewLineBot(cfg.Secrets.Linebot)
	if err != nil {
		return err
	}
	if err := botHandler.Routes(linebot); err != nil {
		return err
	}

	engine, err := server.NewEcho(cfg.HTTPServer, handler.Routes, linebot.HookOnEcho)
	if err != nil {
		return err
	}

	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go func() {
					if err := engine.Run(); err != nil {
						log.Get().Errorf("server: run server failed, err:%+v", err)
					}
				}()
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return engine.Shutdown(ctx)
			},
		},
	)
	return nil
}
