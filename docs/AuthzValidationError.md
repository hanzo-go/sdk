# AuthzValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]AuthzValidationErrorErrorsInner**](AuthzValidationErrorErrorsInner.md) |  | [optional] 

## Methods

### NewAuthzValidationError

`func NewAuthzValidationError() *AuthzValidationError`

NewAuthzValidationError instantiates a new AuthzValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthzValidationErrorWithDefaults

`func NewAuthzValidationErrorWithDefaults() *AuthzValidationError`

NewAuthzValidationErrorWithDefaults instantiates a new AuthzValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *AuthzValidationError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthzValidationError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthzValidationError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthzValidationError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *AuthzValidationError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AuthzValidationError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AuthzValidationError) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AuthzValidationError) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrors

`func (o *AuthzValidationError) GetErrors() []AuthzValidationErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AuthzValidationError) GetErrorsOk() (*[]AuthzValidationErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AuthzValidationError) SetErrors(v []AuthzValidationErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AuthzValidationError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


