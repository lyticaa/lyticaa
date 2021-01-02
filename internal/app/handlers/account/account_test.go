package account

import (
	. "gopkg.in/check.v1"
)

type accountSuite struct{}

var _ = Suite(&accountSuite{})

func (s *accountSuite) SetUpSuite(c *C)    {}
func (s *accountSuite) TestAccount(c *C)   {}
func (s *accountSuite) TearDownSuite(c *C) {}
