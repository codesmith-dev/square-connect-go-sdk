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

// SearchOrdersSortField : Specifies which timestamp to use to sort `SearchOrder` results.
type SearchOrdersSortField string

// List of SearchOrdersSortField
const (
	CREATED_AT_SearchOrdersSortField SearchOrdersSortField = "CREATED_AT"
	UPDATED_AT_SearchOrdersSortField SearchOrdersSortField = "UPDATED_AT"
	CLOSED_AT_SearchOrdersSortField  SearchOrdersSortField = "CLOSED_AT"
)
