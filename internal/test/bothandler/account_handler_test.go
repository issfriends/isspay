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

func (s *AccountSuite) SetupSuite() {
	s.botSuite = &botSuite{}
	s.Require().NoError(s.Start())
}

func (s *AccountSuite) SetupTest() {
	s.Require().NoError(s.TruncateTables("accounts", "wallets"))
}

func (s *AccountSuite) TearDownSuite() {
	s.Require().NoError(s.Shutdown(false))
}

func (s *AccountSuite) TestSignUp() {

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
		s.Run(tc.name, func() {
			acc := tc.acc
			msgs := chatbot.TestForm(view.SignUpCmd, acc.Email, acc.NickName, strconv.Itoa(int(acc.Role)))
			q := &query.GetAccountQuery{
				Email: acc.Email,
			}
			for _, msg := range msgs {
				err := s.Bot.HandleMsg(msg)
				s.Require().NoError(err)
			}
			err := s.Database.GetAccount(s.Ctx, q)
			s.Require().NoError(err)
			s.Assert().NotNil(q.Data, "account cannot be nil")
			s.Equal(q.Data.Email, acc.Email, "email should be equal")
		})
	}
}
