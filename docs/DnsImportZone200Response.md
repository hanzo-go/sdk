# DnsImportZone200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RecordsAdded** | Pointer to **int32** |  | [optional] 
**RecordsUpdated** | Pointer to **int32** |  | [optional] 
**Errors** | Pointer to [**[]DnsImportZone200ResponseErrorsInner**](DnsImportZone200ResponseErrorsInner.md) |  | [optional] 

## Methods

### NewDnsImportZone200Response

`func NewDnsImportZone200Response() *DnsImportZone200Response`

NewDnsImportZone200Response instantiates a new DnsImportZone200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDnsImportZone200ResponseWithDefaults

`func NewDnsImportZone200ResponseWithDefaults() *DnsImportZone200Response`

NewDnsImportZone200ResponseWithDefaults instantiates a new DnsImportZone200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecordsAdded

`func (o *DnsImportZone200Response) GetRecordsAdded() int32`

GetRecordsAdded returns the RecordsAdded field if non-nil, zero value otherwise.

### GetRecordsAddedOk

`func (o *DnsImportZone200Response) GetRecordsAddedOk() (*int32, bool)`

GetRecordsAddedOk returns a tuple with the RecordsAdded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsAdded

`func (o *DnsImportZone200Response) SetRecordsAdded(v int32)`

SetRecordsAdded sets RecordsAdded field to given value.

### HasRecordsAdded

`func (o *DnsImportZone200Response) HasRecordsAdded() bool`

HasRecordsAdded returns a boolean if a field has been set.

### GetRecordsUpdated

`func (o *DnsImportZone200Response) GetRecordsUpdated() int32`

GetRecordsUpdated returns the RecordsUpdated field if non-nil, zero value otherwise.

### GetRecordsUpdatedOk

`func (o *DnsImportZone200Response) GetRecordsUpdatedOk() (*int32, bool)`

GetRecordsUpdatedOk returns a tuple with the RecordsUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordsUpdated

`func (o *DnsImportZone200Response) SetRecordsUpdated(v int32)`

SetRecordsUpdated sets RecordsUpdated field to given value.

### HasRecordsUpdated

`func (o *DnsImportZone200Response) HasRecordsUpdated() bool`

HasRecordsUpdated returns a boolean if a field has been set.

### GetErrors

`func (o *DnsImportZone200Response) GetErrors() []DnsImportZone200ResponseErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DnsImportZone200Response) GetErrorsOk() (*[]DnsImportZone200ResponseErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DnsImportZone200Response) SetErrors(v []DnsImportZone200ResponseErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DnsImportZone200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


