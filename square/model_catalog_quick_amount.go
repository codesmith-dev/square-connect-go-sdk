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

// Represents a Quick Amount in the Catalog.
type CatalogQuickAmount struct {
	// Represents the type of the Quick Amount.
	Type_  string `json:"type"`
	Amount *Money `json:"amount"`
	// Describes the ranking of the Quick Amount provided by machine learning model, in the range [0, 100]. MANUAL type amount will always have score = 100.
	Score int64 `json:"score,omitempty"`
	// The order in which this Quick Amount should be displayed.
	Ordinal int64 `json:"ordinal,omitempty"`
}
