# S3Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** |  | [optional] 
**Message** | Pointer to **string** |  | [optional] 
**Resource** | Pointer to **string** |  | [optional] 
**RequestId** | Pointer to **string** |  | [optional] 

## Methods

### NewS3Error

`func NewS3Error() *S3Error`

NewS3Error instantiates a new S3Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewS3ErrorWithDefaults

`func NewS3ErrorWithDefaults() *S3Error`

NewS3ErrorWithDefaults instantiates a new S3Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *S3Error) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *S3Error) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *S3Error) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *S3Error) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetMessage

`func (o *S3Error) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *S3Error) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *S3Error) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *S3Error) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetResource

`func (o *S3Error) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *S3Error) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *S3Error) SetResource(v string)`

SetResource sets Resource field to given value.

### HasResource

`func (o *S3Error) HasResource() bool`

HasResource returns a boolean if a field has been set.

### GetRequestId

`func (o *S3Error) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *S3Error) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *S3Error) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *S3Error) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


