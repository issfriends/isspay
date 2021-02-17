package database

import (
	"github.com/issfriends/isspay/internal/repository/database/account"
	"github.com/issfriends/isspay/internal/repository/database/inventory"
	"github.com/vx416/gox/dbprovider"
)

func New(gormDB dbprovider.GormProvider) *Database {
	return &Database{
		gormDB: gormDB,
	}
}

type Database struct {
	gormDB dbprovider.GormProvider
}

func (db *Database) Account() *account.AccountDB {
	return &account.AccountDB{GormProvider: db.gormDB}
}

func (db *Database) Inventory() *inventory.InventoryDB {
	return &inventory.InventoryDB{GormProvider: db.gormDB}
}
