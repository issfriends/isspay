package database

import (
	"reflect"
	"testing"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	accDB "github.com/issfriends/isspay/internal/repository/database/account"

	"github.com/issfriends/isspay/pkg/factory"
	"github.com/stretchr/testify/suite"
)

func TestAccountDB(t *testing.T) {
	suite.Run(t, new(AccountSuite))
}

type AccountSuite struct {
	suite.Suite
	*dbSuite
	accountDB *accDB.AccountDB
}

func (su *AccountSuite) SetupSuite() {
	su.dbSuite = &dbSuite{}
	err := su.Start()
	su.Require().NoError(err)
	su.accountDB = su.Database.Account()
	su.SetupAssertion(su.Suite)
}

func (su *AccountSuite) SetupTest() {
	err := su.TruncateTables("accounts", "wallets")
	su.Require().NoError(err)
}

func (su *AccountSuite) TearDownSuite() {
	err := su.Shutdown(false)
	su.Require().NoError(err)
}

func (su *AccountSuite) TestCreateAccount() {
	existed := factory.Account.MustInsert().(*model.Account)

	tcs := []struct {
		name  string
		acc   *model.Account
		exist bool
		succ  bool
	}{
		{
			"create",
			factory.Account.MustBuild().(*model.Account),
			true,
			true,
		},
		{
			"duplicated",
			existed,
			true,
			false,
		},
		{
			"messengerIDIsNil",
			factory.Account.Omit("MessengerID").MustBuild().(*model.Account),
			false,
			false,
		},
	}

	for _, tc := range tcs {
		su.Run(tc.name, func() {
			err := su.accountDB.CreateAccount(su.Ctx, tc.acc)
			q := &query.GetAccountQuery{
				ID: tc.acc.ID,
			}

			if tc.succ {
				su.Require().NoError(err)
				foundAcc := su.AssertHelper.AssertAccountExist(q, tc.exist, false)
				su.Assert().True(reflect.DeepEqual(tc.acc, foundAcc))
			} else {
				su.Require().Error(err)
				su.AssertHelper.AssertAccountExist(q, tc.exist, false)
			}
		})

	}

}

func (su *AccountSuite) TestGetAccount() {
	accounts := factory.Account.MustInsertN(2).([]*model.Account)

	tcs := []struct {
		name   string
		q      *query.GetAccountQuery
		expect *model.Account
	}{
		{
			"getByEmail",
			&query.GetAccountQuery{
				Email: accounts[0].Email,
			},
			accounts[0],
		},
		{
			"getByMsgID",
			&query.GetAccountQuery{
				MessengerID: accounts[1].MessengerID.String,
			},
			accounts[1],
		},
		{
			"getByID",
			&query.GetAccountQuery{
				ID: accounts[1].ID,
			},
			accounts[1],
		},
		{
			"getByEmailAndMsgID",
			&query.GetAccountQuery{
				MessengerID: accounts[1].MessengerID.String,
				Email:       accounts[1].Email,
			},
			accounts[1],
		},
		{
			"getNothing",
			&query.GetAccountQuery{
				MessengerID: accounts[1].MessengerID.String,
				Email:       accounts[0].Email,
			},
			nil,
		},
	}

	for _, tc := range tcs {
		su.Run(tc.name, func() {
			err := su.accountDB.GetAccount(su.Ctx, tc.q)
			if tc.expect != nil {
				su.Require().NoError(err)
				data := tc.q.Data
				su.Require().NotNil(data)
				su.Assert().True(reflect.DeepEqual(data, tc.expect))
			} else {
				su.Require().Error(err)
				su.Assert().True(reflect.DeepEqual(tc.q.Data, &model.Account{}) || tc.q.Data == nil)
			}
		})
	}
}
