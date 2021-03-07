package account

type Databaser interface {
	IdentityDatabaser
	WalletDatabaser
}

type Servicer interface {
	AccountServicer
	WalletServicer
}

func New(db Databaser) Servicer {
	return service{
		db: db,
	}
}

type service struct {
	db Databaser
}
