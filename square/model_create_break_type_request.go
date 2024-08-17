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

// A request to create a new `BreakType`.
type CreateBreakTypeRequest struct {
	// A unique string value to ensure the idempotency of the operation.
	IdempotencyKey string     `json:"idempotency_key,omitempty"`
	BreakType      *BreakType `json:"break_type"`
}
