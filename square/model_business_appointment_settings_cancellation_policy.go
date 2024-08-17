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

// BusinessAppointmentSettingsCancellationPolicy : The category of the seller’s cancellation policy.
type BusinessAppointmentSettingsCancellationPolicy string

// List of BusinessAppointmentSettingsCancellationPolicy
const (
	CANCELLATION_TREATED_AS_NO_SHOW_BusinessAppointmentSettingsCancellationPolicy BusinessAppointmentSettingsCancellationPolicy = "CANCELLATION_TREATED_AS_NO_SHOW"
	CUSTOM_POLICY_BusinessAppointmentSettingsCancellationPolicy                   BusinessAppointmentSettingsCancellationPolicy = "CUSTOM_POLICY"
)
