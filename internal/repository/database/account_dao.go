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

	err := db.Scopes(scope.GetAccount(q)).First(data).Error
	if err != nil {
		return err
	}

	q.Data = data
	return nil
}

func (d AccountDao) UpdateAccount(ctx context.Context, q *query.GetAccountQuery, updateAccount *model.Account) error {
	db := d.GetDB(ctx)

	err := db.Table(updateAccount.TableName()).Scopes(scope.GetAccount(q)).Updates(updateAccount).Error
	if err != nil {
		return err
	}

	return nil
}
