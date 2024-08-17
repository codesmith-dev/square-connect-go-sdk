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

// References a text-based modifier or a list of non text-based modifiers applied to a `CatalogItem` instance and specifies supported behaviors of the application.
type CatalogItemModifierListInfo struct {
	// The ID of the `CatalogModifierList` controlled by this `CatalogModifierListInfo`.
	ModifierListId string `json:"modifier_list_id"`
	// A set of `CatalogModifierOverride` objects that override whether a given `CatalogModifier` is enabled by default.
	ModifierOverrides []CatalogModifierOverride `json:"modifier_overrides,omitempty"`
	// If 0 or larger, the smallest number of `CatalogModifier`s that must be selected from this `CatalogModifierList`. The default value is `-1`.  When  `CatalogModifierList.selection_type` is `MULTIPLE`, `CatalogModifierListInfo.min_selected_modifiers=-1`  and `CatalogModifierListInfo.max_selected_modifier=-1` means that from zero to the maximum number of modifiers of the `CatalogModifierList` can be selected from the `CatalogModifierList`.   When the `CatalogModifierList.selection_type` is `SINGLE`, `CatalogModifierListInfo.min_selected_modifiers=-1` and `CatalogModifierListInfo.max_selected_modifier=-1` means that exactly one modifier must be present in  and can be selected from the `CatalogModifierList`
	MinSelectedModifiers int32 `json:"min_selected_modifiers,omitempty"`
	// If 0 or larger, the largest number of `CatalogModifier`s that can be selected from this `CatalogModifierList`. The default value is `-1`.  When  `CatalogModifierList.selection_type` is `MULTIPLE`, `CatalogModifierListInfo.min_selected_modifiers=-1`  and `CatalogModifierListInfo.max_selected_modifier=-1` means that from zero to the maximum number of modifiers of the `CatalogModifierList` can be selected from the `CatalogModifierList`.   When the `CatalogModifierList.selection_type` is `SINGLE`, `CatalogModifierListInfo.min_selected_modifiers=-1` and `CatalogModifierListInfo.max_selected_modifier=-1` means that exactly one modifier must be present in  and can be selected from the `CatalogModifierList`
	MaxSelectedModifiers int32 `json:"max_selected_modifiers,omitempty"`
	// If `true`, enable this `CatalogModifierList`. The default value is `true`.
	Enabled bool `json:"enabled,omitempty"`
	// The position of this `CatalogItemModifierListInfo` object within the `modifier_list_info` list applied  to a `CatalogItem` instance.
	Ordinal int32 `json:"ordinal,omitempty"`
}
