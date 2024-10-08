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

// The custom attribute filter. Use this filter in a set of [custom attribute filters](https://developer.squareup.com/reference/square_2024-07-17/objects/CustomerCustomAttributeFilters) to search based on the value or last updated date of a customer-related [custom attribute](https://developer.squareup.com/reference/square_2024-07-17/objects/CustomAttribute).
type CustomerCustomAttributeFilter struct {
	// The `key` of the [custom attribute](https://developer.squareup.com/reference/square_2024-07-17/objects/CustomAttribute) to filter by. The key is the identifier of the custom attribute (and the corresponding custom attribute definition) and can be retrieved using the [Customer Custom Attributes API](https://developer.squareup.com/reference/square_2024-07-17/customer-custom-attributes-api).
	Key       string                              `json:"key"`
	Filter    *CustomerCustomAttributeFilterValue `json:"filter,omitempty"`
	UpdatedAt *TimeRange                          `json:"updated_at,omitempty"`
}
