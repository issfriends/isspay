package model

type Order struct {
	UID             string `gorm:"column:uid" json:"uid"`
	OrderedProducts []*OrderedProduct
}

type OrderedProduct struct {
	ProductID int64 `json:"productID"`
	Quantity  int64 `json:"quantity"`
}
