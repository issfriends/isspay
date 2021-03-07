package factory

import (
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gogo-factory/attr"
	"github.com/vx416/gogo-factory/genutil"
)

type OrderFactory struct {
	*gofactory.Factory
}

func (f *OrderFactory) WalletId(walletID int64) *OrderFactory {
	return &OrderFactory{
		f.Attrs(
			attr.Int("WalletID", genutil.FixInt(int(walletID))),
		),
	}
}

var Order = &OrderFactory{gofactory.New(
	&model.Order{},
	attr.Int("ID", genutil.SeqInt(1, 1)),
	attr.Str("UID", genutil.RandUUID()),
	attr.Int("WalletID", genutil.SeqInt(1, 1)),
	attr.Int("Status", genutil.RandInt(1, 2)),
	attr.Float("Amount", genutil.RandFloat(1, 2)),
	attr.Time("CreatedAt", genutil.Now(time.UTC)),
	attr.Time("UpdatedAt", genutil.Now(time.UTC)),
).Table("orders")}

type OrderedProductsFactory struct {
	*gofactory.Factory
}

func (f *OrderedProductsFactory) ProductID(productID int64) *OrderedProductsFactory {
	return &OrderedProductsFactory{
		f.Attrs(
			attr.Int("ProductID", genutil.FixInt(int(productID))),
		),
	}
}

func (f *OrderedProductsFactory) ProductIDs(productIDs ...int64) *OrderedProductsFactory {
	ids := make([]int, len(productIDs))
	for i, id := range productIDs {
		ids[i] = int(id)
	}

	return &OrderedProductsFactory{
		f.Attrs(
			attr.Int("ProductID", genutil.SeqIntSet(ids...)),
		),
	}
}

func (f *OrderedProductsFactory) OrderID(orderID int64) *OrderedProductsFactory {
	return &OrderedProductsFactory{
		f.Attrs(
			attr.Int("OrderID", genutil.FixInt(int(orderID))),
		),
	}
}

func (f *OrderedProductsFactory) OrderIDs(orderIDs ...int64) *OrderedProductsFactory {
	ids := make([]int, len(orderIDs))
	for i, id := range orderIDs {
		ids[i] = int(id)
	}

	return &OrderedProductsFactory{
		f.Attrs(
			attr.Int("OrderID", genutil.SeqIntSet(ids...)),
		),
	}
}

var OrderedProducts = &OrderedProductsFactory{gofactory.New(
	&model.OrderedProduct{},
	attr.Int("ID", genutil.SeqInt(1, 1)),
	attr.Int("ProductID", genutil.SeqInt(1, 1)),
	attr.Int("OrderID", genutil.SeqInt(1, 1)),
	attr.Float("Price", genutil.RandFloat(20, 50)),
	attr.Int("Quantity", genutil.RandInt(1, 10)),
).Table("ordered_products")}
