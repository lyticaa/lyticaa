package auth

import (
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"gopkg.in/boj/redistore.v1"
	. "gopkg.in/check.v1"
)

type authSuite struct {
	a *Auth
}

var _ = Suite(&authSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *authSuite) SetUpSuite(c *C) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	c.Assert(err, IsNil)

	sessionStore, err := redistore.NewRediStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_KEY")),
	)
	c.Assert(err, IsNil)

	s.a = NewAuth(db, sessionStore, log.With().Logger())
}

func (s *authSuite) TestAuth(c *C) {
	c.Assert(s.a.db, NotNil)
	c.Assert(s.a.sessionStore, NotNil)
	c.Assert(s.a.logger, NotNil)
}

func (s *authSuite) TearDownSuite(c *C) {}
