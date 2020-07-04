package metrics

import (
	. "gopkg.in/check.v1"
)

type metricsSuite struct{}

var _ = Suite(&metricsSuite{})

func (s *metricsSuite) SetUpSuite(c *C)    {}
func (s *metricsSuite) TestMetrics(c *C)   {}
func (s *metricsSuite) TearDownSuite(c *C) {}
