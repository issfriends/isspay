package bothandler

import (
	"strconv"
	"testing"

	"github.com/issfriends/isspay/internal/app/account"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/internal/pkg/chatbot"
	"github.com/issfriends/isspay/internal/test/testutil"
	"github.com/issfriends/isspay/pkg/factory"
	"github.com/stretchr/testify/suite"
)

func TestAccountBotHandler(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

type AccountSuite struct {
	suite.Suite
	*botSuite
	assert *testutil.Assertion
}

func (su *AccountSuite) SetupSuite() {
	su.botSuite = &botSuite{}
	su.Require().NoError(su.Start())
	su.assert = &testutil.Assertion{
		Suite:    su.Suite,
		Database: su.Database,
		Ctx:      su.Ctx,
	}
}

func (su *AccountSuite) SetupTest() {
	// err := dbprovider.Truncate(su.DB,
	// 	"accounts", "wallets",
	// )
	// su.Require().NoError(err)
}

func (su *AccountSuite) TearDownSuite() {
	su.Require().NoError(su.Finish(false))
}

func (su *AccountSuite) TestSignUp() {

	tcs := []struct {
		name string
		acc  *model.Account
	}{
		{
			"masterAccount",
			factory.Account.Role(value.Master).MustBuild().(*model.Account),
		},
	}

	for _, tc := range tcs {
		su.Run(tc.name, func() {
			acc := tc.acc
			msgs := chatbot.TestForm(view.SignUpCmd, acc.Email, acc.NickName, strconv.Itoa(int(acc.Role)))
			q := &account.GetAccountQuery{
				Email: acc.Email,
			}
			for _, msg := range msgs {
				err := su.bot.HandleMsg(msg)
				su.Require().NoError(err)
			}
			su.assert.AssertAccountExist(q, true, true)
		})
	}

}
