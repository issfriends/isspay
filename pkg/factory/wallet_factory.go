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

func (f *WalletFactory) ID(id uint64) *WalletFactory {
	return &WalletFactory{
		f.Attrs(
			attr.Uint("ID", genutil.FixUint(uint(id))),
		),
	}
}

func (f *WalletFactory) OwnerID(id uint64) *WalletFactory {
	return &WalletFactory{
		f.Attrs(
			attr.Uint("OwnerID", genutil.FixUint(uint(id))),
		),
	}
}

func (f *WalletFactory) Amount(a float64) *WalletFactory {
	return &WalletFactory{
		f.Attrs(
			attr.Float("Amount", genutil.FixFloat(a)),
		),
	}
}

func (f *WalletFactory) LastPaiedAt(t time.Time) *WalletFactory {
	return &WalletFactory{
		f.Attrs(
			attr.Time("LastPaiedAt", genutil.FixTime(t)),
		),
	}
}

func (f *WalletFactory) BelongAccount() *WalletFactory {
	ass := Account.ToAssociation().ForeignKey("owner_id").ReferColumn("id").
		ForeignField("OwnerID").ReferField("ID").AssociatedField("Wallet")

	return &WalletFactory{f.BelongsTo("Owner", ass)}
}

var Wallet = &WalletFactory{gofactory.New(
	&model.Wallet{},
	attr.Uint("ID", genutil.SeqUint(1, 1)),
	attr.Str("UID", genutil.RandUUID()),
	attr.Float("Amount", genutil.RandFloat(10, 100)),
	attr.Uint("OwnerID", genutil.SeqUint(1, 1)),
	attr.Time("CreatedAt", genutil.Now(time.UTC)),
	attr.Time("UpdatedAt", genutil.Now(time.UTC)),
	attr.Time("LastPaiedAt", genutil.Now(time.UTC)),
).Table("wallets")}
