package invgo_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo"
)

func TestHelpDesksGet(t *testing.T) {
	a := assert.New(t)
	desks := []invgo.HelpDesksGetResponse{}
	for range 4 {
		var desk invgo.HelpDesksGetResponse
		gofakeit.Struct(&desk)
		desks = append(desks, desk)
	}

	server := newTestServer(t, http.MethodGet, "/helpdesks", desks)

	c := newTestClient(t, server, invgo.HelpDesksGet)

	got, err := c.HelpDesks().Get(invgo.HelpDeskGetParams{ID: 0})
	a.NoError(err)
	a.EqualValues(desks, got)
}
