package helpers

import (
	"testing"

	. "gopkg.in/check.v1"
)

type helpersSuite struct{}

var _ = Suite(&helpersSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *helpersSuite) SetUpSuite(c *C)     {}
func (s *helpersSuite) TestDataTables(c *C) {}
func (s *helpersSuite) TestDateRange(c *C)  {}
func (s *helpersSuite) TestEnv(c *C)        {}
func (s *helpersSuite) TestErrors(c *C)     {}
func (s *helpersSuite) TestFilters(c *C)    {}
func (s *helpersSuite) TestFlash(c *C)      {}
func (s *helpersSuite) TestForm(c *C)       {}
func (s *helpersSuite) TestMath(c *C)       {}
func (s *helpersSuite) TestNav(c *C)        {}
func (s *helpersSuite) TestRoutes(c *C)     {}
func (s *helpersSuite) TestSession(c *C)    {}
func (s *helpersSuite) TestTemplates(c *C)  {}
func (s *helpersSuite) TestView(c *C)       {}
func (s *helpersSuite) TearDownSuite(c *C)  {}
