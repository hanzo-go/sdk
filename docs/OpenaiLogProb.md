# OpenaiLogProb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bytes** | Pointer to **string** |  | [optional] 
**Logprob** | Pointer to **float32** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TopLogprobs** | Pointer to [**[]OpenaiTopLogProbs**](OpenaiTopLogProbs.md) |  | [optional] 

## Methods

### NewOpenaiLogProb

`func NewOpenaiLogProb() *OpenaiLogProb`

NewOpenaiLogProb instantiates a new OpenaiLogProb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOpenaiLogProbWithDefaults

`func NewOpenaiLogProbWithDefaults() *OpenaiLogProb`

NewOpenaiLogProbWithDefaults instantiates a new OpenaiLogProb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBytes

`func (o *OpenaiLogProb) GetBytes() string`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *OpenaiLogProb) GetBytesOk() (*string, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *OpenaiLogProb) SetBytes(v string)`

SetBytes sets Bytes field to given value.

### HasBytes

`func (o *OpenaiLogProb) HasBytes() bool`

HasBytes returns a boolean if a field has been set.

### GetLogprob

`func (o *OpenaiLogProb) GetLogprob() float32`

GetLogprob returns the Logprob field if non-nil, zero value otherwise.

### GetLogprobOk

`func (o *OpenaiLogProb) GetLogprobOk() (*float32, bool)`

GetLogprobOk returns a tuple with the Logprob field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogprob

`func (o *OpenaiLogProb) SetLogprob(v float32)`

SetLogprob sets Logprob field to given value.

### HasLogprob

`func (o *OpenaiLogProb) HasLogprob() bool`

HasLogprob returns a boolean if a field has been set.

### GetToken

`func (o *OpenaiLogProb) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *OpenaiLogProb) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *OpenaiLogProb) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *OpenaiLogProb) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTopLogprobs

`func (o *OpenaiLogProb) GetTopLogprobs() []OpenaiTopLogProbs`

GetTopLogprobs returns the TopLogprobs field if non-nil, zero value otherwise.

### GetTopLogprobsOk

`func (o *OpenaiLogProb) GetTopLogprobsOk() (*[]OpenaiTopLogProbs, bool)`

GetTopLogprobsOk returns a tuple with the TopLogprobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopLogprobs

`func (o *OpenaiLogProb) SetTopLogprobs(v []OpenaiTopLogProbs)`

SetTopLogprobs sets TopLogprobs field to given value.

### HasTopLogprobs

`func (o *OpenaiLogProb) HasTopLogprobs() bool`

HasTopLogprobs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


