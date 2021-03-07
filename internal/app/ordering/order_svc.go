package ordering

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model/value"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type OrderDatabaser interface {
	//ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error
	CreateOrder(ctx context.Context, order *model.Order) error
	UpdateOrder(ctx context.Context) error
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
	GetOrder(ctx context.Context, q *query.GetOrderQuery) error
}

type OrderServicer interface {
	CreateOrder(ctx context.Context, order *model.Order) (walletBalance decimal.Decimal, err error)
	CancelOrderByUID(ctx context.Context, orderUID string) (walletBalance decimal.Decimal, err error)
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
}

func (svc orderingSvc) CreateOrder(ctx context.Context, order *model.Order) (walletBalance decimal.Decimal, err error) {
	// 鄭言竹交給你了
	var orderAmount decimal.Decimal

	for _, op := range order.OrderedProducts {
		price := op.Price.Mul(decimal.NewFromInt(op.Quantity))
		orderAmount = orderAmount.Add(price)
		op.OrderID = order.ID
	}

	err = svc.db.CreateOrder(ctx, &model.Order{
		UID:             order.UID,
		WalletID:        order.WalletID,
		Status:          value.Completed,
		Amount:          orderAmount,
		OrderedProducts: order.OrderedProducts,
	})
	if err != nil {
		return decimal.Zero, err
	}

	// TODO 計算餘額...
	return decimal.Zero, nil
}

func (svc orderingSvc) CancelOrderByUID(ctx context.Context, orderUID string) (walletBalance decimal.Decimal, err error) {
	// 鄭言竹交給你了
	return decimal.Zero, nil
}

func (svc orderingSvc) ListOrders(ctx context.Context, q *query.ListOrdersQuery) error {
	// 鄭言竹交給你了
	return nil
}

func (svc orderingSvc) GetOrder(ctx context.Context, q *query.GetOrderQuery) error {
	// 鄭言竹交給你了
	if q == nil {
		q = &query.GetOrderQuery{}
	}
	err := svc.db.GetOrder(ctx, q)
	if err != nil {
		return err
	}
	return nil
}
