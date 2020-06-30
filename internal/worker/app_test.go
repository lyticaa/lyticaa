package app

import (
	"testing"

	. "gopkg.in/check.v1"
)

type appSuite struct {
	a *App
}

var _ = Suite(&appSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *appSuite) SetUpSuite(c *C) {
	s.a = NewApp()
}

func (s *appSuite) TestApp(c *C) {
	c.Assert(s.a.Logger, NotNil)
	c.Assert(s.a.NewRelic, NotNil)
	c.Assert(s.a.Db, NotNil)
}

func (s *appSuite) TearDownSuite(c *C) {}
