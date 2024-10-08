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

// The response object returned by the [ListMerchant](https://developer.squareup.com/reference/square_2024-07-17/merchants-api/list-merchants) endpoint.
type ListMerchantsResponse struct {
	// Information on errors encountered during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// The requested `Merchant` entities.
	Merchant []Merchant `json:"merchant,omitempty"`
	// If the  response is truncated, the cursor to use in next  request to fetch next set of objects.
	Cursor int32 `json:"cursor,omitempty"`
}
