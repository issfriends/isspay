package model

type Order struct {
	OrderedProducts []*OrderedProduct
	Name            string
}

func (o Order) GetName() string {
	return o.Name
}

func (o *Order) UpdateName(n string) {
	o.Name = n
}

var o1 = Order{}
var o2 = &Order{}

type OrderedProduct struct {
}
