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

type Device struct {
	// A synthetic identifier for the device. The identifier includes a standardized prefix and is otherwise an opaque id generated from key device fields.
	Id         string            `json:"id,omitempty"`
	Attributes *DeviceAttributes `json:"attributes"`
	// A list of components applicable to the device.
	Components []Component   `json:"components,omitempty"`
	Status     *DeviceStatus `json:"status,omitempty"`
}
