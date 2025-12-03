package invgo

import (
	"encoding/json"
	"html/template"

	"github.com/tmstorm/invgo/internal/utils"
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

type (
	// BreakingNewsBase is used as a base to map BreakingNews requests and responses.
	// Invgate has different requirements for each call type. This implements the fields they all share
	// and each request must extend this struct as needed.
	BreakingNewsBase struct {
		CreatedByID         int   `json:"created_by_id,omitempty" url:"created_by_id"`
		AffectedHelpDeskIDs []int `json:"affected_helpdesk_ids,omitempty" url:"affected_helpdesk_ids"`
		ResolutionTime      int   `json:"resolution_time,omitempty" url:"resolution_time"`
		StatusID            int   `json:"status_id,omitempty" url:"status_id"`
		CreatedAt           int   `json:"created_at,omitempty" url:"created_at"`
		AffectedGroupIDs    []int `json:"affected_group_ids,omitempty" url:"affected_group_ids"`
	}

	// BreakingNewsGetResponse extends BreakingNewsBase for GET responses
	BreakingNewsGetResponse struct {
		ID     int           `json:"id,omitempty"`
		TypeID int           `json:"type_id,omitempty"`
		Title  string        `json:"title,omitempty"`
		Body   template.HTML `json:"body,omitempty"`
		BreakingNewsBase
	}

	// BreakingNewsGetParams extends BreakingNewsBase for GET requests
	BreakingNewsGetParams struct {
		ID         int    `url:"id,required"`
		DateFormat string `url:"date_format"`
		BreakingNewsBase
	}
)

// Get for BreakingNews
// See https://releases.invgate.com/service-desk/api/#breakingnews-GET
func (b *BreakingNewsMethods) Get(p BreakingNewsGetParams) (BreakingNewsGetResponse, error) {
	news := BreakingNewsGetResponse{}
	err := checkScopes(b.client.CurrentScopes, BreakingNewsGet)
	if err != nil {
		return news, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsGetResponse{}, err
	}
	b.Endpoint.RawQuery = q.Encode()

	m := MethodCall(*b)
	resp, err := m.get()
	if err != nil {
		return news, err
	}

	err = json.Unmarshal(resp, &news)
	if err != nil {
		return news, err
	}
	return news, nil
}

type (
	// BreakingNewsPostParams extends BreakingNewsBase for POST requests
	BreakingNewsPostParams struct {
		TypeID    int           `json:"type_id,omitempty" url:"type_id,required"`
		CreatorID int           `json:"creator_id" url:"creator_id"`
		Title     string        `json:"title,omitempty" url:"title,required"`
		Body      template.HTML `json:"body,omitempty" url:"body,required"`
		BreakingNewsBase
	}

	// BreakingNewsInfoResponse is used to map responses from POST and PUT requests
	BreakingNewsInfoResponse struct {
		Info   string `json:"info,omitempty"`
		ID     string `json:"id,omitempty"`
		Status string `json:"status,omitempty"`
	}
)

// Post creates breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnews-POST
func (b *BreakingNewsMethods) Post(p BreakingNewsPostParams) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsPost)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

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

// BreakingNewsPutParams extends BreakingNewsBase for PUT requests
type BreakingNewsPutParams struct {
	ID     int           `json:"id,omitempty" url:"id,required"`
	TypeID int           `json:"type_id,omitempty" url:"type_id"`
	Title  string        `json:"title,omitempty" url:"title"`
	Body   template.HTML `json:"body,omitempty" url:"body"`
	BreakingNewsBase
}

// Put modifies a set of breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnews-PUT
func (b *BreakingNewsMethods) Put(p BreakingNewsPutParams) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsPut)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

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
func (c *Client) BreakingNewsAll() ([]BreakingNewsGetResponse, error) {
	err := checkScopes(c.CurrentScopes, BreakingNewsAll)
	if err != nil {
		return nil, err
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

	news := []BreakingNewsGetResponse{}
	err = json.Unmarshal(resp, &news)
	if err != nil {
		var s BreakingNewsGetResponse
		err := json.Unmarshal(resp, &s)
		if err != nil {
			return nil, err
		}
		news = append(news, s)
	}

	return news, nil
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

// BreakingNewsStatusGetParams is used to construct GET requests to BreakingNewsStatus
type BreakingNewsStatusGetParams struct {
	ID         int    `url:"id,required"`
	dateFormat string `url:"date_format"`
}

// BreakingNewsStatusGetResponse maps breaking news updates
type BreakingNewsStatusGetResponse struct {
	CreatedAt int    `json:"created_at,omitempty"`
	Body      string `json:"body,omitempty"`
	CreatorID int    `json:"creator_id,omitempty"`
}

// Get returns updates of the requestd breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-GET
func (b *BreakingNewsStatusMethods) Get(p BreakingNewsStatusGetParams) ([]BreakingNewsStatusGetResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsStatusGet)
	if err != nil {
		return []BreakingNewsStatusGetResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
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

// BreakingNewsStatusPostParams is used to construct POST requests to BreakingNewsStatus
type BreakingNewsStatusPostParams struct {
	ID          int    `url:"id,required"`
	Body        string `url:"body,required"`
	CreatorID   int    `url:"creator_id"`
	IsSolutions bool   `url:"is_solution"`
}

// Post creates a new update to the given breaking news
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-POST
func (b *BreakingNewsStatusMethods) Post(p BreakingNewsStatusPostParams) (BreakingNewsInfoResponse, error) {
	err := checkScopes(b.client.CurrentScopes, BreakingNewsStatusPost)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}
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
