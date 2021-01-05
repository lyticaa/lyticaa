package ingest

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/newrelic/go-agent"
	. "gopkg.in/check.v1"
)

type ingestSuite struct {
	i *Ingest
}

var _ = Suite(&ingestSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *ingestSuite) SetUpSuite(c *C) {
	config := newrelic.NewConfig(
		os.Getenv("APP_NAME"),
		os.Getenv("NEW_RELIC_LICENSE_KEY"),
	)

	nr, err := newrelic.NewApplication(config)
	c.Assert(err, IsNil)

	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	c.Assert(err, IsNil)

	s.i = NewIngest(nr, db)
}

func (s *ingestSuite) TestIngest(c *C) {
	c.Assert(s.i.Logger, NotNil)
	c.Assert(s.i.NewRelic, NotNil)
	c.Assert(s.i.Db, NotNil)
}

func (s *ingestSuite) TearDownSuite(c *C) {}
