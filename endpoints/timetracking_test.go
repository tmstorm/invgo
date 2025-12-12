package endpoints_test

import (
	"net/http"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestTimeTrackingGet(t *testing.T) {
	a := assert.New(t)
	times := []endpoints.TimeTrackingGetResponse{}
	var time endpoints.TimeTrackingGetResponse
	gofakeit.Struct(&time)
	times = append(times, time)

	server := newTestServer(t, http.MethodGet, "/timetracking", times)

	c := newTestClient(t, server, scopes.TimeTrackingGet)

	got, err := c.TimeTracking().Get(endpoints.TimeTrackingGetParams{From: "2025-01-01"})
	a.NoError(err)
	a.EqualValues(times, got)
}

func TestTimeTrackingPost(t *testing.T) {
	a := assert.New(t)
	var tm endpoints.TimeTrackingPostResponse
	gofakeit.Struct(&tm)

	server := newTestServer(t, http.MethodPost, "/timetracking", tm)

	c := newTestClient(t, server, scopes.TimeTrackingPost)

	got, err := c.TimeTracking().Post(endpoints.TimeTrackingPostParams{UserID: 1, RequestID: 2, To: time.Now().Second()})
	a.NoError(err)
	a.EqualValues(tm, got)
}

func TestTimeTrackingDelete(t *testing.T) {
	a := assert.New(t)
	var time endpoints.TimeTrackingDeleteResponse
	time.Status = "OK"

	server := newTestServer(t, http.MethodDelete, "/timetracking", time)

	c := newTestClient(t, server, scopes.TimeTrackingDelete)

	got, err := c.TimeTracking().Delete(endpoints.TimeTrackingDeleteParams{TimetrackingID: 1, RequestID: 2, UserID: 3})
	a.NoError(err)
	a.Equal(time.Status, got.Status)

	time.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodDelete, "/timetracking", time)

	cErr := newTestClient(t, serverErr, scopes.TimeTrackingDelete)
	gotErr, err := cErr.TimeTracking().Delete(endpoints.TimeTrackingDeleteParams{TimetrackingID: 1, RequestID: 2, UserID: 3})
	a.Error(err)
	a.Equal(time.Status, gotErr.Status)
}

func TestTimeTrackingAttributesCategoryGet(t *testing.T) {
	a := assert.New(t)
	times := []endpoints.TimeTrackingAttributesCategoryGetResponse{}
	var tm endpoints.TimeTrackingAttributesCategoryGetResponse
	gofakeit.Struct(&tm)
	times = append(times, tm)

	server := newTestServer(t, http.MethodGet, "/timetracking.attributes.category", times)

	c := newTestClient(t, server, scopes.TimeTrackingAttributesCategoryGet)

	got, err := c.TimeTrackingAttributesCategory().Get(endpoints.TimeTrackingAttributesCategoryGetParams{ID: tm.ID})
	a.NoError(err)
	a.EqualValues(times, got)
}
