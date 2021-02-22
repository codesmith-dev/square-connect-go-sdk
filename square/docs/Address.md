# Address

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AddressLine1** | **string** | The first line of the address.  Fields that start with &#x60;address_line&#x60; provide the address&#x27;s most specific details, like street number, street name, and building name. They do *not* provide less specific details like city, state/province, or country (these details are provided in other fields). | [optional] [default to null]
**AddressLine2** | **string** | The second line of the address, if any. | [optional] [default to null]
**AddressLine3** | **string** | The third line of the address, if any. | [optional] [default to null]
**Locality** | **string** | The city or town of the address. | [optional] [default to null]
**Sublocality** | **string** | A civil region within the address&#x27;s &#x60;locality&#x60;, if any. | [optional] [default to null]
**Sublocality2** | **string** | A civil region within the address&#x27;s &#x60;sublocality&#x60;, if any. | [optional] [default to null]
**Sublocality3** | **string** | A civil region within the address&#x27;s &#x60;sublocality_2&#x60;, if any. | [optional] [default to null]
**AdministrativeDistrictLevel1** | **string** | A civil entity within the address&#x27;s country. In the US, this is the state. | [optional] [default to null]
**AdministrativeDistrictLevel2** | **string** | A civil entity within the address&#x27;s &#x60;administrative_district_level_1&#x60;. In the US, this is the county. | [optional] [default to null]
**AdministrativeDistrictLevel3** | **string** | A civil entity within the address&#x27;s &#x60;administrative_district_level_2&#x60;, if any. | [optional] [default to null]
**PostalCode** | **string** | The address&#x27;s postal code. | [optional] [default to null]
**Country** | **string** | The address&#x27;s country, in ISO 3166-1-alpha-2 format. See [Country](#type-country) for possible values | [optional] [default to null]
**FirstName** | **string** | Optional first name when it&#x27;s representing recipient. | [optional] [default to null]
**LastName** | **string** | Optional last name when it&#x27;s representing recipient. | [optional] [default to null]
**Organization** | **string** | Optional organization name when it&#x27;s representing recipient. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

