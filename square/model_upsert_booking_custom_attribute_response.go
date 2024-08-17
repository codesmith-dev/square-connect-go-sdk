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

// Represents an [UpsertBookingCustomAttribute](https://developer.squareup.com/reference/square_2024-07-17/booking-custom-attributes-api/upsert-booking-custom-attribute) response. Either `custom_attribute_definition` or `errors` is present in the response.
type UpsertBookingCustomAttributeResponse struct {
	CustomAttribute *CustomAttribute `json:"custom_attribute,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
