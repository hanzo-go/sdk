# AuthorsError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to [**AuthorsErrorError**](AuthorsErrorError.md) |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **int32** |  | [optional] 

## Methods

### NewAuthorsError

`func NewAuthorsError() *AuthorsError`

NewAuthorsError instantiates a new AuthorsError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorsErrorWithDefaults

`func NewAuthorsErrorWithDefaults() *AuthorsError`

NewAuthorsErrorWithDefaults instantiates a new AuthorsError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *AuthorsError) GetError() AuthorsErrorError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *AuthorsError) GetErrorOk() (*AuthorsErrorError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *AuthorsError) SetError(v AuthorsErrorError)`

SetError sets Error field to given value.

### HasError

`func (o *AuthorsError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *AuthorsError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *AuthorsError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *AuthorsError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *AuthorsError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetStatus

`func (o *AuthorsError) GetStatus() int32`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AuthorsError) GetStatusOk() (*int32, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AuthorsError) SetStatus(v int32)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AuthorsError) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


