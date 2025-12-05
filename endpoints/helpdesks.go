package endpoints

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

// HelpDesksMethods is used to call methods for HelpDesks
type HelpDesksMethods struct{ methods.MethodCall }

// HelpDesksGetResponse maps an individual help desk
type HelpDesksGetResponse struct {
	TotalMembers int    `json:"total_members,omitempty"`
	EngineID     int    `json:"engine_id,omitempty"`
	Name         string `json:"name,omitempty"`
	StatusID     int    `json:"status_id,omitempty"`
	ParentID     int    `json:"parent_id,omitempty"`
	ID           int    `json:"id,omitempty"`
}

type HelpDeskGetParams struct {
	ID             int    `url:"id"`
	Name           string `url:"name"`
	IncludeDeleted bool   `url:"include_deleted"`
}

// Get for HelpDesks
// If an ID of 0 is passed all help desks will be returned
// See https://releases.invgate.com/service-desk/api/#helpdesks-GET
func (h *HelpDesksMethods) Get(p HelpDeskGetParams) ([]HelpDesksGetResponse, error) {
	err := scopes.CheckScopes(h.Client.CurrentScopes, scopes.HelpDesksGet)
	if err != nil {
		return []HelpDesksGetResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	h.Endpoint.RawQuery = q.Encode()

	resp, err := h.RemoteGet()
	if err != nil {
		return []HelpDesksGetResponse{}, err
	}

	var d []HelpDesksGetResponse
	if p.Name != "" || p.ID > 0 {
		var calledDesk HelpDesksGetResponse
		err = json.Unmarshal(resp, &calledDesk)
		if err != nil {
			return []HelpDesksGetResponse{}, err
		}
		d = append(d, calledDesk)
	} else {
		err = json.Unmarshal(resp, &d)
		if err != nil {
			return []HelpDesksGetResponse{}, err
		}
	}

	return d, nil
}
