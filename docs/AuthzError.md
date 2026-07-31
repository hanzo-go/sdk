# AuthzError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** | Always \&quot;error\&quot; for error responses | 
**Msg** | **string** | Human-readable error message | 
**Code** | Pointer to **int32** | Application-specific error code | [optional] 
**Data** | Pointer to **map[string]interface{}** | Additional error context | [optional] 
**Data2** | Pointer to **map[string]interface{}** | Additional error details | [optional] 
**RequestId** | Pointer to **string** | Request ID for debugging | [optional] 

## Methods

### NewAuthzError

`func NewAuthzError(status string, msg string, ) *AuthzError`

NewAuthzError instantiates a new AuthzError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzErrorWithDefaults

`func NewAuthzErrorWithDefaults() *AuthzError`

NewAuthzErrorWithDefaults instantiates a new AuthzError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthzError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthzError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthzError) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetMsg

`func (o *AuthzError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AuthzError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AuthzError) SetMsg(v string)`

SetMsg sets Msg field to given value.


### GetCode

`func (o *AuthzError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AuthzError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AuthzError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *AuthzError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetData

`func (o *AuthzError) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AuthzError) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AuthzError) SetData(v map[string]interface{})`

SetData sets Data field to given value.

### HasData

`func (o *AuthzError) HasData() bool`

HasData returns a boolean if a field has been set.

### GetData2

`func (o *AuthzError) GetData2() map[string]interface{}`

GetData2 returns the Data2 field if non-nil, zero value otherwise.

### GetData2Ok

`func (o *AuthzError) GetData2Ok() (*map[string]interface{}, bool)`

GetData2Ok returns a tuple with the Data2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData2

`func (o *AuthzError) SetData2(v map[string]interface{})`

SetData2 sets Data2 field to given value.

### HasData2

`func (o *AuthzError) HasData2() bool`

HasData2 returns a boolean if a field has been set.

### GetRequestId

`func (o *AuthzError) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *AuthzError) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *AuthzError) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *AuthzError) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


