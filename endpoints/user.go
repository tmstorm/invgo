package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// UserMethods is used to call methods for User
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
		UserName       string `json:"username,omitempty" url:"username"`
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
		LastName string `json:"lastname,omitempty"`
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

	// Invgate returns a bool of false if the user is not found
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
		LastName string `json:"lastname,omitempty" url:"lastname"`
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
		LastName string `json:"lastname,omitempty" url:"lastname,required"`
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

	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}

type (
	// UserByMethods is used to call methods for UserBy
	UserByMethods struct{ methods.MethodCall }

	UserByGetParams struct {
		Email    string `url:"email"`
		Username string `url:"username"`
	}

	// UserByGetResponse is used to map a user returned from the Invgate API
	UserByGetResponse struct{ UserGetResponse }
)

// Get for UserBy
// Requires scope: UserByGet
// See https://releases.invgate.com/service-desk/api/#userby-GET
// At least one param must be provided.
func (c *UserByMethods) Get(p UserByGetParams) (UserByGetResponse, error) {
	u := UserByGetResponse{}

	c.RequiredScope = scopes.UserByGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
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
	// UserConvertMethods is used to call methods for UserConvert
	UserConvertMethods struct{ methods.MethodCall }

	UserConvertPostParams struct {
		ID int `url:"id,required"`
	}

	// UserConvertPostResponse is used to map a user conversion response from the Invgate API
	UserConvertPostResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Post for UserConvert
// Requires scope: UserConvertPost
// See https://releases.invgate.com/service-desk/api/#userconvert-POST
// An ID must be provided
func (c *UserConvertMethods) Post(p UserConvertPostParams) (UserConvertPostResponse, error) {
	u := UserConvertPostResponse{}

	c.RequiredScope = scopes.UserConvertPost

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

	if u.Status == "ERROR" {
		return u, fmt.Errorf("invgate returned a status of %s when converting user (id: %d) ", u.Status, p.ID)
	}

	return u, nil
}

type (
	// UserDisableMethods is used to call methods for UserDisable
	UserDisableMethods struct{ methods.MethodCall }

	UserDisablePutParams struct {
		ID int `url:"id,required"`
	}

	// UserDisablePutResponse is used to map a user disable response from the Invgate API
	UserDisablePutResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for UserDisable
// Requires scope: UserDisablePut
// See https://releases.invgate.com/service-desk/api/#userdisable-PUT
// An ID must be provided
func (c *UserDisableMethods) Put(p UserDisablePutParams) (UserDisablePutResponse, error) {
	u := UserDisablePutResponse{}

	c.RequiredScope = scopes.UserDisablePut

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
		return u, fmt.Errorf("invgate returned a status of %s when disabling user (id: %d) ", u.Status, p.ID)
	}

	return u, nil
}

type (
	// UserEnableMethods is used to call methods for UserDisable
	UserEnableMethods struct{ methods.MethodCall }

	UserEnablePutParams struct {
		ID int `url:"id,required"`
	}

	// UserEnablePutResponse is used to map a user enable response from the Invgate API
	UserEnablePutResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for UserEnable
// Requires scope: UserEnablePut
// See https://releases.invgate.com/service-desk/api/#userenable-PUT
// An ID must be provided
func (c *UserEnableMethods) Put(p UserEnablePutParams) (UserEnablePutResponse, error) {
	u := UserEnablePutResponse{}

	c.RequiredScope = scopes.UserEnablePut

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
		return u, fmt.Errorf("invgate returned a status of %s when enabling user (id: %d) ", u.Status, p.ID)
	}

	return u, nil
}

type (
	// UserPasswordMethods is used to call methods for UserPassword
	UserPasswordMethods struct{ methods.MethodCall }

	UserPasswordPutParams struct {
		// ForcePasswordChange is is true user will be forced to change their password on next login
		ForcePasswordChange bool   `url:"force_password_change"`
		ID                  int    `url:"id,required"`
		Password            string `url:"password,required"`
	}

	// UserPasswordPutResponse is used to map a user disable password change from the Invgate API
	UserPasswordPutResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Put for UserPassword
// Requires scope: UserPasswordPut
// See https://releases.invgate.com/service-desk/api/#userpassword-PUT
// An ID and Password must be provided
func (c *UserPasswordMethods) Put(p UserPasswordPutParams) (UserPasswordPutResponse, error) {
	u := UserPasswordPutResponse{}

	c.RequiredScope = scopes.UserPasswordPut

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
		return u, fmt.Errorf("invgate returned a status of %s when changing user password (id: %d) ", u.Status, p.ID)
	}

	return u, nil
}

type (
	// UserPasswordResetMethods is used to call methods for UserPasswordReset
	UserPasswordResetMethods struct{ methods.MethodCall }

	UserPasswordResetPostParams struct {
		// Type only accepts 'NEW_USER' and 'RESET_PASSWORD'
		Type string `url:"type,required"`
		ID   int    `url:"id,required"`
	}

	// UserPasswordResetPostResponse is used to map a user password reset response from the Invgate API
	UserPasswordResetPostResponse struct {
		// OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status"`
	}
)

// Post for UserPasswordReset
// Requires scope: UserPasswordResetPost
// See https://releases.invgate.com/service-desk/api/#userpasswordreset-POST
// An ID and Type must be provided
// Invgate Accepts two types
// 'NEW_USER': for new users
// 'RESET_PASSWORD': for existing users
func (c *UserPasswordResetMethods) Post(p UserPasswordResetPostParams) (UserPasswordResetPostResponse, error) {
	u := UserPasswordResetPostResponse{}

	c.RequiredScope = scopes.UserPasswordResetPost

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

	if u.Status == "ERROR" {
		return u, fmt.Errorf("invgate returned a status of %s when resetting user password (id: %d) ", u.Status, p.ID)
	}

	return u, nil
}

type (
	// UserTokenMethods is used to call methods for UserToken
	UserTokenMethods struct{ methods.MethodCall }

	UserTokenPostParams struct {
		ID int `url:"id,required"`
	}

	// UserTokenPostResponse is used to map a user token response from the Invgate API
	UserTokenPostResponse struct {
		Token string `json:"token"`
	}
)

// Post for UserToken
// Requires scope: UserTokenPost
// See https://releases.invgate.com/service-desk/api/#usertoken-POST
// An ID and Type must be provided
func (c *UserTokenMethods) Post(p UserTokenPostParams) (UserTokenPostResponse, error) {
	u := UserTokenPostResponse{}

	c.RequiredScope = scopes.UserTokenPost

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
	// UsersMethods is used to call methods for Users
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

	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}

type (
	// UsersByMethods is used to call methods for UsersBy
	UsersByMethods struct{ methods.MethodCall }

	UsersByGetParams struct {
		OfficePhone     string `url:"office_phone"`
		Phone           string `url:"phone"`
		Username        string `url:"username"`
		EmployeeNumber  string `url:"employee_number"`
		OtherPhone      string `url:"other_phone"`
		IncludeDisabled bool   `url:"include_disabled"`
		Email           string `url:"email"`
		ExactMatch      bool   `url:"exact_match"`
		Phones          string `url:"phones"`
		PageKey         string `url:"page_key"`
		FaxPhone        string `url:"fax_phone"`
		MobilePhone     string `url:"mobile_phone"`
	}

	// UsersByGetResponse is used to map users returned from the Invgate API
	// Invgate says this is an array but it is actually a map which looks like
	// map[string]map[int]map[string]any
	// To accomidate this users are bound to UsersByGetResponse.Data map.
	// If you know the user id you can get them from the map like so:
	// i := id
	// user := UsersByGetResponse.Data[i]
	UsersByGetResponse struct {
		Data        map[int]UserGetResponse `json:"data,omitempty"`
		NextPageKey []int                   `json:"next_page_key,omitempty"`
	}
)

// Get for UsersBy
// Requires scope: UsersByGet
// See https://releases.invgate.com/service-desk/api/#usersby-GET
func (c *UsersByMethods) Get(p UsersByGetParams) (UsersByGetResponse, error) {
	u := UsersByGetResponse{}

	c.RequiredScope = scopes.UsersByGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
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
	// UsersGroupsMethods is used to call methods for UsersGroups
	UsersGroupsMethods struct{ methods.MethodCall }

	UsersGroupsGetParams struct {
		IDs []int `url:"ids,required"`
	}

	// UsersGroupsGetResponse is used to map a users group information returned from the Invgate API
	//
	// NOTE: Some of these might not be correct. The Invgate API docs only say array
	// for some return types and not everything has a value in the instance I am testing in.
	// I have already found that some Fields like Groups and Helpdesks return maps instead of an Array.
	// I suspect that the others do the same. With this in mind I have set all type to `any` where the docs
	// declare them as an Array but I do not yet know what is actually returned.
	UsersGroupsGetResponse struct {
		CompaniesObserved any           `json:"companies_observed,omitempty"`
		LocationsObserved any           `json:"locations_observed,omitempty"`
		Companies         any           `json:"companies,omitempty"`
		ID                int           `json:"id,omitempty"`
		GroupsObserved    any           `json:"groups_observed,omitempty"`
		Groups            CollectionMap `json:"groups,omitempty"`
		Helpdesks         CollectionMap `json:"helpdesks,omitempty"`
		HelpdesksObserved any           `json:"helpdesks_observed,omitempty"`
		Username          string        `json:"username,omitempty"`
		Email             string        `json:"email,omitempty"`
		Locations         any           `json:"locations,omitempty"`
	}

	// CollectionMap is used to map both groups and helpdesks for UsersGroups
	// The map index is collections ID e.g. If the group id is 2 then CollectionMap[2] will return this groups collection.
	CollectionMap map[int]Collection

	// Collection is the struct used by CollectionMap
	//
	// WARNING: The Invgate API docs are not clear on what data this returns.
	// This has been constructed using what I know although some types might be wrong
	Collection struct {
		EngineID          int    `json:"engine_id,omitempty"`
		HolidaysCalendar  bool   `json:"holidays_calendar,omitempty"`
		ID                int    `json:"id,omitempty"`
		IsDefault         int    `json:"is_default,omitempty"` // Should be a bool but Invgate sends int
		LevelOrder        int    `json:"level_order,omitempty"`
		Name              string `json:"name,omitempty"`
		OrgID             int    `json:"org_id,omitempty"`
		ParentID          int    `json:"parent_id,omitempty"`
		QueueForNextShift int    `json:"queue_for_next_shift,omitempty"`
		Restricted        int    `json:"restricted,omitempty"` // Probably should be a bool but is sent as int
		ScaleRuleID       int    `json:"scale_rule_id,omitempty"`
		StatusID          int    `json:"status_id,omitempty"`
		TimeZone          int    `json:"time_zone,omitempty"` // Invgate uses an internal 5 digit number for time zones
		TypeID            int    `json:"type_id,omitempty"`
	}
)

// Get for UsersGroups
// Requires scope: UsersGroupsGet
// See https://releases.invgate.com/service-desk/api/#usersgroups-GET
// At least one user ID is required
func (c *UsersGroupsMethods) Get(p UsersGroupsGetParams) ([]UsersGroupsGetResponse, error) {
	u := []UsersGroupsGetResponse{}

	c.RequiredScope = scopes.UsersGroupsGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return u, err
	}
	c.Endpoint.RawQuery = q.Encode()

	resp, err := c.RemoteGet()
	if err != nil {
		return u, err
	}

	err = json.Unmarshal(resp, &u)
	if err != nil {
		return u, err
	}

	return u, nil
}
