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

type ListCashDrawerShiftEventsResponse struct {
	// Opaque cursor for fetching the next page. Cursor is not present in the last page of results.
	Cursor string `json:"cursor,omitempty"`
	// Any errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
	// All of the events (payments, refunds, etc.) for a cash drawer during the shift.
	CashDrawerShiftEvents []CashDrawerShiftEvent `json:"cash_drawer_shift_events,omitempty"`
}
