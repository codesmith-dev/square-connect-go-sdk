# {{classname}}

All URIs are relative to *https://connect.squareup.com/*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DisableEvents**](EventsApi.md#DisableEvents) | **Put** /v2/events/disable | DisableEvents
[**EnableEvents**](EventsApi.md#EnableEvents) | **Put** /v2/events/enable | EnableEvents
[**ListEventTypes**](EventsApi.md#ListEventTypes) | **Get** /v2/events/types | ListEventTypes
[**SearchEvents**](EventsApi.md#SearchEvents) | **Post** /v2/events | SearchEvents

# **DisableEvents**
> DisableEventsResponse DisableEvents(ctx, )
DisableEvents

Disables events to prevent them from being searchable. All events are disabled by default. You must enable events to make them searchable. Disabling events for a specific time period prevents them from being searchable, even if you re-enable them later.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**DisableEventsResponse**](DisableEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **EnableEvents**
> EnableEventsResponse EnableEvents(ctx, )
EnableEvents

Enables events to make them searchable. Only events that occur while in the enabled state are searchable.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**EnableEventsResponse**](EnableEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ListEventTypes**
> ListEventTypesResponse ListEventTypes(ctx, optional)
ListEventTypes

Lists all event types that you can subscribe to as webhooks or query using the Events API.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***EventsApiListEventTypesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a EventsApiListEventTypesOpts struct
Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **apiVersion** | **optional.String**| The API version for which to list event types. Setting this field overrides the default version used by the application. | 

### Return type

[**ListEventTypesResponse**](ListEventTypesResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SearchEvents**
> SearchEventsResponse SearchEvents(ctx, body)
SearchEvents

Search for Square API events that occur within a 28-day timeframe.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SearchEventsRequest**](SearchEventsRequest.md)| An object containing the fields to POST for the request.

See the corresponding object definition for field details. | 

### Return type

[**SearchEventsResponse**](SearchEventsResponse.md)

### Authorization

[oauth2](../README.md#oauth2)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

