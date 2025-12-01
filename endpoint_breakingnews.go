package invgo

import (
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"
)

// BreakingNewsMethods is used to call methods for BreakingNews
type BreakingNewsMethods MethodCall

// BreakingNews manages the /breakingnews endpoint
// Get: Returns the requested Breaking News
// Post: Creates Breaking News
// Put: Modifies a set of Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnews
func (c *Client) BreakingNews() *BreakingNewsMethods {
	ep := c.APIURL.JoinPath("/breakingnews")
	return &BreakingNewsMethods{
		client:   c,
		Endpoint: ep,
	}
}

// BreakingNews is used to construct BreakingNews news posts
// This is used to get a post, create new, or modify existing posts
type BreakingNews struct {
	CreatedByID         int           `json:"created_by_id,omitempty"`
	AffectedHelpDeskIDs []int         `json:"affected_helpdesk_ids,omitempty"`
	ResolutionTime      int           `json:"resolution_time,omitempty"`
	Body                template.HTML `json:"body,omitempty"`
	TypeID              int           `json:"type_id,omitempty"`
	StatusID            int           `json:"status_id,omitempty"`
	CreatedAt           int           `json:"created_at,omitempty"`
	AffectedGroupIDs    []int         `json:"affected_group_ids,omitempty"`
	Title               string        `json:"title,omitempty"`
	ID                  int           `json:"id,omitempty"`
}

// Get for BreakingNews
// See https://releases.invgate.com/service-desk/api/#breakingnews-GET
func (b *BreakingNewsMethods) Get(id int, dateFormat string) (BreakingNews, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsGet)
	if err != nil {
		return BreakingNews{}, err
	}

	q := b.Endpoint.Query()
	q.Add("id", strconv.Itoa(id))
	if dateFormat != "" {
		q.Add("date_format", dateFormat)
	}
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)
	resp, err := m.get()
	if err != nil {
		return BreakingNews{}, err
	}

	var d BreakingNews
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return BreakingNews{}, err
	}
	return d, nil
}

// BreakingNewsInfoResponse is used to map responses from Post and Put requests
type BreakingNewsInfoResponse struct {
	Info   string `json:"info,omitempty"`
	ID     string `json:"id,omitempty"`
	Status string `json:"status,omitempty"`
}

// Post creates breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnews-POST
func (b *BreakingNewsMethods) Post(p BreakingNews) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsPost)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q := b.Endpoint.Query()
	for i := range p.AffectedHelpDeskIDs {
		affected := fmt.Sprintf("affected_helpdesk_ids[%d]", i)
		q.Add(affected, strconv.Itoa(p.AffectedHelpDeskIDs[i]))
	}
	if p.ResolutionTime > 0 {
		q.Add("resolution_time", strconv.Itoa(p.ResolutionTime))
	}
	if p.CreatedByID > 0 {
		q.Add("creator_id", strconv.Itoa(p.CreatedByID))
	}
	for i := range p.AffectedGroupIDs {
		affected := fmt.Sprintf("affected_group_ids[%d]", i)
		q.Add(affected, strconv.Itoa(p.AffectedGroupIDs[i]))
	}

	q.Add("body", string(p.Body))
	q.Add("type_id", strconv.Itoa(p.TypeID))
	q.Add("title", p.Title)
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)

	resp, err := m.post()
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	var d BreakingNewsInfoResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}
	return d, nil
}

// Put modifies a set of breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnews-PUT
func (b *BreakingNewsMethods) Put(id int, CreatorID int, p BreakingNews) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsPut)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q := b.Endpoint.Query()
	for i := range p.AffectedHelpDeskIDs {
		affected := fmt.Sprintf("affected_helpdesk_ids[%d]", i)
		q.Add(affected, strconv.Itoa(p.AffectedHelpDeskIDs[i]))
	}
	if p.ResolutionTime > 0 {
		q.Add("resolution_time", strconv.Itoa(p.ResolutionTime))
	}
	if p.Body != "" {
		q.Add("body", string(p.Body))
	}
	if p.TypeID > 0 {
		q.Add("type_id", strconv.Itoa(p.TypeID))
	}
	if CreatorID > 0 {
		q.Add("creator_id", strconv.Itoa(CreatorID))
	}
	for i := range p.AffectedGroupIDs {
		affected := fmt.Sprintf("affected_group_ids[%d]", i)
		q.Add(affected, strconv.Itoa(p.AffectedGroupIDs[i]))
	}
	if p.Title != "" {
		q.Add("title", p.Title)
	}

	q.Add("id", strconv.Itoa(id))
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)

	resp, err := m.put()
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	var d BreakingNewsInfoResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}
	return d, nil
}

