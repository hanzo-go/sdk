# AiSpeechRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Model** | **string** |  | 
**Input** | **string** |  | 
**Voice** | **string** |  | 
**ResponseFormat** | Pointer to **string** |  | [optional] 
**Speed** | Pointer to **float32** |  | [optional] 

## Methods

### NewAiSpeechRequest

`func NewAiSpeechRequest(model string, input string, voice string, ) *AiSpeechRequest`

NewAiSpeechRequest instantiates a new AiSpeechRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiSpeechRequestWithDefaults

`func NewAiSpeechRequestWithDefaults() *AiSpeechRequest`

NewAiSpeechRequestWithDefaults instantiates a new AiSpeechRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetModel

`func (o *AiSpeechRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiSpeechRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiSpeechRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetInput

`func (o *AiSpeechRequest) GetInput() string`

GetInput returns the Input field if non-nil, zero value otherwise.

### GetInputOk

`func (o *AiSpeechRequest) GetInputOk() (*string, bool)`

GetInputOk returns a tuple with the Input field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInput

`func (o *AiSpeechRequest) SetInput(v string)`

SetInput sets Input field to given value.


### GetVoice

`func (o *AiSpeechRequest) GetVoice() string`

GetVoice returns the Voice field if non-nil, zero value otherwise.

### GetVoiceOk

`func (o *AiSpeechRequest) GetVoiceOk() (*string, bool)`

GetVoiceOk returns a tuple with the Voice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoice

`func (o *AiSpeechRequest) SetVoice(v string)`

SetVoice sets Voice field to given value.


### GetResponseFormat

`func (o *AiSpeechRequest) GetResponseFormat() string`

GetResponseFormat returns the ResponseFormat field if non-nil, zero value otherwise.

### GetResponseFormatOk

`func (o *AiSpeechRequest) GetResponseFormatOk() (*string, bool)`

GetResponseFormatOk returns a tuple with the ResponseFormat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseFormat

`func (o *AiSpeechRequest) SetResponseFormat(v string)`

SetResponseFormat sets ResponseFormat field to given value.

### HasResponseFormat

`func (o *AiSpeechRequest) HasResponseFormat() bool`

HasResponseFormat returns a boolean if a field has been set.

### GetSpeed

`func (o *AiSpeechRequest) GetSpeed() float32`

GetSpeed returns the Speed field if non-nil, zero value otherwise.

### GetSpeedOk

`func (o *AiSpeechRequest) GetSpeedOk() (*float32, bool)`

GetSpeedOk returns a tuple with the Speed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeed

`func (o *AiSpeechRequest) SetSpeed(v float32)`

SetSpeed sets Speed field to given value.

### HasSpeed

`func (o *AiSpeechRequest) HasSpeed() bool`

HasSpeed returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


