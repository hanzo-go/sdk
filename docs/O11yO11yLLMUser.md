# O11yO11yLLMUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | CompletionTokens is their total output tokens. | [optional] 
**Id** | Pointer to **string** | ID is the end user&#39;s id (user.id). | [optional] 
**Observations** | Pointer to **int64** | Observations is how many observations they produced. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is their total input tokens. | [optional] 
**Sessions** | Pointer to **int64** | Sessions is how many conversations they had. | [optional] 
**TotalCost** | Pointer to **float64** | TotalCost is their total cost. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is their total tokens. | [optional] 
**Traces** | Pointer to **int64** | Traces is how many traces they produced. | [optional] 

## Methods

### NewO11yO11yLLMUser

`func NewO11yO11yLLMUser() *O11yO11yLLMUser`

NewO11yO11yLLMUser instantiates a new O11yO11yLLMUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMUserWithDefaults

`func NewO11yO11yLLMUserWithDefaults() *O11yO11yLLMUser`

NewO11yO11yLLMUserWithDefaults instantiates a new O11yO11yLLMUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *O11yO11yLLMUser) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *O11yO11yLLMUser) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *O11yO11yLLMUser) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *O11yO11yLLMUser) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMUser) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMUser) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMUser) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetObservations

`func (o *O11yO11yLLMUser) GetObservations() int64`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *O11yO11yLLMUser) GetObservationsOk() (*int64, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *O11yO11yLLMUser) SetObservations(v int64)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *O11yO11yLLMUser) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetPromptTokens

`func (o *O11yO11yLLMUser) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *O11yO11yLLMUser) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *O11yO11yLLMUser) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *O11yO11yLLMUser) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetSessions

`func (o *O11yO11yLLMUser) GetSessions() int64`

GetSessions returns the Sessions field if non-nil, zero value otherwise.

### GetSessionsOk

`func (o *O11yO11yLLMUser) GetSessionsOk() (*int64, bool)`

GetSessionsOk returns a tuple with the Sessions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessions

`func (o *O11yO11yLLMUser) SetSessions(v int64)`

SetSessions sets Sessions field to given value.

### HasSessions

`func (o *O11yO11yLLMUser) HasSessions() bool`

HasSessions returns a boolean if a field has been set.

### GetTotalCost

`func (o *O11yO11yLLMUser) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *O11yO11yLLMUser) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *O11yO11yLLMUser) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *O11yO11yLLMUser) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTokens

`func (o *O11yO11yLLMUser) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *O11yO11yLLMUser) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *O11yO11yLLMUser) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *O11yO11yLLMUser) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetTraces

`func (o *O11yO11yLLMUser) GetTraces() int64`

GetTraces returns the Traces field if non-nil, zero value otherwise.

### GetTracesOk

`func (o *O11yO11yLLMUser) GetTracesOk() (*int64, bool)`

GetTracesOk returns a tuple with the Traces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraces

`func (o *O11yO11yLLMUser) SetTraces(v int64)`

SetTraces sets Traces field to given value.

### HasTraces

`func (o *O11yO11yLLMUser) HasTraces() bool`

HasTraces returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


