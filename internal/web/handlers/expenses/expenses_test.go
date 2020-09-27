package expenses

import (
	. "gopkg.in/check.v1"
)

type expensesSuite struct{}

var _ = Suite(&expensesSuite{})

func (s *expensesSuite) SetUpSuite(c *C)    {}
func (s *expensesSuite) TestExpenses(c *C)  {}
func (s *expensesSuite) TearDownSuite(c *C) {}
