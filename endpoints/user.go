package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// UserMethods is use to call methods for User
	UserMethods struct{ methods.MethodCall }

	// UserBase is used to map an trigger returned from the Invgate API
	UserBase struct {
		Doc            string `json:"doc,omitempty" url:"doc"`
		IsDisabled     bool   `json:"is_disabled,omitempty" url:"is_disabled"`
		ManagerID      int    `json:"manager_id,omitempty" url:"manager_id"`
		Location       string `json:"location,omitempty" url:"location"`
		IsDeleted      bool   `json:"is_deleted,omitempty" url:"is_deleted"`
		Mobile         string `json:"mobile,omitempty" url:"mobile"`
		Country        string `json:"country,omitempty" url:"country"`
		Address        string `json:"address,omitempty" url:"address"`
		Type           int    `json:"type,omitempty" url:"type"`
		City           string `json:"city,omitempty" url:"city"`
		Department     string `json:"department,omitempty" url:"department"`
		RoleName       string `json:"role_name,omitempty" url:"role_name"`
		UserName       string `json:"user_name,omitempty" url:"user_name"`
		Birthday       string `json:"birthday,omitempty" url:"birthday"`
		Position       string `json:"position,omitempty" url:"position"`
		EmployeeNumber string `json:"employee_number,omitempty" url:"employee_number"`
		Phone          string `json:"phone,omitempty" url:"phone"`
		OtherEmail     string `json:"other_email,omitempty" url:"other_email"`
		UserType       int    `json:"user_type,omitempty" url:"user_type"`
		Other          string `json:"other,omitempty" url:"other"`
		IsExternal     bool   `json:"is_external,omitempty" url:"is_external"`
		Fax            string `json:"fax,omitempty" url:"fax"`
		Office         string `json:"office,omitempty" url:"office"`
	}

	UserGetParams struct {
		ID              int  `url:"id,required"`
		IncludeDisabled bool `url:"include_disabled"`
	}

	// UserGetResponse is used to map an user returned from the Invgate API
	UserGetResponse struct {
		ID       int    `json:"id,omitempty"`
		Email    string `json:"email,omitempty"`
		Name     string `json:"name,omitempty"`
		LastName string `json:"last_name,omitempty"`
		UserBase
	}
)

// Get for User
// Requires scope: UserGet
// See https://releases.invgate.com/service-desk/api/#user-GET
func (c *UserMethods) Get(p UserGetParams) (UserGetResponse, error) {
	u := UserGetResponse{}

	c.RequiredScope = scopes.UserGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
	if err != nil {
		return u, err
	}

	// Ingate returns a bool of false if the user is not found
	err = json.Unmarshal(resp, &u)
	if err != nil {
		var isFalse bool
		err = json.Unmarshal(resp, &isFalse)
		if err != nil {
			return u, err
		} else if !isFalse {
			return u, fmt.Errorf("no user found with id: %d", p.ID)
		}
	}

	return u, nil
}

type (
	UserPutParams struct {
		ID       int    `json:"id,omitempty" url:"id,required"`
		Email    string `json:"email,omitempty" url:"email"`
		Name     string `json:"name,omitempty" url:"name"`
		LastName string `json:"last_name,omitempty" url:"last_name"`
		UserBase
	}

	// UserPutResponse is used to map an user put response returned from the Invgate API
	UserPutResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for User
// Requires scope: UserPut
// See https://releases.invgate.com/service-desk/api/#user-PUT
func (c *UserMethods) Put(p UserPutParams) (UserPutResponse, error) {
	u := UserPutResponse{}

	c.RequiredScope = scopes.UserPut

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemotePut()
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	if u.Status == "ERROR" {
		return u, fmt.Errorf("invgate returned a status of %s when adding user (id: %d, name: %s) ", u.Status, p.ID, p.Name)
	}

	return u, nil
}

type (
	UserPostParams struct {
		ID       int    `json:"id,omitempty" url:"id"`
		Email    string `json:"email,omitempty" url:"email,required"`
		Name     string `json:"name,omitempty" url:"name,required"`
		LastName string `json:"last_name,omitempty" url:"last_name,required"`
		UserBase
	}

	// UserPostResponse is used to map an user returned from the Invgate API
	UserPostResponse struct{ UserGetResponse }
)

// Post for User
// Requires scope: UserPost
// See https://releases.invgate.com/service-desk/api/#user-POST
func (c *UserMethods) Post(p UserPostParams) (UserPostResponse, error) {
	u := UserPostResponse{}

	c.RequiredScope = scopes.UserPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemotePost()
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}

type (
	UserDeleteParams struct {
		// ID is not the user ID. It is the location ID
		ID    int   `url:"id,required"`
		Users []int `url:"users,required"`
	}

	// UserDeleteResponse is used to map an user delete response returned from the Invgate API
	UserDeleteResponse struct {
		// ID is not the user ID. It is the location ID
		ID int `json:"id,omitempty"`
		// Returns true if user was correctly deleted or false if not
		Value bool `json:"value,omitempty"`
	}
)

// Delete for User
// Requires scope: UserDelete
// See https://releases.invgate.com/service-desk/api/#user-DELETE
func (c *UserMethods) Delete(p UserDeleteParams) ([]UserDeleteResponse, error) {
	u := []UserDeleteResponse{}

	c.RequiredScope = scopes.UserDelete

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteDelete()
	if err != nil {
		return u, err
	}

	// Ingate returns a bool of false if the user is not found
	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}

type (
	// UsersMethods is use to call methods for Users
	UsersMethods struct{ methods.MethodCall }

	UsersGetParams struct {
		IncludeDisabled bool  `url:"include_disabled"`
		IDs             []int `url:"ids"`
	}

	// UsersGetResponse is used to map users returned from the Invgate API
	UsersGetResponse struct{ UserGetResponse }
)

// Get for Users
// Requires scope: UsersGet
// See https://releases.invgate.com/service-desk/api/#users-GET
// If no user ids or 0 is provided all users will be returned
func (c *UsersMethods) Get(p UsersGetParams) ([]UsersGetResponse, error) {
	u := []UsersGetResponse{}

	c.RequiredScope = scopes.UsersGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
	if err != nil {
		return u, err
	}

	// Ingate returns a bool of false if the user is not found
	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}
