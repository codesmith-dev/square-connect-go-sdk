# V1PageCell

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PageId** | **string** | The unique identifier of the page the cell is included on. | [optional] [default to null]
**Row** | **int32** | The row of the cell. Always an integer between 0 and 4, inclusive. | [optional] [default to null]
**Column** | **int32** | The column of the cell. Always an integer between 0 and 4, inclusive. | [optional] [default to null]
**ObjectType** | **string** | The type of entity represented in the cell (ITEM, DISCOUNT, CATEGORY, or PLACEHOLDER). See [V1PageCellObjectType](#type-v1pagecellobjecttype) for possible values | [optional] [default to null]
**ObjectId** | **string** | The unique identifier of the entity represented in the cell. Not present for cells with an object_type of PLACEHOLDER. | [optional] [default to null]
**PlaceholderType** | **string** | For a cell with an object_type of PLACEHOLDER, this value indicates the cell&#x27;s special behavior. See [V1PageCellPlaceholderType](#type-v1pagecellplaceholdertype) for possible values | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

