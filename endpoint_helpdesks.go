package invgo

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/utils"
)

// HelpDesksMethods is used to call methods for HelpDesks
type HelpDesksMethods MethodCall

// HelpDesks manages the help desks
// See https://releases.invgate.com/service-desk/api/#helpdesks
func (c *Client) HelpDesks() *HelpDesksMethods {
	ep := c.APIURL.JoinPath("/helpdesks")
	return &HelpDesksMethods{
		client:   c,
		Endpoint: ep,
	}
}

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

// Get returns all active help desks, or one in particular
// See https://releases.invgate.com/service-desk/api/#helpdesks-GET
// If an ID of 0 is passed all help desks will be returned
func (h *HelpDesksMethods) Get(p HelpDeskGetParams) ([]HelpDesksGetResponse, error) {
	err := checkScopes(h.client.CurrentScopes, HelpDesksGet)
	if err != nil {
		return []HelpDesksGetResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	h.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*h)
	resp, err := m.get()
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
