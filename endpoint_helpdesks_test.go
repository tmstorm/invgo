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
	var desks []invgo.HelpDesksGetResponse
	gofakeit.Struct(&desks)

	server := newTestServer(t, http.MethodGet, "/helpdesks", desks)

	c := newTestClient(t, server, invgo.HelpDesksGet)

	got, err := c.HelpDesks().Get(0, "", false)
	a.NoError(err)
	a.EqualValues(desks, got)
}
