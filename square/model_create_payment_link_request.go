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

type CreatePaymentLinkRequest struct {
	// A unique string that identifies this `CreatePaymentLinkRequest` request. If you do not provide a unique string (or provide an empty string as the value), the endpoint treats each request as independent.  For more information, see [Idempotency](https://developer.squareup.com/docs/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
	// A description of the payment link. You provide this optional description that is useful in your application context. It is not used anywhere.
	Description      string            `json:"description,omitempty"`
	QuickPay         *QuickPay         `json:"quick_pay,omitempty"`
	Order            *Order            `json:"order,omitempty"`
	CheckoutOptions  *CheckoutOptions  `json:"checkout_options,omitempty"`
	PrePopulatedData *PrePopulatedData `json:"pre_populated_data,omitempty"`
	// A note for the payment. After processing the payment, Square adds this note to the resulting `Payment`.
	PaymentNote string `json:"payment_note,omitempty"`
}
