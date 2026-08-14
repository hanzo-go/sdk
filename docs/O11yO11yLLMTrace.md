# O11yO11yLLMTrace

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int32** | CompletionTokens is the trace&#39;s total output tokens. | [optional] 
**Id** | Pointer to **string** | ID is the trace id. | [optional] 
**LatencyMs** | Pointer to **float32** | LatencyMs is the trace&#39;s span, in milliseconds. | [optional] 
**Observations** | Pointer to **int32** | Observations is how many observations the trace holds. | [optional] 
**PromptTokens** | Pointer to **int32** | PromptTokens is the trace&#39;s total input tokens. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the app that emitted it. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the conversation the trace belongs to. | [optional] 
**TotalCost** | Pointer to **float32** | TotalCost is the trace&#39;s total cost. | [optional] 
**TotalTokens** | Pointer to **int32** | TotalTokens is the trace&#39;s total tokens. | [optional] 
**UserId** | Pointer to **string** | UserID is the end user the trace is attributed to. | [optional] 

## Methods

### NewO11yO11yLLMTrace

`func NewO11yO11yLLMTrace() *O11yO11yLLMTrace`

NewO11yO11yLLMTrace instantiates a new O11yO11yLLMTrace object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMTraceWithDefaults

`func NewO11yO11yLLMTraceWithDefaults() *O11yO11yLLMTrace`

NewO11yO11yLLMTraceWithDefaults instantiates a new O11yO11yLLMTrace object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *O11yO11yLLMTrace) GetCompletionTokens() int32`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *O11yO11yLLMTrace) GetCompletionTokensOk() (*int32, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *O11yO11yLLMTrace) SetCompletionTokens(v int32)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *O11yO11yLLMTrace) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMTrace) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMTrace) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMTrace) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMTrace) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatencyMs

`func (o *O11yO11yLLMTrace) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *O11yO11yLLMTrace) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *O11yO11yLLMTrace) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *O11yO11yLLMTrace) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetObservations

`func (o *O11yO11yLLMTrace) GetObservations() int32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *O11yO11yLLMTrace) GetObservationsOk() (*int32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *O11yO11yLLMTrace) SetObservations(v int32)`

SetObservations sets Observations field to given value.

### HasObservations

`func (o *O11yO11yLLMTrace) HasObservations() bool`

HasObservations returns a boolean if a field has been set.

### GetPromptTokens

`func (o *O11yO11yLLMTrace) GetPromptTokens() int32`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *O11yO11yLLMTrace) GetPromptTokensOk() (*int32, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *O11yO11yLLMTrace) SetPromptTokens(v int32)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *O11yO11yLLMTrace) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yLLMTrace) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yLLMTrace) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yLLMTrace) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yLLMTrace) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSessionId

`func (o *O11yO11yLLMTrace) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *O11yO11yLLMTrace) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *O11yO11yLLMTrace) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *O11yO11yLLMTrace) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetTotalCost

`func (o *O11yO11yLLMTrace) GetTotalCost() float32`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *O11yO11yLLMTrace) GetTotalCostOk() (*float32, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *O11yO11yLLMTrace) SetTotalCost(v float32)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *O11yO11yLLMTrace) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTokens

`func (o *O11yO11yLLMTrace) GetTotalTokens() int32`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *O11yO11yLLMTrace) GetTotalTokensOk() (*int32, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *O11yO11yLLMTrace) SetTotalTokens(v int32)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *O11yO11yLLMTrace) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetUserId

`func (o *O11yO11yLLMTrace) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *O11yO11yLLMTrace) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *O11yO11yLLMTrace) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *O11yO11yLLMTrace) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


