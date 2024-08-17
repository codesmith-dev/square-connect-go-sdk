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

// Represents a [RetrieveLocationCustomAttribute](https://developer.squareup.com/reference/square_2024-07-17/location-custom-attributes-api/retrieve-location-custom-attribute) request.
type RetrieveLocationCustomAttributeRequest struct {
	// Indicates whether to return the [custom attribute definition](https://developer.squareup.com/reference/square_2024-07-17/objects/CustomAttributeDefinition) in the `definition` field of the custom attribute. Set this parameter to `true` to get the name and description of the custom attribute, information about the data type, or other definition details. The default value is `false`.
	WithDefinition bool `json:"with_definition,omitempty"`
	// The current version of the custom attribute, which is used for strongly consistent reads to guarantee that you receive the most up-to-date data. When included in the request, Square returns the specified version or a higher version if one exists. If the specified version is higher than the current version, Square returns a `BAD_REQUEST` error.
	Version int32 `json:"version,omitempty"`
}
