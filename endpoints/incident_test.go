package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestIncidentGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.Incident, 1)
	var inc endpoints.Incident
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident", incs)

	c := newTestClient(t, server, scopes.IncidentGet)

	resp, err := c.Incident().Get(endpoints.IncidentGetParams{ID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentPost(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentPostResponse
	var inc endpoints.IncidentPostParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPost, "/incident", r)

	c := newTestClient(t, server, scopes.IncidentPost)

	resp, err := c.Incident().Post(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentPut(t *testing.T) {
	a := assert.New(t)

	var put endpoints.IncidentPutParams
	gofakeit.Struct(&put)

	incs := make([]endpoints.Incident, 1)
	var inc endpoints.Incident
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodPut, "/incident", incs)

	c := newTestClient(t, server, scopes.IncidentPut)

	resp, err := c.Incident().Put(endpoints.IncidentPutParams{ID: put.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentApprovalGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.IncidentApprovalGetResponse, 1)
	var inc endpoints.IncidentApprovalGetResponse
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident.approval", incs)

	c := newTestClient(t, server, scopes.IncidentApprovalGet)

	resp, err := c.IncidentApproval().Get(endpoints.IncidentApprovalGetParams{RequestID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentApprovalAcceptPut(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentApprovalAcceptPutResponse
	var inc endpoints.IncidentApprovalAcceptPutParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPut, "/incident.approval.accept", r)

	c := newTestClient(t, server, scopes.IncidentApprovalAcceptPut)

	resp, err := c.IncidentApprovalAccept().Put(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentApprovalAddVotersPost(t *testing.T) {
	a := assert.New(t)

	var inc endpoints.IncidentApprovalAddVoterPostResponse
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPost, "/incident.approval.add_voter", inc)

	c := newTestClient(t, server, scopes.IncidentApprovalAddVoterPost)

	resp, err := c.IncidentApprovalAddVoter().Post(endpoints.IncidentApprovalAddVoterPostParams{UserID: inc.UserID, ApprovalID: inc.ApprovalID})
	a.NoError(err)
	a.Equal(inc, resp)
}

func TestIncidentApprovalCancelPut(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentApprovalCancelPutResponse
	var inc endpoints.IncidentApprovalCancelPutParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPut, "/incident.approval.cancel", r)

	c := newTestClient(t, server, scopes.IncidentApprovalCancelPut)

	resp, err := c.IncidentApprovalCancel().Put(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentApprovalPossibleVotersGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.IncidentApprovalPossibleVotersGetResponse, 1)
	var inc endpoints.IncidentApprovalPossibleVotersGetResponse
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident.approval.possible_voters", incs)

	c := newTestClient(t, server, scopes.IncidentApprovalPossibleVotersGet)

	resp, err := c.IncidentApprovalPossibleVoters().Get(endpoints.IncidentApprovalPossibleVotersGetParams{ApprovalID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentCancelPost(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentCancelPostResponse
	var inc endpoints.IncidentCancelPostParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPost, "/incident.cancel", r)

	c := newTestClient(t, server, scopes.IncidentCancelPost)

	resp, err := c.IncidentCancel().Post(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentCollaboratorGet(t *testing.T) {
	a := assert.New(t)

	incs := []int{1, 2, 3}

	server := newTestServer(t, http.MethodGet, "/incident.collaborator", incs)

	c := newTestClient(t, server, scopes.IncidentCollaboratorGet)

	resp, err := c.IncidentCollaborator().Get(endpoints.IncidentCollaboratorGetParams{RequestID: 1})
	a.NoError(err)
	a.Equal(incs, resp.IDs)
}

func TestIncidentCollaboratorPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentCollaboratorPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.collaborator", u)

	c := newTestClient(t, server, scopes.IncidentCollaboratorPost)

	got, err := c.IncidentCollaborator().Post(endpoints.IncidentCollaboratorPostParams{AuthorID: 1, UserID: 2, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.collaborator", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentCollaboratorPost)
	gotErr, err := cErr.IncidentCollaborator().Post(endpoints.IncidentCollaboratorPostParams{AuthorID: 1, UserID: 2, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentCommentGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.IncidentCommentGetResponse, 1)
	var inc endpoints.IncidentCommentGetResponse
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident.comment", incs)

	c := newTestClient(t, server, scopes.IncidentCommentGet)

	resp, err := c.IncidentComment().Get(endpoints.IncidentCommentGetParams{RequestID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentCustomApprovalGet(t *testing.T) {
	a := assert.New(t)

	incs := make([]endpoints.IncidentCustomApprovalGetResponse, 1)
	var inc endpoints.IncidentCustomApprovalGetResponse
	gofakeit.Struct(&inc)
	incs = append(incs, inc)

	server := newTestServer(t, http.MethodGet, "/incident.custom_approval", incs)

	c := newTestClient(t, server, scopes.IncidentCustomApprovalGet)

	resp, err := c.IncidentCustomApproval().Get(endpoints.IncidentCustomApprovalGetParams{RequestID: inc.ID})
	a.NoError(err)
	a.Equal(incs, resp)
}

func TestIncidentCustomApprovalPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentCustomApprovalPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.custom_approval", u)

	c := newTestClient(t, server, scopes.IncidentCustomApprovalPost)

	got, err := c.IncidentCustomApproval().Post(endpoints.IncidentCustomApprovalPostParams{AuthorID: 1, ApprovalID: 2, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.custom_approval", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentCustomApprovalPost)
	gotErr, err := cErr.IncidentCustomApproval().Post(endpoints.IncidentCustomApprovalPostParams{AuthorID: 1, ApprovalID: 2, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentExternalEntityGet(t *testing.T) {
	a := assert.New(t)

	ents := make([]endpoints.IncidentExternalEntityGetResponse, 1)
	var ent endpoints.IncidentExternalEntityGetResponse
	gofakeit.Struct(&ent)
	ents = append(ents, ent)

	server := newTestServer(t, http.MethodGet, "/incident.external_entity", ents)

	c := newTestClient(t, server, scopes.IncidentExternalEntityGet)

	resp, err := c.IncidentExternalEntity().Get(endpoints.IncidentExternalEntityGetParams{RequestID: 1})
	a.NoError(err)
	a.Equal(ents, resp)
}

func TestIncidentExternalEntityPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentExternalEntityPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.external_entity", u)

	c := newTestClient(t, server, scopes.IncidentExternalEntityPost)

	got, err := c.IncidentExternalEntity().Post(endpoints.IncidentExternalEntityPostParams{ExternalEntityID: 1, ExternalEntityRefID: "2", RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.external_entity", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentExternalEntityPost)
	gotErr, err := cErr.IncidentExternalEntity().Post(endpoints.IncidentExternalEntityPostParams{ExternalEntityID: 1, ExternalEntityRefID: "2", RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentLinkGet(t *testing.T) {
	a := assert.New(t)

	ents := make([]endpoints.IncidentLinkGetResponse, 1)
	var ent endpoints.IncidentLinkGetResponse
	gofakeit.Struct(&ent)
	ents = append(ents, ent)

	server := newTestServer(t, http.MethodGet, "/incident.link", ents)

	c := newTestClient(t, server, scopes.IncidentLinkGet)

	resp, err := c.IncidentLink().Get(endpoints.IncidentLinkGetParams{RequestID: 1})
	a.NoError(err)
	a.Equal(ents, resp)
}

func TestIncidentLinkPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentLinkPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.link", u)

	c := newTestClient(t, server, scopes.IncidentLinkPost)

	got, err := c.IncidentLink().Post(endpoints.IncidentLinkPostParams{RequestID: 101, RequestIDs: []int{1, 2}})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.link", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentLinkPost)
	gotErr, err := cErr.IncidentLink().Post(endpoints.IncidentLinkPostParams{RequestID: 101, RequestIDs: []int{1, 2}})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentLinkedCIsCountersFromGet(t *testing.T) {
	a := assert.New(t)

	ents := make([]endpoints.IncidentLinkedCIsCountersFromGetResponse, 1)
	var ent endpoints.IncidentLinkedCIsCountersFromGetResponse
	gofakeit.Struct(&ent)
	ents = append(ents, ent)

	server := newTestServer(t, http.MethodGet, "/incident.linked_cis.counters.from", ents)

	c := newTestClient(t, server, scopes.IncidentLinkedCIsCountersFromGet)

	resp, err := c.IncidentLinkedCIsCountersFrom().Get(endpoints.IncidentLinkedCIsCountersFromGetParams{From: 1, CIsSourceID: 2})
	a.NoError(err)
	a.Equal(ents, resp)
}

func TestIncidentObserverGet(t *testing.T) {
	a := assert.New(t)

	var b endpoints.IncidentObserverGetResponse
	gofakeit.Struct(&b)

	server := newTestServer(t, http.MethodGet, "/incident.observer", b.UserIDs)

	c := newTestClient(t, server, scopes.IncidentObserverGet)

	resp, err := c.IncidentObserver().Get(endpoints.IncidentObserverGetParams{RequestID: 1})
	a.NoError(err)
	a.Equal(b, resp)
}

func TestIncidentObserverPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentObserverPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.observer", u)

	c := newTestClient(t, server, scopes.IncidentObserverPost)

	got, err := c.IncidentObserver().Post(endpoints.IncidentObserverPostParams{UsersID: []int{1, 2}, AuthorID: 3, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.observer", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentObserverPost)
	gotErr, err := cErr.IncidentObserver().Post(endpoints.IncidentObserverPostParams{UsersID: []int{1, 2}, AuthorID: 3, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentReassignPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentReassignPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.reassign", u)

	c := newTestClient(t, server, scopes.IncidentReassignPost)

	got, err := c.IncidentReassign().Post(endpoints.IncidentReassignPostParams{AuthorID: 1, GroupID: 2, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.reassign", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentReassignPost)
	gotErr, err := cErr.IncidentReassign().Post(endpoints.IncidentReassignPostParams{AuthorID: 1, GroupID: 2, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentRejectPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentRejectPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.reject", u)

	c := newTestClient(t, server, scopes.IncidentRejectPost)

	got, err := c.IncidentReject().Post(endpoints.IncidentRejectPostParams{AuthorID: 1, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.reject", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentRejectPost)
	gotErr, err := cErr.IncidentReject().Post(endpoints.IncidentRejectPostParams{AuthorID: 1, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentReopenPut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentReopenPutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/incident.reopen", u)

	c := newTestClient(t, server, scopes.IncidentReopenPut)

	got, err := c.IncidentReopen().Put(endpoints.IncidentReopenPutParams{AuthorID: 1, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/incident.reopen", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentReopenPut)
	gotErr, err := cErr.IncidentReopen().Put(endpoints.IncidentReopenPutParams{AuthorID: 1, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentSolutionAcceptPut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentSolutionAcceptPutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/incident.solution.accept", u)

	c := newTestClient(t, server, scopes.IncidentSolutionAcceptPut)

	got, err := c.IncidentSolutionAccept().Put(endpoints.IncidentSolutionAcceptPutParams{ID: 1, Rating: 5})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/incident.solution.accept", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentSolutionAcceptPut)
	gotErr, err := cErr.IncidentSolutionAccept().Put(endpoints.IncidentSolutionAcceptPutParams{ID: 1, Rating: 5})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentSolutionRejectPut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentSolutionRejectPutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/incident.solution.reject", u)

	c := newTestClient(t, server, scopes.IncidentSolutionRejectPut)

	got, err := c.IncidentSolutionReject().Put(endpoints.IncidentSolutionRejectPutParams{ID: 1, Comment: "no"})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/incident.solution.reject", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentSolutionRejectPut)
	gotErr, err := cErr.IncidentSolutionReject().Put(endpoints.IncidentSolutionRejectPutParams{ID: 1, Comment: "no"})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentSpontaneousApprovalPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentSpontaneousApprovalPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.spontaneous_approval", u)

	c := newTestClient(t, server, scopes.IncidentSpontaneousApprovalPost)

	got, err := c.IncidentSpontaneousApproval().Post(endpoints.IncidentSpontaneousApprovalPostParams{ApprovalUserID: 1, AuthorID: 2, Description: "approval test", RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.spontaneous_approval", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentSpontaneousApprovalPost)
	gotErr, err := cErr.IncidentSpontaneousApproval().Post(endpoints.IncidentSpontaneousApprovalPostParams{ApprovalUserID: 1, AuthorID: 2, Description: "approval test", RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentTasksGet(t *testing.T) {
	a := assert.New(t)

	body := make([]endpoints.IncidentTasksGetResponse, 1)
	var i endpoints.IncidentTasksGetResponse
	gofakeit.Struct(&i)
	body = append(body, i)

	server := newTestServer(t, http.MethodGet, "/incident.tasks", body)

	c := newTestClient(t, server, scopes.IncidentTasksGet)

	resp, err := c.IncidentTasks().Get(endpoints.IncidentTasksGetParams{RequestID: 1})
	a.NoError(err)
	a.Equal(body, resp)
}

func TestIncidentWaitingForAgentPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentWaitingForAgentPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.waitingfor.agent", u)

	c := newTestClient(t, server, scopes.IncidentWaitingForAgentPost)

	got, err := c.IncidentWaitingForAgent().Post(endpoints.IncidentWaitingForAgentPostParams{RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.waitingfor.agent", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentWaitingForAgentPost)
	gotErr, err := cErr.IncidentWaitingForAgent().Post(endpoints.IncidentWaitingForAgentPostParams{RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentWaitingForExternalEntityPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.IncidentWaitingForExternalEntityPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/incident.waitingfor.external_entity", u)

	c := newTestClient(t, server, scopes.IncidentWaitingForExternalEntityPost)

	got, err := c.IncidentWaitingForExternalEntity().Post(endpoints.IncidentWaitingForExternalEntityPostParams{EntityLinkID: 1, RequestID: 101})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/incident.waitingfor.external_entity", u)

	cErr := newTestClient(t, serverErr, scopes.IncidentWaitingForExternalEntityPost)
	gotErr, err := cErr.IncidentWaitingForExternalEntity().Post(endpoints.IncidentWaitingForExternalEntityPostParams{EntityLinkID: 1, RequestID: 101})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestIncidentApprovalRejectPut(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentApprovalRejectPutResponse
	var inc endpoints.IncidentApprovalRejectPutParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodPut, "/incident.approval.reject", r)

	c := newTestClient(t, server, scopes.IncidentApprovalRejectPut)

	resp, err := c.IncidentApprovalReject().Put(inc)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentApprovalStatusGet(t *testing.T) {
	a := assert.New(t)

	d := make(map[int]string, 3)
	d[0] = "Desc1"
	d[1] = "Desc2"
	d[2] = "Desc3"

	incs := make([]endpoints.IncidentApprovalStatusGetResponse, 0, len(d))
	for k, v := range d {
		s := endpoints.IncidentApprovalStatusGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	server := newTestServer(t, http.MethodGet, "/incident.approval.status", d)

	c := newTestClient(t, server, scopes.IncidentApprovalStatusGet)

	resp, err := c.IncidentApprovalStatus().Get()
	a.NoError(err)
	for i := range resp {
		a.Contains(incs, resp[i])
	}
}

func TestIncidentApprovalTypeGet(t *testing.T) {
	a := assert.New(t)

	d := make(map[int]string, 3)
	d[0] = "Desc1"
	d[1] = "Desc2"
	d[2] = "Desc3"

	incs := make([]endpoints.IncidentApprovalTypeGetResponse, 0, len(d))
	for k, v := range d {
		s := endpoints.IncidentApprovalTypeGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	server := newTestServer(t, http.MethodGet, "/incident.approval.type", d)

	c := newTestClient(t, server, scopes.IncidentApprovalTypeGet)

	resp, err := c.IncidentApprovalType().Get()
	a.NoError(err)
	for i := range resp {
		a.Contains(incs, resp[i])
	}
}

func TestIncidentApprovalVoteStatusGet(t *testing.T) {
	a := assert.New(t)

	d := make(map[int]string, 3)
	d[0] = "Desc1"
	d[1] = "Desc2"
	d[2] = "Desc3"

	incs := make([]endpoints.IncidentApprovalVoteStatusGetResponse, 0, len(d))
	for k, v := range d {
		s := endpoints.IncidentApprovalVoteStatusGetResponse{
			ID:          k,
			Description: v,
		}
		incs = append(incs, s)
	}

	server := newTestServer(t, http.MethodGet, "/incident.approval.vote_status", d)

	c := newTestClient(t, server, scopes.IncidentApprovalVoteStatusGet)

	resp, err := c.IncidentApprovalVoteStatus().Get()
	a.NoError(err)
	for i := range resp {
		a.Contains(incs, resp[i])
	}
}

func TestIncidentAttachmentGet(t *testing.T) {
	a := assert.New(t)

	var att endpoints.IncidentAttachmentGetResponse
	gofakeit.Struct(&att)

	server := newTestServer(t, http.MethodGet, "/incident.attachment", att)

	c := newTestClient(t, server, scopes.IncidentAttachmentGet)

	resp, err := c.IncidentAttachment().Get(endpoints.IncidentAttachmentGetParams{ID: att.ID})
	a.NoError(err)
	a.Equal(att, resp)
}

func TestIncidentCommentPost(t *testing.T) {
	a := assert.New(t)

	var r endpoints.IncidentCommentPostResponse
	var comm endpoints.IncidentCommentPostParams
	gofakeit.Struct(&r)
	gofakeit.Struct(&comm)

	server := newTestServer(t, http.MethodPost, "/incident.comment", r)

	c := newTestClient(t, server, scopes.IncidentCommentPost)

	resp, err := c.IncidentComment().Post(comm)
	a.NoError(err)
	a.Equal(r, resp)
}

func TestIncidentsGet(t *testing.T) {
	a := assert.New(t)

	var inc endpoints.Incident
	gofakeit.Struct(&inc)

	incsMap := make(map[int]endpoints.Incident, 1)
	incsMap[0] = inc

	incsSlice := make([]endpoints.Incident, 1)
	incsSlice[0] = inc

	server := newTestServer(t, http.MethodGet, "/incidents", incsMap)

	c := newTestClient(t, server, scopes.IncidentsGet)

	resp, err := c.Incidents().Get(endpoints.IncidentsGetParams{IDs: []int{inc.ID}})
	a.NoError(err)
	a.Equal(incsSlice, resp)
}

func TestIncidentsByStatusGet(t *testing.T) {
	a := assert.New(t)

	var inc endpoints.IncidentsByStatusResponse
	gofakeit.Struct(&inc)

	server := newTestServer(t, http.MethodGet, "/incidents.by.status", inc)

	c := newTestClient(t, server, scopes.IncidentsByStatusGet)

	resp, err := c.IncidentsByStatus().Get(endpoints.IncidentsByStatusGetParams{StatusIDs: []int{0}})
	a.NoError(err)
	a.Equal(inc, resp)
}
