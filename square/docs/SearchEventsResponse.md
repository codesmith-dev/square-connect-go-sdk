# SearchEventsResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | [**[]ModelError**](Error.md) | Information on errors encountered during the request. | [optional] [default to null]
**Events** | [**[]Event**](Event.md) | The list of [Event](https://developer.squareup.com/reference/square_2024-07-17/objects/Event)s returned by the search. | [optional] [default to null]
**Metadata** | [**[]EventMetadata**](EventMetadata.md) | Contains the metadata of an event. For more information, see [Event](https://developer.squareup.com/reference/square_2024-07-17/objects/Event). | [optional] [default to null]
**Cursor** | **string** | When a response is truncated, it includes a cursor that you can use in a subsequent request to fetch the next set of events. If empty, this is the final response.  For more information, see [Pagination](https://developer.squareup.com/docs/build-basics/common-api-patterns/pagination). | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

