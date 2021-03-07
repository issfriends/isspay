package model

import (
	"database/sql"
	"time"

	"github.com/issfriends/isspay/internal/app/model/value"
	"github.com/shopspring/decimal"
)

type Account struct {
	ID          int64            `gorm:"column:id" json:"id"`
	UID         string           `gorm:"column:uid" json:"uid"`
	Email       string           `gorm:"column:email" json:"email"`
	UserName    string           `gorm:"column:username" json:"username"`
	NickName    string           `gorm:"column:nickname" json:"nickname"`
	MessengerID sql.NullString   `gorm:"column:messenger_id" json:"messengerID"`
	Membership  value.Membership `gorm:"column:membership" json:"membership"`
	Role        value.Role       `gorm:"column:role" json:"role"`
	CreatedAt   time.Time        `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt   time.Time        `gorm:"column:updated_at" json:"updatedAt"`
	Wallet      *Wallet          `gorm:"foreignKey:OwnerID;references:ID"`
}

func (Account) TableName() string {
	return "accounts"
}

type Wallet struct {
	ID        int64           `gorm:"column:id"`
	UID       string          `gorm:"column:uid"`
	Amount    decimal.Decimal `gorm:"column:amount"`
	OwnerID   int64           `gorm:"column:owner_id"`
	CreatedAt time.Time       `gorm:"column:created_at" json:"createdAt"`
	UpdatedAt time.Time       `gorm:"column:updated_at" json:"updatedAt"`
	Owner     *Account        `gorm:"foreignKey:ID;references:OwnerID"`
}

func (Wallet) TableName() string {
	return "wallets"
}
