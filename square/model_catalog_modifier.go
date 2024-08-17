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

// A modifier applicable to items at the time of sale. An example of a modifier is a Cheese add-on to a Burger item.
type CatalogModifier struct {
	// The modifier name.  This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.
	Name       string `json:"name,omitempty"`
	PriceMoney *Money `json:"price_money,omitempty"`
	// Determines where this `CatalogModifier` appears in the `CatalogModifierList`.
	Ordinal int32 `json:"ordinal,omitempty"`
	// The ID of the `CatalogModifierList` associated with this modifier.
	ModifierListId string `json:"modifier_list_id,omitempty"`
	// Location-specific price overrides.
	LocationOverrides []ModifierLocationOverrides `json:"location_overrides,omitempty"`
	// The ID of the image associated with this `CatalogModifier` instance. Currently this image is not displayed by Square, but is free to be displayed in 3rd party applications.
	ImageId string `json:"image_id,omitempty"`
}
