# NotifyValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** |  | [optional] 
**Msg** | Pointer to **string** |  | [optional] 
**Errors** | Pointer to [**[]AuthzValidationErrorErrorsInner**](AuthzValidationErrorErrorsInner.md) |  | [optional] 

## Methods

### NewNotifyValidationError

`func NewNotifyValidationError() *NotifyValidationError`

NewNotifyValidationError instantiates a new NotifyValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNotifyValidationErrorWithDefaults

`func NewNotifyValidationErrorWithDefaults() *NotifyValidationError`

NewNotifyValidationErrorWithDefaults instantiates a new NotifyValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *NotifyValidationError) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NotifyValidationError) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NotifyValidationError) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NotifyValidationError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetMsg

`func (o *NotifyValidationError) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *NotifyValidationError) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *NotifyValidationError) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *NotifyValidationError) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrors

`func (o *NotifyValidationError) GetErrors() []AuthzValidationErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *NotifyValidationError) GetErrorsOk() (*[]AuthzValidationErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *NotifyValidationError) SetErrors(v []AuthzValidationErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *NotifyValidationError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


