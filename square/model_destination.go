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

// Information about the destination against which the payout was made.
type Destination struct {
	// Type of the destination such as a bank account or debit card.
	Type_ string `json:"type,omitempty"`
	// Square issued unique ID (also known as the instrument ID) associated with this destination.
	Id string `json:"id,omitempty"`
}
