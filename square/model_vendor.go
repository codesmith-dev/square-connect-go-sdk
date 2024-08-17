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

// Represents a supplier to a seller.
type Vendor struct {
	// A unique Square-generated ID for the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor). This field is required when attempting to update a [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Id string `json:"id,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor) was created.
	CreatedAt string `json:"created_at,omitempty"`
	// An RFC 3339-formatted timestamp that indicates when the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor) was last updated.
	UpdatedAt string `json:"updated_at,omitempty"`
	// The name of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor). This field is required when attempting to create or update a [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Name    string   `json:"name,omitempty"`
	Address *Address `json:"address,omitempty"`
	// The contacts of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Contacts []VendorContact `json:"contacts,omitempty"`
	// The account number of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	AccountNumber string `json:"account_number,omitempty"`
	// A note detailing information about the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Note string `json:"note,omitempty"`
	// The version of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Version int32 `json:"version,omitempty"`
	// The status of the [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor).
	Status string `json:"status,omitempty"`
}
