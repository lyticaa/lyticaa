package worker

import (
	"testing"

	. "gopkg.in/check.v1"
)

type appSuite struct {
	w *Worker
}

var _ = Suite(&appSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *appSuite) SetUpSuite(c *C) {
	s.w = NewWorker()
}

func (s *appSuite) TestApp(c *C) {
	c.Assert(s.w.Logger, NotNil)
	c.Assert(s.w.NewRelic, NotNil)
	c.Assert(s.w.Db, NotNil)
}

func (s *appSuite) TearDownSuite(c *C) {}
