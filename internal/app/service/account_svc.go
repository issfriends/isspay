package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type AccountDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
	UpdateAccount(ctx context.Context, q *query.GetAccountQuery, updateAccount *model.Account) error
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	UpdateWalletAmount(ctx context.Context, walletID uint64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error)
}

type AccountServicer interface {
	GetAccount(ctx context.Context, q *query.GetAccountQuery) error
	UpdateAccount(ctx context.Context, q *query.GetAccountQuery, updateAccount *model.Account) error
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	MakePayment(ctx context.Context, walletID uint64, amount decimal.Decimal) (balance decimal.Decimal, err error)
}

func NewAccount(db AccountDatabaser) AccountServicer {
	return &accountSvc{
		AccountDatabaser: db,
	}
}

type accountSvc struct {
	AccountDatabaser
}

func (svc accountSvc) MakePayment(ctx context.Context, walletID uint64, amount decimal.Decimal) (decimal.Decimal, error) {

	balance, err := svc.UpdateWalletAmount(ctx, walletID, amount, true)
	if err != nil {
		return decimal.Decimal{}, err
	}

	return balance, nil
}
