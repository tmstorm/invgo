package invgo

import (
	"encoding/json"
	"strconv"
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

// Get returns all active help desks, or one in particular
// See https://releases.invgate.com/service-desk/api/#helpdesks-GET
func (h *HelpDesksMethods) Get(id int, name string, includeDeleted bool) ([]HelpDesksGetResponse, error) {
	err := checkScopes(h.client.CurrentScopes, HelpDesksGet)
	if err != nil {
		return []HelpDesksGetResponse{}, err
	}

	q := h.Endpoint.Query()
	if name != "" {
		q.Add("name", name)
	}
	if id > 0 {
		q.Add("id", strconv.Itoa(id))
	}
	q.Add("include_deleted", strconv.FormatBool(includeDeleted))
	h.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*h)
	resp, err := m.get()
	if err != nil {
		return []HelpDesksGetResponse{}, err
	}

	var d []HelpDesksGetResponse
	if name != "" || id > 0 {
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
