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

// TenderBankAccountDetailsStatus : Indicates the bank account payment's current status.
type TenderBankAccountDetailsStatus string

// List of TenderBankAccountDetailsStatus
const (
	PENDING_TenderBankAccountDetailsStatus   TenderBankAccountDetailsStatus = "PENDING"
	COMPLETED_TenderBankAccountDetailsStatus TenderBankAccountDetailsStatus = "COMPLETED"
	FAILED_TenderBankAccountDetailsStatus    TenderBankAccountDetailsStatus = "FAILED"
)
