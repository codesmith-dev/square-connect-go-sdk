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

// Lists all [Subscription](https://developer.squareup.com/reference/square_2024-07-17/objects/WebhookSubscription)s owned by your application.
type ListWebhookSubscriptionsRequest struct {
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for your original query.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
	// Includes disabled [Subscription](https://developer.squareup.com/reference/square_2024-07-17/objects/WebhookSubscription)s. By default, all enabled [Subscription](https://developer.squareup.com/reference/square_2024-07-17/objects/WebhookSubscription)s are returned.
	IncludeDisabled bool `json:"include_disabled,omitempty"`
	// Sorts the returned list by when the [Subscription](https://developer.squareup.com/reference/square_2024-07-17/objects/WebhookSubscription) was created with the specified order. This field defaults to ASC.
	SortOrder string `json:"sort_order,omitempty"`
	// The maximum number of results to be returned in a single page. It is possible to receive fewer results than the specified limit on a given page. The default value of 100 is also the maximum allowed value.  Default: 100
	Limit int32 `json:"limit,omitempty"`
}
