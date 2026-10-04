# OpenaiEmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dimensions** | Pointer to **int32** |  | [optional] 
**EncodingFormat** | Pointer to **string** |  | [optional] 
**Input** | Pointer to **interface{}** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewOpenaiEmbeddingRequest

`func NewOpenaiEmbeddingRequest() *OpenaiEmbeddingRequest`

NewOpenaiEmbeddingRequest instantiates a new OpenaiEmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiEmbeddingRequestWithDefaults

`func NewOpenaiEmbeddingRequestWithDefaults() *OpenaiEmbeddingRequest`

NewOpenaiEmbeddingRequestWithDefaults instantiates a new OpenaiEmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDimensions

`func (o *OpenaiEmbeddingRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *OpenaiEmbeddingRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *OpenaiEmbeddingRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *OpenaiEmbeddingRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetEncodingFormat

`func (o *OpenaiEmbeddingRequest) GetEncodingFormat() string`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *OpenaiEmbeddingRequest) GetEncodingFormatOk() (*string, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *OpenaiEmbeddingRequest) SetEncodingFormat(v string)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *OpenaiEmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.

### GetInput

`func (o *OpenaiEmbeddingRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *OpenaiEmbeddingRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *OpenaiEmbeddingRequest) SetInput(v interface{})`

SetInput sets Input field to given value.

### HasInput

`func (o *OpenaiEmbeddingRequest) HasInput() bool`

HasInput returns a boolean if a field has been set.

### SetInputNil

`func (o *OpenaiEmbeddingRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *OpenaiEmbeddingRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetModel

`func (o *OpenaiEmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OpenaiEmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OpenaiEmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OpenaiEmbeddingRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetUser

`func (o *OpenaiEmbeddingRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *OpenaiEmbeddingRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *OpenaiEmbeddingRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *OpenaiEmbeddingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


