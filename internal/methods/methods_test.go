package methods_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/scopes"
)

func TestRemoteGet(t *testing.T) {
	a := assert.New(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"id":1}`))
	}))
	defer server.Close()

	uri, err := url.Parse(server.URL + "/test")
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.BreakingNewsGet},
		},
		Endpoint:      uri,
		RequiredScope: scopes.BreakingNewsGet,
	}

	resp, err := m.RemoteGet()
	a.NoError(err)

	a.NotEmpty(resp)
}

func TestRemotePost(t *testing.T) {
	a := assert.New(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":ok}`))
	}))
	defer server.Close()

	uri, err := url.Parse(server.URL + "/test")
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.BreakingNewsPost},
		},
		Endpoint:      uri,
		RequiredScope: scopes.BreakingNewsPost,
	}

	resp, err := m.RemotePost()
	a.NoError(err)

	a.NotEmpty(resp)
}

func TestRemotePut(t *testing.T) {
	a := assert.New(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":ok}`))
	}))
	defer server.Close()

	uri, err := url.Parse(server.URL + "/test")
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.BreakingNewsPut},
		},
		Endpoint:      uri,
		RequiredScope: scopes.BreakingNewsPut,
	}

	resp, err := m.RemotePut()
	a.NoError(err)

	a.NotEmpty(resp)
}

func TestRemotePatch(t *testing.T) {
	a := assert.New(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":ok}`))
	}))
	defer server.Close()

	uri, err := url.Parse(server.URL + "/test")
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.ScopeType("api.v1.test:patch")},
		},
		Endpoint:      uri,
		RequiredScope: scopes.ScopeType("api.v1.test:patch"),
	}

	resp, err := m.RemotePatch()
	a.NoError(err)

	a.NotEmpty(resp)
}

func TestRemoteDelete(t *testing.T) {
	a := assert.New(t)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte(`{"status":ok}`))
	}))
	defer server.Close()

	uri, err := url.Parse(server.URL + "/test")
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.ScopeType("api.v1.test:delete")},
		},
		Endpoint:      uri,
		RequiredScope: scopes.ScopeType("api.v1.test:delete"),
	}

	resp, err := m.RemoteDelete()
	a.NoError(err)

	a.NotEmpty(resp)
}

func TestRemoteGetError(t *testing.T) {
	a := assert.New(t)
	invError := methods.InvgateError{
		Error:  "It broke like it should",
		Status: 501,
	}
	server := newTestServer(t, http.MethodGet, "/", 501, invError)
	defer server.Close()

	uri, err := url.Parse(server.URL)
	a.NoError(err)

	m := &methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    server.Client(),
			CurrentScopes: []scopes.ScopeType{scopes.BreakingNewsGet},
		},
		Endpoint:      uri,
		RequiredScope: scopes.BreakingNewsGet,
	}

	_, err = m.RemoteGet()
	a.Error(err)
}

func newTestServer(t *testing.T, expectedMethod, expectedPath string, status int, response any) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, expectedMethod)
		assert.Equal(t, expectedPath, r.URL.Path)

		w.WriteHeader(status)
		b, err := json.Marshal(&response)
		assert.NoError(t, err)

		w.Write(b)
	}))
}
