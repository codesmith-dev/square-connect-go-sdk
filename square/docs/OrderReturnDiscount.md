# OrderReturnDiscount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uid** | **string** | Unique ID that identifies the return discount only within this order. | [optional] [default to null]
**SourceDiscountUid** | **string** | &#x60;uid&#x60; of the Discount from the Order which contains the original application of this discount. | [optional] [default to null]
**CatalogObjectId** | **string** | The catalog object id referencing &#x60;CatalogDiscount&#x60;. | [optional] [default to null]
**Name** | **string** | The discount&#x27;s name. | [optional] [default to null]
**Type_** | **string** | The type of the discount. If it is created by API, it would be either &#x60;FIXED_PERCENTAGE&#x60; or &#x60;FIXED_AMOUNT&#x60;.  Discounts that don&#x27;t reference a catalog object ID must have a type of &#x60;FIXED_PERCENTAGE&#x60; or &#x60;FIXED_AMOUNT&#x60;. See [OrderLineItemDiscountType](#type-orderlineitemdiscounttype) for possible values | [optional] [default to null]
**Percentage** | **string** | The percentage of the tax, as a string representation of a decimal number. A value of &#x60;7.25&#x60; corresponds to a percentage of 7.25%.  &#x60;percentage&#x60; is not set for amount-based discounts. | [optional] [default to null]
**AmountMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**AppliedMoney** | [***Money**](Money.md) |  | [optional] [default to null]
**Scope** | **string** | Indicates the level at which the &#x60;OrderReturnDiscount&#x60; applies. For &#x60;ORDER&#x60; scoped discounts, the server will generate references in &#x60;applied_discounts&#x60; on all &#x60;OrderReturnLineItem&#x60;s. For &#x60;LINE_ITEM&#x60; scoped discounts, the discount will only apply to &#x60;OrderReturnLineItem&#x60;s with references in their &#x60;applied_discounts&#x60; field. See [OrderLineItemDiscountScope](#type-orderlineitemdiscountscope) for possible values | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

