package database

import (
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/vx416/gox/dbprovider"
)

var (
	_ service.AccountDatabaser = (*Database)(nil)
	_ service.WalletDatabaser  = (*Database)(nil)
	_ service.ProductDatabaser = (*Database)(nil)
)

type Database struct {
	dbprovider.GormProvider
	*AccountDB
	*WalletDB
	*ProductDB
}

func New(gormDB dbprovider.GormProvider) *Database {
	adapter := &DBAdapter{GormProvider: gormDB}
	return &Database{
		GormProvider: gormDB,
		AccountDB:    &AccountDB{DBAdapter: adapter},
		WalletDB:     &WalletDB{DBAdapter: adapter},
		ProductDB:    &ProductDB{DBAdapter: adapter},
	}
}

type DBAdapter struct {
	dbprovider.GormProvider
}
