package invgo

/*
All endpoints must be implemented here to be available in the Invgo public API

To implement a new endpoint add its methods to invgo/endpoints/endpoint_name.go.
Then add it as a method to the Invgo Client using NewPublicMethod.

Example:
	func (c *Client) BreakingNews() *endpoints.BreakingNewsMethods {
		return NewPublicMethod[endpoints.BreakingNewsMethods](c, "/breakingnews")
	}
*/

import (
	"unsafe"

	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/scopes"
)

// NewPublicMethod should be used when adding a new enpoint to the Invgo public API
// T must be a struct whose first field is methods.MethodCall
func NewPublicMethod[T any](c *Client, endpoint string) *T {
	var zero T
	result := &zero

	ep := c.APIURL.JoinPath(endpoint)

	mcPtr := (*methods.MethodCall)(unsafe.Pointer(result))

	*mcPtr = methods.MethodCall{
		Client: &methods.Client{
			HTTPClient:    c.HTTPClient,
			CurrentScopes: c.CurrentScopes,
			APIURL:        c.APIURL,
		},
		Endpoint: ep,
	}

	return result
}

// BreakingNews manages the /breakingnews endpoint
// Get: Returns the requested Breaking News
// Post: Creates Breaking News
// Put: Modifies a set of Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnews
func (c *Client) BreakingNews() *endpoints.BreakingNewsMethods {
	return NewPublicMethod[endpoints.BreakingNewsMethods](c, "/breakingnews")
}

// BreakingNewsAll gets all the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsall
func (c *Client) BreakingNewsAll() *endpoints.BreakingNewsAllMethods {
	return NewPublicMethod[endpoints.BreakingNewsAllMethods](c, "/breakingnews.all")
}

// BreakingNewsStatus manages the updates of the breaking news.
// Get: Returns the updates of the requested Breaking News
// Post: Creates a new update to the given Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus
func (c *Client) BreakingNewsStatus() *endpoints.BreakingNewsStatusMethods {
	return NewPublicMethod[endpoints.BreakingNewsStatusMethods](c, "/breakingnews.status")
}

// BreakingNewsAttributesStatus gets all the possible status for the Breaking News'
// and their descriptions.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributesstatus
func (c *Client) BreakingNewsAttributesStatus() *endpoints.AttributesMethods {
	m := NewPublicMethod[endpoints.AttributesMethods](c, "/breakingnews.attributes.status")
	m.RequiredScope = scopes.BreakingNewsAttributesStatus
	return m
}

// BreakingNewsAttributesType gets all the importance types of the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributestype
func (c *Client) BreakingNewsAttributesType() *endpoints.AttributesMethods {
	m := NewPublicMethod[endpoints.AttributesMethods](c, "/breakingnews.attributes.type")
	m.RequiredScope = scopes.BreakingNewsAttributesType
	return m
}

// Categories is used to get all categories for the current Invgate instance
// See https://releases.invgate.com/service-desk/api/#categories
func (c *Client) Categories() *endpoints.CategoriesMethods {
	return NewPublicMethod[endpoints.CategoriesMethods](c, "/categories")
}

// HelpDesks manages the help desks
// See https://releases.invgate.com/service-desk/api/#helpdesks
func (c *Client) HelpDesks() *endpoints.HelpDesksMethods {
	return NewPublicMethod[endpoints.HelpDesksMethods](c, "/helpdesks")
}

// Incident manages the /incident endpoint
// Get: Returns the information of the given request
// Post: Creates a request
// Put: Change attributes of a request
// See https://releases.invgate.com/service-desk/api/#incident
func (c *Client) Incident() *endpoints.IncidentMethods {
	return NewPublicMethod[endpoints.IncidentMethods](c, "/incident")
}

// Incidents is used to get Incidents from the Invgate API
// See https://releases.invgate.com/service-desk/api/#incidents
func (c *Client) Incidents() *endpoints.IncidentsMethods {
	return NewPublicMethod[endpoints.IncidentsMethods](c, "/incidents")
}

// IncidentsByStatus gets incidents by the given set of status IDs
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus
func (c *Client) IncidentsByStatus() *endpoints.IncidentsByStatusMethods {
	return NewPublicMethod[endpoints.IncidentsByStatusMethods](c, "/incidents.by.status")
}

// IncidentAttributesStatus gets all the status types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesStatus() *endpoints.AttributesMethods {
	m := NewPublicMethod[endpoints.AttributesMethods](c, "/incident.attributes.status")
	m.RequiredScope = scopes.IncidentAttributesStatusGet
	return m
}

// IncidentAttributesType gets all the types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesType() *endpoints.AttributesMethods {
	m := NewPublicMethod[endpoints.AttributesMethods](c, "/incident.attributes.type")
	m.RequiredScope = scopes.IncidentAttributesTypeGet
	return m
}

// ServiceDeskVersion returns the current version of the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#sdversion
func (c *Client) ServiceDeskVersion() *endpoints.ServiceDeskVersionMethods {
	return NewPublicMethod[endpoints.ServiceDeskVersionMethods](c, "/sd.version")
}

// Triggers returns user defined tiggers in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#triggers
func (c *Client) Triggers() *endpoints.TriggersMethods {
	return NewPublicMethod[endpoints.TriggersMethods](c, "/triggers")
}

// TriggersExecutions returns a list of each time a trigger was executed in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#triggers
func (c *Client) TriggersExecutions() *endpoints.TriggersExecutionsMethods {
	return NewPublicMethod[endpoints.TriggersExecutionsMethods](c, "/triggers.executions")
}
