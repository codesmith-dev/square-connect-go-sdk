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

// TenderCardDetailsEntryMethod : Indicates the method used to enter the card's details.
type TenderCardDetailsEntryMethod string

// List of TenderCardDetailsEntryMethod
const (
	SWIPED_TenderCardDetailsEntryMethod      TenderCardDetailsEntryMethod = "SWIPED"
	KEYED_TenderCardDetailsEntryMethod       TenderCardDetailsEntryMethod = "KEYED"
	EMV_TenderCardDetailsEntryMethod         TenderCardDetailsEntryMethod = "EMV"
	ON_FILE_TenderCardDetailsEntryMethod     TenderCardDetailsEntryMethod = "ON_FILE"
	CONTACTLESS_TenderCardDetailsEntryMethod TenderCardDetailsEntryMethod = "CONTACTLESS"
)
