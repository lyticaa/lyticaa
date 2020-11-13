package expenses

import (
	"testing"

	. "gopkg.in/check.v1"
)

type expensesSuite struct{}

var _ = Suite(&expensesSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *expensesSuite) SetUpSuite(c *C)   {}
func (s *expensesSuite) TestExpenses(c *C) {}

func (s *expensesSuite) TestSendMessage(c *C) {

	err := sendMessage("", "", "")
	c.Assert(err, IsNil)
}
func (s *expensesSuite) TearDownSuite(c *C) {}
