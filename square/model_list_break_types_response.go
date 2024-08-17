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

// The response to a request for a set of `BreakType` objects. The response contains the requested `BreakType` objects and might contain a set of `Error` objects if the request resulted in errors.
type ListBreakTypesResponse struct {
	//  A page of `BreakType` results.
	BreakTypes []BreakType `json:"break_types,omitempty"`
	// The value supplied in the subsequent request to fetch the next page of `BreakType` results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
