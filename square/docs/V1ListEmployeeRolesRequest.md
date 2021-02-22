# V1ListEmployeeRolesRequest

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | **string** | The order in which employees are listed in the response, based on their created_at field.Default value: ASC See [SortOrder](#type-sortorder) for possible values | [optional] [default to null]
**Limit** | **int32** | The maximum integer number of employee entities to return in a single response. Default 100, maximum 200. | [optional] [default to null]
**BatchToken** | **string** | A pagination cursor to retrieve the next set of results for your original query to the endpoint. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

