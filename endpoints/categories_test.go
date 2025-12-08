package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestCategoriesGet(t *testing.T) {
	a := assert.New(t)
	recCount := 3

	cats := make([]endpoints.CategoriesGetResponse, recCount)
	for range len(cats) {
		var cat endpoints.CategoriesGetResponse
		gofakeit.Struct(&cat)
		cats = append(cats, cat)
	}

	server := newTestServer(t, http.MethodGet, "/categories", cats)

	c := newTestClient(t, server, scopes.CategoriesGet)

	resp, err := c.Categories().Get(endpoints.CategoriesGetParams{ID: 0})
	a.NoError(err)
	a.Equal(cats, resp)
}
