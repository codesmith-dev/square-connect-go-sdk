/*
 * Square Connect API
 *
 * Client library for accessing the Square Connect APIs
 *
 * API version: 2.0
 * Contact: developers@squareup.com
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package square

// TeamMemberInvitationStatus : Enumerates the possible invitation statuses the team member can have within a business.
type TeamMemberInvitationStatus string

// List of TeamMemberInvitationStatus
const (
	UNINVITED_TeamMemberInvitationStatus TeamMemberInvitationStatus = "UNINVITED"
	PENDING_TeamMemberInvitationStatus   TeamMemberInvitationStatus = "PENDING"
	ACCEPTED_TeamMemberInvitationStatus  TeamMemberInvitationStatus = "ACCEPTED"
)
