package account

import (
	"github.com/issfriends/isspay/internal/app/account"
	"gorm.io/gorm"
)

func GetAccountScope(q *account.GetAccountQuery) func(db *gorm.DB) *gorm.DB {
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
