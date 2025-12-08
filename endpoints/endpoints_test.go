package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

var testScope = scopes.ScopeType("api.v1.test:get")

func testAttrs(c *invgo.Client) *endpoints.AttributesMethods {
	m := invgo.NewPublicMethod[endpoints.AttributesMethods](c, "/attrs.test")
	m.RequiredScope = testScope
	return m
}

func TestAttributesGet(t *testing.T) {
	testScope := scopes.ScopeType("api.v1.test:get")

	a := assert.New(t)

	var attr endpoints.AttributesResponse
	gofakeit.Struct(&attr)

	attrs := make([]endpoints.AttributesResponse, 1)
	attrs[0] = attr

	server := newTestServer(t, http.MethodGet, "/attrs.test", attrs)

	c := newTestClient(t, server, testScope)

	resp, err := testAttrs(c).Get(endpoints.AttributesGetParams{ID: 0})
	a.NoError(err)
	a.Equal(attrs, resp)
}

func TestServiceDeskVersionGet(t *testing.T) {
	a := assert.New(t)

	v := endpoints.ServiceDeskVersionResponse{
		Version: "v1.2.3",
	}

	server := newTestServer(t, http.MethodGet, "/sd.version", v)

	c := newTestClient(t, server, scopes.ServiceDeskVersionGet)

	resp, err := c.ServiceDeskVersion().Get()
	a.NoError(err)
	a.Equal(v.Version, resp)
}
