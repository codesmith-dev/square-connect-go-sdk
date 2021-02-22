# OrderFulfillmentUpdatedUpdate

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FulfillmentUid** | **string** | Unique ID that identifies the fulfillment only within this order. | [optional] [default to null]
**OldState** | **string** | The state of the fulfillment before the change. Will not be populated if the fulfillment is created with this new Order version. See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values | [optional] [default to null]
**NewState** | **string** | The state of the fulfillment after the change. May be equal to old_state if a non-state field was changed on the fulfillment (e.g. tracking number). See [OrderFulfillmentState](#type-orderfulfillmentstate) for possible values | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

