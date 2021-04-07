package database

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/repository/database/scope"
	"gorm.io/gorm"
)

type ProductDao struct {
	*DBAdapter
}

func (d *ProductDao) ListProducts(ctx context.Context, q *query.ListProductsQuery) (int64, error) {
	var (
		products = make([]*model.Product, 0, 1)
		total    int64
	)

	db := d.GetDB(ctx).Table(model.Product{}.TableName())
	db = db.Scopes(scope.ListProductsScope(q))
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

func (d *ProductDao) BatchCreateProducts(ctx context.Context, products []*model.Product) error {
	db := d.GetDB(ctx)

	err := db.CreateInBatches(products, len(products)).Error
	if err != nil {
		return err
	}
	return nil
}

func (d *ProductDao) CreateProduct(ctx context.Context, product *model.Product) error {
	db := d.GetDB(ctx)

	if err := db.Create(product).Error; err != nil {
		return err
	}
	return nil
}

func (d *ProductDao) UpdateProduct(ctx context.Context, q *query.GetProductQuery, updateData *model.Product) error {
	db := d.GetDB(ctx)

	err := db.Table(updateData.TableName()).Scopes(scope.GetProductScope(q)).Updates(updateData).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *ProductDao) GetProduct(ctx context.Context, q *query.GetProductQuery) error {
	db := d.GetDB(ctx)
	q.Data = &model.Product{}

	err := db.Table(model.Product{}.TableName()).Scopes(scope.GetProductScope(q)).First(q.Data).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *ProductDao) DeleteProduct(ctx context.Context, q *query.GetProductQuery) error {
	db := d.GetDB(ctx)

	err := db.Unscoped().Scopes(scope.GetProductScope(q)).Delete(model.Product{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *ProductDao) UpdateProductQuantity(ctx context.Context, productID uint64, delta int64) error {
	var (
		db   = d.GetDB(ctx)
		prod = &model.Product{ID: productID}
	)

	if delta > 0 {
		db = db.Model(prod).Update("quantity", gorm.Expr("quantity + ?", delta))
	} else {
		db = db.Model(prod).Where("quantity >= ?", -delta).
			Update("quantity", gorm.Expr("quantity - ?", -delta))
	}
	err := db.Error
	if err != nil {
		return err
	}
	if db.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}

	return nil
}
