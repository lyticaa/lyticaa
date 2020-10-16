package admin

import (
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/lyticaa/lyticaa-app/internal/models"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gopkg.in/boj/redistore.v1"
	. "gopkg.in/check.v1"
	"syreclabs.com/go/faker"
)

type adminSuite struct {
	a    *Admin
	user *models.User
}

var _ = Suite(&adminSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *adminSuite) SetUpSuite(c *C) {
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

	s.a = NewAdmin(db, sessionStore, log.With().Logger())

	user, err := models.CreateUser(
		faker.RandomString(10),
		faker.Internet().Email(),
		faker.Lorem().Word(),
		faker.Avatar().Url("jpg", 200, 300),
		s.a.db,
	)
	c.Assert(err, IsNil)

	user.Admin = true
	err = user.Save(s.a.db)
	c.Assert(err, IsNil)

	s.user = user
}

func (s *adminSuite) TestAdmin(c *C) {
	c.Assert(s.a.db, NotNil)
	c.Assert(s.a.sessionStore, NotNil)
	c.Assert(s.a.logger, NotNil)
}

func (s *adminSuite) TestPaintUsers(c *C) {
	_, err := models.CreateUser(
		faker.RandomString(10),
		faker.Internet().Email(),
		faker.Lorem().Word(),
		faker.Avatar().Url("jpg", 200, 300),
		s.a.db,
	)
	c.Assert(err, IsNil)

	r, err := http.NewRequest("GET", fmt.Sprintf("%v?start=%v&length=%v", faker.Internet().Url(), 0, 10), nil)
	c.Assert(r, NotNil)
	c.Assert(err, IsNil)

	req := mux.SetURLVars(r, map[string]string{"dateRange": "all_time"})
	c.Assert(req, NotNil)

	users := s.a.paintUsers(req)
	c.Assert(assert.Greater(c, len(users.Data), 1), Equals, true)
	c.Assert(assert.Greater(c, users.RecordsTotal, int64(1)), Equals, true)
}

func (s *adminSuite) TearDownSuite(c *C) {
	err := s.user.Delete(s.a.db)
	c.Assert(err, IsNil)
}
