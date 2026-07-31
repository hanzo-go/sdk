# CloudActionOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]CloudResult**](CloudResult.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudActionOut

`func NewCloudActionOut() *CloudActionOut`

NewCloudActionOut instantiates a new CloudActionOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudActionOutWithDefaults

`func NewCloudActionOutWithDefaults() *CloudActionOut`

NewCloudActionOutWithDefaults instantiates a new CloudActionOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudActionOut) GetData() []CloudResult`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudActionOut) GetDataOk() (*[]CloudResult, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudActionOut) SetData(v []CloudResult)`

SetData sets Data field to given value.

### HasData

`func (o *CloudActionOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudActionOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudActionOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudActionOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudActionOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudActionOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudActionOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudActionOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudActionOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


