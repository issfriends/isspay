package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
)

type ProductDatabaser interface {
	// GetProduct(ctx context.Context, q *GetProductQuery) error
	ListProducts(ctx context.Context, q *query.ListProductsQuery) (total int64, err error)
	BatchCreateProducts(ctx context.Context, products []*model.Product) error
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error
	UpdateProduct(ctx context.Context, q *query.GetProductQuery, updateData *model.Product) error
	DeleteProduct(ctx context.Context, q *query.GetProductQuery) error
}

type InventoryServicer interface {
}
