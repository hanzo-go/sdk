# O11yClusterListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]O11yClusterListRecord**](O11yClusterListRecord.md) |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yClusterListResponse

`func NewO11yClusterListResponse() *O11yClusterListResponse`

NewO11yClusterListResponse instantiates a new O11yClusterListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yClusterListResponseWithDefaults

`func NewO11yClusterListResponseWithDefaults() *O11yClusterListResponse`

NewO11yClusterListResponseWithDefaults instantiates a new O11yClusterListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *O11yClusterListResponse) GetRecords() []O11yClusterListRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yClusterListResponse) GetRecordsOk() (*[]O11yClusterListRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yClusterListResponse) SetRecords(v []O11yClusterListRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yClusterListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yClusterListResponse) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yClusterListResponse) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yClusterListResponse) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yClusterListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yClusterListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yClusterListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yClusterListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yClusterListResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


