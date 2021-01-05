package amazon

import (
	"testing"

	. "gopkg.in/check.v1"
)

type amazonSuite struct{}

var _ = Suite(&amazonSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *amazonSuite) SetUpSuite(c *C)    {}
func (s *amazonSuite) TearDownSuite(c *C) {}
