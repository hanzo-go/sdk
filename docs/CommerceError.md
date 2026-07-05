# CommerceError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 

## Methods

### NewCommerceError

`func NewCommerceError() *CommerceError`

NewCommerceError instantiates a new CommerceError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceErrorWithDefaults

`func NewCommerceErrorWithDefaults() *CommerceError`

NewCommerceErrorWithDefaults instantiates a new CommerceError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *CommerceError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CommerceError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CommerceError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *CommerceError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetMessage

`func (o *CommerceError) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CommerceError) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CommerceError) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *CommerceError) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetCode

`func (o *CommerceError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CommerceError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CommerceError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *CommerceError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


