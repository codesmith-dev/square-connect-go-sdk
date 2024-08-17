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

type RetrieveBookingResponse struct {
	Booking *Booking `json:"booking,omitempty"`
	// Errors that occurred during the request.
	Errors []ModelError `json:"errors,omitempty"`
}
