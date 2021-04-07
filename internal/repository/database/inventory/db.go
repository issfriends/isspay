package inventory

import (
	"github.com/issfriends/isspay/internal/app/inventory"
	"github.com/issfriends/isspay/internal/app/ordering"
	"github.com/vx416/gox/dbprovider"
)

var (
	_ inventory.Databaser         = (*InventoryDB)(nil)
	_ ordering.InventoryDatabaser = (*InventoryDB)(nil)
)

type InventoryDB struct {
	dbprovider.GormProvider
}
