package httphandler

import (
	"github.com/issfriends/isspay/internal/test/testutil"
	"github.com/stretchr/testify/suite"
)

type httpSuite struct {
	suite.Suite
	*testutil.TestInstance
}

func (s *httpSuite) SetupSuite() {
	var err error
	s.TestInstance, err = testutil.New()
	s.Require().NoError(err)

	opts := s.TestInstance.ProvideRestfulHandler()

	err = s.Start(opts)
	s.Require().NoError(err)

	s.Require().NotNil(s.Serv)
}
