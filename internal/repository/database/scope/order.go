package scope

import (
	"github.com/issfriends/isspay/internal/app/query"
	"gorm.io/gorm"
)

func GetOrder(q *query.GetOrderQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q == nil {
			return db
		}

		if q.ID != 0 {
			db = db.Where("id = ?", q.ID)
		}

		if q.UID != "" {
			db = db.Where("uid = ?", q.UID)
		}

		if q.WalletID != 0 {
			db = db.Where("wallet_id = ?", q.WalletID)
		}

		if q.Status != 0 {
			db = db.Where("status = ?", q.Status)
		}
		return db
	}
}
