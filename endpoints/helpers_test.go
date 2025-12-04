package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
	"github.com/tmstorm/invgo/scopes"
)

func newTestClient(t *testing.T, server *httptest.Server, scopes ...scopes.ScopeType) *invgo.Client {
	uri, err := url.Parse(server.URL)
	assert.NoError(t, err)

	return &invgo.Client{
		APIURL:        uri,
		HTTPClient:    server.Client(),
		CurrentScopes: scopes,
	}
}

func newTestServer(t *testing.T, expectedMethod, expectedPath string, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, expectedMethod)
		assert.Equal(t, expectedPath, r.URL.Path)

		w.WriteHeader(http.StatusOK)
		b, err := json.Marshal(&response)
		assert.NoError(t, err)

		w.Write(b)
	}))
}
