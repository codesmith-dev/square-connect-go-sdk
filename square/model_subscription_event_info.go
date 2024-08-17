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

// Provides information about the subscription event.
type SubscriptionEventInfo struct {
	// A human-readable explanation for the event.
	Detail string          `json:"detail,omitempty"`
	Code   json.RawMessage `json:"code,omitempty"`
}
