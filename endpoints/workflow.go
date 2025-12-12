package endpoints

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

// WorkflowDeployMethods is used to call methods for ServiceDeskVersionMethods
type (
	WorkflowDeployMethods struct{ methods.MethodCall }

	// WorkflowDeployPutResponse is used to map a workflow put response from the Invgate API
	WorkflowDeployPutResponse struct {
		// Status OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status,omitempty"`
		// Description Confirmation message NOTE: Invgate says this is an int I am unsure if this is true
		Description int `json:"description,omitempty"`
		// WorkflowID ID of deployed workflow
		WorkflowID int `json:"workflow_id,omitempty"`
	}

	WorkflowDeployPutParams struct {
		WorkflowID int `url:"workflow_id,required"`
	}
)

// Put for WorkflowDeploy
// Requires scope: WorkflowDeployPut
// See https://releases.invgate.com/service-desk/api/#wfdeploy-PUT
func (w *WorkflowDeployMethods) Put(p WorkflowDeployPutParams) (WorkflowDeployPutResponse, error) {
	wf := WorkflowDeployPutResponse{}
	w.RequiredScope = scopes.WorkflowDeployPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return wf, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemotePut()
	if err != nil {
		return wf, err
	}

	err = json.Unmarshal(resp, &wf)
	if err != nil {
		return wf, err
	}
	return wf, nil
}

// WorkflowInitialFieldsByCategoryMethods is used to call methods for ServiceDeskVersionMethods
type (
	WorkflowInitialFieldsByCategoryMethods struct{ methods.MethodCall }

	// WorkflowInitialFieldsByCategoryGetResponse is used to map a workflow put response from the Invgate API
	WorkflowInitialFieldsByCategoryGetResponse struct {
		CategoryID             int    `json:"category_id,omitempty"`
		WorkflowInitialFields  any    `json:"workflow_initial_fields,omitempty"` // I am not yet sure what Invgate returns here
		AssociatedWorkflowName string `json:"associated_workflow_name,omitempty"`
		AssociatedWorkflowID   int    `json:"associated_workflow_id,omitempty"`
	}

	WorkflowInitialFieldsByCategoryGetParams struct {
		CategoryID int `url:"category_id,required"`
	}
)

// Get for WorkflowInitialFieldsByCategory
// Requires scope: WorkflowInitialFieldsByCategoryPut
// See https://releases.invgate.com/service-desk/api/#wfinitialfieldsbycategory-GET
func (w *WorkflowInitialFieldsByCategoryMethods) Get(p WorkflowInitialFieldsByCategoryGetParams) (WorkflowInitialFieldsByCategoryGetResponse, error) {
	wf := WorkflowInitialFieldsByCategoryGetResponse{}
	w.RequiredScope = scopes.WorkflowInitialFieldsByCategoryGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return wf, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemoteGet()
	if err != nil {
		return wf, err
	}

	err = json.Unmarshal(resp, &wf)
	if err != nil {
		return wf, err
	}
	return wf, nil
}
