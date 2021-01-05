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
	s.a.Start(true)
}

func (s *appSuite) TestApp(c *C) {
	c.Assert(s.a.Monitoring.Logger, NotNil)
	c.Assert(s.a.Monitoring.NewRelic, NotNil)
	c.Assert(s.a.HTTP.Server, NotNil)
	c.Assert(s.a.HTTP.Router, NotNil)
	c.Assert(s.a.HTTP.Client, NotNil)
	c.Assert(s.a.Database.Redis, NotNil)
	c.Assert(s.a.Database.PG, NotNil)
}

func (s *appSuite) TearDownSuite(c *C) {
	s.a.Stop()
}
