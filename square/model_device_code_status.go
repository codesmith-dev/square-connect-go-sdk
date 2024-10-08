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

// DeviceCodeStatus : DeviceCode.Status enum.
type DeviceCodeStatus string

// List of DeviceCodeStatus
const (
	UNKNOWN_DeviceCodeStatus  DeviceCodeStatus = "UNKNOWN"
	UNPAIRED_DeviceCodeStatus DeviceCodeStatus = "UNPAIRED"
	PAIRED_DeviceCodeStatus   DeviceCodeStatus = "PAIRED"
	EXPIRED_DeviceCodeStatus  DeviceCodeStatus = "EXPIRED"
)
