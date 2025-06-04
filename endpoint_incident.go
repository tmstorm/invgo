package invgo

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

// IncidentsResponse is used to map an incident returned from the Invgate API
type Incident struct {
	ID                    int                       `json:"id,omitempty"`
	CategoryID            int                       `json:"category_id,omitempty"`
	CreatedAt             int                       `json:"created_at,omitempty"`
	UserID                int                       `json:"user_id,omitempty"`
	CustomFields          interface{}               `json:"custom_fields,omitempty"`
	Description           string                    `json:"description,omitempty"`
	CreatorID             int                       `json:"creator_id,omitempty"`
	SourceID              int                       `json:"source_id,omitempty"`
	Attachments           []int                     `json:"attachments,omitempty"`
	DateOcurred           int                       `json:"date_ocurred,omitempty"`
	StatusID              int                       `json:"status_id,omitempty"`
	ClosedAt              int                       `json:"closed_at,omitempty"`
	SlaIncidentFirstReply string                    `json:"sla_incident_first_reply,omitempty"`
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
	SlaIncidentResolution string                    `json:"sla_incident_resolution,omitempty"`
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
type IncidentMethods MethodCall

// Incident manages the /incident endpoint
// Get: Returns the information of the given request
// Post: Creates a request
// Put: Change attributes of a request
// See https://releases.invgate.com/service-desk/api/#incident
func (c *Client) Incident() *IncidentMethods {
	ep := c.APIURL.JoinPath("/incident")
	return &IncidentMethods{
		client:   c,
		Endpoint: ep,
	}
}

// Get Method is used to get an incident using the given ID
// See https://releases.invgate.com/service-desk/api/#incident-GET
// NOTE: Invgate documentation says it returns and array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Get(id int, decodedSpecialCharacters bool, dateFormat string, comments bool) ([]Incident, error) {
	err := checkScopes(i.client.CurrentScopes, IncidentGet)
	if err != nil {
		return nil, err
	}

	if id == 0 {
		return nil, errors.New("no ID provided for get incident")
	}

	q := i.Endpoint.Query()
	q.Add("id", strconv.Itoa(id))
	if decodedSpecialCharacters {
		q.Add("decoded_special_character", strconv.FormatBool(decodedSpecialCharacters))
	}
	if dateFormat != "" {
		q.Add("date_format", dateFormat)
	}
	if comments {
		q.Add("comments", strconv.FormatBool(comments))
	}

	i.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*i)
	resp, err := m.get()
	if err != nil {
		return nil, err
	}

	var d []Incident
	err = json.Unmarshal(resp, &d)
	if err != nil {
		var inc Incident
		err = json.Unmarshal(resp, &inc)
		if err != nil {
			return nil, err
		}
		d = append(d, inc)
	}
	return d, nil
}

// IncidentPostParams is used to construct a new POST request to create a new incident
type IncidentPostParams struct {
	Date        string                       `json:"date,omitempty"`
	PriorityID  int                          `json:"priority_id,omitempty"`
	CategoryID  int                          `json:"category_id,omitempty"`
	TypeID      int                          `json:"type_id,omitempty"`
	SourceID    int                          `json:"source_id,omitempty"`
	Title       string                       `json:"title,omitempty"`
	LocationID  int                          `json:"location_id,omitempty"`
	CreatorID   int                          `json:"creator_id,omitempty"`
	Description string                       `json:"description,omitempty"`
	RelatedTo   []int                        `json:"related_to,omitempty"`
	CustomerID  int                          `json:"customer_id,omitempty"`
	Attachments []IncidentAttachmentResponse `json:"attachments,omitempty"`
}

// IncidentPostResponse is used to map the response after posting a new incident
type IncidentPostResponse struct {
	RequestID string `json:"request_id,omitempty"`
	Info      string `json:"info,omitempty"`
	Status    string `json:"status,omitempty"`
}

