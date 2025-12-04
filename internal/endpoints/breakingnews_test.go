package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/internal/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestBreakingNewsGet(t *testing.T) {
	a := assert.New(t)
	var news endpoints.BreakingNewsGetResponse
	gofakeit.Struct(&news)

	var req endpoints.BreakingNewsGetParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodGet, "/breakingnews", news)

	c := newTestClient(t, server, scopes.BreakingNewsGet)

	c.BreakingNews()
	got, err := c.BreakingNews().Get(req)
	a.NoError(err)
	a.EqualValues(news, got)
}

func TestBreakingNewsPost(t *testing.T) {
	a := assert.New(t)
	var newPost endpoints.BreakingNewsPostParams
	gofakeit.Struct(&newPost)
	resp := endpoints.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post created",
		ID:     "1",
	}

	server := newTestServer(t, http.MethodPost, "/breakingnews", resp)

	c := newTestClient(t, server, scopes.BreakingNewsPost)

	got, err := c.BreakingNews().Post(newPost)
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsPut(t *testing.T) {
	a := assert.New(t)
	var updatePost endpoints.BreakingNewsPutParams
	gofakeit.Struct(&updatePost)
	resp := endpoints.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post created",
	}

	server := newTestServer(t, http.MethodPut, "/breakingnews", resp)

	c := newTestClient(t, server, scopes.BreakingNewsPut)

	got, err := c.BreakingNews().Put(updatePost)
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsAll(t *testing.T) {
	a := assert.New(t)
	var allNews []endpoints.BreakingNewsGetResponse
	gofakeit.Struct(&allNews)

	server := newTestServer(t, http.MethodGet, "/breakingnews.all", allNews)

	c := newTestClient(t, server, scopes.BreakingNewsAll)

	got, err := c.BreakingNewsAll().Get()
	a.NoError(err)
	a.EqualValues(allNews, got)
}

func TestBreakingNewsAttributesStatus(t *testing.T) {
	a := assert.New(t)
	var resp []endpoints.AttributesResponse
	att1 := endpoints.AttributesResponse{
		ID:   1,
		Name: "BreakingNewsAttribute_1",
	}
	att2 := endpoints.AttributesResponse{
		ID:   2,
		Name: "BreakingNewsAttribute_2",
	}

	resp = append(resp, att1, att2)

	server := newTestServer(t, http.MethodGet, "/breakingnews.attributes.status", resp)

	c := newTestClient(t, server, scopes.BreakingNewsAttributesStatus)

	got, err := c.BreakingNewsAttributesStatus().Get(endpoints.AttributesGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsAttributesType(t *testing.T) {
	a := assert.New(t)
	var resp []endpoints.AttributesResponse
	att1 := endpoints.AttributesResponse{
		ID:   1,
		Name: "BreakingNewsTypeAttribute_1",
	}
	att2 := endpoints.AttributesResponse{
		ID:   2,
		Name: "BreakingNewsTypeAttribute_2",
	}

	resp = append(resp, att1, att2)

	server := newTestServer(t, http.MethodGet, "/breakingnews.attributes.type", resp)

	c := newTestClient(t, server, scopes.BreakingNewsAttributesType)

	got, err := c.BreakingNewsAttributesType().Get(endpoints.AttributesGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsStatusGet(t *testing.T) {
	a := assert.New(t)
	var newsStatus []endpoints.BreakingNewsStatusGetResponse
	gofakeit.Struct(&newsStatus)

	var req endpoints.BreakingNewsStatusGetParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodGet, "/breakingnews.status", newsStatus)

	c := newTestClient(t, server, scopes.BreakingNewsStatusGet)

	got, err := c.BreakingNewsStatus().Get(req)
	a.NoError(err)
	a.EqualValues(newsStatus, got)
}

func TestBreakingNewsStatusPost(t *testing.T) {
	a := assert.New(t)
	resp := endpoints.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post updated",
	}

	var req endpoints.BreakingNewsStatusPostParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodPost, "/breakingnews.status", resp)

	c := newTestClient(t, server, scopes.BreakingNewsStatusPost)

	got, err := c.BreakingNewsStatus().Post(req)
	a.NoError(err)
	a.EqualValues(resp, got)
}
