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

// Represents an [UpdateCustomerCustomAttributeDefinition](https://developer.squareup.com/reference/square_2024-07-17/customer-custom-attributes-api/update-customer-custom-attribute-definition) request.
type UpdateCustomerCustomAttributeDefinitionRequest struct {
	CustomAttributeDefinition *CustomAttributeDefinition `json:"custom_attribute_definition"`
	// A unique identifier for this request, used to ensure idempotency. For more information, see [Idempotency](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency).
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}
