# AdminUserList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Data** | Pointer to [**[]AdminOperatorUser**](AdminOperatorUser.md) |  | [optional] 
**Data2** | Pointer to **int32** |  | [optional] 

## Methods

### NewAdminUserList

`func NewAdminUserList() *AdminUserList`

NewAdminUserList instantiates a new AdminUserList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserListWithDefaults

`func NewAdminUserListWithDefaults() *AdminUserList`

NewAdminUserListWithDefaults instantiates a new AdminUserList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AdminUserList) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminUserList) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminUserList) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminUserList) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AdminUserList) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AdminUserList) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AdminUserList) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AdminUserList) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetData

`func (o *AdminUserList) GetData() []AdminOperatorUser`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AdminUserList) GetDataOk() (*[]AdminOperatorUser, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AdminUserList) SetData(v []AdminOperatorUser)`

SetData sets Data field to given value.

### HasData

`func (o *AdminUserList) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AdminUserList) GetData2() int32`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AdminUserList) GetData2Ok() (*int32, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AdminUserList) SetData2(v int32)`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AdminUserList) HasData2() bool`

HasData2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


