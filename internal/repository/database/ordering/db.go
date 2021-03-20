package ordering

import (
	"github.com/issfriends/isspay/internal/app/ordering"
	"github.com/vx416/gox/dbprovider"
)

var _ ordering.OrderDatabaser = (*OrderingDB)(nil)

type OrderingDB struct {
	dbprovider.GormProvider
}
