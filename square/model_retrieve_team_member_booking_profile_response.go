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

type RetrieveTeamMemberBookingProfileResponse struct {
	TeamMemberBookingProfile *TeamMemberBookingProfile `json:"team_member_booking_profile,omitempty"`
	// Errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
