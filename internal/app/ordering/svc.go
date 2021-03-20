package ordering

type Servicer interface {
	OrderServicer
}

func New(orderDB OrderDatabaser, accountDB AccountDatabaser) Servicer {
	return &orderingSvc{
		orderDB:   orderDB,
		accountDB: accountDB,
	}
}

type orderingSvc struct {
	orderDB   OrderDatabaser
	accountDB AccountDatabaser
}
