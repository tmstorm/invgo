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

func TestUserByGet(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserByGetResponse
	gofakeit.Struct(&u)

	server := newTestServer(t, http.MethodGet, "/user.by", u)

	c := newTestClient(t, server, scopes.UserByGet)

	got, err := c.UserBy().Get(endpoints.UserByGetParams{Username: "john"})
	a.NoError(err)
	a.EqualValues(u, got)
}

func TestUserConvertPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserConvertPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/user.convert", u)

	c := newTestClient(t, server, scopes.UserConvertPost)

	got, err := c.UserConvert().Post(endpoints.UserConvertPostParams{ID: 1})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/user.convert", u)

	cErr := newTestClient(t, serverErr, scopes.UserConvertPost)
	gotErr, err := cErr.UserConvert().Post(endpoints.UserConvertPostParams{ID: 1})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserDisablePut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserDisablePutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/user.disable", u)

	c := newTestClient(t, server, scopes.UserDisablePut)

	got, err := c.UserDisable().Put(endpoints.UserDisablePutParams{ID: 1})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/user.disable", u)

	cErr := newTestClient(t, serverErr, scopes.UserDisablePut)
	gotErr, err := cErr.UserDisable().Put(endpoints.UserDisablePutParams{ID: 1})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserEnablePut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserEnablePutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/user.enable", u)

	c := newTestClient(t, server, scopes.UserEnablePut)

	got, err := c.UserEnable().Put(endpoints.UserEnablePutParams{ID: 1})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/user.enable", u)

	cErr := newTestClient(t, serverErr, scopes.UserEnablePut)
	gotErr, err := cErr.UserEnable().Put(endpoints.UserEnablePutParams{ID: 1})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserPasswordPut(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserPasswordPutResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPut, "/user.password", u)

	c := newTestClient(t, server, scopes.UserPasswordPut)

	got, err := c.UserPassword().Put(endpoints.UserPasswordPutParams{ID: 1, Password: "somenewpass"})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPut, "/user.password", u)

	cErr := newTestClient(t, serverErr, scopes.UserPasswordPut)
	gotErr, err := cErr.UserPassword().Put(endpoints.UserPasswordPutParams{ID: 1, Password: "somenewpass"})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserPasswordResetPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserPasswordResetPostResponse
	u.Status = "OK"

	server := newTestServer(t, http.MethodPost, "/user.password.reset", u)

	c := newTestClient(t, server, scopes.UserPasswordResetPost)

	got, err := c.UserPasswordReset().Post(endpoints.UserPasswordResetPostParams{ID: 1, Type: "NEW_USER"})
	a.NoError(err)
	a.Equal(u.Status, got.Status)

	u.Status = "ERROR"
	serverErr := newTestServer(t, http.MethodPost, "/user.password.reset", u)

	cErr := newTestClient(t, serverErr, scopes.UserPasswordResetPost)
	gotErr, err := cErr.UserPasswordReset().Post(endpoints.UserPasswordResetPostParams{ID: 1, Type: "RESET_PASSWORD"})
	a.Error(err)
	a.Equal(u.Status, gotErr.Status)
}

func TestUserTokenPost(t *testing.T) {
	a := assert.New(t)
	var u endpoints.UserTokenPostResponse
	gofakeit.Struct(&u)

	server := newTestServer(t, http.MethodPost, "/user.token", u)

	c := newTestClient(t, server, scopes.UserTokenPost)

	got, err := c.UserToken().Post(endpoints.UserTokenPostParams{ID: 1})
	a.NoError(err)
	a.Equal(u, got)
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

func TestUsersByGet(t *testing.T) {
	a := assert.New(t)
	users := endpoints.UsersByGetResponse{}
	gofakeit.Struct(&users)

	server := newTestServer(t, http.MethodGet, "/users.by", users)

	c := newTestClient(t, server, scopes.UsersByGet)

	got, err := c.UsersBy().Get(endpoints.UsersByGetParams{Username: "johndoe@example.com"})
	a.NoError(err)
	a.EqualValues(users, got)
}

func TestUsersGroupsGet(t *testing.T) {
	a := assert.New(t)
	users := []endpoints.UsersGroupsGetResponse{}
	var u endpoints.UsersGroupsGetResponse
	gofakeit.Struct(&u)
	users = append(users, u)

	server := newTestServer(t, http.MethodGet, "/users.groups", users)

	c := newTestClient(t, server, scopes.UsersGroupsGet)

	got, err := c.UsersGroups().Get(endpoints.UsersGroupsGetParams{IDs: []int{u.ID}})
	a.NoError(err)
	a.EqualValues(users, got)
}
