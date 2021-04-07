package ordering

import (
	"context"

	"github.com/issfriends/isspay/internal/app/query"

	"github.com/issfriends/isspay/internal/app/model"
)

func (d *OrderingDB) CreateOrder(ctx context.Context, order *model.Order) error {
	db := d.GetDB(ctx)

	if err := db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (d *OrderingDB) UpdateOrder(ctx context.Context, q *query.GetOrderQuery, updateOrder *model.Order) error {
	return nil
}

func (d *OrderingDB) GetOrder(ctx context.Context, q *query.GetOrderQuery) error {
	db := d.GetDB(ctx)
	q.Data = &model.Order{}

	err := db.Scopes(GetOrderScope(q)).First(q.Data).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *OrderingDB) ListOrders(ctx context.Context, q *query.ListOrdersQuery) error {
	return nil
}
