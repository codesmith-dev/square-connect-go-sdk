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

// Represents an individual upsert request in a [BulkUpsertBookingCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/booking-custom-attributes-api/bulk-upsert-booking-custom-attributes) request. An individual request contains a booking ID, the custom attribute to create or update, and an optional idempotency key.
type BookingCustomAttributeUpsertRequest struct {
	// The ID of the target [booking](https://developer.squareup.com/reference/square_2024-07-17/objects/Booking).
	BookingId       string           `json:"booking_id"`
	CustomAttribute *CustomAttribute `json:"custom_attribute"`
	// A unique identifier for this individual upsert request, used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