// BreakingNewsAll gets all the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsall
func (c *Client) BreakingNewsAll() ([]BreakingNews, error) {
	err := checkScopes(c.CurrentScopes, BreakingNewsAll)
	if err != nil {
		return []BreakingNews{}, err
	}

	ep := c.APIURL.JoinPath("/breakingnews.all")
	m := MethodCall{
		client:   c,
		Endpoint: ep,
	}
	resp, err := m.get()
	if err != nil {
		return nil, err
	}

	var d []BreakingNews
	err = json.Unmarshal(resp, &d)
	if err != nil {
		var s BreakingNews
		err := json.Unmarshal(resp, &s)
		if err != nil {
			return nil, err
		}
		d = append(d, s)
	}

	return d, nil
}

// BreakingNewsAttributesStatus gets all the possible status for the Breaking News'
// and their descriptions.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributesstatus
func (c *Client) BreakingNewsAttributesStatus() *AttributesMethods {
	ep := c.APIURL.JoinPath("/breakingnews.attributes.status")

	return &AttributesMethods{
		client:   c,
		Endpoint: ep,
	}
}

// BreakingNewsAttributesType gets all the importance types of the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributestype
func (c *Client) BreakingNewsAttributesType() *AttributesMethods {
	ep := c.APIURL.JoinPath("/breakingnews.attributes.type")
	return &AttributesMethods{
		client:   c,
		Endpoint: ep,
	}
}

// BreakingNewsStatusMethods is used to call methods for BreakingNewsStatus
type BreakingNewsStatusMethods MethodCall

// BreakingNewsStatus manages the updates of the breaking news.
// Get: Returns the updates of the requested Breaking News
// Post: Creates a new update to the given Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus
func (c *Client) BreakingNewsStatus() *BreakingNewsStatusMethods {
	ep := c.APIURL.JoinPath("/breakingnews.status")
	return &BreakingNewsStatusMethods{
		client:   c,
		Endpoint: ep,
	}
}

// BreakingNewsStatusGetResponse maps breaking news updates
type BreakingNewsStatusGetResponse struct {
	CreatedAt int    `json:"created_at,omitempty"`
	Body      string `json:"body,omitempty"`
	CreatorID int    `json:"creator_id,omitempty"`
}

// Get returns updates of the requestd breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-GET
func (b *BreakingNewsStatusMethods) Get(id int, dateFormat string) ([]BreakingNewsStatusGetResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsStatusGet)
	if err != nil {
		return []BreakingNewsStatusGetResponse{}, err
	}

	q := b.Endpoint.Query()
	q.Add("id", strconv.Itoa(id))
	if dateFormat != "" {
		q.Add("date_format", dateFormat)
	}
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)
	resp, err := m.get()
	if err != nil {
		return []BreakingNewsStatusGetResponse{}, err
	}

	var d []BreakingNewsStatusGetResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return []BreakingNewsStatusGetResponse{}, err
	}

	return d, nil
}

// Post creates a new update to the given breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-POST
func (b *BreakingNewsStatusMethods) Post(id int, body string, creatorID int, isSolution bool) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsStatusPost)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q := b.Endpoint.Query()
	q.Add("id", strconv.Itoa(id))
	q.Add("body", body)
	if creatorID > 0 {
		q.Add("creator_id", strconv.Itoa(creatorID))
	}

	solution := 0
	if isSolution {
		solution = 1
	}
	q.Add("is_solution", strconv.Itoa(solution))
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)
	resp, err := m.post()
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	var d BreakingNewsInfoResponse
	err = json.Unmarshal(resp, &d)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	return d, nil
}
