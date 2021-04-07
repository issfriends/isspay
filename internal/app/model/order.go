package model

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

func NewOrder(walletID int64, orderedProducts ...*OrderedProduct) *Order {
	amount := decimal.Zero
	for _, op := range orderedProducts {
		amount = amount.Add(op.GetCost())
	}

	return &Order{
		WalletID:        walletID,
		UID:             uuid.New().String(),
		OrderedProducts: orderedProducts,
		Status:          value.Completed,
		Amount:          amount,
	}
}

type Order struct {
	ID              int64             `gorm:"column:id" json:"id"`
	UID             string            `gorm:"column:uid" json:"uid"`
	WalletID        int64             `gorm:"column:wallet_id" json:"wallet_id"`
	Status          value.OrderStatus `gorm:"column:status" json:"status"`
	Amount          decimal.Decimal   `gorm:"column:amount" json:"amount"`
	CreatedAt       time.Time         `gorm:"column:created_at" json:"created_at"`
	UpdatedAt       time.Time         `gorm:"column:updated_at" json:"updated_at"`
	CanceledAt      sql.NullTime      `gorm:"column:canceled_at" json:"canceled_at"`
	OrderedProducts []*OrderedProduct `gorm:"foreignKey:OrderID;references:ID"`
}

func (model Order) TableName() string {
	return "orders"
}

func (model *Order) Setup(walletID int64) {
	amount := decimal.Zero
	for _, op := range model.OrderedProducts {
		amount = amount.Add(op.GetCost())
	}
	model.UID = uuid.New().String()
	model.Amount = amount
	model.Status = value.Completed
	model.WalletID = walletID
}

type OrderedProduct struct {
	ID        uint64          `gorm:"column:id" json:"id"`
	OrderID   uint64          `gorm:"column:order_id" json:"order_id"`
	ProductID uint64          `gorm:"column:product_id" json:"productID"`
	Quantity  uint64          `gorm:"column:quantity" json:"quantity"`
	Price     decimal.Decimal `gorm:"column:price" json:"price"`
}

func (model OrderedProduct) TableName() string {
	return "ordered_products"
}

func (model OrderedProduct) GetCost() decimal.Decimal {
	return model.Price.Mul(decimal.NewFromInt(int64(model.Quantity)))
}
