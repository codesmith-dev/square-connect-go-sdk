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

// Disables the card, preventing any further updates or charges. Disabling an already disabled card is allowed but has no effect. Accessible via HTTP requests at POST https://connect.squareup.com/v2/cards/{card_id}/disable
type DisableCardRequest struct{}
