// Package endpoints
package endpoints

// NOTE: This is just a catch all for misc endpoints and might be renamed or moved in the future

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// AttributesMethods is used to call methods for all Attributes endpoints
	AttributesMethods struct{ methods.MethodCall }

	// AttributesResponse is used to map all Attributes endpoints
	AttributesResponse struct {
		Name        string  `json:"name,omitempty"`
		ParentID    int     `json:"parent_id,omitempty"`
		ID          int     `json:"id,omitempty"`
		CostPerHour float64 `json:"cost_per_hour,omitempty"`
	}

	// AttributesGetParams is used to construct a GET call to AttributesMethods
	// NOTE: Even though this is required, 0 is a valid input so required is not specified in the tag.
	// A better approach should be taken to this at some point
	AttributesGetParams struct {
		ID int `url:"id"`
	}
)

// Get for Attributes
// Requires scope: This depends on which attributes endpoint you are calling ensure its scope
// is created in invgo/scopes/scopes.go
// This Get method works for all attribute endpoints see the related endpoints documentation
// for ID definition and return definitions.
// If ID > 0 is provided, only one will be listed.
func (b *AttributesMethods) Get(p AttributesGetParams) ([]AttributesResponse, error) {
	// NOTE: RequiredScope type should be set in invgo/endpoint_methods.go when
	// creating the public method for each endpoint attribute since they all share this
	// Get method.
	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemoteGet()
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
type ServiceDeskVersionMethods struct{ methods.MethodCall }

// Get for ServiceDeskVersion
// Requires scope: ServiceDeskVersionGet
// See https://releases.invgate.com/service-desk/api/#sdversion-GET
func (s *ServiceDeskVersionMethods) Get() (string, error) {
	s.RequiredScope = scopes.ServiceDeskVersionGet

	resp, err := s.RemoteGet()
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
