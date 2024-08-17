# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**V1ListOrders**](V1TransactionsApi.md#V1ListOrders) | **Get** /v1/{location_id}/orders | V1ListOrders
[**V1RetrieveOrder**](V1TransactionsApi.md#V1RetrieveOrder) | **Get** /v1/{location_id}/orders/{order_id} | V1RetrieveOrder
[**V1UpdateOrder**](V1TransactionsApi.md#V1UpdateOrder) | **Put** /v1/{location_id}/orders/{order_id} | V1UpdateOrder

# **V1ListOrders**
> []V1Order V1ListOrders(ctx, locationId, optional)
V1ListOrders

Provides summary information for a merchant's online store orders.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the location to list online store orders for. | 
 **optional** | ***V1TransactionsApiV1ListOrdersOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a V1TransactionsApiV1ListOrdersOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **optional.String**| The order in which payments are listed in the response. | 
 **limit** | **optional.Int32**| The maximum number of payments to return in a single response. This value cannot exceed 200. | 
 **batchToken** | **optional.String**| A pagination cursor to retrieve the next set of results for your original query to the endpoint. | 

### Return type

[**[]V1Order**](V1Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1RetrieveOrder**
> V1Order V1RetrieveOrder(ctx, locationId, orderId)
V1RetrieveOrder

Provides comprehensive information for a single online store order, including the order's history.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **locationId** | **string**| The ID of the order&#x27;s associated location. | 
  **orderId** | **string**| The order&#x27;s Square-issued ID. You obtain this value from Order objects returned by the List Orders endpoint | 

### Return type

[**V1Order**](V1Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **V1UpdateOrder**
> V1Order V1UpdateOrder(ctx, body, locationId, orderId)
V1UpdateOrder

Updates the details of an online store order. Every update you perform on an order corresponds to one of three actions:

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**V1UpdateOrderRequest**](V1UpdateOrderRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 
  **locationId** | **string**| The ID of the order&#x27;s associated location. | 
  **orderId** | **string**| The order&#x27;s Square-issued ID. You obtain this value from Order objects returned by the List Orders endpoint | 

### Return type

[**V1Order**](V1Order.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

