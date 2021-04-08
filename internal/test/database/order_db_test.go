package database

// import (
// 	"testing"

// 	"github.com/issfriends/isspay/internal/app/query"

// 	"github.com/shopspring/decimal"

// 	"github.com/google/uuid"
// 	"github.com/issfriends/isspay/internal/app/model"
// 	"github.com/issfriends/isspay/internal/app/model/value"
// 	"github.com/issfriends/isspay/internal/repository/database/ordering"
// 	"github.com/issfriends/isspay/pkg/factory"
// 	"github.com/stretchr/testify/suite"
// )

// type OrderingSuite struct {
// 	*dbSuite
// 	suite.Suite
// 	orderDB *ordering.OrderingDB
// }

// func TestOrderingDB(t *testing.T) {
// 	suite.Run(t, new(OrderingSuite))
// }

// func (s *OrderingSuite) SetupSuite() {
// 	s.dbSuite = &dbSuite{}
// 	err := s.dbSuite.Start()
// 	s.Require().NoError(err)
// 	s.orderDB = s.Database.Ordering()
// }

// func (s *OrderingSuite) SetupTest() {
// 	err := s.TruncateTables("products", "accounts", "wallets", "orders", "ordered_products")
// 	s.Require().NoError(err)
// }

// func (s *OrderingSuite) TearDownSuite() {
// 	err := s.Shutdown(false)
// 	s.Require().NoError(err)
// }

// //	TODO 寫更多測試case
// func (s *OrderingSuite) TestGetOrder() {
// 	factory.Account.MustInsert()
// 	wallet := factory.Wallet.MustInsert().(*model.Wallet)
// 	products := factory.Product.MustInsertN(3).([]*model.Product)
// 	order := factory.Order.WalletId(wallet.ID).MustInsert().(*model.Order)
// 	factory.OrderedProducts.OrderID(order.ID).ProductIDs(products[0].ID, products[1].ID, products[2].ID).MustInsertN(3)

// 	wantOrder := &model.Order{
// 		WalletID: wallet.ID,
// 		OrderedProducts: []*model.OrderedProduct{
// 			{
// 				OrderID:   order.ID,
// 				ProductID: products[0].ID,
// 			},
// 			{
// 				OrderID:   order.ID,
// 				ProductID: products[1].ID,
// 			},
// 			{
// 				OrderID:   order.ID,
// 				ProductID: products[2].ID,
// 			},
// 		},
// 	}

// 	tests := []struct {
// 		name      string
// 		query     *query.GetOrderQuery
// 		wantOrder *model.Order
// 		wantErr   bool
// 	}{
// 		{
// 			name: "get by orderID",
// 			query: &query.GetOrderQuery{
// 				ID: order.ID,
// 			},
// 			wantOrder: wantOrder,
// 			wantErr:   false,
// 		},
// 		{
// 			name: "get by UID",
// 			query: &query.GetOrderQuery{
// 				UID: order.UID,
// 			},
// 			wantOrder: wantOrder,
// 			wantErr:   false,
// 		},
// 		{
// 			name: "get by UID",
// 			query: &query.GetOrderQuery{
// 				WalletID: order.WalletID,
// 			},
// 			wantOrder: wantOrder,
// 			wantErr:   false,
// 		},
// 	}

// 	for _, t := range tests {
// 		s.Run(t.name, func() {
// 			err := s.orderDB.GetOrder(s.Ctx, t.query)
// 			if t.wantErr {
// 				s.Require().Error(err)
// 			}

// 			data := t.query.Data
// 			s.Require().Equal(data.WalletID, wallet.ID)
// 			for i, op := range data.OrderedProducts {
// 				s.Require().Equal(op.ProductID, products[i].ID)
// 				s.Require().Equal(op.OrderID, order.ID)
// 			}
// 		})
// 	}
// }

// //	TODO 寫更多測試case
// func (s *OrderingSuite) TestCreateOrder() {
// 	factory.Account.MustInsert()
// 	wallet := factory.Wallet.MustInsert().(*model.Wallet)
// 	products := factory.Product.MustInsertN(2).([]*model.Product)

// 	tests := []struct {
// 		name      string
// 		wantOrder *model.Order
// 		wantErr   bool
// 	}{
// 		{
// 			name: "insert 1 Order & 2 OrderProduct",
// 			wantOrder: &model.Order{
// 				UID:      uuid.New().String(),
// 				WalletID: wallet.ID,
// 				Status:   value.Completed,
// 				Amount:   products[0].Price.Mul(decimal.NewFromInt(2)).Add(products[1].Price.Mul(decimal.NewFromInt(1))),
// 				OrderedProducts: []*model.OrderedProduct{
// 					{
// 						ProductID: products[0].ID,
// 						Quantity:  2,
// 						Price:     products[0].Price,
// 					},
// 					{
// 						ProductID: products[1].ID,
// 						Quantity:  1,
// 						Price:     products[1].Price,
// 					},
// 				},
// 			},
// 			wantErr: false,
// 		},
// 	}

// 	for _, t := range tests {
// 		s.Run(t.name, func() {
// 			err := s.orderDB.CreateOrder(s.Ctx, t.wantOrder)
// 			if t.wantErr {
// 				s.Require().Error(err)
// 			}
// 			getOrderQ := &query.GetOrderQuery{
// 				ID: t.wantOrder.ID,
// 			}
// 			err = s.orderDB.GetOrder(s.Ctx, getOrderQ)
// 			s.Require().NoError(err)
// 			data := getOrderQ.Data
// 			s.Require().Equal(t.wantOrder.ID, data.ID)
// 			s.Require().Equal(t.wantOrder.UID, data.UID)
// 			s.Require().Equal(t.wantOrder.WalletID, data.WalletID)
// 			s.Require().Equal(t.wantOrder.Status, data.Status)
// 			s.Require().Equal(t.wantOrder.Amount.StringFixed(2), data.Amount.StringFixed(2))
// 			for i, op := range t.wantOrder.OrderedProducts {
// 				s.Require().Equal(op.ProductID, data.OrderedProducts[i].ProductID)
// 				s.Require().Equal(op.Quantity, data.OrderedProducts[i].Quantity)
// 				s.Require().Equal(op.Price.StringFixed(2), data.OrderedProducts[i].Price.StringFixed(2))
// 			}
// 		})
// 	}

// }
