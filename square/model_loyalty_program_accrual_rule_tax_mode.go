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

// LoyaltyProgramAccrualRuleTaxMode : Indicates how taxes should be treated when calculating the purchase amount used for loyalty points accrual.  This setting applies only to `SPEND` accrual rules or `VISIT` accrual rules that have a minimum spend requirement.
type LoyaltyProgramAccrualRuleTaxMode string

// List of LoyaltyProgramAccrualRuleTaxMode
const (
	BEFORE_TAX_LoyaltyProgramAccrualRuleTaxMode LoyaltyProgramAccrualRuleTaxMode = "BEFORE_TAX"
	AFTER_TAX_LoyaltyProgramAccrualRuleTaxMode  LoyaltyProgramAccrualRuleTaxMode = "AFTER_TAX"
)
