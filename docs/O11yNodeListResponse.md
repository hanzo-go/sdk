# O11yNodeListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]O11yNodeListRecord**](O11yNodeListRecord.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yNodeListResponse

`func NewO11yNodeListResponse() *O11yNodeListResponse`

NewO11yNodeListResponse instantiates a new O11yNodeListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yNodeListResponseWithDefaults

`func NewO11yNodeListResponseWithDefaults() *O11yNodeListResponse`

NewO11yNodeListResponseWithDefaults instantiates a new O11yNodeListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *O11yNodeListResponse) GetRecords() []O11yNodeListRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yNodeListResponse) GetRecordsOk() (*[]O11yNodeListRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yNodeListResponse) SetRecords(v []O11yNodeListRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yNodeListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yNodeListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yNodeListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yNodeListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yNodeListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yNodeListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yNodeListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yNodeListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yNodeListResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


