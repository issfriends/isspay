package app

import (
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/issfriends/isspay/internal/repository/database"
)

func New(db *database.Database) *App {
	return &App{
		Account:   &service.AccountSvc{AccountDatabaser: db},
		Inventory: &service.InventorySvc{InventoryDatabaser: db},
	}
}

type App struct {
	Account   service.AccountServicer
	Inventory service.InventoryServicer
}
