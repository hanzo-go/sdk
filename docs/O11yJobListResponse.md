# O11yJobListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]O11yJobListRecord**](O11yJobListRecord.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yJobListResponse

`func NewO11yJobListResponse() *O11yJobListResponse`

NewO11yJobListResponse instantiates a new O11yJobListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yJobListResponseWithDefaults

`func NewO11yJobListResponseWithDefaults() *O11yJobListResponse`

NewO11yJobListResponseWithDefaults instantiates a new O11yJobListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *O11yJobListResponse) GetRecords() []O11yJobListRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yJobListResponse) GetRecordsOk() (*[]O11yJobListRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yJobListResponse) SetRecords(v []O11yJobListRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yJobListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yJobListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yJobListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yJobListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yJobListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yJobListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yJobListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yJobListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yJobListResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


