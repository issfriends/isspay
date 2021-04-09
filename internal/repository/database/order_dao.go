package database

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/repository/database/scope"
)

type OrderDao struct {
	*DBAdapter
}

func (d *OrderDao) CreateOrder(ctx context.Context, order *model.Order) error {
	db := d.GetDB(ctx)

	if err := db.Create(order).Error; err != nil {
		return err
	}
	return nil
}

func (d *OrderDao) GetOrder(ctx context.Context, q *query.GetOrderQuery) error {
	db := d.GetDB(ctx)
	q.Data = &model.Order{}

	err := db.Preload("OrderedProducts").Scopes(scope.GetOrder(q)).First(q.Data).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *OrderDao) UpdateOrder(ctx context.Context, q *query.GetOrderQuery, updateOrder *model.Order) error {
	db := d.GetDB(ctx)

	err := db.Table(updateOrder.TableName()).Scopes(scope.GetOrder(q)).Updates(updateOrder).Error
	if err != nil {
		return err
	}
	return nil
}
