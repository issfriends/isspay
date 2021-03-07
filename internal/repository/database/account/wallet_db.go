package account

import (
	"context"
	"fmt"

	"gorm.io/gorm"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
)

func (d *AccountDB) GetWallet(ctx context.Context, q *query.GetWalletQuery) error {
	var tx *gorm.DB

	data := &model.Wallet{}
	db := d.GetDB(ctx)
	tx = db.Preload("Owner")

	if q == nil {
		q = &query.GetWalletQuery{}
	}

	if q.MessengerID != "" {
		tx.Joins("JOIN accounts AS owner ON wallets.owner_id = Owner.id").
			Where("Owner.messenger_id = ?", q.MessengerID)
	}

	err := tx.Scopes(GetWalletScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	fmt.Println(data)
	return nil
}
