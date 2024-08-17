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

// Defines the fields that are included in the response body of a request to the [ListWebhookSubscriptions](https://developer.squareup.com/reference/square_2024-07-17/webhook-subscriptions-api/list-webhook-subscriptions) endpoint.  Note: if there are errors processing the request, the subscriptions field will not be present.
type ListWebhookSubscriptionsResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The requested list of [Subscription](https://developer.squareup.com/reference/square_2024-07-17/objects/WebhookSubscription)s.
	Subscriptions []WebhookSubscription `json:"subscriptions,omitempty"`
	// The pagination cursor to be used in a subsequent request. If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination).
	Cursor string `json:"cursor,omitempty"`
}
