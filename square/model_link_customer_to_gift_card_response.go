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

// A response that contains the linked `GiftCard` object. If the request resulted in errors,  the response contains a set of `Error` objects.
type LinkCustomerToGiftCardResponse struct {
	// Any errors that occurred during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	GiftCard *GiftCard    `json:"gift_card,omitempty"`
}
