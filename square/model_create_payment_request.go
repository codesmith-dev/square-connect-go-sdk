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

// Describes a request to create a payment using  [CreatePayment](https://developer.squareup.com/reference/square_2024-07-17/payments-api/create-payment).
type CreatePaymentRequest struct {
	// The ID for the source of funds for this payment. This could be a payment token generated by the Web Payments SDK for any of its [supported methods](https://developer.squareup.com/docs/web-payments/overview#explore-payment-methods), including cards, bank transfers, Afterpay or Cash App Pay. If recording a payment that the seller received outside of Square, specify either \"CASH\" or \"EXTERNAL\". For more information, see  [Take Payments](https://developer.squareup.com/docs/payments-api/take-payments).
	SourceId string `json:"source_id"`
	// A unique string that identifies this `CreatePayment` request. Keys can be any valid string but must be unique for every `CreatePayment` request.  Note: The number of allowed characters might be less than the stated maximum, if multi-byte characters are used.  For more information, see [Idempotency](https://developer.squareup.com/docs/working-with-apis/idempotency).
	IdempotencyKey string `json:"idempotency_key"`
	AmountMoney    *Money `json:"amount_money,omitempty"`
	TipMoney       *Money `json:"tip_money,omitempty"`
	AppFeeMoney    *Money `json:"app_fee_money,omitempty"`
	// The duration of time after the payment's creation when Square automatically  either completes or cancels the payment depending on the `delay_action` field value.  For more information, see  [Time threshold](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture#time-threshold).   This parameter should be specified as a time duration, in RFC 3339 format.  Note: This feature is only supported for card payments. This parameter can only be set for a delayed capture payment (`autocomplete=false`).  Default:  - Card-present payments: \"PT36H\" (36 hours) from the creation time. - Card-not-present payments: \"P7D\" (7 days) from the creation time.
	DelayDuration string `json:"delay_duration,omitempty"`
	// The action to be applied to the payment when the `delay_duration` has elapsed. The action must be CANCEL or COMPLETE. For more information, see  [Time Threshold](https://developer.squareup.com/docs/payments-api/take-payments/card-payments/delayed-capture#time-threshold).   Default: CANCEL
	DelayAction string `json:"delay_action,omitempty"`
	// If set to `true`, this payment will be completed when possible. If set to `false`, this payment is held in an approved state until either explicitly completed (captured) or canceled (voided). For more information, see [Delayed capture](https://developer.squareup.com/docs/payments-api/take-payments/card-payments#delayed-capture-of-a-card-payment).  Default: true
	Autocomplete bool `json:"autocomplete,omitempty"`
	// Associates a previously created order with this payment.
	OrderId string `json:"order_id,omitempty"`
	// The [Customer](https://developer.squareup.com/reference/square_2024-07-17/objects/Customer) ID of the customer associated with the payment.  This is required if the `source_id` refers to a card on file created using the Cards API.
	CustomerId string `json:"customer_id,omitempty"`
	// The location ID to associate with the payment. If not specified, the [main location](https://developer.squareup.com/docs/locations-api#about-the-main-location) is used.
	LocationId string `json:"location_id,omitempty"`
	// An optional [TeamMember](https://developer.squareup.com/reference/square_2024-07-17/objects/TeamMember) ID to associate with  this payment.
	TeamMemberId string `json:"team_member_id,omitempty"`
	// A user-defined ID to associate with the payment.  You can use this field to associate the payment to an entity in an external system  (for example, you might specify an order ID that is generated by a third-party shopping cart).
	ReferenceId string `json:"reference_id,omitempty"`
	// An identifying token generated by [payments.verifyBuyer()](https://developer.squareup.com/reference/sdks/web/payments/objects/Payments#Payments.verifyBuyer). Verification tokens encapsulate customer device information and 3-D Secure challenge results to indicate that Square has verified the buyer identity.  For more information, see [SCA Overview](https://developer.squareup.com/docs/sca-overview).
	VerificationToken string `json:"verification_token,omitempty"`
	// If set to `true` and charging a Square Gift Card, a payment might be returned with `amount_money` equal to less than what was requested. For example, a request for $20 when charging a Square Gift Card with a balance of $5 results in an APPROVED payment of $5. You might choose to prompt the buyer for an additional payment to cover the remainder or cancel the Gift Card payment. This field cannot be `true` when `autocomplete = true`.  For more information, see [Partial amount with Square Gift Cards](https://developer.squareup.com/docs/payments-api/take-payments#partial-payment-gift-card).  Default: false
	AcceptPartialAuthorization bool `json:"accept_partial_authorization,omitempty"`
	// The buyer's email address.
	BuyerEmailAddress string   `json:"buyer_email_address,omitempty"`
	BillingAddress    *Address `json:"billing_address,omitempty"`
	ShippingAddress   *Address `json:"shipping_address,omitempty"`
	// An optional note to be entered by the developer when creating a payment.
	Note string `json:"note,omitempty"`
	// Optional additional payment information to include on the customer's card statement as part of the statement description. This can be, for example, an invoice number, ticket number, or short description that uniquely identifies the purchase.  Note that the `statement_description_identifier` might get truncated on the statement description to fit the required information including the Square identifier (SQ *) and name of the seller taking the payment.
	StatementDescriptionIdentifier string                  `json:"statement_description_identifier,omitempty"`
	CashDetails                    *CashPaymentDetails     `json:"cash_details,omitempty"`
	ExternalDetails                *ExternalPaymentDetails `json:"external_details,omitempty"`
	CustomerDetails                *CustomerDetails        `json:"customer_details,omitempty"`
}
