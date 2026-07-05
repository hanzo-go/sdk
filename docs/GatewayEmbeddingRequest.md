# GatewayEmbeddingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | [**GatewayCreateCompletionRequestPrompt**](GatewayCreateCompletionRequestPrompt.md) |  | 
**EncodingFormat** | Pointer to **string** |  | [optional] [default to "float"]
**Dimensions** | Pointer to **int32** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewGatewayEmbeddingRequest

`func NewGatewayEmbeddingRequest(model string, input GatewayCreateCompletionRequestPrompt, ) *GatewayEmbeddingRequest`

NewGatewayEmbeddingRequest instantiates a new GatewayEmbeddingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGatewayEmbeddingRequestWithDefaults

`func NewGatewayEmbeddingRequestWithDefaults() *GatewayEmbeddingRequest`

NewGatewayEmbeddingRequestWithDefaults instantiates a new GatewayEmbeddingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *GatewayEmbeddingRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GatewayEmbeddingRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GatewayEmbeddingRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *GatewayEmbeddingRequest) GetInput() GatewayCreateCompletionRequestPrompt`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *GatewayEmbeddingRequest) GetInputOk() (*GatewayCreateCompletionRequestPrompt, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *GatewayEmbeddingRequest) SetInput(v GatewayCreateCompletionRequestPrompt)`

SetInput sets Input field to given value.


### GetEncodingFormat

`func (o *GatewayEmbeddingRequest) GetEncodingFormat() string`

GetEncodingFormat returns the EncodingFormat field if non-nil, zero value otherwise.

### GetEncodingFormatOk

`func (o *GatewayEmbeddingRequest) GetEncodingFormatOk() (*string, bool)`

GetEncodingFormatOk returns a tuple with the EncodingFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncodingFormat

`func (o *GatewayEmbeddingRequest) SetEncodingFormat(v string)`

SetEncodingFormat sets EncodingFormat field to given value.

### HasEncodingFormat

`func (o *GatewayEmbeddingRequest) HasEncodingFormat() bool`

HasEncodingFormat returns a boolean if a field has been set.

### GetDimensions

`func (o *GatewayEmbeddingRequest) GetDimensions() int32`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *GatewayEmbeddingRequest) GetDimensionsOk() (*int32, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *GatewayEmbeddingRequest) SetDimensions(v int32)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *GatewayEmbeddingRequest) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### GetUser

`func (o *GatewayEmbeddingRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *GatewayEmbeddingRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *GatewayEmbeddingRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *GatewayEmbeddingRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


