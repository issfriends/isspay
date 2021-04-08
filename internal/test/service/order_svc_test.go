package service

// package service

// import (
// 	"testing"
// 	"time"

// 	"github.com/issfriends/isspay/pkg/factory"
// 	"github.com/shopspring/decimal"

// 	"github.com/issfriends/isspay/internal/app/model"
// 	"github.com/issfriends/isspay/internal/app/query"

// 	"github.com/stretchr/testify/suite"
// )

// type OrderSvcSuite struct {
// 	suite.Suite
// 	*svcSuite
// }

// func TestOrderSvc(t *testing.T) {
// 	suite.Run(t, new(OrderSvcSuite))
// }

// func (s *OrderSvcSuite) SetupSuite() {
// 	s.svcSuite = &svcSuite{}
// 	s.Require().NoError(s.Start())
// }

// func (su *OrderSvcSuite) SetupTest() {
// 	err := su.TruncateTables("orders", "ordered_products", "products", "wallets", "accounts")
// 	su.Require().NoError(err)
// }

// func (s *OrderSvcSuite) TestCreateOrder() {
// 	products := factory.Product.Price(100).MustInsertN(20).([]*model.Product)
// 	productsMap := make(map[int64]*model.Product)
// 	for i := range products {
// 		productsMap[products[i].ID] = products[i]
// 	}

// 	tcs := []struct {
// 		name  string
// 		happy bool
// 		setup func() (*model.Wallet, []*model.OrderedProduct)
// 	}{
// 		{
// 			"happy_case", true,
// 			func() (*model.Wallet, []*model.OrderedProduct) {
// 				wallet := factory.Wallet.Amount(0).BelongAccount().MustInsert().(*model.Wallet)
// 				return wallet, []*model.OrderedProduct{
// 					{ProductID: products[0].ID, Quantity: 1, Price: products[0].Price},
// 				}
// 			},
// 		},
// 		{
// 			"fail_case_invalid_quantity", false,
// 			func() (*model.Wallet, []*model.OrderedProduct) {
// 				wallet := factory.Wallet.Amount(0).BelongAccount().MustInsert().(*model.Wallet)
// 				return wallet, []*model.OrderedProduct{
// 					{ProductID: products[1].ID, Quantity: products[1].Quantity + 1, Price: products[1].Price},
// 				}
// 			},
// 		},
// 		{
// 			"faile_case_invalid_wallet", false,
// 			func() (*model.Wallet, []*model.OrderedProduct) {
// 				ti := time.Now().Add(-50 * 24 * time.Hour)
// 				wallet := factory.Wallet.Amount(-100).LastPaiedAt(ti).
// 					BelongAccount().MustInsert().(*model.Wallet)
// 				return wallet, []*model.OrderedProduct{
// 					{ProductID: products[1].ID, Quantity: products[1].Quantity + 1, Price: products[1].Price},
// 				}
// 			},
// 		},
// 	}

// 	for _, tc := range tcs {
// 		err := s.TruncateTables("orders", "ordered_products", "wallets", "accounts")
// 		s.Require().NoError(err)
// 		s.Run(tc.name, func() {
// 			w, ops := tc.setup()

// 			beforeQuantity := make(map[int64]uint64)
// 			productIDs := make([]int64, 0, len(ops))
// 			opMap := make(map[int64]*model.OrderedProduct)
// 			cost := decimal.Zero
// 			for _, op := range ops {
// 				cost = cost.Add(op.GetCost())
// 				beforeQuantity[op.ProductID] = productsMap[op.ProductID].Quantity
// 				productIDs = append(productIDs, op.ProductID)
// 				opMap[op.ProductID] = op
// 			}

// 			q := &query.GetWalletQuery{MessengerID: w.Owner.MessengerID.String}
// 			o := &model.Order{OrderedProducts: ops}
// 			balance, err := s.Svc.Ordering.CreateOrder(s.Ctx, q, o)
// 			if tc.happy {
// 				s.Require().NoError(err)
// 				s.Assert().True(w.Amount.Sub(cost).Equal(balance))
// 			} else {
// 				s.Require().Error(err)
// 			}

// 			listProducts := &query.ListProductsQuery{
// 				IDs: productIDs,
// 			}
// 			_, err = s.Database.Inventory().ListProducts(s.Ctx, listProducts)
// 			s.Require().NoError(err)

// 			for _, p := range listProducts.Data {
// 				beforeQ := beforeQuantity[p.ID]
// 				orderQ := opMap[p.ID].Quantity
// 				if tc.happy {
// 					s.Assert().Equal(beforeQ-orderQ, p.Quantity, "product(%d) quantity not correct", p.ID)
// 				}
// 			}

// 		})
// 	}
// }

// func (s *OrderSvcSuite) TestCancelOrder() {
// 	products := factory.Product.Price(100).MustInsertN(20).([]*model.Product)
// 	productsMap := make(map[int64]*model.Product)
// 	for i := range products {
// 		productsMap[products[i].ID] = products[i]
// 	}

// 	tcs := []struct {
// 		name  string
// 		happy bool
// 		setup func() (*model.Wallet, *model.Order)
// 	}{
// 		{
// 			"happy_case", true,
// 			func() (*model.Wallet, *model.Order) {
// 				p, _ := products[0].Price.Float64()
// 				w := factory.Wallet.Amount(-p * 2).BelongAccount().MustInsert().(*model.Wallet)
// 				o := factory.Order.WalletId(w.ID).Amount(p * 2).MustInsert().(*model.Order)
// 				op := factory.OrderedProducts.OrderID(o.ID).ProductID(products[0].ID).PriceQuantity(p, 2).MustInsert().(*model.OrderedProduct)
// 				o.OrderedProducts = []*model.OrderedProduct{op}
// 				return w, o
// 			},
// 		},
// 	}

// 	for _, tc := range tcs {
// 		s.Run(tc.name, func() {
// 			w, o := tc.setup()
// 			wq := &query.GetWalletQuery{MessengerID: w.Owner.MessengerID.String}
// 			balance, err := s.Svc.Ordering.CancelOrder(s.Ctx, wq, o.UID)

// 			if tc.happy {
// 				s.Require().NoError(err)
// 				s.True(w.Amount.Equal(balance.Sub(o.Amount)))
// 			}

// 			quantity := make(map[int64]uint64)
// 			productIDs := make([]int64, 0, 1)
// 			for _, op := range o.OrderedProducts {
// 				productIDs = append(productIDs, op.ProductID)
// 				quantity[op.ProductID] = op.Quantity
// 			}
// 			listProducts := &query.ListProductsQuery{
// 				IDs: productIDs,
// 			}
// 			_, err = s.Database.Inventory().ListProducts(s.Ctx, listProducts)
// 			s.Require().NoError(err)

// 			if tc.happy {
// 				for _, prod := range listProducts.Data {
// 					beforeQ := productsMap[prod.ID].Quantity
// 					expect := beforeQ + quantity[prod.ID]
// 					s.Equal(int(expect), int(prod.Quantity), "product(%d) quantity not equal", prod.ID)
// 				}
// 			}
// 		})
// 	}
// }
