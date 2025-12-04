package endpoints

import (
	"encoding/json"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

// Incident is used to map an incident returned from the Invgate API
type Incident struct {
	ID                    int                       `json:"id,omitempty"`
	CategoryID            int                       `json:"category_id,omitempty"`
	CreatedAt             int                       `json:"created_at,omitempty"`
	UserID                int                       `json:"user_id,omitempty"`
	CustomFields          any                       `json:"custom_fields,omitempty"`
	Description           string                    `json:"description,omitempty"`
	CreatorID             int                       `json:"creator_id,omitempty"`
	SourceID              int                       `json:"source_id,omitempty"`
	Attachments           []int                     `json:"attachments,omitempty"`
	DateOcurred           int                       `json:"date_ocurred,omitempty"` // NOTE: The misspelling here is from the Invgate API
	StatusID              int                       `json:"status_id,omitempty"`
	ClosedAt              int                       `json:"closed_at,omitempty"`
	SLAIncidentFirstReply string                    `json:"sla_incident_first_reply,omitempty"`
	Comments              []IncidentCommentResponse `json:"comments,omitempty"`
	TypeID                int                       `json:"type_id,omitempty"`
	LastUpdate            int                       `json:"last_update,omitempty"`
	ClosedReason          int                       `json:"closed_reason,omitempty"`
	AssignedID            int                       `json:"assigned_id,omitempty"`
	Rating                int                       `json:"rating,omitempty"`
	AssignedGroupID       int                       `json:"assigned_group_id,omitempty"`
	Title                 string                    `json:"title,omitempty"`
	ProcessID             int                       `json:"process_id,omitempty"`
	PrettyID              string                    `json:"pretty_id,omitempty"`
	PriorityID            int                       `json:"priority_id,omitempty"`
	SolvedAt              int                       `json:"solved_at,omitempty"`
	SLAIncidentResolution string                    `json:"sla_incident_resolution,omitempty"`
}

// IncidentCommentResponse is used to map an comment returned from the Invgate API
type IncidentCommentResponse struct {
	ID              int    `json:"id,omitempty"`
	Attachments     []int  `json:"attachments,omitempty"`
	Message         string `json:"message,omitempty"`
	CreatedAt       int    `json:"created_at,omitempty"`
	Reference       int    `json:"reference,omitempty"`
	AuthorID        int    `json:"author_id,omitempty"`
	MsgNum          int    `json:"msg_num,omitempty"`
	IncidentID      int    `json:"incident_id,omitempty"`
	CustomerVisible int    `json:"customer_visible,omitempty"` // this is a bool but is returned as an 0-1 int
	IsSolution      bool   `json:"is_solution,omitempty"`
}

// IncidentAttachmentResponse is use to map an attachment returned from the Invgate API
type IncidentAttachmentResponse struct {
	ID        int    `json:"id,omitempty"`
	Extension string `json:"extension,omitempty"`
	Name      string `json:"name,omitempty"`
	Hash      string `json:"hash,omitempty"`
	URL       string `json:"url,omitempty"`
}

// IncidentMethods is use to call methods for Incident
type IncidentMethods struct{ methods.MethodCall }

type IncidentGetParams struct {
	ID                       int    `url:"id,required"`
	DecodedSpecialCharacters bool   `url:"decoded_special_character"`
	DateFormat               string `url:"date_format"`
	Comments                 bool   `url:"comments"`
}

// Get Method is used to get an incident using the given ID
// See https://releases.invgate.com/service-desk/api/#incident-GET
// NOTE: Invgate documentation says it returns and array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Get(p IncidentGetParams) ([]Incident, error) {
	err := scopes.CheckScopes(i.Client.CurrentScopes, scopes.IncidentGet)
	if err != nil {
		return nil, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return nil, err
	}

	var incs []Incident
	err = json.Unmarshal(resp, &incs)
	if err != nil {
		var inc Incident
		err = json.Unmarshal(resp, &inc)
		if err != nil {
			return nil, err
		}
		incs = append(incs, inc)
	}
	return incs, nil
}

// IncidentPostParams is used to construct a new POST request to create a new incident
type IncidentPostParams struct {
	Title       string                       `url:"title,required"`
	TypeID      int                          `url:"type_id,required"`
	CreatorID   int                          `url:"creator_id,required"`
	PriorityID  int                          `url:"priority_id,required"`
	CustomerID  int                          `url:"customer_id,required"`
	Date        string                       `url:"date"`
	CategoryID  int                          `url:"category_id"`
	SourceID    int                          `url:"source_id"`
	LocationID  int                          `url:"location_id"`
	Description string                       `url:"description"`
	RelatedTo   []int                        `url:"related_to"`
	Attachments []IncidentAttachmentResponse `url:"attachments"`
}

