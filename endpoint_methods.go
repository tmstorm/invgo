package invgo

/*
All endpoints must be implemented here to be available in the Invgo public API

To implement a new endpoint add its methods to invgo/endpoints/endpoint_name.go.
Then add it as a method to the Invgo Client using newPublicMethod.

Example:
	func (c *Client) BreakingNews() *endpoints.BreakingNewsMethods {
		return newPublicMethod[endpoints.BreakingNewsMethods](c, "/breakingnews")
	}
*/

import (
	"reflect"

	"github.com/tmstorm/invgo/endpoints"
	"github.com/tmstorm/invgo/internal/methods"
	"github.com/tmstorm/invgo/scopes"
)

// newPublicMethod should be used when adding a new enpoint to the Invgo public API
// T must be a struct whose first field is methods.MethodCall
func newPublicMethod[T any](c *Client, endpoint string) *T {
	var result T
	v := reflect.ValueOf(&result).Elem()

	ep := c.APIURL.JoinPath(endpoint)

	for i := range v.NumField() {
		if v.Field(i).Type() == reflect.TypeOf(methods.MethodCall{}) {
			mc := methods.MethodCall{
				Client: &methods.Client{
					HTTPClient:    c.HTTPClient,
					CurrentScopes: c.CurrentScopes,
					APIURL:        c.APIURL,
				},
				Endpoint: ep,
			}
			v.Field(i).Set(reflect.ValueOf(mc))
			break
		}
	}

	return &result
}

// BreakingNews manages the /breakingnews endpoint
// Get: Returns the requested Breaking News
// Post: Creates Breaking News
// Put: Modifies a set of Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnews
func (c *Client) BreakingNews() *endpoints.BreakingNewsMethods {
	return newPublicMethod[endpoints.BreakingNewsMethods](c, "/breakingnews")
}

// BreakingNewsAll gets all the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsall
func (c *Client) BreakingNewsAll() *endpoints.BreakingNewsAllMethods {
	return newPublicMethod[endpoints.BreakingNewsAllMethods](c, "/breakingnews.all")
}

// BreakingNewsStatus manages the updates of the breaking news.
// Get: Returns the updates of the requested Breaking News
// Post: Creates a new update to the given Breaking News
// See https://releases.invgate.com/service-desk/api/#breakingnewsstatus
func (c *Client) BreakingNewsStatus() *endpoints.BreakingNewsStatusMethods {
	return newPublicMethod[endpoints.BreakingNewsStatusMethods](c, "/breakingnews.status")
}

// BreakingNewsAttributesStatus gets all the possible status for the Breaking News'
// and their descriptions.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributesstatus
func (c *Client) BreakingNewsAttributesStatus() *endpoints.AttributesMethods {
	m := newPublicMethod[endpoints.AttributesMethods](c, "/breakingnews.attributes.status")
	m.RequiredScope = scopes.BreakingNewsAttributesStatus
	return m
}

// BreakingNewsAttributesType gets all the importance types of the Breaking News.
// See https://releases.invgate.com/service-desk/api/#breakingnewsattributestype
func (c *Client) BreakingNewsAttributesType() *endpoints.AttributesMethods {
	m := newPublicMethod[endpoints.AttributesMethods](c, "/breakingnews.attributes.type")
	m.RequiredScope = scopes.BreakingNewsAttributesType
	return m
}

// Categories is used to get all categories for the current Invgate instance
// See https://releases.invgate.com/service-desk/api/#categories
func (c *Client) Categories() *endpoints.CategoriesMethods {
	return newPublicMethod[endpoints.CategoriesMethods](c, "/categories")
}

// HelpDesks manages the help desks
// See https://releases.invgate.com/service-desk/api/#helpdesks
func (c *Client) HelpDesks() *endpoints.HelpDesksMethods {
	return newPublicMethod[endpoints.HelpDesksMethods](c, "/helpdesks")
}

// Incident manages the /incident endpoint
// Get: Returns the information of the given request
// Post: Creates a request
// Put: Change attributes of a request
// See https://releases.invgate.com/service-desk/api/#incident
func (c *Client) Incident() *endpoints.IncidentMethods {
	return newPublicMethod[endpoints.IncidentMethods](c, "/incident")
}

// IncidentApproval returns an incidents approvals
// See https://releases.invgate.com/service-desk/api/#incidentapproval
func (c *Client) IncidentApproval() *endpoints.IncidentApprovalMethods {
	return newPublicMethod[endpoints.IncidentApprovalMethods](c, "/incident.approval")
}

// IncidentApprovalAccept returns an incidents approval accept
// See https://releases.invgate.com/service-desk/api/#incidentapprovalaccept
func (c *Client) IncidentApprovalAccept() *endpoints.IncidentApprovalAcceptMethods {
	return newPublicMethod[endpoints.IncidentApprovalAcceptMethods](c, "/incident.approval.accept")
}

