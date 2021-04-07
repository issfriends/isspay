package model

import (
	"time"

	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

type Product struct {
	ID        uint64                `gorm:"column:id" json:"id"`
	UID       string                `gorm:"column:uid;default:uuid_generate_v4()" json:"uid"`
	Name      string                `gorm:"column:name" json:"name"`
	Price     decimal.Decimal       `gorm:"column:price" json:"price"`
	Cost      decimal.Decimal       `gorm:"column:cost" json:"cost"`
	Quantity  uint64                `gorm:"column:quantity" json:"quantity"`
	ImageURL  string                `gorm:"column:image_url" json:"imageURL"`
	Category  value.ProductCategory `gorm:"column:category" json:"category"`
	CreatedAt time.Time             `gorm:"column:created_at" json:"createdAt"`
}

func (Product) TableName() string {
	return "products"
}
