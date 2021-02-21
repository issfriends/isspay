package chatbot

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (bot *lineBot) HookOnEcho(e *echo.Echo) {
	e.POST("/webhook", func(c echo.Context) error {
		bot.Webhook(c.Response().Writer, c.Request())
		return nil
	})
}

func (bot *lineBot) Webhook(resp http.ResponseWriter, req *http.Request) {
	events, err := bot.ParseRequest(req)
	if err != nil {
		panic(err)
	}

	ctx := req.Context()

	for _, event := range events {
		log.Printf("%+v", event)
		msg, err := parseMsg(ctx, bot, event)

		if err != nil {
			bot.errHandle(err, msg)
			continue
		}

		if err := bot.HandleMsg(msg); err != nil {
			bot.errHandle(err, msg)
		}
	}
	resp.WriteHeader(http.StatusOK)
}
