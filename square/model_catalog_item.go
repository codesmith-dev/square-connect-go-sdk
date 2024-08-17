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

import "encoding/json"

// A [CatalogObject](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogObject) instance of the `ITEM` type, also referred to as an item, in the catalog.
type CatalogItem struct {
	// The item's name. This is a searchable attribute for use in applicable query filters, its value must not be empty, and the length is of Unicode code points.
	Name string `json:"name,omitempty"`
	// The item's description. This is a searchable attribute for use in applicable query filters, and its value length is of Unicode code points.  Deprecated at 2022-07-20, this field is planned to retire in 6 months. You should migrate to use `description_html` to set the description of the [CatalogItem](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogItem) instance.  The `description` and `description_html` field values are kept in sync. If you try to set the both fields, the `description_html` text value overwrites the `description` value. Updates in one field are also reflected in the other, except for when you use an early version before Square API 2022-07-20 and `description_html` is set to blank, setting the `description` value to null does not nullify `description_html`.
	Description string `json:"description,omitempty"`
	// The text of the item's display label in the Square Point of Sale app. Only up to the first five characters of the string are used. This attribute is searchable, and its value length is of Unicode code points.
	Abbreviation string `json:"abbreviation,omitempty"`
	// The color of the item's display label in the Square Point of Sale app. This must be a valid hex color code.
	LabelColor string `json:"label_color,omitempty"`
	// Indicates whether the item is taxable (`true`) or non-taxable (`false`). Default is `true`.
	IsTaxable bool `json:"is_taxable,omitempty"`
	// If `true`, the item can be added to shipping orders from the merchant's online store.
	AvailableOnline bool `json:"available_online,omitempty"`
	// If `true`, the item can be added to pickup orders from the merchant's online store.
	AvailableForPickup bool `json:"available_for_pickup,omitempty"`
	// If `true`, the item can be added to electronically fulfilled orders from the merchant's online store.
	AvailableElectronically bool `json:"available_electronically,omitempty"`
	// The ID of the item's category, if any. Deprecated since 2023-12-13. Use `CatalogItem.categories`, instead.
	CategoryId string `json:"category_id,omitempty"`
	// A set of IDs indicating the taxes enabled for this item. When updating an item, any taxes listed here will be added to the item. Taxes may also be added to or deleted from an item using `UpdateItemTaxes`.
	TaxIds []string `json:"tax_ids,omitempty"`
	// A set of `CatalogItemModifierListInfo` objects representing the modifier lists that apply to this item, along with the overrides and min and max limits that are specific to this item. Modifier lists may also be added to or deleted from an item using `UpdateItemModifierLists`.
	ModifierListInfo []CatalogItemModifierListInfo `json:"modifier_list_info,omitempty"`
	// A list of [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogItemVariation) objects for this item. An item must have at least one variation.
	Variations []CatalogObject `json:"variations,omitempty"`
	// The product type of the item. Once set, the `product_type` value cannot be modified.  Items of the `LEGACY_SQUARE_ONLINE_SERVICE` and `LEGACY_SQUARE_ONLINE_MEMBERSHIP` product types can be updated but cannot be created using the API.
	ProductType string `json:"product_type,omitempty"`
	// If `false`, the Square Point of Sale app will present the `CatalogItem`'s details screen immediately, allowing the merchant to choose `CatalogModifier`s before adding the item to the cart.  This is the default behavior.  If `true`, the Square Point of Sale app will immediately add the item to the cart with the pre-selected modifiers, and merchants can edit modifiers by drilling down onto the item's details.  Third-party clients are encouraged to implement similar behaviors.
	SkipModifierScreen bool `json:"skip_modifier_screen,omitempty"`
	// List of item options IDs for this item. Used to manage and group item variations in a specified order.  Maximum: 6 item options.
	ItemOptions []CatalogItemOptionForItem `json:"item_options,omitempty"`
	// The IDs of images associated with this `CatalogItem` instance. These images will be shown to customers in Square Online Store. The first image will show up as the icon for this item in POS.
	ImageIds []string `json:"image_ids,omitempty"`
	// A name to sort the item by. If this name is unspecified, namely, the `sort_name` field is absent, the regular `name` field is used for sorting. Its value must not be empty.  It is currently supported for sellers of the Japanese locale only.
	SortName string `json:"sort_name,omitempty"`
	// The list of categories.
	Categories []CatalogObjectCategory `json:"categories,omitempty"`
	// The item's description as expressed in valid HTML elements. The length of this field value, including those of HTML tags, is of Unicode points. With application query filters, the text values of the HTML elements and attributes are searchable. Invalid or unsupported HTML elements or attributes are ignored.  Supported HTML elements include: - `a`: Link. Supports linking to website URLs, email address, and telephone numbers. - `b`, `strong`:  Bold text - `br`: Line break - `code`: Computer code - `div`: Section - `h1-h6`: Headings - `i`, `em`: Italics - `li`: List element - `ol`: Numbered list - `p`: Paragraph - `ul`: Bullet list - `u`: Underline   Supported HTML attributes include: - `align`: Alignment of the text content - `href`: Link destination - `rel`: Relationship between link's target and source - `target`: Place to open the linked document
	DescriptionHtml string `json:"description_html,omitempty"`
	// A server-generated plaintext version of the `description_html` field, without formatting tags.
	DescriptionPlaintext string `json:"description_plaintext,omitempty"`
	// A list of IDs representing channels, such as a Square Online site, where the item can be made visible or available.
	Channels []string `json:"channels,omitempty"`
	// Indicates whether this item is archived (`true`) or not (`false`).
	IsArchived             bool                   `json:"is_archived,omitempty"`
	EcomSeoData            *CatalogEcomSeoData    `json:"ecom_seo_data,omitempty"`
	FoodAndBeverageDetails json.RawMessage        `json:"food_and_beverage_details,omitempty"`
	ReportingCategory      *CatalogObjectCategory `json:"reporting_category,omitempty"`
}
