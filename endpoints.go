package invgo

import (
	"encoding/json"
	"strconv"
)

type (
	// AttributesMethods is used to call methods for all Attributes endpoints
	AttributesMethods MethodCall

	// AttributesResponse is used to map all Attributes endpoints
	AttributesResponse struct {
		Name        string  `json:"name,omitempty"`
		ParentID    int     `json:"parent_id,omitempty"`
		ID          int     `json:"id,omitempty"`
		CostPerHour float64 `json:"cost_per_hour,omitempty"`
	}
)

// Get method for Attribute endpoints
// If ID > 0 is provided, only one will be listed.
// This Get method works for all attribute endpoints see the related endpoints documentation
// for ID definition and return definitions
func (b *AttributesMethods) Get(id int) ([]AttributesResponse, error) {
	// NOTE: Scopes are not checked here because this method can be called
	// by every Attributes endpoint. Scopes will need to be checked in a different way.
	if id > 0 {
		q := b.Endpoint.Query()
		q.Add("id", strconv.Itoa(id))
		b.Endpoint.RawQuery = q.Encode()
	}

	m := MethodCall(*b)

	resp, err := m.get()
	if err != nil {
		return []AttributesResponse{}, err
	}

	var d []AttributesResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		var s AttributesResponse
		err := json.Unmarshal(resp, &s)
		if err != nil {
			return []AttributesResponse{}, err
		}
		d = append(d, s)
	}
	return d, nil
}

// ServiceDeskVersionMethods is used to call methods for ServiceDeskVersionMethods
type ServiceDeskVersionMethods MethodCall

// ServiceDeskVersion returns the current version of the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#sdversion
func (c *Client) ServiceDeskVersion() *ServiceDeskVersionMethods {
	ep := c.APIURL.JoinPath("/sd.version")
	return &ServiceDeskVersionMethods{
		client:   c,
		Endpoint: ep,
	}
}

// Get returns the current version of the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#sdversion-GET
func (s *ServiceDeskVersionMethods) Get() (string, error) {
	err := checkScopes(s.client.CurrentScopes, ServiceDeskVersionGet)
	if err != nil {
		return "", err
	}

	m := MethodCall(*s)
	resp, err := m.get()
	if err != nil {
		return "", err
	}

	type version struct {
		Version string `json:"version,omitempty"`
	}
	var d version
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return "", err
	}
	return d.Version, nil
}
