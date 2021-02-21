package database

import (
	"fmt"

	"github.com/issfriends/isspay/internal/test/testutil"
	"go.uber.org/fx"
)

type dbSuite struct {
	*testutil.TestInstance
}

func (su *dbSuite) Start() error {
	ti, err := testutil.New()
	if err != nil {
		return err
	}

	su.TestInstance = ti

	err = su.TestInstance.Start(fx.Options(
		su.ProvideDB(),
	))
	if err != nil {
		return err
	}

	if su.Database == nil {
		return fmt.Errorf("database is nil")
	}
	return nil
}
