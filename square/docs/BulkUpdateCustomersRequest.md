# BulkUpdateCustomersRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Customers** | [**map[string]BulkUpdateCustomerData**](BulkUpdateCustomerData.md) | A map of 1 to 100 individual update requests, represented by &#x60;customer ID: { customer data }&#x60; key-value pairs.  Each key is the ID of the [customer profile](https://developer.squareup.com/reference/square_2024-07-17/objects/Customer) to update. To update a customer profile that was created by merging existing profiles, provide the ID of the newly created profile.  Each value contains the updated customer data. Only new or changed fields are required. To add or update a field, specify the new value. To remove a field, specify &#x60;null&#x60;. | [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

