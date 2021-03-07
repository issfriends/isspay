package inventory

import (
	"github.com/issfriends/isspay/internal/app/query"
	"gorm.io/gorm"
)

func ListProductsScope(q *query.ListProductsQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q == nil {
			return db
		}

		if q.Category != 0 {
			db = db.Where("category = ?", q.Category)
		}

		if q.NameLike != "" {
			db = db.Where("name LIKE ?", "%"+q.NameLike+"%")
		}

		if !q.PriceGte.IsZero() {
			db = db.Where("price >= ?", q.PriceGte)
		}

		if !q.PriceLte.IsZero() {
			db = db.Where("price <= ?", q.PriceGte)
		}

		if q.QuantityGte > 0 {
			db = db.Where("quantity >= ?", q.QuantityGte)
		}

		if q.QuantityLte > 0 {
			db = db.Where("quantity <= ?", q.QuantityLte)
		}

		if len(q.Names) > 0 {
			db = db.Where("name IN (?)", q.Names)
		}

		if len(q.IDs) > 0 {
			db = db.Where("id IN (?)", q.IDs)
		}
		return db
	}
}

func GetProductScope(q *query.GetProductQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q == nil {
			return db
		}

		if q.ID != 0 {
			db = db.Where("id = ?", q.ID)
		}

		if q.Name != "" {
			db = db.Where("name = ?", q.Name)
		}

		if q.UID != "" {
			db = db.Where("uid = ?", q.UID)
		}

		return db
	}
}
