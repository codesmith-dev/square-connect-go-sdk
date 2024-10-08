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

// Represents a response for an individual upsert request in a [BulkDeleteBookingCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/booking-custom-attributes-api/bulk-delete-booking-custom-attributes) operation.
type BookingCustomAttributeDeleteResponse struct {
	// The ID of the [booking](https://developer.squareup.com/reference/square_2024-07-17/objects/Booking) associated with the custom attribute.
	BookingId string `json:"booking_id,omitempty"`
	// Any errors that occurred while processing the individual request.
	Errors []ModelError `json:"errors,omitempty"`
}
