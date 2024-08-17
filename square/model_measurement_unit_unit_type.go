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

// MeasurementUnitUnitType : Describes the type of this unit and indicates which field contains the unit information. This is an ‘open’ enum.
type MeasurementUnitUnitType string

// List of MeasurementUnitUnitType
const (
	CUSTOM_MeasurementUnitUnitType  MeasurementUnitUnitType = "TYPE_CUSTOM"
	AREA_MeasurementUnitUnitType    MeasurementUnitUnitType = "TYPE_AREA"
	LENGTH_MeasurementUnitUnitType  MeasurementUnitUnitType = "TYPE_LENGTH"
	VOLUME_MeasurementUnitUnitType  MeasurementUnitUnitType = "TYPE_VOLUME"
	WEIGHT_MeasurementUnitUnitType  MeasurementUnitUnitType = "TYPE_WEIGHT"
	GENERIC_MeasurementUnitUnitType MeasurementUnitUnitType = "TYPE_GENERIC"
)
