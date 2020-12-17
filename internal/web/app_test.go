package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"github.com/lyticaa/lyticaa-app/internal/web/types"

	. "gopkg.in/check.v1"
)

type appSuite struct {
	a *App
}

var _ = Suite(&appSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *appSuite) SetUpSuite(c *C) {
	s.a = NewApp()
	s.a.Start()
}

func (s *appSuite) TestApp(c *C) {
	c.Assert(s.a.Monitoring.Logger, NotNil)
	c.Assert(s.a.Monitoring.NewRelic, NotNil)
	c.Assert(s.a.HTTP.Server, NotNil)
	c.Assert(s.a.HTTP.Router, NotNil)
	c.Assert(s.a.HTTP.Client, NotNil)
	c.Assert(s.a.Data.SessionStore, NotNil)
	c.Assert(s.a.Data.Db, NotNil)
}

func (s *appSuite) TestHealth(c *C) {
	url := fmt.Sprintf("%v/api/health_check", os.Getenv("BASE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	c.Assert(err, IsNil)

	resp, err := s.a.HTTP.Client.Do(req)
	c.Assert(err, IsNil)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	c.Assert(err, IsNil)

	var health types.Health
	err = json.Unmarshal(body, &health)

	c.Assert(err, IsNil)
	c.Assert(health.Status, Equals, "OK")
}

func (s *appSuite) TearDownSuite(c *C) {
	s.a.Stop()
}
