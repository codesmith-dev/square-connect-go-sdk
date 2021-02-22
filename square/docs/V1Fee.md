# V1Fee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The fee&#x27;s unique ID. | [optional] [default to null]
**Name** | **string** | The fee&#x27;s name. | [optional] [default to null]
**Rate** | **string** | The rate of the fee, as a string representation of a decimal number. A value of 0.07 corresponds to a rate of 7%. | [optional] [default to null]
**CalculationPhase** | **string** | Forthcoming See [V1FeeCalculationPhase](#type-v1feecalculationphase) for possible values | [optional] [default to null]
**AdjustmentType** | **string** | The type of adjustment the fee applies to a payment. Currently, this value is TAX for all fees. See [V1FeeAdjustmentType](#type-v1feeadjustmenttype) for possible values | [optional] [default to null]
**AppliesToCustomAmounts** | **bool** | If true, the fee applies to custom amounts entered into Square Point of Sale that are not associated with a particular item. | [optional] [default to null]
**Enabled** | **bool** | If true, the fee is applied to all appropriate items. If false, the fee is not applied at all. | [optional] [default to null]
**InclusionType** | **string** | Whether the fee is ADDITIVE or INCLUSIVE. See [V1FeeInclusionType](#type-v1feeinclusiontype) for possible values | [optional] [default to null]
**Type_** | **string** | In countries with multiple classifications for sales taxes, indicates which classification the fee falls under. Currently relevant only to Canadian merchants. See [V1FeeType](#type-v1feetype) for possible values | [optional] [default to null]
**V2Id** | **string** | The ID of the CatalogObject in the Connect v2 API. Objects that are shared across multiple locations share the same v2 ID. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

