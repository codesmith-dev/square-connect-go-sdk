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

// The set of search requirements.
type SearchLoyaltyRewardsRequestLoyaltyRewardQuery struct {
	// The ID of the [loyalty account](https://developer.squareup.com/reference/square_2024-07-17/objects/LoyaltyAccount) to which the loyalty reward belongs.
	LoyaltyAccountId string `json:"loyalty_account_id"`
	// The status of the loyalty reward.
	Status string `json:"status,omitempty"`
}
