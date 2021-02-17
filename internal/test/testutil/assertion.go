package testutil

import (
	"context"

	"github.com/shopspring/decimal"
	"github.com/vx416/gox/converter"

	"github.com/issfriends/isspay/internal/app/account"
	"github.com/issfriends/isspay/internal/app/model"

	"github.com/issfriends/isspay/internal/repository/database"
	"github.com/stretchr/testify/suite"
)

type Assertion struct {
	suite.Suite
	Ctx      context.Context
	Database *database.Database
}

func (su *Assertion) AssertAccountExist(q *account.GetAccountQuery, exist, hasWallet bool) *model.Account {
	err := su.Database.Account().GetAccount(su.Ctx, q)
	if exist {
		su.Require().NoError(err)
		su.Require().NotNil(q.Data)
		su.Require().NotZero(q.Data.ID)
		su.Assert().NotZero(q.Data.ID)
		su.Assert().NotEmpty(q.Data.Email)
		su.Assert().NotEmpty(q.Data.UID)
		su.Assert().NotEmpty(q.Data.MessengerID.String)
		su.Assert().False(q.Data.CreatedAt.IsZero())
		su.Assert().False(q.Data.UpdatedAt.IsZero())

		if hasWallet {
			su.Require().NotNil(q.Data.Wallet)
			su.Assert().Equal(q.Data.Wallet.OwnerID, q.Data.ID)
			su.Assert().NotZero(q.Data.Wallet.ID)
			su.Assert().GreaterOrEqual(q.Data.Wallet.Amount, decimal.Zero)
		}

		return q.Data
	}

	su.Require().Error(err)
	su.Require().Nil(q.Data)
	return nil
}

func AssertProductsEq(su suite.Suite, a, b []*model.Product, eq bool) {
	productsMap := converter.SliceToMap(a, "Name").(map[string]*model.Product)
	result := true

	if len(a) == len(b) {
		for _, product := range b {
			product.Price = product.Price.Round(2)
			productsMap[product.Name].Price = productsMap[product.Name].Price.Round(2)
			product.Cost = product.Cost.Round(2)
			productsMap[product.Name].Cost = productsMap[product.Name].Cost.Round(2)

			aP := productsMap[product.Name]
			bP := product
			if !aP.Price.Equal(bP.Price) {
				result = false
				break
			}
			if !aP.Cost.Equal(bP.Cost) {
				result = false
				break
			}
			if aP.Name != bP.Name {
				result = false
				break
			}
			if aP.Category != bP.Category {
				result = false
				break
			}
			if aP.ImageURL != bP.ImageURL {
				result = false
				break
			}
		}
	} else {
		result = false
	}

	su.Assert().Equalf(result, eq, "products equal should be %+v, but %+v", eq, result)
}
