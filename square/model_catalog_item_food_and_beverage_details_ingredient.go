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

import "encoding/json"

// Describes the ingredient used in a `FOOD_AND_BEV` item.
type CatalogItemFoodAndBeverageDetailsIngredient struct {
	Type_ json.RawMessage `json:"type,omitempty"`
	// The name of the ingredient from a standard pre-defined list. This should be null if it's a custom dietary preference.
	StandardName string `json:"standard_name,omitempty"`
	// The name of a custom user-defined ingredient. This should be null if it's a standard dietary preference.
	CustomName string `json:"custom_name,omitempty"`
}
