package app

import (
	"github.com/issfriends/isspay/internal/app/account"
	"github.com/issfriends/isspay/internal/app/inventory"
	"github.com/issfriends/isspay/internal/app/ordering"
	"github.com/issfriends/isspay/internal/repository/database"
)

func New(db *database.Database) *Service {
	return &Service{
		Account:   account.New(db.Account()),
		Inventory: inventory.New(db.Inventory()),
		Ordering:  ordering.New(db.Ordering()),
	}
}

type Service struct {
	Account   account.Servicer
	Inventory inventory.Servicer
	Ordering  ordering.Servicer
}
