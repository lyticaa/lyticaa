package types

import (
	"testing"

	. "gopkg.in/check.v1"
)

var (
	valid = []string{
		"text/csv",
		"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	}
	invalid = []string{
		"application/vnd.ms-excel",
		"application/x-msexcel",
		"application/x-ms-excel",
		"application/x-excel",
		"application/x-dos_ms_excel",
		"application/xls",
		"application/x-xls",
	}
)

type typesSuite struct{}

var _ = Suite(&typesSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *typesSuite) SetUpSuite(c *C) {}

func (s *typesSuite) TestMime(c *C) {
	for _, mimeType := range valid {
		ok := ValidMime(mimeType)
		c.Assert(ok, Equals, true)
	}

	for _, mimeType := range invalid {
		ok := ValidMime(mimeType)
		c.Assert(ok, Equals, false)
	}

	ok := IsCSV(valid[0])
	c.Assert(ok, Equals, true)

	ok = IsXLSX(valid[1])
	c.Assert(ok, Equals, true)
}

func (s *typesSuite) TestSignature(c *C) {
	for _, row := range Ignore {
		ok := ShouldIgnore(row)
		c.Assert(ok, Equals, true)
	}
}

func (s *typesSuite) TearDownSuite(c *C) {}
