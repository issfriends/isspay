package scope

import (
	"github.com/issfriends/isspay/internal/app/query"
	"gorm.io/gorm"
)

func GetAccountScope(q *query.GetAccountQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q.MessengerID != "" {
			db = db.Where("messenger_id = ?", q.MessengerID)
		}

		if q.Email != "" {
			db = db.Where("email = ?", q.Email)
		}

		if q.ID != 0 {
			db = db.Where("id = ?", q.ID)
		}

		return db
	}
}

func GetWalletScope(q *query.GetWalletQuery) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if q.AccountID != 0 {
			db = db.Where("wallets.owner_id = ?", q.AccountID)
		}

		if q.ID != 0 {
			db = db.Where("wallets.id = ?", q.ID)
		}

		if q.MessengerID != "" {
			db = db.Joins("JOIN accounts AS owner ON wallets.owner_id = owner.id").
				Where("owner.messenger_id = ?", q.MessengerID)
		}

		return db
	}
}
