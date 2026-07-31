# CloudGrantsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudGrantRow**](CloudGrantRow.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Total** | Pointer to **int32** |  | [optional] 

## Methods

### NewCloudGrantsOut

`func NewCloudGrantsOut() *CloudGrantsOut`

NewCloudGrantsOut instantiates a new CloudGrantsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudGrantsOutWithDefaults

`func NewCloudGrantsOutWithDefaults() *CloudGrantsOut`

NewCloudGrantsOutWithDefaults instantiates a new CloudGrantsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudGrantsOut) GetData() []CloudGrantRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudGrantsOut) GetDataOk() (*[]CloudGrantRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudGrantsOut) SetData(v []CloudGrantRow)`

SetData sets Data field to given value.

### HasData

`func (o *CloudGrantsOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudGrantsOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudGrantsOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudGrantsOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudGrantsOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudGrantsOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudGrantsOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudGrantsOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudGrantsOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTotal

`func (o *CloudGrantsOut) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *CloudGrantsOut) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *CloudGrantsOut) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *CloudGrantsOut) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


