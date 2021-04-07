package ordering

type Servicer interface {
	OrderServicer
}

func New(
	orderDB OrderDatabaser,
	accountDB AccountDatabaser,
	inventoryDB InventoryDatabaser,
) Servicer {
	return &orderingSvc{
		orderDB:     orderDB,
		accountDB:   accountDB,
		inventoryDB: inventoryDB,
	}
}

type orderingSvc struct {
	orderDB     OrderDatabaser
	accountDB   AccountDatabaser
	inventoryDB InventoryDatabaser
}
