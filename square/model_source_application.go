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

// Represents information about the application used to generate a change.
type SourceApplication struct {
	// __Read only__ The [product](https://developer.squareup.com/reference/square_2024-07-17/enums/Product) type of the application.
	Product string `json:"product,omitempty"`
	// __Read only__ The Square-assigned ID of the application. This field is used only if the [product](https://developer.squareup.com/reference/square_2024-07-17/enums/Product) type is `EXTERNAL_API`.
	ApplicationId string `json:"application_id,omitempty"`
	// __Read only__ The display name of the application (for example, `\"Custom Application\"` or `\"Square POS 4.74 for Android\"`).
	Name string `json:"name,omitempty"`
}
