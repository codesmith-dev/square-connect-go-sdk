# BulkCreateVendorsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Any errors that occurred during the request. | [optional] [default to null]
**Responses** | [**map[string]CreateVendorResponse**](CreateVendorResponse.md) | A set of [CreateVendorResponse](https://developer.squareup.com/reference/square_2024-07-17/objects/CreateVendorResponse) objects encapsulating successfully created [Vendor](https://developer.squareup.com/reference/square_2024-07-17/objects/Vendor) objects or error responses for failed attempts. The set is represented by  a collection of idempotency-key/&#x60;Vendor&#x60;-object or idempotency-key/error-object pairs. The idempotency keys correspond to those specified in the input. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

