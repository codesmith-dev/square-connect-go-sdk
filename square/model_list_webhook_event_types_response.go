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

// Defines the fields that are included in the response body of a request to the [ListWebhookEventTypes](https://developer.squareup.com/reference/square_2024-07-17/webhook-subscriptions-api/list-webhook-event-types) endpoint.  Note: if there are errors processing the request, the event types field will not be present.
type ListWebhookEventTypesResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The list of event types.
	EventTypes []string `json:"event_types,omitempty"`
	// Contains the metadata of a webhook event type. For more information, see [EventTypeMetadata](https://developer.squareup.com/reference/square_2024-07-17/objects/EventTypeMetadata).
	Metadata []EventTypeMetadata `json:"metadata,omitempty"`
}
