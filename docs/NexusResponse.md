# NexusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]interface{}** | Response data (string, object, or array) | [optional] 
**Data2** | Pointer to **map[string]interface{}** | Secondary response data | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusResponse

`func NewNexusResponse() *NexusResponse`

NewNexusResponse instantiates a new NexusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusResponseWithDefaults

`func NewNexusResponseWithDefaults() *NexusResponse`

NewNexusResponseWithDefaults instantiates a new NexusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *NexusResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *NexusResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *NexusResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *NexusResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *NexusResponse) GetData2() map[string]interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *NexusResponse) GetData2Ok() (*map[string]interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *NexusResponse) SetData2(v map[string]interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *NexusResponse) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### GetMsg

`func (o *NexusResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *NexusResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *NexusResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *NexusResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetStatus

`func (o *NexusResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NexusResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NexusResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NexusResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


