# AiImageGenerationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | Pointer to **string** |  | [optional] 
**Prompt** | **string** |  | 
**N** | Pointer to **int32** |  | [optional] [default to 1]
**Size** | Pointer to **string** |  | [optional] 
**Quality** | Pointer to **string** |  | [optional] 
**ResponseFormat** | Pointer to **string** |  | [optional] 

## Methods

### NewAiImageGenerationRequest

`func NewAiImageGenerationRequest(prompt string, ) *AiImageGenerationRequest`

NewAiImageGenerationRequest instantiates a new AiImageGenerationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiImageGenerationRequestWithDefaults

`func NewAiImageGenerationRequestWithDefaults() *AiImageGenerationRequest`

NewAiImageGenerationRequestWithDefaults instantiates a new AiImageGenerationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiImageGenerationRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiImageGenerationRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiImageGenerationRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiImageGenerationRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetPrompt

`func (o *AiImageGenerationRequest) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiImageGenerationRequest) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiImageGenerationRequest) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.


### GetN

`func (o *AiImageGenerationRequest) GetN() int32`

GetN returns the N field if non-nil, zero value otherwise.

### GetNOk

`func (o *AiImageGenerationRequest) GetNOk() (*int32, bool)`

GetNOk returns a tuple with the N field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetN

`func (o *AiImageGenerationRequest) SetN(v int32)`

SetN sets N field to given value.

### HasN

`func (o *AiImageGenerationRequest) HasN() bool`

HasN returns a boolean if a field has been set.

### GetSize

`func (o *AiImageGenerationRequest) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AiImageGenerationRequest) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AiImageGenerationRequest) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *AiImageGenerationRequest) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetQuality

`func (o *AiImageGenerationRequest) GetQuality() string`

GetQuality returns the Quality field if non-nil, zero value otherwise.

### GetQualityOk

`func (o *AiImageGenerationRequest) GetQualityOk() (*string, bool)`

GetQualityOk returns a tuple with the Quality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuality

`func (o *AiImageGenerationRequest) SetQuality(v string)`

SetQuality sets Quality field to given value.

### HasQuality

`func (o *AiImageGenerationRequest) HasQuality() bool`

HasQuality returns a boolean if a field has been set.

### GetResponseFormat

`func (o *AiImageGenerationRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AiImageGenerationRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AiImageGenerationRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AiImageGenerationRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


