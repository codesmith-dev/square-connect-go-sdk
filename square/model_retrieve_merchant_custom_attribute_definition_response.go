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

// Represents a [RetrieveMerchantCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-07-17/merchant-custom-attributes-api/retrieve-merchant-custom-attributeDefinition) response. Either `custom_attribute_definition` or `errors` is present in the response.
type RetrieveMerchantCustomAttributeDefinitionResponse struct {
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
