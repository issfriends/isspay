package database

import (
	"context"
	"fmt"
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/issfriends/isspay/internal/repository/database/scope"
	"github.com/shopspring/decimal"
	"github.com/vx416/gox/log"
)

var _ service.AccountDatabaser = (*AccountDB)(nil)

type AccountDB struct {
	*DBAdapter
}

func (d AccountDB) CreateAccount(ctx context.Context, account *model.Account) error {
	db := d.GetDB(ctx)

	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (d AccountDB) GetAccount(ctx context.Context, q *query.GetAccountQuery) error {
	data := &model.Account{}
	db := d.GetDB(ctx)

	err := db.Preload("Wallet").Scopes(scope.GetAccountScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}

func (d *AccountDB) GetWallet(ctx context.Context, q *query.GetWalletQuery) error {
	data := &model.Wallet{}
	db := d.GetDB(ctx)
	db = db.Preload("Owner")

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

func (d *AccountDB) UpdateWallet(ctx context.Context, q *query.GetWalletQuery, wallet *model.Wallet) error {
	db := d.GetDB(ctx)

	err := db.Scopes(scope.GetWalletScope(q)).Updates(wallet).Error
	if err != nil {
		return err
	}

	return nil
}

func (d *AccountDB) UpdateWalletAmount(ctx context.Context, id int64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error) {
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
	args = append(args, id)

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
		log.Ctx(ctx).Debugf(statement+" delta:%s, id:%d", delta.String(), id)
		return decimal.Zero, err
	}

	log.Ctx(ctx).Debugf(statement+" delta:%s, id:%d", delta.String(), id)
	if err = row.Scan(&amount); err != nil {
		return amount, err
	}

	return amount, nil
}
