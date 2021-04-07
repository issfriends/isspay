package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
)

type InventoryDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	// Product
	GetProduct(ctx context.Context, q *query.GetProductQuery) error
	ListProducts(ctx context.Context, q *query.ListProductsQuery) (total int64, err error)
	BatchCreateProducts(ctx context.Context, products []*model.Product) error
	UpdateProduct(ctx context.Context, q *query.GetProductQuery, updateData *model.Product) error
	DeleteProduct(ctx context.Context, q *query.GetProductQuery) error

	// Report
}

type InventoryServicer interface {
	ListProducts(ctx context.Context, q *query.ListProductsQuery) (total int64, err error)
	BatchCreateProducts(ctx context.Context, products []*model.Product) error
	UpdateProduct(ctx context.Context, q *query.GetProductQuery, updateData *model.Product) error
	DeleteProduct(ctx context.Context, q *query.GetProductQuery) error
}

type InventorySvc struct {
	InventoryDatabaser
}
