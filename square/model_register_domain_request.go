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

// Defines the parameters that can be included in the body of a request to the [RegisterDomain](https://developer.squareup.com/reference/square_2024-07-17/apple-pay-api/register-domain) endpoint.
type RegisterDomainRequest struct {
	// A domain name as described in RFC-1034 that will be registered with ApplePay.
	DomainName string `json:"domain_name"`
}
