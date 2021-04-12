package bothandler

import (
	"errors"
	"strconv"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/internal/test/testutil"
	"github.com/issfriends/isspay/pkg/chatbot"
	"go.uber.org/fx"
)

type botSuite struct {
	*testutil.TestInstance
}

func (su *botSuite) Start() error {
	var err error
	if su.TestInstance, err = testutil.New(); err != nil {
		return nil
	}

	err = su.TestInstance.Start(
		fx.Options(
			su.ProvideBotHandler(),
		),
	)
	if err != nil {
		return err
	}

	if su.Bot == nil {
		return errors.New("bot is nil")
	}
	return nil
}

func (su *botSuite) signUp(acc *model.Account) error {
	msgs := chatbot.TestForm(view.SignUpCmd, acc.Email, acc.NickName, strconv.Itoa(int(acc.Role)))
	for _, msg := range msgs {
		chatbot.SetTestMsgID(msg, acc.MessengerID.String)
		err := su.Bot.HandleMsg(msg)
		if err != nil {
			return err
		}
	}
	return nil
}
