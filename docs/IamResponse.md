# IamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Code is a STABLE machine-readable reason, where the human &#x60;msg&#x60; is deliberately generic. &#x60;msg&#x60; is prose for a person and several distinct causes legitimately share one sentence; a caller that must BRANCH on the cause — or tell its own user which of them happened — cannot parse prose. Optional, so every existing envelope is byte-identical and no SDK changes. | [optional] 
**Data** | Pointer to **map[string]interface{}** |  | [optional] 
**Data2** | Pointer to **map[string]interface{}** |  | [optional] 
**Data3** | Pointer to **map[string]interface{}** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Sub** | Pointer to **string** |  | [optional] 

## Methods

### NewIamResponse

`func NewIamResponse() *IamResponse`

NewIamResponse instantiates a new IamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIamResponseWithDefaults

`func NewIamResponseWithDefaults() *IamResponse`

NewIamResponseWithDefaults instantiates a new IamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *IamResponse) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *IamResponse) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *IamResponse) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *IamResponse) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *IamResponse) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *IamResponse) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *IamResponse) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *IamResponse) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *IamResponse) GetData2() map[string]interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *IamResponse) GetData2Ok() (*map[string]interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *IamResponse) SetData2(v map[string]interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *IamResponse) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### GetData3

`func (o *IamResponse) GetData3() map[string]interface{}`

GetData3 returns the Data3 field if non-nil, zero value otherwise.

### GetData3Ok

`func (o *IamResponse) GetData3Ok() (*map[string]interface{}, bool)`

GetData3Ok returns a tuple with the Data3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData3

`func (o *IamResponse) SetData3(v map[string]interface{})`

SetData3 sets Data3 field to given value.

### HasData3

`func (o *IamResponse) HasData3() bool`

HasData3 returns a boolean if a field has been set.

### GetMsg

`func (o *IamResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *IamResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *IamResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *IamResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetName

`func (o *IamResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IamResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IamResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IamResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *IamResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IamResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IamResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IamResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSub

`func (o *IamResponse) GetSub() string`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *IamResponse) GetSubOk() (*string, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *IamResponse) SetSub(v string)`

SetSub sets Sub field to given value.

### HasSub

`func (o *IamResponse) HasSub() bool`

HasSub returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


