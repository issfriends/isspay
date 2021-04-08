package app

import (
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/issfriends/isspay/internal/repository/cache"
	"github.com/issfriends/isspay/internal/repository/database"
)

func New(db *database.Database, cache *cache.Cache) *App {
	return &App{
		Account:   service.NewAccount(db),
		Auth:      service.NewAuth(db, cache),
		Inventory: service.NewInventory(db),
		Order:     service.NewOrder(db),
	}
}

type App struct {
	Account   service.AccountServicer
	Inventory service.InventoryServicer
	Auth      service.AuthServicer
	Order     service.OrderServicer
}
