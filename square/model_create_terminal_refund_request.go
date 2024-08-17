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

type CreateTerminalRefundRequest struct {
	// A unique string that identifies this `CreateRefund` request. Keys can be any valid string but must be unique for every `CreateRefund` request.  See [Idempotency keys](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) for more information.
	IdempotencyKey string          `json:"idempotency_key"`
	Refund         *TerminalRefund `json:"refund,omitempty"`
}
