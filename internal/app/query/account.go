package query

import (
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/shopspring/decimal"
)

type GetAccountQuery struct {
	Email       string
	ID          int64
	MessengerID string
	Data        *model.Account

	HasWallet bool
}

type GetWalletQuery struct {
	ID          int64
	AccountID   int64
	MessengerID string

	HasOwner bool
	Data     *model.Wallet
	Lock     LockType
}

type MakePaymentQuery struct {
	AccountID   int64
	MessengerID string
	Amount      decimal.Decimal

	Data *model.Wallet
}
