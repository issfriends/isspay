package bothandler

import (
	"errors"

	"github.com/issfriends/isspay/internal/test/testutil"
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

	if su.Bot == nil {
		return errors.New("bot is nil")
	}
	return nil
}
