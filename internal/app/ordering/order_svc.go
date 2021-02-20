package ordering

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
)

type OrderServicer interface {
	CreateOrder(ctx context.Context, order *model.Order) error
	CancelOrderByUID(ctx context.Context, orderUID string) error
	ListOrders(ctx context.Context, q *query.ListOrdersQuery) error
}
