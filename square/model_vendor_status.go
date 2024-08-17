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

// VendorStatus : The status of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor), whether a [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor) is active or inactive.
type VendorStatus string

// List of VendorStatus
const (
	ACTIVE_VendorStatus   VendorStatus = "ACTIVE"
	INACTIVE_VendorStatus VendorStatus = "INACTIVE"
)
