# GuardError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 

## Methods

### NewGuardError

`func NewGuardError() *GuardError`

NewGuardError instantiates a new GuardError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuardErrorWithDefaults

`func NewGuardErrorWithDefaults() *GuardError`

NewGuardErrorWithDefaults instantiates a new GuardError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *GuardError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *GuardError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *GuardError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *GuardError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCode

`func (o *GuardError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GuardError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GuardError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *GuardError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


