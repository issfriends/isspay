package account

type AccountDatabaser interface {
	IdentityDatabaser
}

type Servicer interface {
	AccountServicer
}

func New(db AccountDatabaser) Servicer {
	return service{
		db: db,
	}
}

type service struct {
	db AccountDatabaser
}
