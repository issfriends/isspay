package query

import (
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
)

type ListOrdersQuery struct {
	StatusIn   []value.OrderStatus
	PaiedAtGte time.Time
	PaideAtLte time.Time

	Pagination
	Sort

	Data []*model.Order
}
