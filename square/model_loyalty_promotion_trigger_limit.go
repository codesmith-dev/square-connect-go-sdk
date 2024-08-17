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

// Represents the number of times a buyer can earn points during a [loyalty promotion](https://developer.squareup.com/reference/square_2024-07-17/objects/LoyaltyPromotion). If this field is not set, buyers can trigger the promotion an unlimited number of times to earn points during the time that the promotion is available.  A purchase that is disqualified from earning points because of this limit might qualify for another active promotion.
type LoyaltyPromotionTriggerLimit struct {
	// The maximum number of times a buyer can trigger the promotion during the specified `interval`.
	Times int32 `json:"times"`
	// The time period the limit applies to.
	Interval string `json:"interval,omitempty"`
}
