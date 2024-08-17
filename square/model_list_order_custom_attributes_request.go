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

// Represents a list request for order custom attributes.
type ListOrderCustomAttributesRequest struct {
	// Requests that all of the custom attributes be returned, or only those that are read-only or read-write.
	VisibilityFilter string `json:"visibility_filter,omitempty"`
	// The cursor returned in the paged response from the previous call to this endpoint.  Provide this cursor to retrieve the next page of results for your original request.  For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The maximum number of results to return in a single paged response. This limit is advisory.  The response might contain more or fewer results. The minimum value is 1 and the maximum value is 100.  The default value is 20. For more information, see [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination).
	Limit int32 `json:"limit,omitempty"`
	// Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-07-17/objects/CustomAttributeDefinition) in the `definition` field of each custom attribute. Set this parameter to `true` to get the name and description of each custom attribute,  information about the data type, or other definition details. The default value is `false`.
	WithDefinitions bool `json:"with_definitions,omitempty"`
}
