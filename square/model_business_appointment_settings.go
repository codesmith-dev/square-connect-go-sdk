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

// The service appointment settings, including where and how the service is provided.
type BusinessAppointmentSettings struct {
	// Types of the location allowed for bookings.
	LocationTypes []string `json:"location_types,omitempty"`
	// The time unit of the service duration for bookings.
	AlignmentTime string `json:"alignment_time,omitempty"`
	// The minimum lead time in seconds before a service can be booked. A booking must be created at least this amount of time before its starting time.
	MinBookingLeadTimeSeconds int32 `json:"min_booking_lead_time_seconds,omitempty"`
	// The maximum lead time in seconds before a service can be booked. A booking must be created at most this amount of time before its starting time.
	MaxBookingLeadTimeSeconds int32 `json:"max_booking_lead_time_seconds,omitempty"`
	// Indicates whether a customer can choose from all available time slots and have a staff member assigned automatically (`true`) or not (`false`).
	AnyTeamMemberBookingEnabled bool `json:"any_team_member_booking_enabled,omitempty"`
	// Indicates whether a customer can book multiple services in a single online booking.
	MultipleServiceBookingEnabled bool `json:"multiple_service_booking_enabled,omitempty"`
	// Indicates whether the daily appointment limit applies to team members or to business locations.
	MaxAppointmentsPerDayLimitType string `json:"max_appointments_per_day_limit_type,omitempty"`
	// The maximum number of daily appointments per team member or per location.
	MaxAppointmentsPerDayLimit int32 `json:"max_appointments_per_day_limit,omitempty"`
	// The cut-off time in seconds for allowing clients to cancel or reschedule an appointment.
	CancellationWindowSeconds int32  `json:"cancellation_window_seconds,omitempty"`
	CancellationFeeMoney      *Money `json:"cancellation_fee_money,omitempty"`
	// The cancellation policy adopted by the seller.
	CancellationPolicy string `json:"cancellation_policy,omitempty"`
	// The free-form text of the seller's cancellation policy.
	CancellationPolicyText string `json:"cancellation_policy_text,omitempty"`
	// Indicates whether customers has an assigned staff member (`true`) or can select s staff member of their choice (`false`).
	SkipBookingFlowStaffSelection bool `json:"skip_booking_flow_staff_selection,omitempty"`
}
