# TrainSampleIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID names the client. | [optional] 
**NumSamples** | Pointer to **int64** | NumSamples is how many sequences to decode. | [optional] 
**Prompt** | Pointer to **string** | Prompt is a text prompt; exactly one of prompt and tokens. | [optional] 
**SamplingParams** | Pointer to **interface{}** |  | [optional] 
**Tokens** | Pointer to **[]int32** | Tokens is a tokenized prompt. | [optional] 

## Methods

### NewTrainSampleIn

`func NewTrainSampleIn() *TrainSampleIn`

NewTrainSampleIn instantiates a new TrainSampleIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainSampleInWithDefaults

`func NewTrainSampleInWithDefaults() *TrainSampleIn`

NewTrainSampleInWithDefaults instantiates a new TrainSampleIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrainSampleIn) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrainSampleIn) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrainSampleIn) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TrainSampleIn) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNumSamples

`func (o *TrainSampleIn) GetNumSamples() int64`

GetNumSamples returns the NumSamples field if non-nil, zero value otherwise.

### GetNumSamplesOk

`func (o *TrainSampleIn) GetNumSamplesOk() (*int64, bool)`

GetNumSamplesOk returns a tuple with the NumSamples field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumSamples

`func (o *TrainSampleIn) SetNumSamples(v int64)`

SetNumSamples sets NumSamples field to given value.

### HasNumSamples

`func (o *TrainSampleIn) HasNumSamples() bool`

HasNumSamples returns a boolean if a field has been set.

### GetPrompt

`func (o *TrainSampleIn) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *TrainSampleIn) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *TrainSampleIn) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *TrainSampleIn) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetSamplingParams

`func (o *TrainSampleIn) GetSamplingParams() interface{}`

GetSamplingParams returns the SamplingParams field if non-nil, zero value otherwise.

### GetSamplingParamsOk

`func (o *TrainSampleIn) GetSamplingParamsOk() (*interface{}, bool)`

GetSamplingParamsOk returns a tuple with the SamplingParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSamplingParams

`func (o *TrainSampleIn) SetSamplingParams(v interface{})`

SetSamplingParams sets SamplingParams field to given value.

### HasSamplingParams

`func (o *TrainSampleIn) HasSamplingParams() bool`

HasSamplingParams returns a boolean if a field has been set.

### SetSamplingParamsNil

`func (o *TrainSampleIn) SetSamplingParamsNil(b bool)`

 SetSamplingParamsNil sets the value for SamplingParams to be an explicit nil

### UnsetSamplingParams
`func (o *TrainSampleIn) UnsetSamplingParams()`

UnsetSamplingParams ensures that no value is present for SamplingParams, not even an explicit nil
### GetTokens

`func (o *TrainSampleIn) GetTokens() []int32`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *TrainSampleIn) GetTokensOk() (*[]int32, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *TrainSampleIn) SetTokens(v []int32)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *TrainSampleIn) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


