package database

import (
	"context"
	"fmt"
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/repository/database/scope"
	"github.com/shopspring/decimal"
	"github.com/vx416/gox/log"
)

type WalletDao struct {
	*DBAdapter
}

func (d *WalletDao) GetWallet(ctx context.Context, q *query.GetWalletQuery) error {
	data := &model.Wallet{}
	db := d.GetDB(ctx)

	if q == nil {
		q = &query.GetWalletQuery{}
	}

	err := db.Scopes(scope.GetWalletScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}

func (d *WalletDao) UpdateWallet(ctx context.Context, q *query.GetWalletQuery, wallet *model.Wallet) error {
	db := d.GetDB(ctx)

	err := db.Scopes(scope.GetWalletScope(q)).Updates(wallet).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *WalletDao) UpdateWalletAmount(ctx context.Context, walletID uint64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error) {
	var (
		amount     decimal.Decimal
		db         = d.GetDB(ctx)
		updateStmt = "amount = w.amount - $1, updated_at = $2"
		whereStmt  = "id = $3"
		args       = []interface{}{delta.Abs(), time.Now().UTC()}
	)

	if delta.IsPositive() {
		updateStmt = "amount = w.amount + $1, updated_at = $2"
	}
	if isPay {
		updateStmt = "amount = w.amount + $1, updated_at = $2, last_paied_at = $3"
		whereStmt = "id = $4"
		args = append(args, time.Now().UTC())
	}
	args = append(args, walletID)

	statement := fmt.Sprintf(`
		UPDATE wallets as w SET %s
		WHERE %s RETURNING w.amount
	`, updateStmt, whereStmt)

	sqlDB, err := db.DB()
	if err != nil {
		return decimal.Zero, err
	}

	row := sqlDB.QueryRow(statement, args...)
	err = row.Err()
	if err != nil {
		log.Ctx(ctx).Errorf(statement+" delta:%s, id:%d", delta.String(), walletID)
		return decimal.Zero, err
	}

	log.Ctx(ctx).Debugf(statement+" delta:%s, id:%d", delta.String(), walletID)
	if err = row.Scan(&amount); err != nil {
		return amount, err
	}

	return amount, nil
}
