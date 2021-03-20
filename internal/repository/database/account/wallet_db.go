package account

import (
	"context"

	"github.com/shopspring/decimal"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
)

func (d *AccountDB) GetWallet(ctx context.Context, q *query.GetWalletQuery) error {
	data := &model.Wallet{}
	db := d.GetDB(ctx)
	db = db.Preload("Owner")

	if q == nil {
		q = &query.GetWalletQuery{}
	}

	err := db.Scopes(GetWalletScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}

func (d *AccountDB) UpdateWallet(ctx context.Context, q *query.GetWalletQuery, wallet *model.Wallet) error {
	db := d.GetDB(ctx)

	err := db.Scopes(GetWalletScope(q)).Updates(wallet).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *AccountDB) UpdateWalletAmount(ctx context.Context, msgID string, wallet *model.Wallet) (balance decimal.Decimal, err error) {
	var amount decimal.Decimal
	db := d.GetDB(ctx)

	statement := `
		UPDATE wallets AS w SET amount = w.amount - $1 
		FROM accounts AS a
		WHERE w.owner_id = a.id AND a.messenger_id = $2
		RETURNING w.amount
		`

	sqlDB, err := db.DB()
	if err != nil {
		return decimal.Zero, err
	}

	row := sqlDB.QueryRow(statement, wallet.Amount, msgID)
	err = row.Err()
	if err != nil {
		return decimal.Zero, err
	}
	row.Scan(&amount)

	return amount, nil
}
