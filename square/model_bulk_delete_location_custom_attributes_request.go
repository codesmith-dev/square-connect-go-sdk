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

// Represents a [BulkDeleteLocationCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/location-custom-attributes-api/bulk-delete-location-custom-attributes) request.
type BulkDeleteLocationCustomAttributesRequest struct {
	// The data used to update the `CustomAttribute` objects. The keys must be unique and are used to map to the corresponding response.
	Values map[string]json.RawMessage `json:"values"`
}
