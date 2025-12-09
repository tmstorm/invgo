package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestTriggersGet(t *testing.T) {
	a := assert.New(t)
	trigs := []endpoints.TriggersGetResponse{}
	for range 4 {
		var trig endpoints.TriggersGetResponse
		gofakeit.Struct(&trig)
		trigs = append(trigs, trig)
	}

	server := newTestServer(t, http.MethodGet, "/triggers", trigs)

	c := newTestClient(t, server, scopes.TriggersGet)

	got, err := c.Triggers().Get(endpoints.TriggersGetParams{TriggerID: 0})
	a.NoError(err)
	a.EqualValues(trigs, got)
}

func TestTriggersExecutionsGet(t *testing.T) {
	a := assert.New(t)
	trigs := []endpoints.TriggersExecutionsGetResponse{}
	for range 4 {
		var trig endpoints.TriggersExecutionsGetResponse
		gofakeit.Struct(&trig)
		trigs = append(trigs, trig)
	}

	server := newTestServer(t, http.MethodGet, "/triggers.executions", trigs)

	c := newTestClient(t, server, scopes.TriggersExecutionsGet)

	got, err := c.TriggersExecutions().Get(endpoints.TriggersGetParams{TriggerID: 0})
	a.NoError(err)
	a.EqualValues(trigs, got)
}
