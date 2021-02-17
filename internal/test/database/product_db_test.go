package database

import (
	"testing"

	"github.com/issfriends/isspay/pkg/factory"
	"github.com/shopspring/decimal"

	"github.com/issfriends/isspay/internal/app/inventory"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/issfriends/isspay/internal/test/testutil"
	"github.com/stretchr/testify/suite"
)

func TestInventoryDB(t *testing.T) {
	suite.Run(t, new(InventorySuite))
}

type InventorySuite struct {
	suite.Suite
	*dbSuite
	assert *testutil.Assertion
}

func (su *InventorySuite) SetupSuite() {
	su.dbSuite = &dbSuite{}
	err := su.Start()
	su.Require().NoError(err)
	su.assert = &testutil.Assertion{
		Suite:    su.Suite,
		Database: su.Database,
		Ctx:      su.Ctx,
	}
}

func (su *InventorySuite) SetupTest() {
	// err := dbprovider.Truncate(su.DB,
	// 	"products",
	// )
	// su.Require().NoError(err)
}

func (su *InventorySuite) TearDownSuite() {
	err := su.Finish(false)
	su.Require().NoError(err)
}

func (su *InventorySuite) TestListProducts() {
	tcs := []struct {
		name  string
		setup func() []*model.Product
		q     *inventory.ListProductsQuery
	}{
		{
			"snake",
			func() []*model.Product {
				factory.Product.Category(value.Drink).MustInsertN(5)
				return factory.Product.Category(value.Snake).MustInsertN(5).([]*model.Product)
			},
			&inventory.ListProductsQuery{
				Category: value.Snake,
			},
		},
		{
			"drink",
			func() []*model.Product {
				factory.Product.Category(value.Snake).MustInsertN(5)
				return factory.Product.Category(value.Drink).MustInsertN(5).([]*model.Product)
			},
			&inventory.ListProductsQuery{
				Category: value.Drink,
			},
		},
		{
			"price >= 10",
			func() []*model.Product {
				factory.Product.Price(5).MustInsertN(5)
				return factory.Product.Price(10).MustInsertN(5).([]*model.Product)
			},
			&inventory.ListProductsQuery{
				PriceGte: decimal.NewFromInt(10),
			},
		},
		{
			"quantity > 0",
			func() []*model.Product {
				factory.Product.Quantity(0).MustInsertN(5)
				return factory.Product.Quantity(10).MustInsertN(6).([]*model.Product)
			},
			&inventory.ListProductsQuery{
				QuantityGte: 1,
			},
		},
	}

	for _, tc := range tcs {
		su.Run(tc.name, func() {
			su.SetupTest()
			products := tc.setup()
			err := su.Database.Inventory().ListProducts(su.Ctx, tc.q)
			su.Require().NoError(err)
			su.Require().NotNil(tc.q.Data)
			su.Require().Len(tc.q.Data, len(products))
			testutil.AssertProductsEq(su.Suite, tc.q.Data, products, true)
		})

	}
}

func (su *InventorySuite) TestUpdateProduct() {
	// tcs := []struct {
	// 	name string
	// }{}

	// for _, tc := range tcs {

	// }

}
