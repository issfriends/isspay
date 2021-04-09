package model

import (
	"database/sql"
	"time"

	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

const (
	MaxArrearsDuration = 30 * 24 * time.Hour
)

type Account struct {
	ID          uint64              `gorm:"column:id" json:"id"`
	UID         string              `gorm:"column:uid" json:"uid"`
	Email       string              `gorm:"column:email" json:"email"`
	UserName    string              `gorm:"column:username" json:"username"`
	NickName    string              `gorm:"column:nickname" json:"nickname"`
	MessengerID sql.NullString      `gorm:"column:messenger_id" json:"messengerID"`
	Membership  value.Membership    `gorm:"column:membership" json:"membership"`
	Role        value.Role          `gorm:"column:role" json:"role"`
	Status      value.AccountStatus `gorm:"column:status" json:"status"`
	CreatedAt   time.Time           `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time           `gorm:"column:updated_at" json:"updatedAt"`
	Wallet      *Wallet             `gorm:"foreignKey:OwnerID;references:ID"`
}

func (Account) TableName() string {
	return "accounts"
}

func (model Account) IsIssEmail() bool {
	return false
}

type Wallet struct {
	ID          uint64          `gorm:"column:id"`
	UID         string          `gorm:"column:uid"`
	Amount      decimal.Decimal `gorm:"column:amount"`
	OwnerID     uint64          `gorm:"column:owner_id"`
	CreatedAt   time.Time       `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time       `gorm:"column:updated_at" json:"updatedAt"`
	LastPaiedAt time.Time       `gorm:"column:last_paied_at"`
	Owner       *Account        `gorm:"foreignKey:ID;references:OwnerID"`
}

func (Wallet) TableName() string {
	return "wallets"
}

func (model Wallet) CanPurchase(t time.Time) bool {
	if model.Amount.LessThan(decimal.Zero) {
		diff := t.Sub(model.LastPaiedAt)
		return diff < MaxArrearsDuration
	}

	return true
}
