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

// Represents a [DeleteCustomerCustomAttribute](https://developer.squareup.com/reference/square_2024-07-17/customer-custom-attributes-api/delete-customer-custom-attribute) response. Either an empty object `{}` (for a successful deletion) or `errors` is present in the response.
type DeleteCustomerCustomAttributeResponse struct {
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
