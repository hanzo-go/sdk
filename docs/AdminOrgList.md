# AdminOrgList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AdminOrgRow**](AdminOrgRow.md) |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminOrgList

`func NewAdminOrgList() *AdminOrgList`

NewAdminOrgList instantiates a new AdminOrgList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminOrgListWithDefaults

`func NewAdminOrgListWithDefaults() *AdminOrgList`

NewAdminOrgListWithDefaults instantiates a new AdminOrgList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminOrgList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminOrgList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminOrgList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminOrgList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AdminOrgList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminOrgList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminOrgList) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminOrgList) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AdminOrgList) GetData() []AdminOrgRow`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminOrgList) GetDataOk() (*[]AdminOrgRow, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminOrgList) SetData(v []AdminOrgRow)`

SetData sets Data field to given value.

### HasData

`func (o *AdminOrgList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AdminOrgList) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AdminOrgList) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AdminOrgList) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AdminOrgList) HasData2() bool`

HasData2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


