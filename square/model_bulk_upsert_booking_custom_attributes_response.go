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

// Represents a [BulkUpsertBookingCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/booking-custom-attributes-api/bulk-upsert-booking-custom-attributes) response, which contains a map of responses that each corresponds to an individual upsert request.
type BulkUpsertBookingCustomAttributesResponse struct {
	// A map of responses that correspond to individual upsert requests. Each response has the same ID as the corresponding request and contains either a `booking_id` and `custom_attribute` or an `errors` field.
	Values map[string]BookingCustomAttributeUpsertResponse `json:"values,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
