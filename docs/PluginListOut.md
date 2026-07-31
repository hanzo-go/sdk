# PluginListOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]PluginHost**](PluginHost.md) |  | [optional] 
**Drift** | Pointer to [**[]PluginDrift**](PluginDrift.md) |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 

## Methods

### NewPluginListOut

`func NewPluginListOut() *PluginListOut`

NewPluginListOut instantiates a new PluginListOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginListOutWithDefaults

`func NewPluginListOutWithDefaults() *PluginListOut`

NewPluginListOutWithDefaults instantiates a new PluginListOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *PluginListOut) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PluginListOut) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PluginListOut) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PluginListOut) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *PluginListOut) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *PluginListOut) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *PluginListOut) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *PluginListOut) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *PluginListOut) GetData() []PluginHost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PluginListOut) GetDataOk() (*[]PluginHost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PluginListOut) SetData(v []PluginHost)`

SetData sets Data field to given value.

### HasData

`func (o *PluginListOut) HasData() bool`

HasData returns a boolean if a field has been set.

### GetDrift

`func (o *PluginListOut) GetDrift() []PluginDrift`

GetDrift returns the Drift field if non-nil, zero value otherwise.

### GetDriftOk

`func (o *PluginListOut) GetDriftOk() (*[]PluginDrift, bool)`

GetDriftOk returns a tuple with the Drift field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDrift

`func (o *PluginListOut) SetDrift(v []PluginDrift)`

SetDrift sets Drift field to given value.

### HasDrift

`func (o *PluginListOut) HasDrift() bool`

HasDrift returns a boolean if a field has been set.

### GetData2

`func (o *PluginListOut) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *PluginListOut) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *PluginListOut) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *PluginListOut) HasData2() bool`

HasData2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


