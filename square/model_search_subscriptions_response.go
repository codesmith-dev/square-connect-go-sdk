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

// Defines output parameters in a response from the [SearchSubscriptions](https://developer.squareup.com/reference/square_2024-07-17/subscriptions-api/search-subscriptions) endpoint.
type SearchSubscriptionsResponse struct {
	// Errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The subscriptions matching the specified query expressions.
	Subscriptions []Subscription `json:"subscriptions,omitempty"`
	// When the total number of resulting subscription exceeds the limit of a paged response,  the response includes a cursor for you to use in a subsequent request to fetch the next set of results. If the cursor is unset, the response contains the last page of the results.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
