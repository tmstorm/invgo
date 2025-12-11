package endpoints_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

func newTestClient(t *testing.T, server *httptest.Server, scopes ...scopes.ScopeType) *invgo.Client {
	uri, err := utils.ParseURL(server.URL, "", true)
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

// newPublicMethod should be used when adding a new enpoint to the Invgo public API
// T must be a struct whose first field is methods.MethodCall
func newPublicMethod[T any](c *invgo.Client, endpoint string) *T {
	var zero T
	result := &zero

	ep := c.APIURL.JoinPath(endpoint)

	mcPtr := (*methods.MethodCall)(unsafe.Pointer(result))

	*mcPtr = methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    c.HTTPClient,
			CurrentScopes: c.CurrentScopes,
			APIURL:        c.APIURL,
		},
		Endpoint: ep,
	}

	return result
}
