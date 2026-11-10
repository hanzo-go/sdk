# AiPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completion** | Pointer to **string** | Completion is the price of one output token, a decimal string. | [optional] 
**InputPerMillion** | Pointer to **float64** | InputPerMillion is the price of a million input tokens. | [optional] 
**OutputPerMillion** | Pointer to **float64** | OutputPerMillion is the price of a million output tokens. | [optional] 
**Overrides** | Pointer to **[]map[string]interface{}** | Overrides are the vendor&#39;s conditional rates as billed, each with its condition: min_prompt_tokens, or utc_start, utc_end and utc_days. | [optional] 
**Prompt** | Pointer to **string** | Prompt is the price of one input token, a decimal string. | [optional] 
**Rates** | Pointer to **map[string]string** | Rates is every rate a resold model&#39;s vendor states, in the vendor&#39;s names and units (USD per token for prompt, completion, input_cache_read, input_cache_write, internal_reasoning; per request; per search for web_search), as billed: the vendor&#39;s rate times our margin, exact. | [optional] 
**Variable** | Pointer to **bool** | Variable is true for a router that bills each answer at the price of the model that gave it; the figures above are then that router&#39;s ceiling. | [optional] 

## Methods

### NewAiPrice

`func NewAiPrice() *AiPrice`

NewAiPrice instantiates a new AiPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiPriceWithDefaults

`func NewAiPriceWithDefaults() *AiPrice`

NewAiPriceWithDefaults instantiates a new AiPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletion

`func (o *AiPrice) GetCompletion() string`

GetCompletion returns the Completion field if non-nil, zero value otherwise.

### GetCompletionOk

`func (o *AiPrice) GetCompletionOk() (*string, bool)`

GetCompletionOk returns a tuple with the Completion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletion

`func (o *AiPrice) SetCompletion(v string)`

SetCompletion sets Completion field to given value.

### HasCompletion

`func (o *AiPrice) HasCompletion() bool`

HasCompletion returns a boolean if a field has been set.

### GetInputPerMillion

`func (o *AiPrice) GetInputPerMillion() float64`

GetInputPerMillion returns the InputPerMillion field if non-nil, zero value otherwise.

### GetInputPerMillionOk

`func (o *AiPrice) GetInputPerMillionOk() (*float64, bool)`

GetInputPerMillionOk returns a tuple with the InputPerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPerMillion

`func (o *AiPrice) SetInputPerMillion(v float64)`

SetInputPerMillion sets InputPerMillion field to given value.

### HasInputPerMillion

`func (o *AiPrice) HasInputPerMillion() bool`

HasInputPerMillion returns a boolean if a field has been set.

### GetOutputPerMillion

`func (o *AiPrice) GetOutputPerMillion() float64`

GetOutputPerMillion returns the OutputPerMillion field if non-nil, zero value otherwise.

### GetOutputPerMillionOk

`func (o *AiPrice) GetOutputPerMillionOk() (*float64, bool)`

GetOutputPerMillionOk returns a tuple with the OutputPerMillion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPerMillion

`func (o *AiPrice) SetOutputPerMillion(v float64)`

SetOutputPerMillion sets OutputPerMillion field to given value.

### HasOutputPerMillion

`func (o *AiPrice) HasOutputPerMillion() bool`

HasOutputPerMillion returns a boolean if a field has been set.

### GetOverrides

`func (o *AiPrice) GetOverrides() []map[string]interface{}`

GetOverrides returns the Overrides field if non-nil, zero value otherwise.

### GetOverridesOk

`func (o *AiPrice) GetOverridesOk() (*[]map[string]interface{}, bool)`

GetOverridesOk returns a tuple with the Overrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOverrides

`func (o *AiPrice) SetOverrides(v []map[string]interface{})`

SetOverrides sets Overrides field to given value.

### HasOverrides

`func (o *AiPrice) HasOverrides() bool`

HasOverrides returns a boolean if a field has been set.

### GetPrompt

`func (o *AiPrice) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *AiPrice) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *AiPrice) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *AiPrice) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetRates

`func (o *AiPrice) GetRates() map[string]string`

GetRates returns the Rates field if non-nil, zero value otherwise.

### GetRatesOk

`func (o *AiPrice) GetRatesOk() (*map[string]string, bool)`

GetRatesOk returns a tuple with the Rates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRates

`func (o *AiPrice) SetRates(v map[string]string)`

SetRates sets Rates field to given value.

### HasRates

`func (o *AiPrice) HasRates() bool`

HasRates returns a boolean if a field has been set.

### GetVariable

`func (o *AiPrice) GetVariable() bool`

GetVariable returns the Variable field if non-nil, zero value otherwise.

### GetVariableOk

`func (o *AiPrice) GetVariableOk() (*bool, bool)`

GetVariableOk returns a tuple with the Variable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariable

`func (o *AiPrice) SetVariable(v bool)`

SetVariable sets Variable field to given value.

### HasVariable

`func (o *AiPrice) HasVariable() bool`

HasVariable returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


