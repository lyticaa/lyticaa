package forecast

import (
	. "gopkg.in/check.v1"
)

type forecastSuite struct{}

var _ = Suite(&forecastSuite{})

func (s *forecastSuite) SetUpSuite(c *C)    {}
func (s *forecastSuite) TestForecast(c *C)  {}
func (s *forecastSuite) TearDownSuite(c *C) {}
