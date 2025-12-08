package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestIncidentGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.Incident, 1)
	var inc endpoints.Incident
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident", incs)

	c := newTestClient(t, server, scopes.IncidentGet)

	resp, err := c.Incident().Get(endpoints.IncidentGetParams{ID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentPost(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentPostResponse
	var inc endpoints.IncidentPostParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPost, "/incident", r)

	c := newTestClient(t, server, scopes.IncidentPost)

	resp, err := c.Incident().Post(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentPut(t *testing.T) {
	a := assert.New(t)

	var put endpoints.IncidentPutParams
	gofakeit.Struct(&put)

	incs := make([]endpoints.Incident, 1)
	var inc endpoints.Incident
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodPut, "/incident", incs)

	c := newTestClient(t, server, scopes.IncidentPut)

	resp, err := c.Incident().Put(endpoints.IncidentPutParams{ID: put.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentsGet(t *testing.T) {
	a := assert.New(t)

	var inc endpoints.Incident
	gofakeit.Struct(&inc)

	incsMap := make(map[int]endpoints.Incident, 1)
	incsMap[0] = inc

	incsSlice := make([]endpoints.Incident, 1)
	incsSlice[0] = inc

	server := newTestServer(t, http.MethodGet, "/incidents", incsMap)

	c := newTestClient(t, server, scopes.IncidentsGet)

	resp, err := c.Incidents().Get(endpoints.IncidentsGetParams{IDs: []int{inc.ID}})
	a.NoError(err)
	a.Equal(incsSlice, resp)
}

func TestIncidentsByStatusGet(t *testing.T) {
	a := assert.New(t)

	var inc endpoints.IncidentsByStatusResponse
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodGet, "/incidents.by.status", inc)

	c := newTestClient(t, server, scopes.IncidentsByStatusGet)

	resp, err := c.IncidentsByStatus().Get(endpoints.IncidentsByStatusGetParams{StatusIDs: []int{0}})
	a.NoError(err)
	a.Equal(inc, resp)
}
