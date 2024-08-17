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

// PayoutStatus : Payout status types
type PayoutStatus string

// List of PayoutStatus
const (
	SENT_PayoutStatus   PayoutStatus = "SENT"
	FAILED_PayoutStatus PayoutStatus = "FAILED"
	PAID_PayoutStatus   PayoutStatus = "PAID"
)
