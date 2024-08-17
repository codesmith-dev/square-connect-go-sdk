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

// Represents an input into a call to [SearchVendors](https://developer.squareup.com/reference/square_2024-07-17/vendors-api/search-vendors).
type SearchVendorsRequest struct {
	Filter json.RawMessage `json:"filter,omitempty"`
	Sort   json.RawMessage `json:"sort,omitempty"`
	// A pagination cursor returned by a previous call to this endpoint. Provide this to retrieve the next set of results for the original query.  See the [Pagination](https://developer.squareup.com/docs/working-with-apis/pagination) guide for more information.
	Cursor string `json:"cursor,omitempty"`
}
