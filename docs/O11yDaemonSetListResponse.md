# O11yDaemonSetListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Records** | Pointer to [**[]O11yDaemonSetListRecord**](O11yDaemonSetListRecord.md) |  | [optional] 
**Total** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewO11yDaemonSetListResponse

`func NewO11yDaemonSetListResponse() *O11yDaemonSetListResponse`

NewO11yDaemonSetListResponse instantiates a new O11yDaemonSetListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yDaemonSetListResponseWithDefaults

`func NewO11yDaemonSetListResponseWithDefaults() *O11yDaemonSetListResponse`

NewO11yDaemonSetListResponseWithDefaults instantiates a new O11yDaemonSetListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecords

`func (o *O11yDaemonSetListResponse) GetRecords() []O11yDaemonSetListRecord`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *O11yDaemonSetListResponse) GetRecordsOk() (*[]O11yDaemonSetListRecord, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *O11yDaemonSetListResponse) SetRecords(v []O11yDaemonSetListRecord)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *O11yDaemonSetListResponse) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetTotal

`func (o *O11yDaemonSetListResponse) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *O11yDaemonSetListResponse) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *O11yDaemonSetListResponse) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *O11yDaemonSetListResponse) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetType

`func (o *O11yDaemonSetListResponse) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yDaemonSetListResponse) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yDaemonSetListResponse) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yDaemonSetListResponse) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


