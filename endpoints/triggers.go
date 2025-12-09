package endpoints

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// TriggersMethods is use to call methods for triggers
	TriggersMethods struct{ methods.MethodCall }

	TriggersGetParams struct {
		TriggerID int `url:"trigger_id"`
	}

	// TriggersGetResponse is used to map an trigger returned from the Invgate API
	TriggersGetResponse struct {
		TriggerName string `url:"trigger_name"`
		ID          int    `url:"id"`
	}
)

// Get for Triggers
// Requires scope: TriggersGet
// See https://releases.invgate.com/service-desk/api/#triggers-GET
func (c *TriggersMethods) Get(p TriggersGetParams) ([]TriggersGetResponse, error) {
	c.RequiredScope = scopes.TriggersGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
	if err != nil {
		return nil, err
	}

	var t []TriggersGetResponse
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}

type (
	// TriggersExecutionsMethods is use to call methods for TriggersExecutions
	TriggersExecutionsMethods struct{ methods.MethodCall }

	// TriggersExecutionsGetResponse is used to map an trigger returned from the Invgate API
	TriggersExecutionsGetResponse struct {
		ExecutedAt  int `url:"executed_at"`
		RequestedID int `url:"requested_id"`
		TriggerID   int `url:"trigger_id"`
		ID          int `url:"id"`
	}
)

// Get for Triggers Executions
// Requires scope: TriggersExecutionsGet
// See https://releases.invgate.com/service-desk/api/#triggersexecutions-GET
func (c *TriggersExecutionsMethods) Get(p TriggersGetParams) ([]TriggersExecutionsGetResponse, error) {
	c.RequiredScope = scopes.TriggersExecutionsGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
	if err != nil {
		return nil, err
	}

	var t []TriggersExecutionsGetResponse
	err = json.Unmarshal(resp, &t)
	if err != nil {
		return nil, err
	}

	return t, nil
}
