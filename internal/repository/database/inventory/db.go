package inventory

import (
	"github.com/issfriends/isspay/internal/app/inventory"
	"github.com/vx416/gox/dbprovider"
)

var _ inventory.Databaser = (*InventoryDB)(nil)

type InventoryDB struct {
	dbprovider.GormProvider
}
