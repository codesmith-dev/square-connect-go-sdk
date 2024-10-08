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

// BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType : Types of daily appointment limits.
type BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType string

// List of BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType
const (
	TEAM_MEMBER_BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType = "PER_TEAM_MEMBER"
	LOCATION_BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType    BusinessAppointmentSettingsMaxAppointmentsPerDayLimitType = "PER_LOCATION"
)
