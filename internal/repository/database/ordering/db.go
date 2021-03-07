package ordering

import (
	"github.com/issfriends/isspay/internal/app/ordering"
	"github.com/vx416/gox/dbprovider"
)

var _ ordering.Databaser = (*OrderingDB)(nil)

type OrderingDB struct {
	dbprovider.GormProvider
}
