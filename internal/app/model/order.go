package model

import (
	"database/sql"
	"time"

	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

type Order struct {
	ID              int64             `gorm:"column:id" json:"id"`
	UID             string            `gorm:"column:uid" json:"uid"`
	WalletID        int64             `gorm:"column:wallet_id" json:"wallet_id"`
	Status          value.OrderStatus `gorm:"column:status" json:"status"`
	Amount          decimal.Decimal   `gorm:"column:amount" json:"amount"`
	CreatedAt       time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"column:updated_at" json:"updated_at"`
	CanceledAt      sql.NullTime      `gorm:"column:canceled_at" json:"canceled_at"`
	OrderedProducts []*OrderedProduct `gorm:"foreignKey:OrderID";references:ID`
}

type OrderedProduct struct {
	ID        int64           `gorm:"column:id" json:"id"`
	OrderID   int64           `gorm:"column:order_id" json:"order_id"`
	ProductID int64           `gorm:"column:product_id" json:"productID"`
	Quantity  int64           `gorm:"column:quantity" json:"quantity"`
	Price     decimal.Decimal `gorm:"column:price" json:"price"`
}
