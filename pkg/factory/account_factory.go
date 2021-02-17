package factory

import (
	"time"

	"github.com/Pallinder/go-randomdata"
	"github.com/issfriends/isspay/internal/app/model"
	"github.com/issfriends/isspay/internal/app/model/value"
	gofactory "github.com/vx416/gogo-factory"
	"github.com/vx416/gogo-factory/attr"
	"github.com/vx416/gogo-factory/genutil"
)

type AccountFactory struct {
	*gofactory.Factory
}

func (f *AccountFactory) Role(role value.Role) *AccountFactory {
	return &AccountFactory{
		f.Attrs(
			attr.Int("Role", genutil.FixInt(int(role))),
		),
	}
}

var Account = &AccountFactory{gofactory.New(
	&model.Account{},
	attr.Int("ID", genutil.SeqInt(1, 1)),
	attr.Str("UID", genutil.RandUUID()),
	attr.Str("Email", randomdata.Email),
	attr.Str("MessengerID", genutil.RandAlph(15)),
	attr.Str("UserName", genutil.RandName(3)),
	attr.Str("NickName", genutil.RandName(15)),
	attr.Time("CreatedAt", genutil.Now(time.UTC)),
	attr.Time("UpdatedAt", genutil.Now(time.UTC)),
).Table("accounts")}
