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

// Represents an input to a call to [UpdateVendor](https://developer.squareup.com/reference/square_2024-07-17/vendors-api/update-vendor).
type UpdateVendorRequest struct {
	// A client-supplied, universally unique identifier (UUID) for the request.  See [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) in the [API Development 101](https://developer.squareup.com/docs/buildbasics) section for more information.
	IdempotencyKey string  `json:"idempotency_key,omitempty"`
	Vendor         *Vendor `json:"vendor"`
}
