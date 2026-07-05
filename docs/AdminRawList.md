# AdminRawList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminRawList

`func NewAdminRawList() *AdminRawList`

NewAdminRawList instantiates a new AdminRawList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminRawListWithDefaults

`func NewAdminRawListWithDefaults() *AdminRawList`

NewAdminRawListWithDefaults instantiates a new AdminRawList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminRawList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminRawList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminRawList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminRawList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AdminRawList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminRawList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminRawList) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminRawList) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AdminRawList) GetData() []map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminRawList) GetDataOk() (*[]map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminRawList) SetData(v []map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AdminRawList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AdminRawList) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AdminRawList) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AdminRawList) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AdminRawList) HasData2() bool`

HasData2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