// IncidentApprovalAddVoter add voter to incident
// See https://releases.invgate.com/service-desk/api/#incidentapprovaladd_voter
func (c *Client) IncidentApprovalAddVoter() *endpoints.IncidentApprovalAddVoterMethods {
	return newPublicMethod[endpoints.IncidentApprovalAddVoterMethods](c, "/incident.approval.add_voter")
}

// IncidentApprovalCancel returns an incidents approval cancel
// See https://releases.invgate.com/service-desk/api/#incidentapprovalcancel
func (c *Client) IncidentApprovalCancel() *endpoints.IncidentApprovalCancelMethods {
	return newPublicMethod[endpoints.IncidentApprovalCancelMethods](c, "/incident.approval.cancel")
}

// IncidentApprovalPossibleVoters returns an incidents approval possible voters
// See https://releases.invgate.com/service-desk/api/#incidentapprovalpossible_voters
func (c *Client) IncidentApprovalPossibleVoters() *endpoints.IncidentApprovalPossibleVotersMethods {
	return newPublicMethod[endpoints.IncidentApprovalPossibleVotersMethods](c, "/incident.approval.possible_voters")
}

// IncidentApprovalReject returns an incidents approval reject
// See https://releases.invgate.com/service-desk/api/#incidentapprovalcancel
func (c *Client) IncidentApprovalReject() *endpoints.IncidentApprovalRejectMethods {
	return newPublicMethod[endpoints.IncidentApprovalRejectMethods](c, "/incident.approval.reject")
}

// IncidentApprovalStatus returns an incidents approval statuses
// See https://releases.invgate.com/service-desk/api/#incidentapprovalstatus
func (c *Client) IncidentApprovalStatus() *endpoints.IncidentApprovalStatusMethods {
	return newPublicMethod[endpoints.IncidentApprovalStatusMethods](c, "/incident.approval.status")
}

// IncidentApprovalType returns an incidents approval types
// See https://releases.invgate.com/service-desk/api/#incidentapprovaltype
func (c *Client) IncidentApprovalType() *endpoints.IncidentApprovalTypeMethods {
	return newPublicMethod[endpoints.IncidentApprovalTypeMethods](c, "/incident.approval.type")
}

// IncidentApprovalVoteStatus returns an incidents approval vote statues
// See https://releases.invgate.com/service-desk/api/#incidentapprovalvote_status
func (c *Client) IncidentApprovalVoteStatus() *endpoints.IncidentApprovalVoteStatusMethods {
	return newPublicMethod[endpoints.IncidentApprovalVoteStatusMethods](c, "/incident.approval.vote_status")
}

// IncidentAttachment returns an incidents attachment
// See https://releases.invgate.com/service-desk/api/#incidentattachment
func (c *Client) IncidentAttachment() *endpoints.IncidentAttachmentMethods {
	return newPublicMethod[endpoints.IncidentAttachmentMethods](c, "/incident.attachment")
}

// IncidentComment manages an incidents comments
// See https://releases.invgate.com/service-desk/api/#incidentcomment
func (c *Client) IncidentComment() *endpoints.IncidentCommentMethods {
	return newPublicMethod[endpoints.IncidentCommentMethods](c, "/incident.comment")
}

// Incidents is used to get Incidents from the Invgate API
// See https://releases.invgate.com/service-desk/api/#incidents
func (c *Client) Incidents() *endpoints.IncidentsMethods {
	return newPublicMethod[endpoints.IncidentsMethods](c, "/incidents")
}

// IncidentsByStatus gets incidents by the given set of status IDs
// See https://releases.invgate.com/service-desk/api/#incidentsbystatus
func (c *Client) IncidentsByStatus() *endpoints.IncidentsByStatusMethods {
	return newPublicMethod[endpoints.IncidentsByStatusMethods](c, "/incidents.by.status")
}

// IncidentAttributesStatus gets all the status types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesStatus() *endpoints.AttributesMethods {
	m := newPublicMethod[endpoints.AttributesMethods](c, "/incident.attributes.status")
	m.RequiredScope = scopes.IncidentAttributesStatusGet
	return m
}

// IncidentAttributesType gets all the types usable for an incident
// See https://releases.invgate.com/service-desk/api/#incidentattributestype
func (c *Client) IncidentAttributesType() *endpoints.AttributesMethods {
	m := newPublicMethod[endpoints.AttributesMethods](c, "/incident.attributes.type")
	m.RequiredScope = scopes.IncidentAttributesTypeGet
	return m
}

// ServiceDeskVersion returns the current version of the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#sdversion
func (c *Client) ServiceDeskVersion() *endpoints.ServiceDeskVersionMethods {
	return newPublicMethod[endpoints.ServiceDeskVersionMethods](c, "/sd.version")
}

