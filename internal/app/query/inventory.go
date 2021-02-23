package query

import (
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

type ListProductsQuery struct {
	Category    value.ProductCategory `json:"category" query:"category"`
	QuantityGte uint64                `json:"quantityGte" query:"quantityGte"`
	QuantityLte uint64                `json:"quantityLte" query:"quantityLte"`
	PriceGte    decimal.Decimal       `json:"priceGte" query:"priceGte"`
	PriceLte    decimal.Decimal       `json:"priceLte" query:"priceLte"`
	Names       []string              `json:"names" query:"names"`
	NameLike    string                `json:"nameLike" query:"nameLike"`
	Pagination
	Sort

	Data []*model.Product
}

type GetProductQuery struct {
	ID        int64
	Name      string
	UID       string
	ReadLock  bool
	WriteLock bool

	Data *model.Product
}
