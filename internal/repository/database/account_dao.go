package database

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/issfriends/isspay/internal/repository/database/scope"
)

type AccountDao struct {
	*DBAdapter
}

func (d AccountDao) CreateAccount(ctx context.Context, account *model.Account) error {
	db := d.GetDB(ctx)

	if err := db.Create(account).Error; err != nil {
		return err
	}
	return nil
}

func (d AccountDao) GetAccount(ctx context.Context, q *query.GetAccountQuery) error {
	data := &model.Account{}
	db := d.GetDB(ctx)

	err := db.Preload("Wallet").Scopes(scope.GetAccountScope(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}
