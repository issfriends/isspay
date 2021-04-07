package query

import (
	"time"

	"github.com/shopspring/decimal"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
)

type GetOrderQuery struct {
	ID       int64
	UID      string
	WalletID int64
	Status   value.OrderStatus

	HasOrderedProducts bool
	Data               *model.Order
}

type ListOrdersQuery struct {
	StatusIn   []value.OrderStatus
	PaiedAtGte time.Time
	PaideAtLte time.Time

	Pagination
	Sort

	Data []*model.Order
}

type CreateGetOrderQuery struct {
	WalletID int64
	Amount   decimal.Decimal
}
