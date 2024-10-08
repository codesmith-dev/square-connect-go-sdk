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

// Represents details about an `ADJUST_DECREMENT` [gift card activity type](https://developer.squareup.com/reference/square_2024-07-17/objects/GiftCardActivityType).
type GiftCardActivityAdjustDecrement struct {
	AmountMoney *Money `json:"amount_money"`
	// The reason the gift card balance was adjusted.
	Reason string `json:"reason"`
}
