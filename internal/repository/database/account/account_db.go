package account

import (
	"context"

	"github.com/issfriends/isspay/internal/app/account"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/vx416/gox/dbprovider"
)

var _ account.Databaser = (*AccountDB)(nil)

type AccountDB struct {
	dbprovider.GormProvider
}

func (d *AccountDB) CreateAccount(ctx context.Context, account *model.Account) error {
	db := d.GetDB(ctx)

	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (d *AccountDB) GetAccount(ctx context.Context, q *query.GetAccountQuery) error {
	data := &model.Account{}
	db := d.GetDB(ctx)

	err := db.Preload("Wallet").Scopes(GetAccountScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}
