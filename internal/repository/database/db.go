package database

import (
	"github.com/issfriends/isspay/internal/app/service"
	"github.com/vx416/gox/dbprovider"
)

var (
	_ service.AccountDatabaser = (*Database)(nil)
	_ service.OrderDatabaser   = (*Database)(nil)
	_ service.AuthDatabaser    = (*Database)(nil)
)

type Database struct {
	dbprovider.GormProvider
	*AccountDao
	*WalletDao
	*ProductDao
	*OrderDao
}

func New(gormDB dbprovider.GormProvider) *Database {
	adapter := &DBAdapter{GormProvider: gormDB}
	return &Database{
		GormProvider: gormDB,
		AccountDao:   &AccountDao{DBAdapter: adapter},
		WalletDao:    &WalletDao{DBAdapter: adapter},
		ProductDao:   &ProductDao{DBAdapter: adapter},
		OrderDao:     &OrderDao{DBAdapter: adapter},
	}
}

type DBAdapter struct {
	dbprovider.GormProvider
}
