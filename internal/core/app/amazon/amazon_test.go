package amazon

import (
	. "gopkg.in/check.v1"
)

type amazonSuite struct{}

var _ = Suite(&amazonSuite{})

func (s *amazonSuite) SetUpSuite(c *C)             {}
func (s *amazonSuite) TestTotalSales(c *C)         {}
func (s *amazonSuite) TestUnitsSold(c *C)          {}
func (s *amazonSuite) TestAmazonCosts(c *C)        {}
func (s *amazonSuite) TestProductCosts(c *C)       {}
func (s *amazonSuite) TestAdvertisingSpend(c *C)   {}
func (s *amazonSuite) TestRefunds(c *C)            {}
func (s *amazonSuite) TestShippingCredits(c *C)    {}
func (s *amazonSuite) TestPromotionalRebates(c *C) {}
func (s *amazonSuite) TestTotalCosts(c *C)         {}
func (s *amazonSuite) TestGrossMargin(c *C)        {}
func (s *amazonSuite) TestNetMargin(c *C)          {}
func (s *amazonSuite) TearDownSuite(c *C)          {}
