# PubsubError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **int32** |  | [optional] 

## Methods

### NewPubsubError

`func NewPubsubError() *PubsubError`

NewPubsubError instantiates a new PubsubError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPubsubErrorWithDefaults

`func NewPubsubErrorWithDefaults() *PubsubError`

NewPubsubErrorWithDefaults instantiates a new PubsubError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *PubsubError) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *PubsubError) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *PubsubError) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *PubsubError) HasError() bool`

HasError returns a boolean if a field has been set.

### GetCode

`func (o *PubsubError) GetCode() int32`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *PubsubError) GetCodeOk() (*int32, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *PubsubError) SetCode(v int32)`

SetCode sets Code field to given value.

### HasCode

`func (o *PubsubError) HasCode() bool`

HasCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


