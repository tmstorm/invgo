package endpoints

import (
	"encoding/json"
	"fmt"

	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/internal/utils"
	"github.com/tmstorm/invgo/scopes"
)

type (
	// TimeTrackingMethods is used to call methods for TimeTracking
	TimeTrackingMethods struct{ methods.MethodCall }

	// TimeTrackingGetResponse is used to map a time tracking GET response from the Invgate API
	TimeTrackingGetResponse struct {
		// Status 1 if Enabled, 0 if Deleted
		Status int `json:"status,omitempty"`
		// Comment of the logged period.
		Comment                string `json:"comment,omitempty"`
		TimetrackingCategoryID int    `json:"timetracking_category_id,omitempty"`
		// Incident is the request ID
		Incident int `json:"incident,omitempty"`
		// From	Initial date and time of the interval (in ISO 8601 format).
		From           string `json:"from,omitempty"`
		TimetrackingID int    `json:"timetracking_id,omitempty"`
		UserID         int    `json:"user_id,omitempty"`
		// Total ammount of time in seconds.
		Total int `json:"total,omitempty"` // Invgate says this is a string but golang json.Unmarshal detects it as an int
		// Ending date and time of the interval (in ISO 8601 format).
		To string `json:"to,omitempty"`
	}

	TimeTrackingGetParams struct {
		// Indicate the date format. The available formats are 'iso8601noT' or 'iso8601'. If null,
		// 'iso8601noT' format is returned, which is ISO-8601 no T (YYYY-mm-dd H:i).
		DateFormat string `url:"date_format"`
		// Ending date and time of the interval (must be in ISO-8601 format).
		// If it's not specified, the current time will be used.
		To string `url:"to"`
		// RequestID is required if from parameter is not provided
		RequestID int `url:"request_id"`
		// From Initial date and time of the interval (must be in ISO-8601 format).
		// Required if the request_id parameter is not provided.
		From string `url:"from"`
	}
)

// Get for TimeTracking
// Requires scope: TimeTrackingGet
// See https://releases.invgate.com/service-desk/api/#timetracking-Get
func (w *TimeTrackingMethods) Get(p TimeTrackingGetParams) ([]TimeTrackingGetResponse, error) {
	r := []TimeTrackingGetResponse{}
	w.RequiredScope = scopes.TimeTrackingGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemoteGet()
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
	// TimeTrackingPostResponse is used to map a timetracking POST response from the Invgate API
	TimeTrackingPostResponse struct {
		// Status OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status,omitempty"`
		// TimetrackingID the ID of the created tracking entry
		TimetrackingID int `json:"timetracking_id,omitempty"`
	}

	TimeTrackingPostParams struct {
		CategoryID int    `url:"category_id"`
		Comment    string `url:"comment"`
		UserID     int    `url:"user_id,required"`
		// From Initial date and time of the interval (timestamp format).
		// If it's not specified, the current time will be used.
		From      int `url:"from"`
		RequestID int `url:"request_id,required"`
		To        int `url:"to,required"`
	}
)

// Post for TimeTracking
// Requires scope: TimeTrackingPost
// See https://releases.invgate.com/service-desk/api/#timetracking-POST
func (w *TimeTrackingMethods) Post(p TimeTrackingPostParams) (TimeTrackingPostResponse, error) {
	r := TimeTrackingPostResponse{}
	w.RequiredScope = scopes.TimeTrackingPost

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemotePost()
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
	// TimeTrackingDeleteResponse is used to map a timetracking DELETE response from the Invgate API
	TimeTrackingDeleteResponse struct {
		// Status OK if user was correctly updated, ERROR if something went wrong
		Status string `json:"status,omitempty"`
	}

	TimeTrackingDeleteParams struct {
		TimetrackingID int `url:"timetracking_id,required"`
		RequestID      int `url:"request_id,required"`
		UserID         int `url:"user_id,required"`
	}
)

// Delete for TimeTracking
// Requires scope: TimeTrackingDelete
// See https://releases.invgate.com/service-desk/api/#timetracking-DELETE
func (w *TimeTrackingMethods) Delete(p TimeTrackingDeleteParams) (TimeTrackingDeleteResponse, error) {
	r := TimeTrackingDeleteResponse{}
	w.RequiredScope = scopes.TimeTrackingDelete

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemoteDelete()
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}

	if r.Status == "ERROR" {
		return r, fmt.Errorf("invgate returned a status of %s when deleting time tracking (timetracking_id: %d) ", r.Status, p.TimetrackingID)
	}

	return r, nil
}

type (
	// TimeTrackingAttributesCategoryMethods is used to call methods for TimeTrackingAttributesCategory
	TimeTrackingAttributesCategoryMethods struct{ methods.MethodCall }

	// TimeTrackingAttributesCategoryGetResponse is used to map a timetracking GET response from the Invgate API
	TimeTrackingAttributesCategoryGetResponse struct {
		Name        string  `json:"name,omitempty"`
		CostPerHour float64 `json:"cost_per_hour,omitempty"`
		ID          int     `json:"id,omitempty"`
		ParentID    int     `json:"parent_id,omitempty"`
	}

	TimeTrackingAttributesCategoryGetParams struct {
		ID int `url:"id"`
	}
)

// Get for TimeTrackingAttributesCategory
// Requires scope: TimeTrackingAttributesCategoryGet
// See https://releases.invgate.com/service-desk/api/#timetrackingattributescategory-GET
// If no ID or 0 is provided all records will be returned
func (w *TimeTrackingAttributesCategoryMethods) Get(p TimeTrackingAttributesCategoryGetParams) ([]TimeTrackingAttributesCategoryGetResponse, error) {
	r := []TimeTrackingAttributesCategoryGetResponse{}
	w.RequiredScope = scopes.TimeTrackingAttributesCategoryGet

	q, err := utils.StructToQuery(p)
	if err != nil {
		return r, err
	}
	w.Endpoint.RawQuery = q.Encode()

	resp, err := w.RemoteGet()
	if err != nil {
		return r, err
	}

	err = json.Unmarshal(resp, &r)
	if err != nil {
		return r, err
	}
	return r, nil
}
