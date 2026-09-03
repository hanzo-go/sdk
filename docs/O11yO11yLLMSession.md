# O11yO11yLLMSession

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the conversation&#39;s total output tokens. | [optional] 
**Id** | Pointer to **string** | ID is the session id. | [optional] 
**Observations** | Pointer to **int64** | Observations is how many observations the conversation holds. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the conversation&#39;s total input tokens. | [optional] 
**TotalCost** | Pointer to **float64** | TotalCost is the conversation&#39;s total cost. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the conversation&#39;s total tokens. | [optional] 
**Traces** | Pointer to **int64** | Traces is how many traces the conversation holds. | [optional] 
**UserId** | Pointer to **string** | UserID is the end user the conversation is attributed to. | [optional] 

## Methods

### NewO11yO11yLLMSession

`func NewO11yO11yLLMSession() *O11yO11yLLMSession`

NewO11yO11yLLMSession instantiates a new O11yO11yLLMSession object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMSessionWithDefaults

`func NewO11yO11yLLMSessionWithDefaults() *O11yO11yLLMSession`

NewO11yO11yLLMSessionWithDefaults instantiates a new O11yO11yLLMSession object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *O11yO11yLLMSession) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *O11yO11yLLMSession) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *O11yO11yLLMSession) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *O11yO11yLLMSession) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMSession) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMSession) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMSession) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMSession) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObservations

`func (o *O11yO11yLLMSession) GetObservations() int64`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *O11yO11yLLMSession) GetObservationsOk() (*int64, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *O11yO11yLLMSession) SetObservations(v int64)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *O11yO11yLLMSession) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetPromptTokens

`func (o *O11yO11yLLMSession) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *O11yO11yLLMSession) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *O11yO11yLLMSession) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *O11yO11yLLMSession) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetTotalCost

`func (o *O11yO11yLLMSession) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *O11yO11yLLMSession) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *O11yO11yLLMSession) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *O11yO11yLLMSession) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTokens

`func (o *O11yO11yLLMSession) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *O11yO11yLLMSession) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *O11yO11yLLMSession) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *O11yO11yLLMSession) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetTraces

`func (o *O11yO11yLLMSession) GetTraces() int64`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *O11yO11yLLMSession) GetTracesOk() (*int64, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *O11yO11yLLMSession) SetTraces(v int64)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *O11yO11yLLMSession) HasTraces() bool`

HasTraces returns a boolean if a field has been set.

### GetUserId

`func (o *O11yO11yLLMSession) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *O11yO11yLLMSession) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *O11yO11yLLMSession) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *O11yO11yLLMSession) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