// Post method creates a new incident and returns a success response
// See https://releases.invgate.com/service-desk/api/#incident-POST
func (i *IncidentMethods) Post(inc IncidentPostParams) (IncidentPostResponse, error) {
	err := checkScopes(i.client.CurrentScopes, IncidentPost)
	if err != nil {
		return IncidentPostResponse{}, err
	}

	q := i.Endpoint.Query()
	if inc.Date != "" {
		q.Add("date", inc.Date)
	}

	if inc.PriorityID < 1 || inc.PriorityID > 5 {
		return IncidentPostResponse{}, fmt.Errorf("a valid priority id must be provided. Got %v wanted a number between 1 and 5", inc.PriorityID)
	} else {
		q.Add("priority_id", strconv.Itoa(inc.PriorityID))
	}

	if inc.CategoryID == 0 {
		return IncidentPostResponse{}, fmt.Errorf("a valid category id must be provided but got %v", inc.CategoryID)
	}
	_, err = i.client.Categories().Get(inc.CategoryID)
	if err != nil {
		return IncidentPostResponse{}, err
	}
	q.Add("category_id", strconv.Itoa(inc.CategoryID))

	if inc.TypeID < 1 || inc.TypeID > 6 {
		return IncidentPostResponse{}, fmt.Errorf("a valid type id must be provided. Got %v wanted a number between 1 and 6", inc.TypeID)
	} else {
		q.Add("type_id", strconv.Itoa(inc.TypeID))
	}

	if inc.SourceID > 0 {
		q.Add("source_id", strconv.Itoa(inc.SourceID))
	}

	if inc.Title == "" {
		return IncidentPostResponse{}, errors.New("a title is required when creating a new incident")
	}
	q.Add("title", inc.Title)

	if inc.LocationID > 0 {
		q.Add("location_id", strconv.Itoa(inc.LocationID))
	}

	if inc.CreatorID == 0 {
		return IncidentPostResponse{}, errors.New("a creator id is required")
	}
	q.Add("creator_id", strconv.Itoa(inc.CreatorID))

	if inc.Description != "" {
		q.Add("description", inc.Description)
	}

	if len(inc.RelatedTo) > 0 {
		for k := range inc.RelatedTo {
			id := strconv.Itoa(inc.RelatedTo[k])
			related := fmt.Sprintf("related_to[%d]", k)
			q.Add(related, id)
		}
	}

	if inc.CustomerID == 0 {
		return IncidentPostResponse{}, errors.New("a customer id is required")
	}
	q.Add("customer_id", strconv.Itoa(inc.CustomerID))

	if len(inc.Attachments) > 0 {
		for k := range inc.Attachments {
			at, err := json.Marshal(inc.Attachments[k])
			if err != nil {
				return IncidentPostResponse{}, err
			}
			attachment := fmt.Sprintf("attachments[%d]", k)
			q.Add(attachment, string(at))
		}
	}

	i.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*i)
	resp, err := m.post()
	if err != nil {
		return IncidentPostResponse{}, err
	}

	var d IncidentPostResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return IncidentPostResponse{}, err
	}

	return d, nil
}

// IncidentPutParams is used to construct a PUT request to update an incident
type IncidentPutParams struct {
	Date         string `json:"date,omitempty"`
	PriorityID   int    `json:"priority_id,omitempty"`
	TypeID       int    `json:"type_id,omitempty"`
	ID           int    `json:"id,omitempty"`
	SourceID     int    `json:"source_id,omitempty"`
	Title        string `json:"title,omitempty"`
	LocationID   int    `json:"location_id,omitempty"`
	CategoryID   int    `json:"category_id,omitempty"`
	Description  string `json:"description,omitempty"`
	Reassignment bool   `json:"reassignment,omitempty"`
	DateFormat   string `json:"date_format,omitempty"`
	CustomerID   int    `json:"customer_id,omitempty"`
}

// Put is used to update an incident using the given changes and ID is requried
// See https://releases.invgate.com/service-desk/api/#incident-PUT
// NOTE: Invgate documentation says it returns and array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Put(params IncidentPutParams) ([]Incident, error) {
	err := checkScopes(i.client.CurrentScopes, IncidentPut)
	if err != nil {
		return []Incident{}, err
	}

	if params.ID == 0 {
		return []Incident{}, errors.New("no ID provided for incident to be updated")
	}
	q := i.Endpoint.Query()
	q.Add("id", strconv.Itoa(params.ID))
	if params.Date != "" {
		q.Add("date", params.Date)
	}
	if params.PriorityID > 5 {
		return []Incident{}, fmt.Errorf("a valid priority id must be provided. Got %v wanted a number between 1 and 5", params.PriorityID)
	} else if params.PriorityID != 0 {
		q.Add("priority_id", strconv.Itoa(params.PriorityID))
	}
	if params.TypeID > 6 {
		return []Incident{}, fmt.Errorf("a valid type id must be provided. Got %v wanted a number between 1 and 6", params.TypeID)
	} else if params.TypeID != 0 {
		q.Add("type_id", strconv.Itoa(params.TypeID))
	}
	if params.SourceID > 0 {
		q.Add("source_id", strconv.Itoa(params.SourceID))
	}
	if params.Title != "" {
		q.Add("title", params.Title)
	}
	if params.LocationID > 0 {
		q.Add("location_id", strconv.Itoa(params.LocationID))
	}
	if params.CategoryID > 0 {
		q.Add("category_id", strconv.Itoa(params.CategoryID))
	}
	if params.Description != "" {
		q.Add("description", params.Description)
	}
	if params.Reassignment {
		q.Add("reassignment", strconv.FormatBool(params.Reassignment))
	}
	if params.DateFormat != "" {
		q.Add("date_format", params.DateFormat)
	}
	if params.CustomerID > 0 {
		q.Add("customer_id", strconv.Itoa(params.CustomerID))
	}
	i.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*i)
	resp, err := m.put()
	if err != nil {
		return []Incident{}, err
	}

	var d []Incident
	err = json.Unmarshal(resp, &d)
	if err != nil {
		var inc Incident
		err = json.Unmarshal(resp, &inc)
		if err != nil {
			return nil, err
		}
		d = append(d, inc)
	}
	return d, nil
}

