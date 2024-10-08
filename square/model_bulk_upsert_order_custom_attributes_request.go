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

// Represents a bulk upsert request for one or more order custom attributes.
type BulkUpsertOrderCustomAttributesRequest struct {
	// A map of requests that correspond to individual upsert operations for custom attributes.
	Values map[string]json.RawMessage `json:"values"`
}
