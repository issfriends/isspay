package bothandler

import (
	"strconv"
	"testing"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/delivery/bot/view"
	"github.com/issfriends/isspay/pkg/chatbot"
	"github.com/issfriends/isspay/pkg/factory"
	"github.com/stretchr/testify/suite"
)

func TestAccountBotHandler(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

type AccountSuite struct {
	suite.Suite
	*botSuite
}

func (su *AccountSuite) SetupSuite() {
	su.botSuite = &botSuite{}
	su.Require().NoError(su.Start())
	su.SetupAssertion(su.Suite)
}

func (su *AccountSuite) SetupTest() {
	su.Require().NoError(su.TruncateTables("accounts", "wallets"))
}

func (su *AccountSuite) TearDownSuite() {
	su.Require().NoError(su.Shutdown(false))
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
			q := &query.GetAccountQuery{
				Email: acc.Email,
			}
			for _, msg := range msgs {
				err := su.Bot.HandleMsg(msg)
				su.Require().NoError(err)
			}
			su.AssertHelper.AssertAccountExist(q, true, true)
		})
	}
}
