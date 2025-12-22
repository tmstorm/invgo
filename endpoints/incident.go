package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// Incident is used to map an incident returned from the Invgate API
	Incident struct {
		ID                              int                       `json:"id,omitempty"`
		CategoryID                      int                       `json:"category_id,omitempty"`
		CreatedAt                       int                       `json:"created_at,omitempty"`
		UserID                          int                       `json:"user_id,omitempty"`
		CustomFields                    any                       `json:"custom_fields,omitempty"`
		Description                     string                    `json:"description,omitempty"`
		CreatorID                       int                       `json:"creator_id,omitempty"`
		SourceID                        int                       `json:"source_id,omitempty"`
		Attachments                     []int                     `json:"attachments,omitempty"`
		DateOcurred                     int                       `json:"date_ocurred,omitempty"` // NOTE: The misspelling here is from the Invgate API
		StatusID                        int                       `json:"status_id,omitempty"`
		ClosedAt                        int                       `json:"closed_at,omitempty"`
		SLAIncidentFirstReply           string                    `json:"sla_incident_first_reply,omitempty"`
		Comments                        []IncidentCommentResponse `json:"comments,omitempty"`
		TypeID                          int                       `json:"type_id,omitempty"`
		LastUpdate                      int                       `json:"last_update,omitempty"`
		ClosedReason                    int                       `json:"closed_reason,omitempty"`
		AssignedID                      int                       `json:"assigned_id,omitempty"`
		Rating                          int                       `json:"rating,omitempty"`
		AssignedGroupID                 int                       `json:"assigned_group_id,omitempty"`
		Title                           string                    `json:"title,omitempty"`
		ProcessID                       int                       `json:"process_id,omitempty"`
		PrettyID                        string                    `json:"pretty_id,omitempty"`
		PriorityID                      int                       `json:"priority_id,omitempty"`
		SolvedAt                        int                       `json:"solved_at,omitempty"`
		SLAIncidentResolution           string                    `json:"sla_incident_resolution,omitempty"`
		RequestCustomerSentimentInitial string                    `json:"request_customer_sentiment_initial,omitempty"`
		RequestCustomerSentimentCurrent string                    `json:"request_customer_sentiment_current,omitempty"`
	}

	// IncidentCommentResponse is used to map an comment returned from the Invgate API
	//
	// NOTE: An incident returns a slightly different response structure for comments than
	// calling the /incident.comment endpoint so a they have different structs defined.
	// For example customer_visible returns 0-1 for a bool here but in /incident.comment it
	// returns a false-true for bool.
	IncidentCommentResponse struct {
		AuthorID        int    `json:"author_id,omitempty"`
		Reference       int    `json:"reference,omitempty"`
		IsSolution      bool   `json:"is_solution,omitempty"`
		ID              int    `json:"id,omitempty"`
		CreatedAt       int    `json:"created_at,omitempty"`
		CustomerVisible int    `json:"customer_visible,omitempty"` // NOTE: this is a bool but is returned as an 0-1 int
		Attachments     []int  `json:"attached_files,omitempty"`
		MsgNum          int    `json:"msg_num,omitempty"`
		IncidentID      int    `json:"incident_id,omitempty"`
		Message         string `json:"message,omitempty"`
	}

	// IncidentMethods is use to call methods for Incident
	IncidentMethods struct{ methods.MethodCall }

	IncidentGetParams struct {
		ID                       int    `url:"id,required"`
		DecodedSpecialCharacters bool   `url:"decoded_special_character"`
		DateFormat               string `url:"date_format"`
		Comments                 bool   `url:"comments"`
	}
)

// Get for Incident
// Requires scope: IncidentGet
// See https://releases.invgate.com/service-desk/api/#incident-GET
// NOTE: Invgate documentation says it returns and array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Get(p IncidentGetParams) ([]Incident, error) {
	i.RequiredScope = scopes.IncidentGet

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

type (
	// IncidentPostParams is used to construct a new POST request to create a new incident
	IncidentPostParams struct {
		Title       string                          `url:"title,required"`
		TypeID      int                             `url:"type_id,required"`
		CreatorID   int                             `url:"creator_id,required"`
		PriorityID  int                             `url:"priority_id,required"`
		CustomerID  int                             `url:"customer_id,required"`
		Date        string                          `url:"date"`
		CategoryID  int                             `url:"category_id"`
		SourceID    int                             `url:"source_id"`
		LocationID  int                             `url:"location_id"`
		Description string                          `url:"description"`
		RelatedTo   []int                           `url:"related_to"`
		Attachments []IncidentAttachmentGetResponse `url:"attachments"`
	}

	// IncidentPostResponse is used to map the response after posting a new incident
	IncidentPostResponse struct {
		RequestID string `json:"request_id,omitempty"`
		Info      string `json:"info,omitempty"`
		Status    string `json:"status,omitempty"`
	}
)