// TimeTracking manages time tracking records in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#timetracking
func (c *Client) TimeTracking() *endpoints.TimeTrackingMethods {
	return newPublicMethod[endpoints.TimeTrackingMethods](c, "/timetracking")
}

// TimeTrackingAttributesCategory manages time tracking records attributes by category in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#timetracking
func (c *Client) TimeTrackingAttributesCategory() *endpoints.TimeTrackingAttributesCategoryMethods {
	return newPublicMethod[endpoints.TimeTrackingAttributesCategoryMethods](c, "/timetracking.attributes.category")
}

// Triggers returns user defined tiggers in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#triggers
func (c *Client) Triggers() *endpoints.TriggersMethods {
	return newPublicMethod[endpoints.TriggersMethods](c, "/triggers")
}

// TriggersExecutions returns a list of each time a trigger was executed in the Service Desk instance
// See https://releases.invgate.com/service-desk/api/#triggers
func (c *Client) TriggersExecutions() *endpoints.TriggersExecutionsMethods {
	return newPublicMethod[endpoints.TriggersExecutionsMethods](c, "/triggers.executions")
}

// User manages the /user endpoint
// Get: Returns the requested user
// Post: Creates a user
// Put: Modifies a user
// See https://releases.invgate.com/service-desk/api/#user
func (c *Client) User() *endpoints.UserMethods {
	return newPublicMethod[endpoints.UserMethods](c, "/user")
}

// UserBy gets a user by username or email
// See https://releases.invgate.com/service-desk/api/#userby
func (c *Client) UserBy() *endpoints.UserByMethods {
	return newPublicMethod[endpoints.UserByMethods](c, "/user.by")
}

// UserConvert converts external user to internal user
// See https://releases.invgate.com/service-desk/api/#userconvert
func (c *Client) UserConvert() *endpoints.UserConvertMethods {
	return newPublicMethod[endpoints.UserConvertMethods](c, "/user.convert")
}

// UserDisable disables a user
// See https://releases.invgate.com/service-desk/api/#userdisable
func (c *Client) UserDisable() *endpoints.UserDisableMethods {
	return newPublicMethod[endpoints.UserDisableMethods](c, "/user.disable")
}

// UserEnable enables a user
// See https://releases.invgate.com/service-desk/api/#userenable
func (c *Client) UserEnable() *endpoints.UserEnableMethods {
	return newPublicMethod[endpoints.UserEnableMethods](c, "/user.enable")
}

// UserPassword changes a users password
// See https://releases.invgate.com/service-desk/api/#userpassword
func (c *Client) UserPassword() *endpoints.UserPasswordMethods {
	return newPublicMethod[endpoints.UserPasswordMethods](c, "/user.password")
}

// UserPasswordReset forces a user to do a password reset
// See https://releases.invgate.com/service-desk/api/#userpasswordreset
func (c *Client) UserPasswordReset() *endpoints.UserPasswordResetMethods {
	return newPublicMethod[endpoints.UserPasswordResetMethods](c, "/user.password.reset")
}

// UserToken creates a session token for a user
// See https://releases.invgate.com/service-desk/api/#usertoken
func (c *Client) UserToken() *endpoints.UserTokenMethods {
	return newPublicMethod[endpoints.UserTokenMethods](c, "/user.token")
}

// Users returns a list of each of the specified list or all users
// See https://releases.invgate.com/service-desk/api/#users
func (c *Client) Users() *endpoints.UsersMethods {
	return newPublicMethod[endpoints.UsersMethods](c, "/users")
}

// UsersBy returns a list of users matching the filter
// See https://releases.invgate.com/service-desk/api/#usersby
func (c *Client) UsersBy() *endpoints.UsersByMethods {
	return newPublicMethod[endpoints.UsersByMethods](c, "/users.by")
}

// UsersGroups returns a users groups, companies, helpdesks, and locations
// See https://releases.invgate.com/service-desk/api/#usersgroups
func (c *Client) UsersGroups() *endpoints.UsersGroupsMethods {
	return newPublicMethod[endpoints.UsersGroupsMethods](c, "/users.groups")
}

// WorkflowDeploy is used to deploy a workflow
// See https://releases.invgate.com/service-desk/api/#wfdeploy
func (c *Client) WorkflowDeploy() *endpoints.WorkflowDeployMethods {
	return newPublicMethod[endpoints.WorkflowDeployMethods](c, "/wf.deploy")
}

// WorkflowInitialFieldsByCategory is used to get the initial fields needed to create a workflow
// See https://releases.invgate.com/service-desk/api/#wfinitialfieldsbycategory
func (c *Client) WorkflowInitialFieldsByCategory() *endpoints.WorkflowInitialFieldsByCategoryMethods {
	return newPublicMethod[endpoints.WorkflowInitialFieldsByCategoryMethods](c, "/wf.initialfields.by.category")
}
