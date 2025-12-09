// Package scopes defines all scopes available for use in the Invgate API.
// If a new endpoint is added its related scope must be added here or
// the API call will fail and Invgo will throw an error
package scopes

import (
	"errors"
	"fmt"
	"slices"
)

// Scopes defines all scopes for the Invgate API.
// These must be requested on the initial connection to the API.
// These are checked against each call to ensure the current requested
// scope matches that of the end point being called.
//
// The format for any given Invgate scope is as follows:
// base.endpointname:method
//
// For example when making a call to /breakingnews.all
// the scope would look like: base = api.v1 + endpointname = .breakingnews.all + method = :get
// Resulting in the scope: api.v1.breakingnews.all:get

type (
	// ScopeType used to safely set and get scopes
	ScopeType string

	// MethodTypes used to safely define and get methods for each scope
	MethodTypes struct {
		Get    string
		Post   string
		Patch  string
		Put    string
		Delete string
	}
)

// Define all the reusable scope variables to be used when building a full scope
var (
	// base defines the base value for a scope
	base = "api.v1"
	// attributes is used when building a scope for attributes
	attributes = ".attributes"
	// status is used when building a scope for status
	status = ".status"
	// all is used when getting a scope to get all rows
	all = ".all"

	// methods defines each method for use in a scope
	methods = MethodTypes{
		Get:    ":get",
		Post:   ":post",
		Patch:  ":patch",
		Put:    ":put",
		Delete: ":delete",
	}
)

// Breaking News
var (
	breakingnews                           = ".breakingnews"
	BreakingNewsGet              ScopeType = ScopeType(base + breakingnews + methods.Get)
	BreakingNewsPost             ScopeType = ScopeType(base + breakingnews + methods.Post)
	BreakingNewsPut              ScopeType = ScopeType(base + breakingnews + methods.Put)
	BreakingNewsAll              ScopeType = ScopeType(base + breakingnews + all + methods.Get)
	BreakingNewsAttributesStatus ScopeType = ScopeType(base + breakingnews + attributes + status + methods.Get)
	BreakingNewsAttributesType   ScopeType = ScopeType(base + breakingnews + attributes + ".type" + methods.Get)
	BreakingNewsStatusGet        ScopeType = ScopeType(base + breakingnews + status + methods.Get)
	BreakingNewsStatusPost       ScopeType = ScopeType(base + breakingnews + status + methods.Post)
)

// CategoriesGet
var CategoriesGet ScopeType = ScopeType(base + ".categories" + methods.Get)

// Help Desks
var (
	helpdesks                          = ".helpdesks"
	HelpDesksGet             ScopeType = ScopeType(base + helpdesks + methods.Get)
	HelpDesksObserversGet    ScopeType = ScopeType(base + helpdesks + ".observers" + methods.Get)
	HelpDesksObserversPost   ScopeType = ScopeType(base + helpdesks + ".observers" + methods.Post)
	HelpDesksObserversDelete ScopeType = ScopeType(base + helpdesks + ".observers" + methods.Delete)
	HelpDesksAndLevelsGet    ScopeType = ScopeType(base + ".helpdesksandlevels" + methods.Get)
)

// Incident
var (
	incident                                = ".incident"
	IncidentGet                   ScopeType = ScopeType(base + incident + methods.Get)
	IncidentPost                  ScopeType = ScopeType(base + incident + methods.Post)
	IncidentPut                   ScopeType = ScopeType(base + incident + methods.Put)
	IncidentAttributesPriorityGet ScopeType = ScopeType(base + incident + attributes + ".priority" + methods.Get)
	IncidentAttributesSourceGet   ScopeType = ScopeType(base + incident + attributes + ".source" + methods.Get)
	IncidentAttributesStatusGet   ScopeType = ScopeType(base + incident + attributes + status + methods.Get)
	IncidentAttributesTypeGet     ScopeType = ScopeType(base + incident + attributes + ".type" + methods.Get)
)

// Incidents
var (
	incidents                      = ".incidents"
	IncidentsGet         ScopeType = ScopeType(base + incidents + methods.Get)
	IncidentsByStatusGet ScopeType = ScopeType(base + incidents + ".by" + ".status" + methods.Get)
)

// ServiceDeskVersionGet
var ServiceDeskVersionGet ScopeType = ScopeType(base + ".sd.version" + methods.Get)

// Triggers
var (
	triggers                        = ".triggers"
	TriggersGet           ScopeType = ScopeType(base + triggers + methods.Get)
	TriggersExecutionsGet ScopeType = ScopeType(base + triggers + ".executions" + methods.Get)
)

// CreateScopes is used to take all scopes provided in the config and convert them to strings
// for creating the initial Invgate connection
func CreateScopes(scopes []ScopeType) []string {
	strings := []string{}
	for i := range len(scopes) {
		scope := string(scopes[i])
		strings = append(strings, scope)
	}
	return strings
}

// CheckScopes is an internal function to check if the current set of scopes
// match the required scope for an end point.
func CheckScopes(scopes []ScopeType, requestScopes ...ScopeType) error {
	if requestScopes == nil {
		return errors.New("a request scope was not provided")
	}

	if len(scopes) == 0 {
		return errors.New("current scopes were not provided")
	}

	for _, v := range scopes {
		if ok := slices.Contains(requestScopes, v); ok {
			return nil
		}
	}
	return fmt.Errorf("the scope for the current request has not been acquired for the current client: %s", requestScopes)
}
