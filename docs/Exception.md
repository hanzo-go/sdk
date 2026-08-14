# Exception

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Frames** | Pointer to [**[]Frame**](Frame.md) |  | [optional] 
**Handled** | Pointer to **bool** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Stack** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewException

`func NewException() *Exception`

NewException instantiates a new Exception object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExceptionWithDefaults

`func NewExceptionWithDefaults() *Exception`

NewExceptionWithDefaults instantiates a new Exception object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrames

`func (o *Exception) GetFrames() []Frame`

GetFrames returns the Frames field if non-nil, zero value otherwise.

### GetFramesOk

`func (o *Exception) GetFramesOk() (*[]Frame, bool)`

GetFramesOk returns a tuple with the Frames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrames

`func (o *Exception) SetFrames(v []Frame)`

SetFrames sets Frames field to given value.

### HasFrames

`func (o *Exception) HasFrames() bool`

HasFrames returns a boolean if a field has been set.

### GetHandled

`func (o *Exception) GetHandled() bool`

GetHandled returns the Handled field if non-nil, zero value otherwise.

### GetHandledOk

`func (o *Exception) GetHandledOk() (*bool, bool)`

GetHandledOk returns a tuple with the Handled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandled

`func (o *Exception) SetHandled(v bool)`

SetHandled sets Handled field to given value.

### HasHandled

`func (o *Exception) HasHandled() bool`

HasHandled returns a boolean if a field has been set.

### GetMessage

`func (o *Exception) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *Exception) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *Exception) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *Exception) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStack

`func (o *Exception) GetStack() string`

GetStack returns the Stack field if non-nil, zero value otherwise.

### GetStackOk

`func (o *Exception) GetStackOk() (*string, bool)`

GetStackOk returns a tuple with the Stack field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStack

`func (o *Exception) SetStack(v string)`

SetStack sets Stack field to given value.

### HasStack

`func (o *Exception) HasStack() bool`

HasStack returns a boolean if a field has been set.

### GetType

`func (o *Exception) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Exception) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Exception) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Exception) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


