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

type PaymentBalanceActivityTaxOnFeeDetail struct {
	// The ID of the payment associated with this activity.
	PaymentId string `json:"payment_id,omitempty"`
	// The description of the tax rate being applied. For example: \"GST\", \"HST\".
	TaxRateDescription string `json:"tax_rate_description,omitempty"`
}
