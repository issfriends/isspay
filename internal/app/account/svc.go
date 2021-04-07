package account

type Servicer interface {
	AccountServicer
	WalletServicer
}

func New(
	accountDB IdentityDatabaser,
	walletDB WalletDatabaser,
) Servicer {
	return service{
		accountDB: accountDB,
		walletDB:  walletDB,
	}
}

type service struct {
	accountDB IdentityDatabaser
	walletDB  WalletDatabaser
}
