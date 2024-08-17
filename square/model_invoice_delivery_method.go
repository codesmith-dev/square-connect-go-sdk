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

// InvoiceDeliveryMethod : Indicates how Square delivers the [invoice](https://developer.squareup.com/reference/square_2024-07-17/objects/Invoice) to the customer.
type InvoiceDeliveryMethod string

// List of InvoiceDeliveryMethod
const (
	EMAIL_InvoiceDeliveryMethod          InvoiceDeliveryMethod = "EMAIL"
	SHARE_MANUALLY_InvoiceDeliveryMethod InvoiceDeliveryMethod = "SHARE_MANUALLY"
	SMS_InvoiceDeliveryMethod            InvoiceDeliveryMethod = "SMS"
)
