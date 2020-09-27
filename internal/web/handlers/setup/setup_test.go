package setup

import (
	. "gopkg.in/check.v1"
)

type setupSuite struct{}

var _ = Suite(&setupSuite{})

func (s *setupSuite) SetUpSuite(c *C)    {}
func (s *setupSuite) TestSetup(c *C)     {}
func (s *setupSuite) TearDownSuite(c *C) {}
