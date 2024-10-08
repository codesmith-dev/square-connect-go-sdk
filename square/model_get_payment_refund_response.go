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

// Defines the response returned by [GetRefund](https://developer.squareup.com/reference/square_2024-07-17/refunds-api/get-payment-refund).  Note: If there are errors processing the request, the refund field might not be present or it might be present in a FAILED state.
type GetPaymentRefundResponse struct {
	// Information about errors encountered during the request.
	Errors []ModelError   `json:"errors,omitempty"`
	Refund *PaymentRefund `json:"refund,omitempty"`
}
