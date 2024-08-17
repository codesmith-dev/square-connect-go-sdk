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

// SearchVendorsRequestSortField : The field to sort the returned [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor) objects by.
type SearchVendorsRequestSortField string

// List of SearchVendorsRequestSortField
const (
	NAME_SearchVendorsRequestSortField       SearchVendorsRequestSortField = "NAME"
	CREATED_AT_SearchVendorsRequestSortField SearchVendorsRequestSortField = "CREATED_AT"
)
