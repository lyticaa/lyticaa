package admin

import (
	"testing"

	. "gopkg.in/check.v1"
)

type adminSuite struct{}

var _ = Suite(&adminSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *adminSuite) SetUpSuite(c *C)     {}
func (s *adminSuite) TestAdmin(c *C)      {}
func (s *adminSuite) TestPaintUsers(c *C) {}
func (s *adminSuite) TearDownSuite(c *C)  {}
