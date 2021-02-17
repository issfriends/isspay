package inventory

type InventoryDatabaser interface {
	ProductDatabaser
}

type Servicer interface {
	ProductServicer
}

func New(db InventoryDatabaser) Servicer {
	return &inventorySvc{
		db: db,
	}
}

type inventorySvc struct {
	db InventoryDatabaser
}