// IncidentPostResponse is used to map the response after posting a new incident
type IncidentPostResponse struct {
	RequestID string `json:"request_id,omitempty"`
	Info      string `json:"info,omitempty"`
	Status    string `json:"status,omitempty"`
}

// Post method creates a new incident and returns a success response
// See https://releases.invgate.com/service-desk/api/#incident-POST
func (i *IncidentMethods) Post(p IncidentPostParams) (IncidentPostResponse, error) {
	err := scopes.CheckScopes(i.Client.CurrentScopes, scopes.IncidentPost)
	if err != nil {
		return IncidentPostResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return IncidentPostResponse{}, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return IncidentPostResponse{}, err
	}

	var inc IncidentPostResponse
	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return IncidentPostResponse{}, err
	}

	return inc, nil
}

// IncidentPutParams is used to construct a PUT request to update an incident
type IncidentPutParams struct {
	ID           int    `url:"id,required"`
	Date         string `url:"date"`
	PriorityID   int    `url:"priority_id"`
	TypeID       int    `url:"type_id"`
	SourceID     int    `url:"source_id"`
	Title        string `url:"title"`
	LocationID   int    `url:"location_id"`
	CategoryID   int    `url:"category_id"`
	Description  string `url:"description"`
	Reassignment bool   `url:"reassignment"`
	DateFormat   string `url:"date_format"`
	CustomerID   int    `url:"customer_id"`
}

// Put is used to update an incident using the given changes and ID is required
// See https://releases.invgate.com/service-desk/api/#incident-PUT
// NOTE: Invgate documentation says it returns an array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Put(p IncidentPutParams) ([]Incident, error) {
	err := scopes.CheckScopes(i.Client.CurrentScopes, scopes.IncidentPut)
	if err != nil {
		return []Incident{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return []Incident{}, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return []Incident{}, err
	}

	var incs []Incident
	err = json.Unmarshal(resp, &incs)
	if err != nil {
		var inc Incident
		err = json.Unmarshal(resp, &inc)
		if err != nil {
			return nil, err
		}
		incs = append(incs, inc)
	}
	return incs, nil
}

// IncidentsMethods is use to call methods for Incidents
type IncidentsMethods struct{ methods.MethodCall }

type IncidentsGetParams struct {
	IDs             []int  `url:"ids,required"`
	IncludeComments bool   `url:"comments"`
	DateFormat      string `url:"date_format"`
}

// Get method is used to get a incindents the match the given incident ids
// At least one incident must be provided
// See https://releases.invgate.com/service-desk/api/#incidents-GET
func (i *IncidentsMethods) Get(p IncidentsGetParams) ([]Incident, error) {
	err := scopes.CheckScopes(i.Client.CurrentScopes, scopes.IncidentsGet)
	if err != nil {
		return nil, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return nil, err
	}

	var d map[int]Incident
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return nil, err
	}

	var incs []Incident
	for k := range d {
		incs = append(incs, d[k])
	}

	return incs, nil
}

// IncidentsByStatusMethods is used to call methods for IncidentsByStatus
type IncidentsByStatusMethods struct{ methods.MethodCall }

// IncidentsByStatusResponse is used to map responses from GET requests for IncidentsByStatus
type IncidentsByStatusResponse struct {
	Info       string `json:"info,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	RequestIDs []int  `json:"requestIds,omitempty"`
	Total      int    `json:"total,omitempty"`
	Status     string `json:"status,omitempty"`
}

type IncidentsByStatusGetParams struct {
	StatusIDs []int `url:"status_ids"`
	Limit     int   `url:"limit"`
	Offset    int   `url:"offset"`
}

// Get method for IncidentsByStatus at least one Status ID must be provided
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus-GET
func (i *IncidentsByStatusMethods) Get(p IncidentsByStatusGetParams) (IncidentsByStatusResponse, error) {
	err := scopes.CheckScopes(i.Client.CurrentScopes, scopes.IncidentsByStatusGet)
	if err != nil {
		return IncidentsByStatusResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return IncidentsByStatusResponse{}, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return IncidentsByStatusResponse{}, err
	}

	var d IncidentsByStatusResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return IncidentsByStatusResponse{}, err
	}

	return d, nil
}
