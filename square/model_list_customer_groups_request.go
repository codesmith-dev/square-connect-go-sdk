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

// Defines the query parameters that can be included in a request to the [ListCustomerGroups](https://developer.squareup.com/reference/square_2024-07-17/customer-groups-api/list-customer-groups) endpoint.
type ListCustomerGroupsRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this cursor to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// The maximum number of results to return in a single page. This limit is advisory. The response might contain more or fewer results. If the limit is less than 1 or greater than 50, Square returns a `400 VALUE_TOO_LOW` or `400 VALUE_TOO_HIGH` error. The default value is 50.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Limit int32 `json:"limit,omitempty"`
}
