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

// LoyaltyProgramRewardDefinitionType : The type of discount the reward tier offers. DEPRECATED at version 2020-12-16. Discount details are now defined using a catalog pricing rule and other catalog objects. For more information, see [Getting discount details for a reward tier](https://developer.squareup.com/docs/loyalty-api/loyalty-rewards#get-discount-details).
type LoyaltyProgramRewardDefinitionType string

// List of LoyaltyProgramRewardDefinitionType
const (
	AMOUNT_LoyaltyProgramRewardDefinitionType     LoyaltyProgramRewardDefinitionType = "FIXED_AMOUNT"
	PERCENTAGE_LoyaltyProgramRewardDefinitionType LoyaltyProgramRewardDefinitionType = "FIXED_PERCENTAGE"
)
