package helpers

import (
	"database/sql"
	"encoding/gob"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/lyticaa/lyticaa-app/internal/models"
	"github.com/lyticaa/lyticaa-app/internal/web/types"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"gopkg.in/boj/redistore.v1"
	. "gopkg.in/check.v1"
	"syreclabs.com/go/faker"
)

const (
	successMsg = "Success"
	errorMsg   = "Error"
	warningMsg = "Warning"
	infoMsg    = "Info"
)

type validateTest struct {
	Value string `validate:"required,min=10"`
}

type helpersSuite struct {
	db           *sqlx.DB
	sessionStore *redistore.RediStore
}

var _ = Suite(&helpersSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *helpersSuite) SetUpSuite(c *C) {
	db, err := sqlx.Connect("postgres", os.Getenv("DATABASE_URL"))
	c.Assert(err, IsNil)

	s.db = db

	gob.Register(types.Flash{})
	gob.Register(models.User{})

	store, err := redistore.NewRediStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_KEY")))
	c.Assert(store, NotNil)
	c.Assert(err, IsNil)

	s.sessionStore = store
}

func (s *helpersSuite) TestDataTables(c *C) {
	url := fmt.Sprintf("/?draw=%v&start=%v&length=%v&order[0][column]=%v&order[0][dir]=%v", 1, 0, 10, 0, "asc")
	r, err := http.NewRequest(http.MethodGet, url, nil)
	c.Assert(r, NotNil)
	c.Assert(err, IsNil)

	draw := DtDraw(r)
	c.Assert(draw, Equals, int64(1))

	start := DtStart(r)
	c.Assert(start, Equals, int64(0))

	length := DtLength(r)
	c.Assert(length, Equals, int64(10))

	sort := DtSort(r)
	c.Assert(sort, Equals, int64(0))

	dir := DtDir(r)
	c.Assert(dir, Equals, "ASC")
}

func (s *helpersSuite) TestDateRange(c *C) {
	start, end := today()
	date := time.Now()
	c.Assert(start.Day(), Equals, date.Day())
	c.Assert(end.Day(), Equals, date.Day())

	start, end = yesterday()
	date = time.Now().AddDate(0, 0, -1)
	c.Assert(start.Day(), Equals, date.Day())
	c.Assert(end.Day(), Equals, date.Day())

	start, end = lastThirtyDays()
	date = time.Now().AddDate(0, 0, -30)
	c.Assert(start.Day(), Equals, date.Day())
	c.Assert(end.Day(), Equals, time.Now().Day())

	start, end = previousThirtyDays()
	date = time.Now().AddDate(0, 0, -60)
	c.Assert(start.Day(), Equals, date.Day())
	date = time.Now().AddDate(0, 0, -30)
	c.Assert(end.Day(), Equals, date.Day())

	start, end = thisMonth()
	date = time.Now()
	c.Assert(start.Month(), Equals, date.Month())
	c.Assert(end.Month(), Equals, date.Month())

	start, end = lastMonth()
	date = time.Now().AddDate(0, -1, 0)
	c.Assert(start.Month(), Equals, date.Month())
	c.Assert(end.Month(), Equals, time.Now().Month())

	start, end = monthBeforeLast()
	date = time.Now().AddDate(0, -2, 0)
	c.Assert(start.Month(), Equals, date.Month())
	c.Assert(end.Month(), Equals, date.Month())

	start, end = lastThreeMonths()
	date = time.Now().AddDate(0, -3, 0)
	c.Assert(start.Month(), Equals, date.Month())
	c.Assert(end.Month(), Equals, time.Now().Month())

	start, end = previousThreeMonths()
	date = time.Now().AddDate(0, -6, 0)
	c.Assert(start.Month(), Equals, date.Month())
	date = time.Now().AddDate(0, -3, 0)
	c.Assert(end.Month(), Equals, date.Month())

	start, end = lastSixMonths()
	date = time.Now().AddDate(0, -6, 0)
	c.Assert(start.Month(), Equals, date.Month())
	c.Assert(end.Month(), Equals, time.Now().Month())

	start, end = previousSixMonths()
	date = time.Now().AddDate(0, -12, 0)
	c.Assert(start.Month(), Equals, date.Month())
	date = time.Now().AddDate(0, -6, 0)
	c.Assert(end.Month(), Equals, date.Month())

	start, end = thisYear()
	c.Assert(start.Year(), Equals, time.Now().Year())
	c.Assert(end.Year(), Equals, time.Now().Year())

	start, end = lastYear()
	date = time.Now().AddDate(-1, 0, 0)
	c.Assert(start.Year(), Equals, date.Year())
	c.Assert(end.Year(), Equals, time.Now().Year())

	start, end = allTime()
	date = time.Date(1970, 1, 1, 0, 0, 0, 0, time.Now().Location())
	c.Assert(start.Year(), Equals, date.Year())
	c.Assert(end.Year(), Equals, time.Now().Year())
}

