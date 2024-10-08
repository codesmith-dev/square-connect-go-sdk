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

// Defines the fields that are included in the response body of a request to the [CreateCustomer](https://developer.squareup.com/reference/square_2024-07-17/customers-api/create-customer) or [BulkCreateCustomers](https://developer.squareup.com/reference/square_2024-07-17/customers-api/bulk-create-customers) endpoint.  Either `errors` or `customer` is present in a given response (never both).
type CreateCustomerResponse struct {
	// Any errors that occurred during the request.
	Errors   []ModelError `json:"errors,omitempty"`
	Customer *Customer    `json:"customer,omitempty"`
}
