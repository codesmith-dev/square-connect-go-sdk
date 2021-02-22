# V1Discount

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The discount&#x27;s unique ID. | [optional] [default to null]
**Name** | **string** | The discount&#x27;s name. | [optional] [default to null]
**Rate** | **string** | The rate of the discount, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%. This rate is 0 if discount_type is VARIABLE_PERCENTAGE. | [optional] [default to null]
**AmountMoney** | [***V1Money**](V1Money.md) |  | [optional] [default to null]
**DiscountType** | **string** | Indicates whether the discount is a FIXED value or entered at the time of sale. See [V1DiscountDiscountType](#type-v1discountdiscounttype) for possible values | [optional] [default to null]
**PinRequired** | **bool** | Indicates whether a mobile staff member needs to enter their PIN to apply the discount to a payment. | [optional] [default to null]
**Color** | **string** | The color of the discount&#x27;s display label in Square Point of Sale, if not the default color. The default color is 9da2a6. See [V1DiscountColor](#type-v1discountcolor) for possible values | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

