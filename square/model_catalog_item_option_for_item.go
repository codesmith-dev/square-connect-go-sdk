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

// An option that can be assigned to an item. For example, a t-shirt item may offer a color option or a size option.
type CatalogItemOptionForItem struct {
	// The unique id of the item option, used to form the dimensions of the item option matrix in a specified order.
	ItemOptionId string `json:"item_option_id,omitempty"`
}
