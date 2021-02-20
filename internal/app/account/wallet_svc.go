package account

import (
	"context"

	"github.com/issfriends/isspay/internal/app/query"
	"github.com/shopspring/decimal"
)

type WalletServicer interface {
	GetWallet(ctx context.Context, q *query.GetWalletQuery) error
	MakePayment(ctx context.Context, walletID int64, amount decimal.Decimal, unpaidOrderIDs []int64) error
}
