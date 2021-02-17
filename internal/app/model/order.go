package model

import (
	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID           int64                 `gorm:"column:id"`
	UID          string                `gorm:"column:uid"`
	BeforeAmount decimal.Decimal       `gorm:"column:beforeAmount"`
	AfterAmount  decimal.Decimal       `gorm:"column:afterAmount"`
	Type         value.TransactionType `gorm:"column:type"`
}
