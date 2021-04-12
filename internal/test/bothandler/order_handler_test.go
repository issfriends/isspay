package bothandler

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestOrderBotHandler(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}

type OrderSuite struct {
	suite.Suite
	*botSuite
}

func (s *OrderSuite) SetupSuite() {
	s.botSuite = &botSuite{}
	s.Require().NoError(s.Start())
}

func (s *OrderSuite) SetupTest() {
	s.Require().NoError(s.TruncateTables(
		"orders", "ordered_products", "products", "wallets", "accounts",
	))
}

func (s *OrderSuite) TearDownSuite() {
	s.Require().NoError(s.Shutdown(false))
}
