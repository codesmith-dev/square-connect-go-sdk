# V1Employee

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The employee&#x27;s unique ID. | [optional] [default to null]
**FirstName** | **string** | The employee&#x27;s first name. | [default to null]
**LastName** | **string** | The employee&#x27;s last name. | [default to null]
**RoleIds** | **[]string** | The ids of the employee&#x27;s associated roles. Currently, you can specify only one or zero roles per employee. | [optional] [default to null]
**AuthorizedLocationIds** | **[]string** | The IDs of the locations the employee is allowed to clock in at. | [optional] [default to null]
**Email** | **string** | The employee&#x27;s email address. | [optional] [default to null]
**Status** | **string** | CWhether the employee is ACTIVE or INACTIVE. Inactive employees cannot sign in to Square Register.Merchants update this field from the Square Dashboard. See [V1EmployeeStatus](#type-v1employeestatus) for possible values | [optional] [default to null]
**ExternalId** | **string** | An ID the merchant can set to associate the employee with an entity in another system. | [optional] [default to null]
**CreatedAt** | **string** | The time when the employee entity was created, in ISO 8601 format. | [optional] [default to null]
**UpdatedAt** | **string** | The time when the employee entity was most recently updated, in ISO 8601 format. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

