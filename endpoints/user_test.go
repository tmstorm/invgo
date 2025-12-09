package endpoints_test

import (
	"net/http"
	"testing"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/scopes"
)

func TestUserGet(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserGetResponse
	gofakeit.Struct(&u)

	server := newTestServer(t, http.MethodGet, "/user", u)

	c := newTestClient(t, server, scopes.UserGet)

	got, err := c.User().Get(endpoints.UserGetParams{ID: u.ID})
	a.NoError(err)
	a.EqualValues(u, got)
}

func TestUserPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserPostResponse
	gofakeit.Struct(&u)

	server := newTestServer(t, http.MethodPost, "/user", u)

	c := newTestClient(t, server, scopes.UserPost)

	got, err := c.User().Post(endpoints.UserPostParams{Email: u.Email, Name: u.Name, LastName: u.LastName})
	a.NoError(err)
	a.EqualValues(u, got)
}

func TestUserPut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserPutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/user", u)

	c := newTestClient(t, server, scopes.UserPut)

	got, err := c.User().Put(endpoints.UserPutParams{ID: 1})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/user", u)

	cErr := newTestClient(t, serverErr, scopes.UserPut)
	gotErr, err := cErr.User().Put(endpoints.UserPutParams{ID: 1})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserDelete(t *testing.T) {
	a := assert.New(t)
	users := []endpoints.UserDeleteResponse{}
	var u endpoints.UserDeleteResponse
	gofakeit.Struct(&u)
	users = append(users, u)

	server := newTestServer(t, http.MethodDelete, "/user", users)

	c := newTestClient(t, server, scopes.UserDelete)

	got, err := c.User().Delete(endpoints.UserDeleteParams{ID: 1, Users: []int{u.ID}})
	a.NoError(err)
	a.EqualValues(users, got)
}

func TestUsersGet(t *testing.T) {
	a := assert.New(t)
	users := []endpoints.UsersGetResponse{}
	var u endpoints.UsersGetResponse
	gofakeit.Struct(&u)
	users = append(users, u)

	server := newTestServer(t, http.MethodGet, "/users", users)

	c := newTestClient(t, server, scopes.UsersGet)

	got, err := c.Users().Get(endpoints.UsersGetParams{IDs: []int{u.ID}})
	a.NoError(err)
	a.EqualValues(users, got)
}
