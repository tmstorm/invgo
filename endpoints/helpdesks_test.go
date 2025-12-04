package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestHelpDesksGet(t *testing.T) {
	a := assert.New(t)
	desks := []endpoints.HelpDesksGetResponse{}
	for range 4 {
		var desk endpoints.HelpDesksGetResponse
		gofakeit.Struct(&desk)
		desks = append(desks, desk)
	}

	server := newTestServer(t, http.MethodGet, "/helpdesks", desks)

	c := newTestClient(t, server, scopes.HelpDesksGet)

	got, err := c.HelpDesks().Get(endpoints.HelpDeskGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(desks, got)
}
