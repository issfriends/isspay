package service

import (
	"context"

	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type WalletDatabaser interface {
	ExecuteTx(ctx context.Context, fn func(txCtx context.Context) error) error

	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	UpdateWalletAmount(ctx context.Context, walletID int64, delta decimal.Decimal, isPay bool) (balance decimal.Decimal, err error)
}

type TransactionServicer interface {
	MakePaymentByMessagerID(ctx context.Context, msgID string, amount decimal.Decimal) (walletBalance decimal.Decimal, err error)
}

type TransactionSvc struct {
	walletDB WalletDatabaser
}

func (svc TransactionSvc) MakePaymentByMessagerID(ctx context.Context, msgID string, amount decimal.Decimal) (walletBalance decimal.Decimal, err error) {
	err = svc.walletDB.ExecuteTx(ctx, func(txCtx context.Context) error {
		getWallet := &query.GetWalletQuery{
			MessengerID: msgID,
		}
		err := svc.walletDB.GetWallet(txCtx, getWallet)
		if err != nil {
			return err
		}
		wallet := getWallet.Data

		walletBalance, err = svc.walletDB.UpdateWalletAmount(txCtx, wallet.ID, amount, true)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return decimal.Decimal{}, err
	}

	return walletBalance, nil
}
