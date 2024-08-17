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

// Represents a line item in an order. Each line item describes a different product to purchase, with its own quantity and price details.
type OrderLineItem struct {
	// A unique ID that identifies the line item only within this order.
	Uid string `json:"uid,omitempty"`
	// The name of the line item.
	Name string `json:"name,omitempty"`
	// The count, or measurement, of a line item being purchased:  If `quantity` is a whole number, and `quantity_unit` is not specified, then `quantity` denotes an item count.  For example: `3` apples.  If `quantity` is a whole or decimal number, and `quantity_unit` is also specified, then `quantity` denotes a measurement.  For example: `2.25` pounds of broccoli.  For more information, see [Specify item quantity and measurement unit](https://developer.squareup.com/docs/orders-api/create-orders#specify-item-quantity-and-measurement-unit).  Line items with a quantity of `0` are automatically removed when paying for or otherwise completing the order.
	Quantity     string             `json:"quantity"`
	QuantityUnit *OrderQuantityUnit `json:"quantity_unit,omitempty"`
	// An optional note associated with the line item.
	Note string `json:"note,omitempty"`
	// The [CatalogItemVariation](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogItemVariation) ID applied to this line item.
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this line item references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The name of the variation applied to this line item.
	VariationName string `json:"variation_name,omitempty"`
	// The type of line item: an itemized sale, a non-itemized sale (custom amount), or the activation or reloading of a gift card.
	ItemType string `json:"item_type,omitempty"`
	// Application-defined data attached to this line item. Metadata fields are intended to store descriptive references or associations with an entity in another system or store brief information about the object. Square does not process this field; it only stores and returns it in relevant API calls. Do not use metadata to store any sensitive information (such as personally identifiable information or card details).  Keys written by applications must be 60 characters or less and must be in the character set `[a-zA-Z0-9_-]`. Entries can also include metadata generated by Square. These keys are prefixed with a namespace, separated from the key with a ':' character.  Values have a maximum length of 255 characters.  An application can have up to 10 entries per metadata field.  Entries written by applications are private and can only be read or modified by the same application.  For more information, see [Metadata](https://developer.squareup.com/docs/build-basics/metadata).
	Metadata map[string]string `json:"metadata,omitempty"`
	// The [CatalogModifier](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogModifier)s applied to this line item.
	Modifiers []OrderLineItemModifier `json:"modifiers,omitempty"`
	// The list of references to taxes applied to this line item. Each `OrderLineItemAppliedTax` has a `tax_uid` that references the `uid` of a top-level `OrderLineItemTax` applied to the line item. On reads, the amount applied is populated.  An `OrderLineItemAppliedTax` is automatically created on every line item for all `ORDER` scoped taxes added to the order. `OrderLineItemAppliedTax` records for `LINE_ITEM` scoped taxes must be added in requests for the tax to apply to any line items.  To change the amount of a tax, modify the referenced top-level tax.
	AppliedTaxes []OrderLineItemAppliedTax `json:"applied_taxes,omitempty"`
	// The list of references to discounts applied to this line item. Each `OrderLineItemAppliedDiscount` has a `discount_uid` that references the `uid` of a top-level `OrderLineItemDiscounts` applied to the line item. On reads, the amount applied is populated.  An `OrderLineItemAppliedDiscount` is automatically created on every line item for all `ORDER` scoped discounts that are added to the order. `OrderLineItemAppliedDiscount` records for `LINE_ITEM` scoped discounts must be added in requests for the discount to apply to any line items.  To change the amount of a discount, modify the referenced top-level discount.
	AppliedDiscounts []OrderLineItemAppliedDiscount `json:"applied_discounts,omitempty"`
	// The list of references to service charges applied to this line item. Each `OrderLineItemAppliedServiceCharge` has a `service_charge_id` that references the `uid` of a top-level `OrderServiceCharge` applied to the line item. On reads, the amount applied is populated.  To change the amount of a service charge, modify the referenced top-level service charge.
	AppliedServiceCharges    []OrderLineItemAppliedServiceCharge `json:"applied_service_charges,omitempty"`
	BasePriceMoney           *Money                              `json:"base_price_money,omitempty"`
	VariationTotalPriceMoney *Money                              `json:"variation_total_price_money,omitempty"`
	GrossSalesMoney          *Money                              `json:"gross_sales_money,omitempty"`
	TotalTaxMoney            *Money                              `json:"total_tax_money,omitempty"`
	TotalDiscountMoney       *Money                              `json:"total_discount_money,omitempty"`
	TotalMoney               *Money                              `json:"total_money,omitempty"`
	PricingBlocklists        *OrderLineItemPricingBlocklists     `json:"pricing_blocklists,omitempty"`
	TotalServiceChargeMoney  *Money                              `json:"total_service_charge_money,omitempty"`
}
