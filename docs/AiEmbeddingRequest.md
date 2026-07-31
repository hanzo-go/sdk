# AiEmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | **interface{}** |  | 
**EncodingFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewAiEmbeddingRequest

`func NewAiEmbeddingRequest(model string, input interface{}, ) *AiEmbeddingRequest`

NewAiEmbeddingRequest instantiates a new AiEmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiEmbeddingRequestWithDefaults

`func NewAiEmbeddingRequestWithDefaults() *AiEmbeddingRequest`

NewAiEmbeddingRequestWithDefaults instantiates a new AiEmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiEmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiEmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiEmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *AiEmbeddingRequest) GetInput() interface{}`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiEmbeddingRequest) GetInputOk() (*interface{}, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiEmbeddingRequest) SetInput(v interface{})`

SetInput sets Input field to given value.


### SetInputNil

`func (o *AiEmbeddingRequest) SetInputNil(b bool)`

 SetInputNil sets the value for Input to be an explicit nil

### UnsetInput
`func (o *AiEmbeddingRequest) UnsetInput()`

UnsetInput ensures that no value is present for Input, not even an explicit nil
### GetEncodingFormat

`func (o *AiEmbeddingRequest) GetEncodingFormat() string`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *AiEmbeddingRequest) GetEncodingFormatOk() (*string, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *AiEmbeddingRequest) SetEncodingFormat(v string)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *AiEmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


