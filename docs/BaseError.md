# BaseError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 

## Methods

### NewBaseError

`func NewBaseError() *BaseError`

NewBaseError instantiates a new BaseError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseErrorWithDefaults

`func NewBaseErrorWithDefaults() *BaseError`

NewBaseErrorWithDefaults instantiates a new BaseError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *BaseError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *BaseError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *BaseError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *BaseError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *BaseError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *BaseError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *BaseError) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *BaseError) HasMsg() bool`

HasMsg returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


