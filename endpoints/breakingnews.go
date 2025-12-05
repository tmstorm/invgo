package endpoints

import (
	"encoding/json"
	"html/template"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// BreakingNewsMethods is used to call methods for BreakingNews
	BreakingNewsMethods struct{ methods.MethodCall }

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
// Requires scope: BreakingNewsGet
// See https://releases.invgate.com/service-desk/api/#breakingnews-GET
func (b *BreakingNewsMethods) Get(p BreakingNewsGetParams) (BreakingNewsGetResponse, error) {
	news := BreakingNewsGetResponse{}

	b.RequiredScope = scopes.BreakingNewsGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsGetResponse{}, err
	}
	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemoteGet()
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

// Post for BreakingNews
// Requires scope: BreakingNewsPost
// See https://releases.invgate.com/service-desk/api/#breakingnews-POST
func (b *BreakingNewsMethods) Post(p BreakingNewsPostParams) (BreakingNewsInfoResponse, error) {
	b.RequiredScope = scopes.BreakingNewsPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemotePost()
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

// Put for BreakingNews
// Requires scope: BreakingNewsPut
// See https://releases.invgate.com/service-desk/api/#breakingnews-PUT
func (b *BreakingNewsMethods) Put(p BreakingNewsPutParams) (BreakingNewsInfoResponse, error) {
	b.RequiredScope = scopes.BreakingNewsPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}

	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemotePut()
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

// BreakingNewsAllMethods is used to get all BreakingNews
type BreakingNewsAllMethods struct{ methods.MethodCall }

// Get for BreakingNewsAll
// Requires scope: BreakingNewsAll
// See https://releases.invgate.com/service-desk/api/#breakingnewsall
func (c *BreakingNewsAllMethods) Get() ([]BreakingNewsGetResponse, error) {
	c.RequiredScope = scopes.BreakingNewsAll

	resp, err := c.RemoteGet()
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

type (
	// BreakingNewsStatusMethods is used to call methods for BreakingNewsStatus
	BreakingNewsStatusMethods struct{ methods.MethodCall }

	// BreakingNewsStatusGetParams is used to construct GET requests to BreakingNewsStatus
	BreakingNewsStatusGetParams struct {
		ID         int    `url:"id,required"`
		dateFormat string `url:"date_format"`
	}

	// BreakingNewsStatusGetResponse maps breaking news updates
	BreakingNewsStatusGetResponse struct {
		CreatedAt int    `json:"created_at,omitempty"`
		Body      string `json:"body,omitempty"`
		CreatorID int    `json:"creator_id,omitempty"`
	}
)

// Get for BreakingNewsStatus
// Requires scope: BreakingNewsStatusGet
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-GET
func (b *BreakingNewsStatusMethods) Get(p BreakingNewsStatusGetParams) ([]BreakingNewsStatusGetResponse, error) {
	b.RequiredScope = scopes.BreakingNewsStatusGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return nil, err
	}
	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemoteGet()
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

// Post for BreakingNewsStatus
// Requires scope: BreakingNewsStatusPost
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus-POST
func (b *BreakingNewsStatusMethods) Post(p BreakingNewsStatusPostParams) (BreakingNewsInfoResponse, error) {
	b.RequiredScope = scopes.BreakingNewsStatusPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return BreakingNewsInfoResponse{}, err
	}
	b.Endpoint.RawQuery = q.Encode()

	resp, err := b.RemotePost()
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