// IncidentAttributesStatus gets all the status types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesStatus() *AttributesMethods {
	ep := c.APIURL.JoinPath("/incident.attributes.status")
	return &AttributesMethods{
		client:   c,
		Endpoint: ep,
	}
}

// IncidentAttributesType gets all the types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesType() *AttributesMethods {
	ep := c.APIURL.JoinPath("/incident.attributes.type")
	return &AttributesMethods{
		client:   c,
		Endpoint: ep,
	}
}

// IncidentsMethods is use to call methods for Incidents
type IncidentsMethods MethodCall

// Incidents is used to get Incidents from the Invgate API
// See https://releases.invgate.com/service-desk/api/#incidents
func (c *Client) Incidents() *IncidentsMethods {
	ep := c.APIURL.JoinPath("/incidents")
	return &IncidentsMethods{
		client:   c,
		Endpoint: ep,
	}
}

// Get method is used to get a incindents the match the given incident ids
// At least one incident must be provided
// See https://releases.invgate.com/service-desk/api/#incidents-GET
func (i *IncidentsMethods) Get(ids []int, includeComments bool, dateFormat string) ([]Incident, error) {
	if len(ids) == 0 {
		return nil, fmt.Errorf("expected at least one incident ID but got: %v", ids)
	}
	err := checkScopes(i.client.CurrentScopes, IncidentsGet)
	if err != nil {
		return nil, err
	}

	q := i.Endpoint.Query()
	for i := range ids {
		id := fmt.Sprintf("ids[%d]", i)
		q.Add(id, strconv.Itoa(ids[i]))
	}
	switch includeComments {
	case false:
		q.Add("comments", strconv.Itoa(0))
	case true:
		q.Add("comments", strconv.Itoa(1))
	}
	if dateFormat != "" {
		q.Add("date_format", dateFormat)
	}
	i.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*i)
	resp, err := m.get()
	if err != nil {
		return nil, err
	}

	var d map[int]Incident
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return nil, err
	}

	var incidents []Incident
	for k := range d {
		incidents = append(incidents, d[k])
	}

	return incidents, nil
}

// IncidentsByStatusMethods is used to call methods for IncidentsByStatus
type IncidentsByStatusMethods MethodCall

// IncidentsByStatusResponse is used to map responses from GET requests for IncidentsByStatus
type IncidentsByStatusResponse struct {
	Info       string `json:"info,omitempty"`
	Limit      int    `json:"limit,omitempty"`
	Offset     int    `json:"offset,omitempty"`
	RequestIDs []int  `json:"requestIds,omitempty"`
	Total      int    `json:"total,omitempty"`
	Status     string `json:"status,omitempty"`
}

// IncidentsByStatus gets incidents by the given set of status IDs
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus
func (c *Client) IncidentsByStatus() *IncidentsByStatusMethods {
	ep := c.APIURL.JoinPath("/incidents.by.status")
	return &IncidentsByStatusMethods{
		client:   c,
		Endpoint: ep,
	}
}

// Get method for IncidentsByStatus at least one Status ID must be provided
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus-GET
func (i *IncidentsByStatusMethods) Get(statusIDs []int, limit int, offset int) (IncidentsByStatusResponse, error) {
	err := checkScopes(i.client.CurrentScopes, IncidentsByStatusGet)
	if err != nil {
		return IncidentsByStatusResponse{}, err
	}

	q := i.Endpoint.Query()
	if len(statusIDs) == 0 {
		return IncidentsByStatusResponse{}, fmt.Errorf("expected at least one status ID but got: %v", statusIDs)
	}
	for i := range statusIDs {
		status = fmt.Sprintf("status_ids[%d]", i)
		q.Add(status, strconv.Itoa(statusIDs[i]))
	}
	if limit > 0 {
		q.Add("limit", strconv.Itoa(limit))
	}
	if offset > 0 {
		q.Add("offset", strconv.Itoa(offset))
	}
	i.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*i)
	resp, err := m.get()
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
