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

type RetrieveLocationSettingsResponse struct {
	// Any errors that occurred during the request.
	Errors           []ModelError              `json:"errors,omitempty"`
	LocationSettings *CheckoutLocationSettings `json:"location_settings,omitempty"`
}
