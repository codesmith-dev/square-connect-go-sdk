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

// Represents an additional recipient (other than the merchant) receiving a portion of this tender.
type AdditionalRecipient struct {
	// The location ID for a recipient (other than the merchant) receiving a portion of this tender.
	LocationId string `json:"location_id"`
	// The description of the additional recipient.
	Description string `json:"description,omitempty"`
	AmountMoney *Money `json:"amount_money"`
	// The unique ID for the RETIRED `AdditionalRecipientReceivable` object. This field should be empty for any `AdditionalRecipient` objects created after the retirement.
	ReceivableId string `json:"receivable_id,omitempty"`
}
