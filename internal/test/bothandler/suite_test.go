package bothandler

import (
	"errors"

	"github.com/issfriends/isspay/internal/delivery/bot"
	"github.com/issfriends/isspay/internal/pkg/chatbot"
	"github.com/issfriends/isspay/internal/repository/database"
	"github.com/issfriends/isspay/internal/test/testutil"
	"go.uber.org/fx"
)

type botSuite struct {
	*testutil.TestInstance
	handler  *bot.Handler
	bot      chatbot.ChatBot
	Database *database.Database
}

func (su *botSuite) Start(ㄊ) error {
	var err error
	if su.TestInstance, err = testutil.New(); err != nil {
		return nil
	}

	su.bot = chatbot.TestBot()

	err = su.TestInstance.Start(
		fx.Options(
			fx.Supply(su.bot),
			su.ProvideBotHandler(),
			fx.Populate(&su.handler),
			fx.Populate(&su.Database),
		),
	)

	if su.handler == nil {
		return errors.New("handler is nil")
	}

	bot.Routes(su.bot, su.handler)
	if err != nil {
		return err
	}

	return nil
}
