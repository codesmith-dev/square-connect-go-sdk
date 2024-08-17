# BulkCreateCustomersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**map[string]BulkCreateCustomerData**](BulkCreateCustomerData.md) | A map of 1 to 100 individual create requests, represented by &#x60;idempotency key: { customer data }&#x60; key-value pairs.  Each key is an [idempotency key](https://developer.squareup.com/docs/build-basics/common-api-patterns/idempotency) that uniquely identifies the create request. Each value contains the customer data used to create the customer profile. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

