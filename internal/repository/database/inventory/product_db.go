package inventory

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/repository/database/scope"
)

func (d *InventoryDB) ListProducts(ctx context.Context, q *query.ListProductsQuery) (int64, error) {
	var (
		products = make([]*model.Product, 0, 1)
		total    int64
	)

	db := d.GetDB(ctx).Table(model.Product{}.TableName())
	db = db.Scopes(ListProductsScope(q))
	if err := db.Count(&total).Error; err != nil {
		return 0, err
	}

	err := db.Scopes(scope.Pagination(q.Pagination), scope.Sort(q.Sort)).Find(&products).Error
	if err != nil {
		return 0, err
	}

	q.Data = products
	return total, nil
}

func (d *InventoryDB) BatchCreateProducts(ctx context.Context, products []*model.Product) error {
	db := d.GetDB(ctx)

	err := db.CreateInBatches(products, len(products)).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *InventoryDB) CreateProduct(ctx context.Context, product *model.Product) error {
	db := d.GetDB(ctx)

	if err := db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (d *InventoryDB) UpdateProduct(ctx context.Context, q *query.GetProductQuery, updateData *model.Product) error {
	db := d.GetDB(ctx)

	err := db.Table(updateData.TableName()).Scopes(GetProductScope(q)).Updates(updateData).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *InventoryDB) GetProduct(ctx context.Context, q *query.GetProductQuery) error {
	db := d.GetDB(ctx)
	q.Data = &model.Product{}

	err := db.Table(model.Product{}.TableName()).Scopes(GetProductScope(q)).First(q.Data).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *InventoryDB) DeleteProduct(ctx context.Context, q *query.GetProductQuery) error {
	db := d.GetDB(ctx)

	err := db.Unscoped().Scopes(GetProductScope(q)).Delete(model.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}
