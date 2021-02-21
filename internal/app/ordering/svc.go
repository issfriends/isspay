package ordering

type Databaser interface {
	OrderDatabaser
}

func New() Servicer {
	return &orderingSvc{}
}

type Servicer interface {
	OrderServicer
}

type orderingSvc struct {
	db Databaser
}
