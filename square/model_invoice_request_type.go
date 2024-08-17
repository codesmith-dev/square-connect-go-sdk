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

// InvoiceRequestType : Indicates the type of the payment request. For more information, see  [Configuring payment requests](https://developer.squareup.com/docs/invoices-api/create-publish-invoices#payment-requests).
type InvoiceRequestType string

// List of InvoiceRequestType
const (
	BALANCE_InvoiceRequestType     InvoiceRequestType = "BALANCE"
	DEPOSIT_InvoiceRequestType     InvoiceRequestType = "DEPOSIT"
	INSTALLMENT_InvoiceRequestType InvoiceRequestType = "INSTALLMENT"
)
