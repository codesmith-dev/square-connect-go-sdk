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

// Represents an output from a call to [UpdateVendor](https://developer.squareup.com/reference/square_2024-07-17/vendors-api/update-vendor).
type UpdateVendorResponse struct {
	// Errors occurred when the request fails.
	Errors []ModelError `json:"errors,omitempty"`
	Vendor *Vendor      `json:"vendor,omitempty"`
}
