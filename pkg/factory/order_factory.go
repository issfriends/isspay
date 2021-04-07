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

func (f *OrderFactory) HasProduct(p *model.Product, quantity int64) *OrderFactory {
	price, _ := p.Price.Float64()
	ass := OrderedProducts.ProductID(p.ID).PriceQuantity(price, quantity).ToAssociation().
		ForeignKey("order_id").ForeignField("OrderID").ReferField("id").ReferField("ID")

	return &OrderFactory{
		f.HasMany("OrderedProducts", ass, 2),
	}
}

func (f *OrderFactory) Amount(a float64) *OrderFactory {
	return &OrderFactory{
		f.Attrs(
			attr.Float("Amount", genutil.FixFloat(a)),
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

func (f *OrderedProductsFactory) PriceQuantity(price float64, quantity int64) *OrderedProductsFactory {
	return &OrderedProductsFactory{
		f.Attrs(
			attr.Float("Price", genutil.FixFloat(price)),
			attr.Uint("Quantity", genutil.FixUint(uint(quantity))),
		),
	}
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
	attr.Uint("Quantity", genutil.RandUint(1, 10)),
).Table("ordered_products")}
