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

// Represents a tax being returned that applies to one or more return line items in an order.  Fixed-amount, order-scoped taxes are distributed across all non-zero return line item totals. The amount distributed to each return line item is relative to that item’s contribution to the order subtotal.
type OrderReturnTax struct {
	// A unique ID that identifies the returned tax only within this order.
	Uid string `json:"uid,omitempty"`
	// The tax `uid` from the order that contains the original tax charge.
	SourceTaxUid string `json:"source_tax_uid,omitempty"`
	// The catalog object ID referencing [CatalogTax](https://developer.squareup.com/reference/square_2024-07-17/objects/CatalogTax).
	CatalogObjectId string `json:"catalog_object_id,omitempty"`
	// The version of the catalog object that this tax references.
	CatalogVersion int64 `json:"catalog_version,omitempty"`
	// The tax's name.
	Name string `json:"name,omitempty"`
	// Indicates the calculation method used to apply the tax.
	Type_ string `json:"type,omitempty"`
	// The percentage of the tax, as a string representation of a decimal number. For example, a value of `\"7.25\"` corresponds to a percentage of 7.25%.
	Percentage   string `json:"percentage,omitempty"`
	AppliedMoney *Money `json:"applied_money,omitempty"`
	// Indicates the level at which the `OrderReturnTax` applies. For `ORDER` scoped taxes, Square generates references in `applied_taxes` on all `OrderReturnLineItem`s. For `LINE_ITEM` scoped taxes, the tax is only applied to `OrderReturnLineItem`s with references in their `applied_discounts` field.
	Scope string `json:"scope,omitempty"`
}
