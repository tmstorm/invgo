package invgo_test

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/scopes"
)

var (
	testScope    = scopes.ScopeType("api.v1.test:get")
	testEndpoint = "/test.endpoint"
)

type testEndpointMethods struct{ methods.MethodCall }

func (c *testEndpointMethods) get() error {
	if c.RequiredScope != testScope {
		return errors.New("scopes don't match")
	}
	return nil
}

func TestNewPublicMethod(t *testing.T) {
	a := assert.New(t)

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	defer server.Close()

	u, err := url.Parse("https://test.com")
	a.NoError(err)

	c := &invgo.Client{
		HTTPClient:    server.Client(),
		CurrentScopes: []scopes.ScopeType{scopes.BreakingNewsGet},
		APIURL:        u,
	}

	m := invgo.NewPublicMethod[testEndpointMethods](c, testEndpoint)
	m.RequiredScope = testScope

	err = m.get()
	a.NoError(err)
}
