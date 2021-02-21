package inventory

type Databaser interface {
	ProductDatabaser
}

type Servicer interface {
	ProductServicer
}

func New(db Databaser) Servicer {
	return &inventorySvc{
		db: db,
	}
}

type inventorySvc struct {
	db Databaser
}