// Post for Incident
// Requires scope: IncidentPost
// See https://releases.invgate.com/service-desk/api/#incident-POST
func (i *IncidentMethods) Post(p IncidentPostParams) (IncidentPostResponse, error) {
	i.RequiredScope = scopes.IncidentPost

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

// Put for Incident
// Requires scope: IncidentPut
// See https://releases.invgate.com/service-desk/api/#incident-PUT
// NOTE: Invgate documentation says it returns an array. This does not appear to be the case.
// However this method still accounts for that if it is ever the case.
func (i *IncidentMethods) Put(p IncidentPutParams) ([]Incident, error) {
	i.RequiredScope = scopes.IncidentPut

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

type (
	// IncidentApprovalMethods is use to call methods for Incident
	IncidentApprovalMethods struct{ methods.MethodCall }

	IncidentApprovalGetParams struct {
		OnlyPending bool   `url:"only_pending"`
		DateFormat  string `url:"date_format"`
		RequestID   int    `url:"request_id,required"`
	}

	IncidentApprovalGetResponse struct {
		// Status ID of the approval. -2: Cancelled, -1: Waiting, 0: Rejected, 1: Approved.
		Status int `json:"status,omitempty"`
		// Type ID of the approval. 1: Predefined approval, 2: Spontaneous approval.
		Type                       int    `json:"type,omitempty"`
		ApprovalRequestDescription string `json:"approval_request_description,omitempty"`
		ID                         int    `json:"id,omitempty"`
		// CreatedAt Date when the approval was triggered
		// in epoch or ISO-8601 format depending on the date_format parameter.
		CreatedAt         int `json:"created_at,omitempty"` // NOTE: Docs say this returns a string but returns an int
		ApprovalRequestID int `json:"approval_request_id,omitempty"`
		AuthorID          int `json:"author_id,omitempty"`

		// NOTE: The fields below are not in the docs but are sent from the API
		// so the types are best guesses
		IsUsingApprovalManager int `json:"is_using_approval_manager,omitempty"` // I think this is a bool but int is returned
		ReminderCount          int `json:"reminder_count,omitempty"`
		Reassigned             int `json:"reassigned"`     // I think this is a bool but int is returned
		RemindersSent          int `json:"reminders_sent"` // I think this is a bool but int is returned
		LastReminderDate       int `json:"last_reminder_date"`
	}
)

// Get for IncidentApproval
// Requires scope: IncidentApprovalGet
// See https://releases.invgate.com/service-desk/api/#incidentapproval-GET
func (i *IncidentApprovalMethods) Get(p IncidentApprovalGetParams) ([]IncidentApprovalGetResponse, error) {
	incs := []IncidentApprovalGetResponse{}
	i.RequiredScope = scopes.IncidentApprovalGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return incs, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return incs, err
	}

	err = json.Unmarshal(resp, &incs)
	if err != nil {
		return nil, err
	}
	return incs, nil
}

type (
	// IncidentApprovalAcceptMethods is use to call methods for IncidentApprovalAccept
	IncidentApprovalAcceptMethods struct{ methods.MethodCall }

	IncidentApprovalAcceptPutParams struct {
		RequestID int    `url:"request_id,required"`
		UserID    int    `url:"user_id,required"`
		Note      string `url:"note"`
	}

	IncidentApprovalAcceptPutResponse struct {
		// OK if approval was accepted, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for IncidentApprovalAccept
// Requires scope: IncidentApprovalAcceptPut
// See https://releases.invgate.com/service-desk/api/#incidentapprovalaccept-PUT
func (i *IncidentApprovalAcceptMethods) Put(p IncidentApprovalAcceptPutParams) (IncidentApprovalAcceptPutResponse, error) {
	inc := IncidentApprovalAcceptPutResponse{}
	i.RequiredScope = scopes.IncidentApprovalAcceptPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentApprovalAddVoterMethods is use to call methods for IncidentApprovalAddVoter
	IncidentApprovalAddVoterMethods struct{ methods.MethodCall }

	IncidentApprovalAddVoterPostParams struct {
		UserID     int `url:"user_id,required"`
		ApprovalID int `url:"approval_id,required"`
	}

	IncidentApprovalAddVoterPostResponse struct {
		UserID     int `json:"user_id,omitempty"`
		ApprovalID int `json:"approval_id,omitempty"`
	}
)

// Post for IncidentApprovalAddVoter
// Requires scope: IncidentApprovalAddVoterPost
// See https://releases.invgate.com/service-desk/api/#incidentapprovaladd_voter-POST
func (i *IncidentApprovalAddVoterMethods) Post(p IncidentApprovalAddVoterPostParams) (IncidentApprovalAddVoterPostResponse, error) {
	inc := IncidentApprovalAddVoterPostResponse{}
	i.RequiredScope = scopes.IncidentApprovalAddVoterPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentApprovalCancelMethods is use to call methods for IncidentApprovalCancel
	IncidentApprovalCancelMethods struct{ methods.MethodCall }

	IncidentApprovalCancelPutParams struct {
		RequestID int `url:"request_id,required"`
		UserID    int `url:"user_id,required"`
	}

	IncidentApprovalCancelPutResponse struct {
		// OK if approval was canceled, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for IncidentApprovalCancel
// Requires scope: IncidentApprovalCancelPut
// See https://releases.invgate.com/service-desk/api/#incidentapprovalcancel-PUT
func (i *IncidentApprovalCancelMethods) Put(p IncidentApprovalCancelPutParams) (IncidentApprovalCancelPutResponse, error) {
	inc := IncidentApprovalCancelPutResponse{}
	i.RequiredScope = scopes.IncidentApprovalCancelPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentApprovalPossibleVotersMethods is use to call methods for IncidentApprovalPossibleVoters
	IncidentApprovalPossibleVotersMethods struct{ methods.MethodCall }

	IncidentApprovalPossibleVotersGetParams struct {
		ApprovalID  int  `url:"approval_id,required"`
		OnlyPending bool `url:"only_pending"`
	}

	IncidentApprovalPossibleVotersGetResponse struct {
		ID int `json:"user_id,omitempty"` // NOTE: invgate says it returns id but is actually user_id
		// Approval status ID. -2: Annulled, -1: Waiting, 0: Rejected, 1: Approved, 2:Expired.
		Status int `json:"status"`
	}
)

// Get for IncidentApprovalPossibleVoters
// Requires scope: IncidentApprovalPossibleVotersGet
// See https://releases.invgate.com/service-desk/api/#incidentapprovalpossible_voters-GET
func (i *IncidentApprovalPossibleVotersMethods) Get(p IncidentApprovalPossibleVotersGetParams) ([]IncidentApprovalPossibleVotersGetResponse, error) {
	inc := []IncidentApprovalPossibleVotersGetResponse{}
	i.RequiredScope = scopes.IncidentApprovalPossibleVotersGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentApprovalRejectMethods is use to call methods for IncidentApprovalReject
	IncidentApprovalRejectMethods struct{ methods.MethodCall }

	IncidentApprovalRejectPutParams struct {
		RequestID int `url:"request_id,required"`
		UserID    int `url:"user_id,required"`
		Note      int `url:"note"`
	}

	IncidentApprovalRejectPutResponse struct {
		// OK if approval was rejected, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for IncidentApprovalReject
// Requires scope: IncidentApprovalRejectPut
// See https://releases.invgate.com/service-desk/api/#incidentapprovalreject-PUT
func (i *IncidentApprovalRejectMethods) Put(p IncidentApprovalRejectPutParams) (IncidentApprovalRejectPutResponse, error) {
	inc := IncidentApprovalRejectPutResponse{}
	i.RequiredScope = scopes.IncidentApprovalRejectPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentApprovalStatusMethods is use to call methods for IncidentApprovalStatus
	IncidentApprovalStatusMethods struct{ methods.MethodCall }

	// IncidentApprovalStatusGetResponse map approval statuses
	IncidentApprovalStatusGetResponse struct {
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

// Get for IncidentApprovalStatus
// Requires scope: IncidentApprovalstatusGet
// See https://releases.invgate.com/service-desk/api/#incidentapprovalstatus-GET
//
// NOTE: Invgate says it returns an array but it is actually an map
// like map[int]string where int is the status and string is the description.
// To make this easier to access it has been converted into an array of type IncidentApprovalStatusGetResponse
// containing the ID and Description for each status.
func (i *IncidentApprovalStatusMethods) Get() ([]IncidentApprovalStatusGetResponse, error) {
	i.RequiredScope = scopes.IncidentApprovalStatusGet

	resp, err := i.RemoteGet()
	if err != nil {
		return nil, err
	}

	var b map[int]string
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return nil, err
	}

	incs := make([]IncidentApprovalStatusGetResponse, 0, len(b))
	for k, v := range b {
		s := IncidentApprovalStatusGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	return incs, nil
}

type (
	// IncidentApprovalTypeMethods is use to call methods for IncidentApprovalStatus
	IncidentApprovalTypeMethods struct{ methods.MethodCall }

	// IncidentApprovalTypeGetResponse map approval types
	IncidentApprovalTypeGetResponse struct {
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

// Get for IncidentApprovalType
// Requires scope: IncidentApprovalTypeGet
// See https://releases.invgate.com/service-desk/api/#incidentapprovaltype-GET
//
// NOTE: Invgate says it returns an array but it is actually an map
// like map[int]string where int is the status and string is the description.
// To make this easier to access it has been converted into an array of type IncidentApprovalTypeGetResponse
// containing the ID and Description for each type.
func (i *IncidentApprovalTypeMethods) Get() ([]IncidentApprovalTypeGetResponse, error) {
	i.RequiredScope = scopes.IncidentApprovalTypeGet

	resp, err := i.RemoteGet()
	if err != nil {
		return nil, err
	}

	var b map[int]string
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return nil, err
	}

	incs := make([]IncidentApprovalTypeGetResponse, 0, len(b))
	for k, v := range b {
		s := IncidentApprovalTypeGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	return incs, nil
}

type (
	// IncidentApprovalVoteStatusMethods is use to call methods for IncidentApprovalVoteStatus
	IncidentApprovalVoteStatusMethods struct{ methods.MethodCall }

	// IncidentApprovalVoteStatusGetResponse map approval types
	IncidentApprovalVoteStatusGetResponse struct {
		ID          int    `json:"id,omitempty"`
		Description string `json:"description,omitempty"`
	}
)

// Get for IncidentApprovalVoteStatus
// Requires scope: IncidentApprovalVoteStatusGet
// See https://releases.invgate.com/service-desk/api/#incidentapprovalvote_status-GET
//
// NOTE: Invgate says it returns an array but it is actually an map
// like map[int]string where int is the status and string is the description.
// To make this easier to access it has been converted into an array of type IncidentApprovalVoteStatusGetResponse
// containing the ID and Description for each type.
func (i *IncidentApprovalVoteStatusMethods) Get() ([]IncidentApprovalVoteStatusGetResponse, error) {
	i.RequiredScope = scopes.IncidentApprovalVoteStatusGet

	resp, err := i.RemoteGet()
	if err != nil {
		return nil, err
	}

	var b map[int]string
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return nil, err
	}

	incs := make([]IncidentApprovalVoteStatusGetResponse, 0, len(b))
	for k, v := range b {
		s := IncidentApprovalVoteStatusGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	return incs, nil
}

type (
	// IncidentAttachmentMethods is use to call methods for IncidentAttachment
	IncidentAttachmentMethods struct{ methods.MethodCall }

	IncidentAttachmentGetParams struct {
		ID int `url:"id,required"`
	}

	// IncidentAttachmentGetResponse map approval types
	IncidentAttachmentGetResponse struct {
		Name      string `json:"name,omitempty" url:"name"`
		URL       string `json:"url,omitempty" url:"url"`
		Hash      string `json:"hash,omitempty" url:"hash"`
		ID        int    `json:"id,omitempty" url:"id"`
		Extension string `json:"extension,omitempty" url:"extension"`
	}
)

// Get for IncidentAttachment
// Requires scope: IncidentAttachmentGet
// See https://releases.invgate.com/service-desk/api/#incidentattachment-GET
func (i *IncidentAttachmentMethods) Get(p IncidentAttachmentGetParams) (IncidentAttachmentGetResponse, error) {
	att := IncidentAttachmentGetResponse{}
	i.RequiredScope = scopes.IncidentAttachmentGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return att, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return att, err
	}

	err = json.Unmarshal(resp, &att)
	if err != nil {
		return att, err
	}

	return att, nil
}

type (
	// IncidentCancelMethods is use to call methods for IncidentCancel
	IncidentCancelMethods struct{ methods.MethodCall }

	IncidentCancelPostParams struct {
		Comment   string `url:"comment"`
		AuthorID  int    `url:"author_id,required"`
		RequestID int    `url:"request_id,required"`
	}

	IncidentCancelPostResponse struct {
		// OK if incident was canceled, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Post for IncidentCancel
// Requires scope: IncidentCancelPost
// See https://releases.invgate.com/service-desk/api/#incidentcancel-POST
func (i *IncidentCancelMethods) Post(p IncidentCancelPostParams) (IncidentCancelPostResponse, error) {
	inc := IncidentCancelPostResponse{}
	i.RequiredScope = scopes.IncidentCancelPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}
	return inc, nil
}

type (
	// IncidentCollaboratorMethods is use to call methods for IncidentCollaborator
	IncidentCollaboratorMethods struct{ methods.MethodCall }

	IncidentCollaboratorGetParams struct {
		RequestID int `url:"request_id,required"`
	}

	IncidentCollaboratorGetResponse struct {
		IDs []int `json:"ids"`
	}
)

// Get for IncidentCollaborator
// Requires scope: IncidentCollaboratorGet
// See https://releases.invgate.com/service-desk/api/#incidentcancel-Get
func (i *IncidentCollaboratorMethods) Get(p IncidentCollaboratorGetParams) (IncidentCollaboratorGetResponse, error) {
	inc := IncidentCollaboratorGetResponse{}
	i.RequiredScope = scopes.IncidentCollaboratorGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return inc, err
	}

	var b []int
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return inc, err
	}

	for i := range len(b) {
		coll := b[i]
		inc.IDs = append(inc.IDs, coll)
	}

	return inc, nil
}

type (
	IncidentCollaboratorPostParams struct {
		UserID    int   `url:"user_id"`
		AuthorID  int   `url:"author_id,required"`
		UsersID   []int `url:"users_id"`
		RequestID int   `url:"request_id,required"`
	}

	IncidentCollaboratorPostResponse struct {
		// OK if incident was canceled, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Post for IncidentCollaborator
// Requires scope: IncidentCollaboratorPost
// See https://releases.invgate.com/service-desk/api/#incidentcancel-POST
func (i *IncidentCollaboratorMethods) Post(p IncidentCollaboratorPostParams) (IncidentCollaboratorPostResponse, error) {
	inc := IncidentCollaboratorPostResponse{}
	i.RequiredScope = scopes.IncidentCollaboratorPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return inc, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return inc, err
	}

	err = json.Unmarshal(resp, &inc)
	if err != nil {
		return inc, err
	}

	if inc.Status == "ERROR" {
		return inc, fmt.Errorf("invgate returned a status of %s when adding collaborator (id: %d or ids: %d) ", inc.Status, p.UserID, p.UsersID)
	}

	return inc, nil
}

type (
	// IncidentCommentMethods is use to call methods for IncidentComment
	IncidentCommentMethods struct{ methods.MethodCall }

	IncidentCommentGetParams struct {
		RequestID int `url:"request_id,required"`
		// Indicate the date format. The available formats are 'epoch' or 'iso8601'.
		// If null, epoch format is returned.
		DateFormat               string `url:"date_format"`
		IsSolution               bool   `url:"is_solution"`
		DecodedSpecialCharacters bool   `url:"decoded_special_characters"`
	}

	// IncidentCommentGetResponse is used to map an comment returned from the Invgate API
	IncidentCommentGetResponse struct {
		AuthorID        int    `json:"author_id,omitempty"`
		Reference       int    `json:"reference,omitempty"`
		IsSolution      bool   `json:"is_solution,omitempty"`
		ID              int    `json:"id,omitempty"`
		CreatedAt       int    `json:"created_at,omitempty"`
		CustomerVisible bool   `json:"customer_visible,omitempty"`
		Attachments     []int  `json:"attached_files,omitempty"`
		MsgNum          int    `json:"msg_num,omitempty"`
		IncidentID      int    `json:"incident_id,omitempty"`
		Message         string `json:"message,omitempty"`
	}
)

// Get for IncidentComment
// Requires scope: IncidentCommentGet
// See https://releases.invgate.com/service-desk/api/#incidentcomment-GET
func (i *IncidentCommentMethods) Get(p IncidentCommentGetParams) ([]IncidentCommentGetResponse, error) {
	comms := []IncidentCommentGetResponse{}
	i.RequiredScope = scopes.IncidentCommentGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return comms, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return comms, err
	}

	err = json.Unmarshal(resp, &comms)
	if err != nil {
		return nil, err
	}
	return comms, nil
}

type (
	IncidentCommentPostParams struct {
		IsSolution      bool   `url:"is_solution"`
		AuthorID        int    `url:"author_id,required"`
		RequestID       int    `url:"request_id,required"`
		CustomerVisible bool   `url:"customer_visible"`
		Comment         string `url:"comment,required"`
		IsPropagation   bool   `url:"is_propagation"`
		Attachments     []int  `url:"attached_files"`
	}

	// IncidentCommentPostResponse is used to map an comment post response returned from the Invgate API
	IncidentCommentPostResponse struct {
		// OK if comment was added, ERROR if something went wrong
		Status string `json:"status"`
		Error  string `json:"error"`
	}
)

// Post for IncidentComment
// Requires scope: IncidentCommentPost
// See https://releases.invgate.com/service-desk/api/#incidentcomment-POST
func (i *IncidentCommentMethods) Post(p IncidentCommentPostParams) (IncidentCommentPostResponse, error) {
	com := IncidentCommentPostResponse{}
	i.RequiredScope = scopes.IncidentCommentPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return com, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return com, err
	}

	err = json.Unmarshal(resp, &com)
	if err != nil {
		return com, err
	}
	return com, nil
}

type (
	// IncidentCustomApprovalMethods is use to call methods for IncidentCustomApproval
	IncidentCustomApprovalMethods struct{ methods.MethodCall }

	IncidentCustomApprovalGetParams struct {
		// Indicate the date format. The available formats are 'epoch' or 'iso8601'.
		// If null, epoch format is returned.
		DateFormat string `url:"date_format"`
		RequestID  int    `url:"request_id,required"`
	}

	// IncidentCustomApprovalGetResponse is used to map an custom approval returned from the Invgate API
	IncidentCustomApprovalGetResponse struct {
		ExpiredIn           int    `json:"expired_in,omitempty"`
		Title               string `json:"title,omitempty"`
		DescriptionPrompt   string `json:"description_prompt,omitempty"`
		Status              int    `json:"status,omitempty"`
		WfItemID            int    `json:"wf_item_id,omitempty"`
		CreatedAt           int    `json:"created_at,omitempty"`
		ExpiredApproved     int    `json:"expired_approved,omitempty"`
		WfProcessID         int    `json:"wf_process_id,omitempty"`
		PauseSLA            bool   `json:"pause_sla,omitempty"`
		Description         string `json:"description,omitempty"`
		ID                  int    `json:"id,omitempty"`
		DescriptionRequired bool   `json:"description_required,omitempty"`
	}
)

// Get for IncidentCustomApproval
// Requires scope: IncidentCustomApprovalGet
// See https://releases.invgate.com/service-desk/api/#incidentcustom_approval-GET
func (i *IncidentCustomApprovalMethods) Get(p IncidentCustomApprovalGetParams) ([]IncidentCustomApprovalGetResponse, error) {
	cust := []IncidentCustomApprovalGetResponse{}
	i.RequiredScope = scopes.IncidentCustomApprovalGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

type (
	IncidentCustomApprovalPostParams struct {
		AuthorID    int    `url:"author_id,required"`
		Description string `url:"description"`
		RequestID   int    `url:"request_id,required"`
		ApprovalID  int    `url:"approval_id,required"`
	}

	// IncidentCustomApprovalPostResponse is used to map an custom approval returned from the Invgate API
	IncidentCustomApprovalPostResponse struct {
		// OK if custom approval was added, ERROR if something went wrong
		Status string `json:"status"`
		Error  string `json:"error"`
	}
)

// Post for IncidentCustomApproval
// Requires scope: IncidentCustomApprovalPost
// See https://releases.invgate.com/service-desk/api/#incidentcustom_approval-POST
func (i *IncidentCustomApprovalMethods) Post(p IncidentCustomApprovalPostParams) (IncidentCustomApprovalPostResponse, error) {
	cust := IncidentCustomApprovalPostResponse{}
	i.RequiredScope = scopes.IncidentCustomApprovalPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when adding custom approval (id: %d) ", cust.Status, p.ApprovalID)
	}

	return cust, nil
}

type (
	// IncidentExternalEntityMethods is use to call methods for IncidentExternalEntity
	IncidentExternalEntityMethods struct{ methods.MethodCall }

	IncidentExternalEntityGetParams struct {
		RequestID int `url:"request_id,required"`
	}

	// IncidentExternalEntityGetResponse is used to map an external entity returned from the Invgate API
	IncidentExternalEntityGetResponse struct {
		Type     int    `json:"type,omitempty"` // NOTE: Docs say string api returns int
		Name     string `json:"name,omitempty"`
		ExtRefID int    `json:"ext_ref_id,omitempty"`
		RefID    int    `json:"ref_id,omitempty"`
		LinkID   int    `json:"link_id,omitempty"`
		Status   bool   `json:"status,omitempty"`
	}
)

// Get for IncidentExternalEntity
// Requires scope: IncidentExternalEntityGet
// See https://releases.invgate.com/service-desk/api/#incidentexternal_entity-GET
func (i *IncidentExternalEntityMethods) Get(p IncidentExternalEntityGetParams) ([]IncidentExternalEntityGetResponse, error) {
	cust := []IncidentExternalEntityGetResponse{}
	i.RequiredScope = scopes.IncidentExternalEntityGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

type (
	IncidentExternalEntityPostParams struct {
		ExternalEntityID    int    `url:"external_entity_id,required"`
		ExternalEntityRefID string `url:"external_entity_ref_id"`
		RequestID           int    `url:"request_id,required"`
	}

	// IncidentExternalEntityPostResponse is used to map an external entity returned from the Invgate API
	IncidentExternalEntityPostResponse struct {
		Info string `json:"info,omitempty"`
		// OK if external entity was added, ERROR if something went wrong
		Status string `json:"status,omitempty"`
		LinkID string `json:"link_id,omitempty"`
	}
)

// Post for IncidentExternalEntity
// Requires scope: IncidentExternalEntityPost
// See https://releases.invgate.com/service-desk/api/#incidentexternal_entity-POST
func (i *IncidentExternalEntityMethods) Post(p IncidentExternalEntityPostParams) (IncidentExternalEntityPostResponse, error) {
	cust := IncidentExternalEntityPostResponse{}
	i.RequiredScope = scopes.IncidentExternalEntityPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when adding external entity (id: %d) ", cust.Status, p.RequestID)
	}

	return cust, nil
}

type (
	// IncidentLinkMethods is use to call methods for IncidentLink
	IncidentLinkMethods struct{ methods.MethodCall }

	IncidentLinkGetParams struct {
		RequestID int `url:"request_id,required"`
	}

	// IncidentLinkGetResponse is used to map an incident link returned from the Invgate API
	IncidentLinkGetResponse struct {
		ID    int    `json:"id,omitempty"`
		Title string `json:"title,omitempty"`
	}
)

// Get for IncidentLink
// Requires scope: IncidentLinkGet
// See https://releases.invgate.com/service-desk/api/#incidentlink-GET
func (i *IncidentLinkMethods) Get(p IncidentLinkGetParams) ([]IncidentLinkGetResponse, error) {
	cust := []IncidentLinkGetResponse{}
	i.RequiredScope = scopes.IncidentLinkGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

type (
	IncidentLinkPostParams struct {
		// RequestIDs the ids of all requests to be linked
		RequestIDs []int `url:"request_ids,required"`
		// RequestID the id of the request being linked to
		RequestID int `url:"request_id,required"`
	}

	// IncidentLinkPostResponse is used to map an inicdent link returned from the Invgate API
	IncidentLinkPostResponse struct {
		// OK if incident(s) were linked, ERROR if something went wrong
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
		Info   string `json:"info,omitempty"`
	}
)

// Post for IncidentLink
// Requires scope: IncidentLinkPost
// See https://releases.invgate.com/service-desk/api/#incidentlink-POST
func (i *IncidentLinkMethods) Post(p IncidentLinkPostParams) (IncidentLinkPostResponse, error) {
	cust := IncidentLinkPostResponse{}
	i.RequiredScope = scopes.IncidentLinkPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when linking incident(s) (ids: %d) to (id: %d) ", cust.Status, p.RequestIDs, p.RequestID)
	}

	return cust, nil
}

type (
	// IncidentLinkedCIsCountersFromMethods is use to call methods for IncidentLinkedCIsCountersFrom
	IncidentLinkedCIsCountersFromMethods struct{ methods.MethodCall }

	IncidentLinkedCIsCountersFromGetParams struct {
		From        int `url:"from,required"`
		CIsSourceID int `url:"cis_source_id,required"`
	}

	// IncidentLinkedCIsCountersFromGetResponse is used to map an incidents linked CIs counters returned from the Invgate API
	IncidentLinkedCIsCountersFromGetResponse struct {
		Group    string `json:"group,omitempty"`
		Requests any    `json:"requests,omitempty"` // NOTE: I am not currently sure exactly what this returns
		CiID     int    `json:"ci_id,omitempty"`
	}
)

// Get for IncidentLinkedCIsCounterFrom
// Requires scope: IncidentLinkedCIsCountersFromGet
// See https://releases.invgate.com/service-desk/api/#incidentlinked_ciscountersfrom-GET
func (i *IncidentLinkedCIsCountersFromMethods) Get(p IncidentLinkedCIsCountersFromGetParams) ([]IncidentLinkedCIsCountersFromGetResponse, error) {
	cust := []IncidentLinkedCIsCountersFromGetResponse{}
	i.RequiredScope = scopes.IncidentLinkedCIsCountersFromGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return nil, err
	}
	return cust, nil
}

type (
	// IncidentObserverMethods is use to call methods for IncidentObserver
	IncidentObserverMethods struct{ methods.MethodCall }

	IncidentObserverGetParams struct {
		RequestID int `url:"request_id,required"`
	}

	// IncidentObserverGetResponse ...
	// NOTE: Invgate returns an array of integers
	// with no json key value pairs. To make this easier they are mapped into the
	// UserIDs slice of ints.
	IncidentObserverGetResponse struct {
		UserIDs []int `json:"user_ids,omitempty"`
	}
)

// Get for IncidentObserver
// Requires scope: IncidentObserverGet
// See https://releases.invgate.com/service-desk/api/#incidentobserver-GET
func (i *IncidentObserverMethods) Get(p IncidentObserverGetParams) (IncidentObserverGetResponse, error) {
	var r IncidentObserverGetResponse
	i.RequiredScope = scopes.IncidentObserverGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return r, err
	}

	var b []int
	err = json.Unmarshal(resp, &b)
	if err != nil {
		return r, err
	}

	r.UserIDs = append(r.UserIDs, b...)

	return r, nil
}

type (
	// IncidentObserverPostParams is used to construct a new POST request to add incident observers
	// NOTE: While Invgate accepts a user_id or users_id param, Invgo only uses the users_id field.
	// This field is also required. This was done to simplify calls and make it easier to understand,
	// since an array can by used for a sigle user this doesn't break the Invgate API.
	IncidentObserverPostParams struct {
		UsersID   []int `url:"users_id,required"`
		AuthorID  int   `url:"author_id,required"`
		RequestID int   `url:"request_id,required"`
	}

	// IncidentObserverPostResponse is used to map the response after adding observer(s) to an incident
	IncidentObserverPostResponse struct {
		// OK if observers were added, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Post for IncidentObserver
// Requires scope: IncidentObserverPost
// See https://releases.invgate.com/service-desk/api/#incidentobserver-POST
func (i *IncidentObserverMethods) Post(p IncidentObserverPostParams) (IncidentObserverPostResponse, error) {
	i.RequiredScope = scopes.IncidentObserverPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return IncidentObserverPostResponse{}, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return IncidentObserverPostResponse{}, err
	}

	var r IncidentObserverPostResponse
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	if r.Status == "ERROR" {
		return r, fmt.Errorf("invgate returned a status of %s when adding observer(s) to request (id: %d) ", r.Status, p.RequestID)
	}

	return r, nil
}

type (
	// IncidentReassignMethods is use to call methods for IncidentReassign
	IncidentReassignMethods struct{ methods.MethodCall }

	IncidentReassignPostParams struct {
		AgentID   int `url:"agent_id"`
		AuthorID  int `url:"author_id,required"`
		GroupID   int `url:"group_id,required"`
		RequestID int `url:"request_id,required"`
	}

	// IncidentReassignPostResponse is used to map an incident reassignment returned from the Invgate API
	IncidentReassignPostResponse struct {
		// OK if incident was reassigned, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Post for IncidentReassign
// Requires scope: IncidentReassignPost
// See https://releases.invgate.com/service-desk/api/#incidentreassign-POST
func (i *IncidentReassignMethods) Post(p IncidentReassignPostParams) (IncidentReassignPostResponse, error) {
	cust := IncidentReassignPostResponse{}
	i.RequiredScope = scopes.IncidentReassignPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when re-assigning request (id: %d) ", cust.Status, p.RequestID)
	}

	return cust, nil
}

type (
	// IncidentRejectMethods is use to call methods for IncidentReject
	IncidentRejectMethods struct{ methods.MethodCall }

	IncidentRejectPostParams struct {
		AuthorID  int `url:"author_id,required"`
		RequestID int `url:"request_id,required"`
	}

	// IncidentRejectPostResponse is used to map an external rejection returned from the Invgate API
	IncidentRejectPostResponse struct {
		// OK if incident was rejected, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Post for IncidentReject
// Requires scope: IncidentRejectPost
// See https://releases.invgate.com/service-desk/api/#incidentreject-POST
func (i *IncidentRejectMethods) Post(p IncidentRejectPostParams) (IncidentRejectPostResponse, error) {
	cust := IncidentRejectPostResponse{}
	i.RequiredScope = scopes.IncidentRejectPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when rejecting request (id: %d) ", cust.Status, p.RequestID)
	}

	return cust, nil
}

type (
	// IncidentSpontaneousApprovalMethods is use to call methods for IncidentSpontaneousApproval
	IncidentSpontaneousApprovalMethods struct{ methods.MethodCall }

	// IncidentSpontaneousApprovalPostParams is used to construct a new POST request for incident spontaneous aprroval
	IncidentSpontaneousApprovalPostParams struct {
		ApprovalUserID int    `url:"approval_user_id,required"`
		AuthorID       int    `url:"author_id,required"`
		Description    string `url:"description,required"`
		RequestID      int    `url:"request_id,required"`
	}

	// IncidentSpontaneousApprovalPostResponse is used to map the response after posting a incidents spontaneous approval
	IncidentSpontaneousApprovalPostResponse struct {
		// OK if the approval was created, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Post for IncidentSpontaneousApproval
// Requires scope: IncidentSpontaneousApprovalPost
// See https://releases.invgate.com/service-desk/api/#incidentspontaneous_approval-POST
func (i *IncidentSpontaneousApprovalMethods) Post(p IncidentSpontaneousApprovalPostParams) (IncidentSpontaneousApprovalPostResponse, error) {
	i.RequiredScope = scopes.IncidentSpontaneousApprovalPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return IncidentSpontaneousApprovalPostResponse{}, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return IncidentSpontaneousApprovalPostResponse{}, err
	}

	var r IncidentSpontaneousApprovalPostResponse
	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	if r.Status == "ERROR" {
		return r, fmt.Errorf("invgate returned a status of %s when adding observer(s) to request (id: %d) ", r.Status, p.RequestID)
	}

	return r, nil
}

type (
	// IncidentTasksMethods is use to call methods for IncidentTasks
	IncidentTasksMethods struct{ methods.MethodCall }

	IncidentTasksGetParams struct {
		RequestID int `url:"request_id,required"`
	}

	IncidentTasksGetResponse struct {
		Name            string `json:"name,omitempty"`
		CreateAt        int    `json:"create_at,omitempty"`
		IsPredefined    bool   `json:"is_predefined,omitempty"`
		WfStagedID      string `json:"wf_staged_id,omitempty"`
		Status          int    `json:"status,omitempty"`
		Description     string `json:"description,omitempty"`
		AgentID         int    `json:"agent_id,omitempty"`
		HelpdeskID      int    `json:"helpdesk_id,omitempty"`
		IsRequired      bool   `json:"is_required,omitempty"`
		AssignmentType  int    `json:"assignment_type,omitempty"`
		LinkedRequestID int    `json:"linked_request_id,omitempty"`
		TaskID          int    `json:"task_id,omitempty"`
		ExpirationDate  int    `json:"expiration_date,omitempty"`
		CompletedAt     int    `json:"completed_at,omitempty"`
	}
)

// Get for IncidentTasks
// Requires scope: IncidentTasksGet
// See https://releases.invgate.com/service-desk/api/#incidenttasks-GET
func (i *IncidentTasksMethods) Get(p IncidentTasksGetParams) ([]IncidentTasksGetResponse, error) {
	var r []IncidentTasksGetResponse
	i.RequiredScope = scopes.IncidentTasksGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemoteGet()
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	return r, nil
}

type (
	// IncidentReopenMethods is use to call methods for IncidentReopen
	IncidentReopenMethods struct{ methods.MethodCall }

	IncidentReopenPutParams struct {
		AuthorID  int `url:"author_id"`
		RequestID int `url:"request_id,required"`
	}

	// IncidentReopenPutResponse is used to map an re-opening returned from the Invgate API
	IncidentReopenPutResponse struct {
		Info string `json:"info,omitempty"`
		// OK if incident was reopened, ERROR if something went wrong
		Status string `json:"status,omitempty"`
		Error  string `json:"error,omitempty"`
	}
)

// Put for IncidentReopen
// Requires scope: IncidentReopenPut
// See https://releases.invgate.com/service-desk/api/#incidentreopen-PUT
func (i *IncidentReopenMethods) Put(p IncidentReopenPutParams) (IncidentReopenPutResponse, error) {
	cust := IncidentReopenPutResponse{}
	i.RequiredScope = scopes.IncidentReopenPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when re-opening request (id: %d) ", cust.Status, p.RequestID)
	}

	return cust, nil
}

type (
	// IncidentSolutionAcceptMethods is use to call methods for IncidentSolutionAccept
	IncidentSolutionAcceptMethods struct{ methods.MethodCall }

	IncidentSolutionAcceptPutParams struct {
		ID      int    `url:"id,required"`
		Rating  int    `url:"rating,required"`
		Comment string `url:"comment"`
	}

	// IncidentSolutionAcceptPutResponse is used to map an solution acceptance returned from the Invgate API
	IncidentSolutionAcceptPutResponse struct {
		// OK if solution was accepted, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Put for IncidentSolutionAccept
// Requires scope: IncidentSolutionAcceptPut
// See https://releases.invgate.com/service-desk/api/#incidentsolutionaccept-PUT
func (i *IncidentSolutionAcceptMethods) Put(p IncidentSolutionAcceptPutParams) (IncidentSolutionAcceptPutResponse, error) {
	cust := IncidentSolutionAcceptPutResponse{}
	i.RequiredScope = scopes.IncidentSolutionAcceptPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when accepting solution (id: %d) ", cust.Status, p.ID)
	}

	return cust, nil
}

type (
	// IncidentSolutionRejectMethods is use to call methods for IncidentSolutionReject
	IncidentSolutionRejectMethods struct{ methods.MethodCall }

	IncidentSolutionRejectPutParams struct {
		ID      int    `url:"id,required"`
		Comment string `url:"comment"`
	}

	// IncidentSolutionRejectPutResponse is used to map an solution rejection returned from the Invgate API
	IncidentSolutionRejectPutResponse struct {
		// OK if solution was accepted, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Put for IncidentSolutionReject
// Requires scope: IncidentSolutionRejectPut
// See https://releases.invgate.com/service-desk/api/#incidentsolutionareject-PUT
func (i *IncidentSolutionRejectMethods) Put(p IncidentSolutionRejectPutParams) (IncidentSolutionRejectPutResponse, error) {
	cust := IncidentSolutionRejectPutResponse{}
	i.RequiredScope = scopes.IncidentSolutionRejectPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePut()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when rejecting solution (id: %d) ", cust.Status, p.ID)
	}

	return cust, nil
}

type (
	// IncidentWaitingForExternalEntityMethods is use to call methods for IncidentWaitingForExternalEntity
	IncidentWaitingForExternalEntityMethods struct{ methods.MethodCall }

	IncidentWaitingForExternalEntityPostParams struct {
		RequestID    int `url:"request_id,required"`
		EntityLinkID int `url:"entity_link_id,required"`
	}

	// IncidentWaitingForExternalEntityPostResponse is used to map an external entity wait for returned from the Invgate API
	IncidentWaitingForExternalEntityPostResponse struct {
		Info string `json:"info,omitempty"`
		// OK if comment was added, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}
)

// Post for IncidentWaitingForExternalEntity
// Requires scope: IncidentWaitingForExternalEntityPost
// See https://releases.invgate.com/service-desk/api/#incidentwaitingforexternal_entity-POST
func (i *IncidentWaitingForExternalEntityMethods) Post(p IncidentWaitingForExternalEntityPostParams) (IncidentWaitingForExternalEntityPostResponse, error) {
	cust := IncidentWaitingForExternalEntityPostResponse{}
	i.RequiredScope = scopes.IncidentWaitingForExternalEntityPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return cust, err
	}
	i.Endpoint.RawQuery = q.Encode()

	resp, err := i.RemotePost()
	if err != nil {
		return cust, err
	}

	err = json.Unmarshal(resp, &cust)
	if err != nil {
		return cust, err
	}

	if cust.Status == "ERROR" {
		return cust, fmt.Errorf("invgate returned a status of %s when adding external entity (id: %d) ", cust.Status, p.RequestID)
	}

	return cust, nil
}

// IncidentsMethods is used to call methods for Incidents
type IncidentsMethods struct{ methods.MethodCall }

type IncidentsGetParams struct {
	IDs             []int  `url:"ids,required"`
	IncludeComments bool   `url:"comments"`
	DateFormat      string `url:"date_format"`
}

// Get for Incidents
// Requires scope: IncidentsGet
// At least one incident must be provided
// See https://releases.invgate.com/service-desk/api/#incidents-GET
func (i *IncidentsMethods) Get(p IncidentsGetParams) ([]Incident, error) {
	i.RequiredScope = scopes.IncidentsGet

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

// Get for IncidentsByStatus
// Requires scope: IncidentsByStatusGet
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus-GET
func (i *IncidentsByStatusMethods) Get(p IncidentsByStatusGetParams) (IncidentsByStatusResponse, error) {
	i.RequiredScope = scopes.IncidentsByStatusGet

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
