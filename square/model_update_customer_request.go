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

// Defines the body parameters that can be included in a request to the `UpdateCustomer` endpoint.
type UpdateCustomerRequest struct {
	// The given name (that is, the first name) associated with the customer profile.  The maximum length for this value is 300 characters.
	GivenName string `json:"given_name,omitempty"`
	// The family name (that is, the last name) associated with the customer profile.  The maximum length for this value is 300 characters.
	FamilyName string `json:"family_name,omitempty"`
	// A business name associated with the customer profile.  The maximum length for this value is 500 characters.
	CompanyName string `json:"company_name,omitempty"`
	// A nickname for the customer profile.  The maximum length for this value is 100 characters.
	Nickname string `json:"nickname,omitempty"`
	// The email address associated with the customer profile.  The maximum length for this value is 254 characters.
	EmailAddress string   `json:"email_address,omitempty"`
	Address      *Address `json:"address,omitempty"`
	// The phone number associated with the customer profile. The phone number must be valid and can contain 9–16 digits, with an optional `+` prefix and country code. For more information, see [Customer phone numbers](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#phone-number).
	PhoneNumber string `json:"phone_number,omitempty"`
	// An optional second ID used to associate the customer profile with an entity in another system.  The maximum length for this value is 100 characters.
	ReferenceId string `json:"reference_id,omitempty"`
	// A custom note associated with the customer profile.
	Note string `json:"note,omitempty"`
	// The birthday associated with the customer profile, in `YYYY-MM-DD` or `MM-DD` format. For example, specify `1998-09-21` for September 21, 1998, or `09-21` for September 21. Birthdays are returned in `YYYY-MM-DD` format, where `YYYY` is the specified birth year or `0000` if a birth year is not specified.
	Birthday string `json:"birthday,omitempty"`
	// The current version of the customer profile.  As a best practice, you should include this field to enable [optimistic concurrency](https://developer.squareup.com/docs/build-basics/common-api-patterns/optimistic-concurrency) control. For more information, see [Update a customer profile](https://developer.squareup.com/docs/customers-api/use-the-api/keep-records#update-a-customer-profile).
	Version int64           `json:"version,omitempty"`
	TaxIds  *CustomerTaxIds `json:"tax_ids,omitempty"`
}
