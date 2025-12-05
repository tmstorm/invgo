package invgo

import (
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/internal/methods"
)

// BreakingNews manages the /breakingnews endpoint
// Get: Returns the requested Breaking News
// Post: Creates Breaking News
// Put: Modifies a set of Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnews
func (c *Client) BreakingNews() *endpoints.BreakingNewsMethods {
	ep := c.APIURL.JoinPath("/breakingnews")
	return &endpoints.BreakingNewsMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// BreakingNewsAll gets all the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsall
func (c *Client) BreakingNewsAll() *endpoints.BreakingNewsAllMethods {
	ep := c.APIURL.JoinPath("/breakingnews.all")
	return &endpoints.BreakingNewsAllMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// BreakingNewsStatus manages the updates of the breaking news.
// Get: Returns the updates of the requested Breaking News
// Post: Creates a new update to the given Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus
func (c *Client) BreakingNewsStatus() *endpoints.BreakingNewsStatusMethods {
	ep := c.APIURL.JoinPath("/breakingnews.status")
	return &endpoints.BreakingNewsStatusMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// BreakingNewsAttributesStatus gets all the possible status for the Breaking News'
// and their descriptions.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributesstatus
func (c *Client) BreakingNewsAttributesStatus() *endpoints.AttributesMethods {
	ep := c.APIURL.JoinPath("/breakingnews.attributes.status")

	return &endpoints.AttributesMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// BreakingNewsAttributesType gets all the importance types of the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributestype
func (c *Client) BreakingNewsAttributesType() *endpoints.AttributesMethods {
	ep := c.APIURL.JoinPath("/breakingnews.attributes.type")
	return &endpoints.AttributesMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// HelpDesks manages the help desks
// See https://releases.invgate.com/service-desk/api/#helpdesks
func (c *Client) HelpDesks() *endpoints.HelpDesksMethods {
	ep := c.APIURL.JoinPath("/helpdesks")
	return &endpoints.HelpDesksMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// Categories is used to get all categories for the current Invgate instance
// See https://releases.invgate.com/service-desk/api/#categories
func (c *Client) Categories() *endpoints.CategoriesMethods {
	ep := c.APIURL.JoinPath("/categories")
	return &endpoints.CategoriesMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// Incident manages the /incident endpoint
// Get: Returns the information of the given request
// Post: Creates a request
// Put: Change attributes of a request
// See https://releases.invgate.com/service-desk/api/#incident
func (c *Client) Incident() *endpoints.IncidentMethods {
	ep := c.APIURL.JoinPath("/incident")
	return &endpoints.IncidentMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// Incidents is used to get Incidents from the Invgate API
// See https://releases.invgate.com/service-desk/api/#incidents
func (c *Client) Incidents() *endpoints.IncidentsMethods {
	ep := c.APIURL.JoinPath("/incidents")
	return &endpoints.IncidentsMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// IncidentsByStatus gets incidents by the given set of status IDs
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus
func (c *Client) IncidentsByStatus() *endpoints.IncidentsByStatusMethods {
	ep := c.APIURL.JoinPath("/incidents.by.status")
	return &endpoints.IncidentsByStatusMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// IncidentAttributesStatus gets all the status types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesStatus() *endpoints.AttributesMethods {
	ep := c.APIURL.JoinPath("/incident.attributes.status")
	return &endpoints.AttributesMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// IncidentAttributesType gets all the types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesType() *endpoints.AttributesMethods {
	ep := c.APIURL.JoinPath("/incident.attributes.type")
	return &endpoints.AttributesMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}

// ServiceDeskVersion returns the current version of the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#sdversion
func (c *Client) ServiceDeskVersion() *endpoints.ServiceDeskVersionMethods {
	ep := c.APIURL.JoinPath("/sd.version")
	return &endpoints.ServiceDeskVersionMethods{
		MethodCall: methods.MethodCall{
			Client: &methods.Client{
				HTTPClient:    c.HTTPClient,
				CurrentScopes: c.CurrentScopes,
				APIURL:        c.APIURL,
			},
			Endpoint: ep,
		},
	}
}
