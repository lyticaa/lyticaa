package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"

	"gitlab.com/getlytica/lytica-app/internal/core/app/types"

	. "gopkg.in/check.v1"
)

type AppSuite struct {
	a *App
}

var _ = Suite(&AppSuite{})

func Test(t *testing.T) { TestingT(t) }

func (s *AppSuite) SetUpSuite(c *C) {
	s.a = NewApp()
	s.a.Start()
}

func (s *AppSuite) TestApp(c *C) {
	c.Assert(s.a.Logger, NotNil)
	c.Assert(s.a.NewRelic, NotNil)
	c.Assert(s.a.Srv, NotNil)
	c.Assert(s.a.Router, NotNil)
	c.Assert(s.a.Client, NotNil)
	c.Assert(s.a.SessionStore, NotNil)
	c.Assert(s.a.Db, NotNil)
}

func (s *AppSuite) TestHealth(c *C) {
	url := fmt.Sprintf("%v/api/health_check", os.Getenv("BASE_URL"))
	req, err := http.NewRequest("GET", url, nil)
	c.Assert(err, IsNil)

	resp, err := s.a.Client.Do(req)
	c.Assert(err, IsNil)

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	c.Assert(err, IsNil)

	var health types.Health
	err = json.Unmarshal(body, &health)

	c.Assert(err, IsNil)
	c.Assert(health.Status, Equals, "OK")
}

func (s *AppSuite) TearDownSuite(c *C) {
	s.a.Stop()
}
