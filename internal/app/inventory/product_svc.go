package inventory

import (
	"context"
	"fmt"

	"github.com/vx416/gox/converter"
	"github.com/vx416/gox/resperr"

	goerr "github.com/issfriends/isspay/internal/test/errors"

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

type ProductServicer interface {
	BatchCreateProducts(ctx context.Context, products []*model.Product) error
	UpdateProduct(ctx context.Context, opt *query.GetProductQuery, updateData *model.Product) error
	ListProducts(ctx context.Context, q *query.ListProductsQuery) (total int64, err error)
	DeleteProduct(ctx context.Context, q *query.GetProductQuery) error
}

func (svc inventorySvc) BatchCreateProducts(ctx context.Context, products []*model.Product) error {
	productNames := make([]string, 0, len(products))
	for _, p := range products {
		productNames = append(productNames, p.Name)
	}
	productsQ := query.ListProductsQuery{
		Names: productNames,
	}

	var duplicatedErr error

	err := svc.db.ExecuteTx(ctx, func(txCtx context.Context) error {
		if _, err := svc.db.ListProducts(txCtx, &productsQ); err != nil {
			return err
		}
		if len(productsQ.Data) > 0 {
			errMsg := make(map[string]interface{})
			for _, duplicatedP := range productsQ.Data {
				errMsg[duplicatedP.Name] = fmt.Sprintf("product name(%s) is duplicated", duplicatedP.Name)
			}
			duplicatedErr = resperr.WithDetails(goerr.ErrUnprocessableEntity, errMsg)
			products = svc.diffProducts(products, productsQ.Data)
		}

		if err := svc.db.BatchCreateProducts(txCtx, products); err != nil {
			return err
		}
		return nil
	})

	if err != nil {
		return err
	}
	if duplicatedErr != nil {
		return duplicatedErr
	}

	return nil
}

func (svc inventorySvc) UpdateProduct(ctx context.Context, opt *query.GetProductQuery, updateData *model.Product) error {
	if err := svc.db.UpdateProduct(ctx, opt, updateData); err != nil {
		return err
	}
	return nil
}

func (svc inventorySvc) ListProducts(ctx context.Context, q *query.ListProductsQuery) (int64, error) {
	return svc.db.ListProducts(ctx, q)
}

func (svc inventorySvc) DeleteProduct(ctx context.Context, q *query.GetProductQuery) error {
	if err := svc.db.DeleteProduct(ctx, q); err != nil {
		return err
	}
	return nil
}

func (svc inventorySvc) diffProducts(products []*model.Product, removed []*model.Product) []*model.Product {
	newProudcts := make([]*model.Product, 0, len(products))
	duplicatedP := converter.SliceToMap(removed, "Name").(map[string]*model.Product)

	for _, p := range products {
		if _, ok := duplicatedP[p.Name]; !ok {
			newProudcts = append(newProudcts, p)
		}
	}

	return newProudcts
}
