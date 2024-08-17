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

// PaymentOptionsDelayAction : Describes the action to be applied to a delayed capture payment when the delay_duration has elapsed.
type PaymentOptionsDelayAction string

// List of PaymentOptionsDelayAction
const (
	CANCEL_PaymentOptionsDelayAction   PaymentOptionsDelayAction = "CANCEL"
	COMPLETE_PaymentOptionsDelayAction PaymentOptionsDelayAction = "COMPLETE"
)
