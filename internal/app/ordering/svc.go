package ordering

type Databaser interface {
	OrderDatabaser
}

type Servicer interface {
	OrderServicer
}

func New(db Databaser) Servicer {
	return &orderingSvc{
		db: db,
	}
}

type orderingSvc struct {
	db Databaser
}
