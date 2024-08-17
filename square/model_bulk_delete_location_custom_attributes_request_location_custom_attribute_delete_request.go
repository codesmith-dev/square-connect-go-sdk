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

// Represents an individual delete request in a [BulkDeleteLocationCustomAttributes](https://developer.squareup.com/reference/square_2024-07-17/location-custom-attributes-api/bulk-delete-location-custom-attributes) request. An individual request contains an optional ID of the associated custom attribute definition and optional key of the associated custom attribute definition.
type BulkDeleteLocationCustomAttributesRequestLocationCustomAttributeDeleteRequest struct {
	// The key of the associated custom attribute definition. Represented as a qualified key if the requesting app is not the definition owner.
	Key string `json:"key,omitempty"`
}
