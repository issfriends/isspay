package service

import (
	"fmt"

	"github.com/issfriends/isspay/internal/test/testutil"
	"go.uber.org/fx"
)

type svcSuite struct {
	*testutil.TestInstance
}

func (su *svcSuite) Start() error {
	ti, err := testutil.New()
	if err != nil {
		return err
	}

	su.TestInstance = ti

	err = su.TestInstance.Start(fx.Options(
		su.ProvideSvc(),
	))
	if err != nil {
		return err
	}

	if su.Svc == nil {
		return fmt.Errorf("database is nil")
	}
	return nil
}
