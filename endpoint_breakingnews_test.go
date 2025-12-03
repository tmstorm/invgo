package invgo_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
)

func TestBreakingNewsGet(t *testing.T) {
	a := assert.New(t)
	var news invgo.BreakingNewsGetResponse
	gofakeit.Struct(&news)

	var req invgo.BreakingNewsGetParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodGet, "/breakingnews", news)

	c := newTestClient(t, server, invgo.BreakingNewsGet)

	got, err := c.BreakingNews().Get(req)
	a.NoError(err)
	a.EqualValues(news, got)
}

func TestBreakingNewsPost(t *testing.T) {
	a := assert.New(t)
	var newPost invgo.BreakingNewsPostParams
	gofakeit.Struct(&newPost)
	resp := invgo.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post created",
		ID:     "1",
	}

	server := newTestServer(t, http.MethodPost, "/breakingnews", resp)

	c := newTestClient(t, server, invgo.BreakingNewsPost)

	got, err := c.BreakingNews().Post(newPost)
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsPut(t *testing.T) {
	a := assert.New(t)
	var updatePost invgo.BreakingNewsPutParams
	gofakeit.Struct(&updatePost)
	resp := invgo.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post created",
	}

	server := newTestServer(t, http.MethodPut, "/breakingnews", resp)

	c := newTestClient(t, server, invgo.BreakingNewsPut)

	got, err := c.BreakingNews().Put(updatePost)
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsAll(t *testing.T) {
	a := assert.New(t)
	var allNews []invgo.BreakingNewsGetResponse
	gofakeit.Struct(&allNews)

	server := newTestServer(t, http.MethodGet, "/breakingnews.all", allNews)

	c := newTestClient(t, server, invgo.BreakingNewsAll)

	got, err := c.BreakingNewsAll()
	a.NoError(err)
	a.EqualValues(allNews, got)
}

func TestBreakingNewsAttributesStatus(t *testing.T) {
	a := assert.New(t)
	var resp []invgo.AttributesResponse
	att1 := invgo.AttributesResponse{
		ID:   1,
		Name: "BreakingNewsAttribute_1",
	}
	att2 := invgo.AttributesResponse{
		ID:   2,
		Name: "BreakingNewsAttribute_2",
	}

	resp = append(resp, att1, att2)

	server := newTestServer(t, http.MethodGet, "/breakingnews.attributes.status", resp)

	c := newTestClient(t, server, invgo.BreakingNewsAttributesStatus)

	got, err := c.BreakingNewsAttributesStatus().Get(invgo.AttributesGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsAttributesType(t *testing.T) {
	a := assert.New(t)
	var resp []invgo.AttributesResponse
	att1 := invgo.AttributesResponse{
		ID:   1,
		Name: "BreakingNewsTypeAttribute_1",
	}
	att2 := invgo.AttributesResponse{
		ID:   2,
		Name: "BreakingNewsTypeAttribute_2",
	}

	resp = append(resp, att1, att2)

	server := newTestServer(t, http.MethodGet, "/breakingnews.attributes.type", resp)

	c := newTestClient(t, server, invgo.BreakingNewsAttributesType)

	got, err := c.BreakingNewsAttributesType().Get(invgo.AttributesGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(resp, got)
}

func TestBreakingNewsStatusGet(t *testing.T) {
	a := assert.New(t)
	var newsStatus []invgo.BreakingNewsStatusGetResponse
	gofakeit.Struct(&newsStatus)

	var req invgo.BreakingNewsStatusGetParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodGet, "/breakingnews.status", newsStatus)

	c := newTestClient(t, server, invgo.BreakingNewsStatusGet)

	got, err := c.BreakingNewsStatus().Get(req)
	a.NoError(err)
	a.EqualValues(newsStatus, got)
}

func TestBreakingNewsStatusPost(t *testing.T) {
	a := assert.New(t)
	resp := invgo.BreakingNewsInfoResponse{
		Status: "OK",
		Info:   "post updated",
	}

	var req invgo.BreakingNewsStatusPostParams
	gofakeit.Struct(&req)

	server := newTestServer(t, http.MethodPost, "/breakingnews.status", resp)

	c := newTestClient(t, server, invgo.BreakingNewsStatusPost)

	got, err := c.BreakingNewsStatus().Post(req)
	a.NoError(err)
	a.EqualValues(resp, got)
}