func (s *helpersSuite) TestFilters(c *C) {
	url := fmt.Sprintf("/?draw=%v&start=%v&length=%v&order[0][column]=%v&order[0][dir]=%v", 1, 0, 10, 0, "asc")
	r, err := http.NewRequest(http.MethodGet, url, nil)
	c.Assert(r, NotNil)
	c.Assert(err, IsNil)

	filter := BuildFilter(r)
	c.Assert(filter.Start, Equals, int64(0))
	c.Assert(filter.Length, Equals, int64(10))
	c.Assert(filter.Sort, Equals, int64(0))
	c.Assert(filter.Dir, Equals, "ASC")
}

func (s *helpersSuite) TestFlash(c *C) {
	gob.Register(types.Flash{})
	store, err := redistore.NewRediStore(
		10,
		"tcp",
		os.Getenv("REDIS_URL"),
		os.Getenv("REDIS_PASSWORD"),
		[]byte(os.Getenv("SESSION_KEY")))
	c.Assert(store, NotNil)
	c.Assert(err, IsNil)

	r, err := http.NewRequest(http.MethodGet, faker.Internet().Url(), nil)
	c.Assert(r, NotNil)
	c.Assert(err, IsNil)

	session, err := store.Get(r, "auth-session")
	c.Assert(session, NotNil)
	c.Assert(err, IsNil)

	setFlashSuccess(successMsg, session)
	flash := session.Values["Flash"].(types.Flash)
	c.Assert(flash.Success, Equals, successMsg)

	setFlashError(errorMsg, session)
	flash = session.Values["Flash"].(types.Flash)
	c.Assert(flash.Error, Equals, errorMsg)

	setFlashWarning(warningMsg, session)
	flash = session.Values["Flash"].(types.Flash)
	c.Assert(flash.Warning, Equals, warningMsg)

	setFlashInfo(infoMsg, session)
	flash = session.Values["Flash"].(types.Flash)
	c.Assert(flash.Info, Equals, infoMsg)
}

func (s *helpersSuite) TestForm(c *C) {
	logger := log.With().Logger()

	v := validateTest{Value: faker.RandomString(10)}
	ok, _ := ValidateInput(v, &logger)
	c.Assert(ok, Equals, true)

	v = validateTest{Value: ""}
	ok, _ = ValidateInput(v, &logger)
	c.Assert(ok, Equals, false)
}

func (s *helpersSuite) TestSession(c *C) {
	url := faker.Internet().Url()
	r, err := http.NewRequest(http.MethodGet, url, nil)
	c.Assert(r, NotNil)
	c.Assert(err, IsNil)

	session, err := s.sessionStore.Get(r, "auth-session")
	c.Assert(session, NotNil)
	c.Assert(err, IsNil)

	setFlashSuccess(successMsg, session)
	flash := session.Values["Flash"].(types.Flash)
	c.Assert(flash.Success, Equals, successMsg)

	resetFlash(session)
	c.Assert(session.Values["Flash"], IsNil)

	session = GetSession(s.sessionStore, log.Logger, httptest.NewRecorder(), r)
	c.Assert(session, NotNil)

	user, err := models.CreateUser(
		faker.RandomString(10),
		faker.Internet().Email(),
		faker.Lorem().Word(),
		faker.Avatar().Url("jpg", 200, 300),
		s.db,
	)
	c.Assert(err, IsNil)

	user.Admin = true
	err = user.Save(s.db)
	c.Assert(err, IsNil)

	session.Values["User"] = *user
	sessionUser := GetSessionUser(session)
	c.Assert(sessionUser.UserId, Equals, user.UserId)

	impersonate, err := models.CreateUser(
		faker.RandomString(10),
		faker.Internet().Email(),
		faker.Lorem().Word(),
		faker.Avatar().Url("jpg", 200, 300),
		s.db,
	)
	c.Assert(err, IsNil)

	user.Impersonate = impersonate
	session.Values["User"] = *user
	sessionUser = GetSessionUser(session)
	c.Assert(sessionUser.UserId, Equals, impersonate.UserId)

	SetSessionUser(*impersonate, session, httptest.NewRecorder(), r)
	sessionUser = GetSessionUser(session)
	c.Assert(sessionUser.UserId, Equals, impersonate.UserId)

	user, err = models.CreateUser(
		faker.RandomString(10),
		faker.Internet().Email(),
		faker.Lorem().Word(),
		faker.Avatar().Url("jpg", 200, 300),
		s.db,
	)
	c.Assert(err, IsNil)

	SetSessionUser(*user, session, httptest.NewRecorder(), r)
	sessionUser = GetSessionUser(session)
	c.Assert(sessionUser.UserId, Equals, user.UserId)

	var subscription sql.NullString
	err = subscription.Scan(fmt.Sprintf("sub_%s", faker.RandomString(64)))
	c.Assert(err, IsNil)

	user.StripeSubscriptionId = subscription
	SetSessionUser(*user, session, httptest.NewRecorder(), r)
	err = session.Save(r, httptest.NewRecorder())
	c.Assert(err, IsNil)

	ok := IsSubscribed(session)
	c.Assert(ok, Equals, true)
}

func (s *helpersSuite) TestTemplates(c *C) {
	templates := compileList([]string{})
	c.Assert(assert.Greater(c, len(templates), 0), Equals, true)
}

func (s *helpersSuite) TearDownSuite(c *C) {}
