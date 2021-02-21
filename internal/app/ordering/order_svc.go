package ordering

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type OrderDatabaser interface {
}

type OrderServicer interface {
	CreateOrder(ctx context.Context, order *model.Order) (walletBalance decimal.Decimal, err error)
	CancelOrderByUID(ctx context.Context, orderUID string) (walletBalance decimal.Decimal, err error)
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
}

func (svc orderingSvc) CreateOrder(ctx context.Context, order *model.Order) (walletBalance decimal.Decimal, err error) {
	// 鄭言竹交給你了
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
