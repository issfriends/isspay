package inventory

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"gorm.io/gorm"
)

func (d *InventoryDB) UpdateProductQuantity(ctx context.Context, productID int64, delta int64) error {
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
