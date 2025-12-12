package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestWorkflowDeployPut(t *testing.T) {
	a := assert.New(t)
	var wf endpoints.WorkflowDeployPutResponse
	gofakeit.Struct(&wf)

	server := newTestServer(t, http.MethodPut, "/wf.deploy", wf)

	c := newTestClient(t, server, scopes.WorkflowDeployPut)

	got, err := c.WorkflowDeploy().Put(endpoints.WorkflowDeployPutParams{WorkflowID: 1})
	a.NoError(err)
	a.EqualValues(wf, got)
}

func TestWorkflowInitialFieldsByCategoryGet(t *testing.T) {
	a := assert.New(t)
	var wfCat endpoints.WorkflowInitialFieldsByCategoryGetResponse
	gofakeit.Struct(&wfCat)

	server := newTestServer(t, http.MethodGet, "/wf.initialfields.by.category", wfCat)

	c := newTestClient(t, server, scopes.WorkflowInitialFieldsByCategoryGet)

	got, err := c.WorkflowInitialFieldsByCategory().Get(endpoints.WorkflowInitialFieldsByCategoryGetParams{CategoryID: 1})
	a.NoError(err)
	a.EqualValues(wfCat, got)
}
