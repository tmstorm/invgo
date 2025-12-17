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
