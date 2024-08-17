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

// Defines an appointment segment of a booking.
type AppointmentSegment struct {
	// The time span in minutes of an appointment segment.
	DurationMinutes int32 `json:"duration_minutes,omitempty"`
	// The ID of the [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogItemVariation) object representing the service booked in this segment.
	ServiceVariationId string `json:"service_variation_id,omitempty"`
	// The ID of the [TeamMember](https://developer.squareup.com/reference/square_2024-07-17/objects/TeamMember) object representing the team member booked in this segment.
	TeamMemberId string `json:"team_member_id"`
	// The current version of the item variation representing the service booked in this segment.
	ServiceVariationVersion int64 `json:"service_variation_version,omitempty"`
	// Time between the end of this segment and the beginning of the subsequent segment.
	IntermissionMinutes int32 `json:"intermission_minutes,omitempty"`
	// Whether the customer accepts any team member, instead of a specific one, to serve this segment.
	AnyTeamMember bool `json:"any_team_member,omitempty"`
	// The IDs of the seller-accessible resources used for this appointment segment.
	ResourceIds []string `json:"resource_ids,omitempty"`
}
