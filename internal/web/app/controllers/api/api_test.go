package api

import (
	"testing"

	. "gopkg.in/check.v1"
)

type apiSuite struct {
	a *API
}

var _ = Suite(&apiSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *apiSuite) SetUpSuite(c *C) {
	s.a = NewAPI()
}

func (s *apiSuite) TestHealth(c *C)    {}
func (s *apiSuite) TearDownSuite(c *C) {}
