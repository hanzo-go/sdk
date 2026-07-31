# CloudAccessOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**CloudAccessChange**](CloudAccessChange.md) |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewCloudAccessOut

`func NewCloudAccessOut() *CloudAccessOut`

NewCloudAccessOut instantiates a new CloudAccessOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccessOutWithDefaults

`func NewCloudAccessOutWithDefaults() *CloudAccessOut`

NewCloudAccessOutWithDefaults instantiates a new CloudAccessOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *CloudAccessOut) GetData() CloudAccessChange`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *CloudAccessOut) GetDataOk() (*CloudAccessChange, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *CloudAccessOut) SetData(v CloudAccessChange)`

SetData sets Data field to given value.

### HasData

`func (o *CloudAccessOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetMsg

`func (o *CloudAccessOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *CloudAccessOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *CloudAccessOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *CloudAccessOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *CloudAccessOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CloudAccessOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CloudAccessOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CloudAccessOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


