package helpers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupTemplateTests(t *testing.T) *httptest.Server {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		store := testSessionStore(t)
		session, err := store.Get(r, "auth-session")
		if err != nil {
			t.Error()
		}

		t := []string{
			"partials/nav/_main",
			"dashboard/overview",
			"partials/_filters",
		}
		RenderTemplate(w, t, session.Values)
	})
	s := httptest.NewServer(h)

	return s
}

func TestTemplateList(t *testing.T) {
	templates := templateList([]string{})
	if len(templates) == 0 {
		t.Error()
	}
}

func TestRenderTemplate(t *testing.T) {
	s := SetupTemplateTests(t)

	r, err := http.NewRequest("GET", s.URL, nil)
	if err != nil {
		t.Error()
	}

	client := http.Client{}
	response, err := client.Do(r)
	if err != nil {
		t.Error()
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		t.Error()
	}

	if string(body) == "" {
		t.Error()
	}
}
