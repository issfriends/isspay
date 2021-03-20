package service

import (
	"testing"

	"github.com/issfriends/isspay/pkg/factory"

	"github.com/issfriends/isspay/internal/app/model"

	"github.com/stretchr/testify/suite"
)

type OrderSvcSuite struct {
	suite.Suite
	*svcSuite
}

func TestOrderSvc(t *testing.T) {
	suite.Run(t, new(OrderSvcSuite))
}

func (s *OrderSvcSuite) SetupSuite() {
	s.svcSuite = &svcSuite{}
	s.Require().NoError(s.Start())
}

func (su *OrderSvcSuite) SetupTest() {
	err := su.TruncateTables("orders")
	su.Require().NoError(err)
}

func (s *OrderSvcSuite) TestCreateOrder() {
	order := factory.Order.MustBuild().(*model.Order)
	msgID := ""
	_, err := s.Svc.Ordering.CreateOrderByMsgID(s.Ctx, msgID, order)
	s.Require().NoError(err)
}
