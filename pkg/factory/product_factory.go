package factory

import (
	"fmt"

	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gogo-factory/attr"
	"github.com/vx416/gogo-factory/genutil"
)

type ProductFactory struct {
	*gofactory.Factory
}

func (f *ProductFactory) Name(name string) *ProductFactory {
	return &ProductFactory{
		f.Attrs(
			attr.Str("Name", genutil.FixStr(name)),
		),
	}
}

func (f *ProductFactory) Category(c value.ProductCategory) *ProductFactory {
	return &ProductFactory{
		f.Attrs(
			attr.Int("Category", genutil.FixInt(int(c))),
		),
	}
}

func (f *ProductFactory) Price(price float64) *ProductFactory {
	return &ProductFactory{
		f.Attrs(
			attr.Float("Price", genutil.FixFloat(price)),
		),
	}
}

func (f *ProductFactory) Quantity(quantity uint) *ProductFactory {
	return &ProductFactory{
		f.Attrs(
			attr.Uint("Quantity", genutil.FixUint(quantity)),
		),
	}
}

var Product = &ProductFactory{gofactory.New(
	&model.Product{},
	attr.Uint("ID", genutil.SeqUint(1, 1)),
	attr.Str("UID", genutil.RandUUID()),
	attr.Str("Name", genutil.RandName(3)).Process(func(a attr.Attributer) error {
		p := a.GetObject().(*model.Product)
		if p.ID == 0 {
			return nil
		}
		return a.SetVal(fmt.Sprintf("product_%d", p.ID))
	}),
	attr.Uint("Quantity", genutil.RandUint(1, 100)),
	attr.Float("Price", genutil.RandFloat(100, 1000)),
	attr.Float("Cost", genutil.RandFloat(100, 1000)).Process(func(a attr.Attributer) error {
		p := a.GetObject().(*model.Product)
		val := a.GetVal().(float64)
		price, _ := p.Price.Float64()
		if val > price {
			if price-1 < 0 {
				return a.SetVal(price)
			}
			return a.SetVal(price - 1)
		}
		return nil
	}),
	attr.Int("Category", genutil.RandInt(1, 2)),
	attr.Str("ImageURL", genutil.FixStr("https://i.imgur.com/Lnc1bJx.png")),
).Table("products")}
