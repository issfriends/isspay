package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/query"
)

type AccountDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
}

type AccountServicer interface {
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
}

func NewAccount(db AccountDatabaser) AccountServicer {
	return &accountSvc{
		AccountDatabaser: db,
	}
}

type accountSvc struct {
	AccountDatabaser
}
