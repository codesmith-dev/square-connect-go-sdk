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

// The query expression to specify the key to sort search results.
type CatalogQuerySortedAttribute struct {
	// The attribute whose value is used as the sort key.
	AttributeName string `json:"attribute_name"`
	// The first attribute value to be returned by the query. Ascending sorts will return only objects with this value or greater, while descending sorts will return only objects with this value or less. If unset, start at the beginning (for ascending sorts) or end (for descending sorts).
	InitialAttributeValue string `json:"initial_attribute_value,omitempty"`
	// The desired sort order, `\"ASC\"` (ascending) or `\"DESC\"` (descending).
	SortOrder string `json:"sort_order,omitempty"`
}
