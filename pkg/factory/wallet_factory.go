package factory

import (
	"time"

	"github.com/issfriends/isspay/internal/app/model"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gogo-factory/attr"
	"github.com/vx416/gogo-factory/genutil"
)

type WalletFactory struct {
	*gofactory.Factory
}

func (f *WalletFactory) ID(id int64) *AccountFactory {
	return &AccountFactory{
		f.Attrs(
			attr.Int("ID", genutil.FixInt(int(id))),
		),
	}
}

func (f *WalletFactory) OwnerID(id int64) *AccountFactory {
	return &AccountFactory{
		f.Attrs(
			attr.Int("OwnerID", genutil.FixInt(int(id))),
		),
	}
}

var Wallet = &WalletFactory{gofactory.New(
	&model.Wallet{},
	attr.Int("ID", genutil.SeqInt(1, 1)),
	attr.Str("UID", genutil.RandUUID()),
	attr.Float("Amount", genutil.RandFloat(10, 100)),
	attr.Int("OwnerID", genutil.SeqInt(1, 1)),
	attr.Time("CreatedAt", genutil.Now(time.UTC)),
	attr.Time("UpdatedAt", genutil.Now(time.UTC)),
).Table("wallets")}
