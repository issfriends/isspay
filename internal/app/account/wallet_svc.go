package account

import (
	"context"

	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type WalletServicer interface {
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	MakePaymentByMessagerID(ctx context.Context, msgID string, amount decimal.Decimal) (walletBalance decimal.Decimal, err error)
}

func (svc service) GetWallet(ctx context.Context, q *query.GetWalletQuery) error {
	// 靠你了鄭言竹
	return nil
}

func (svc service) MakePaymentByMessagerID(ctx context.Context, msgID string, amount decimal.Decimal) (decimal.Decimal, error) {
	// 靠你了鄭言竹
	return decimal.Zero, nil
}
