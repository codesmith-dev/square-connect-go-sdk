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

// Represents a contact of a [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
type VendorContact struct {
	// A unique Square-generated ID for the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact). This field is required when attempting to update a [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact).
	Id string `json:"id,omitempty"`
	// The name of the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact). This field is required when attempting to create a [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Name string `json:"name,omitempty"`
	// The email address of the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact).
	EmailAddress string `json:"email_address,omitempty"`
	// The phone number of the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact).
	PhoneNumber string `json:"phone_number,omitempty"`
	// The state of the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact).
	Removed bool `json:"removed,omitempty"`
	// The ordinal of the [VendorContact](https://developer.squareup.com/reference/square_2024-07-17/objects/VendorContact).
	Ordinal int32 `json:"ordinal"`
}
