# O11yO11yLLMObservation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompletionTokens** | Pointer to **int64** | CompletionTokens is the output token count. | [optional] 
**Id** | Pointer to **string** | ID is the observation&#39;s id (the span id). | [optional] 
**LatencyMs** | Pointer to **float64** | LatencyMs is how long it took, in milliseconds. | [optional] 
**Model** | Pointer to **string** | Model is the model that served it. | [optional] 
**Name** | Pointer to **string** | Name is the observation&#39;s name. | [optional] 
**ParentObservationId** | Pointer to **string** | ParentID is the parent observation, when the span has one. | [optional] 
**PromptTokens** | Pointer to **int64** | PromptTokens is the input token count. | [optional] 
**Provider** | Pointer to **string** | Provider is the model&#39;s provider. | [optional] 
**ServiceName** | Pointer to **string** | ServiceName is the app that emitted it. | [optional] 
**SessionId** | Pointer to **string** | SessionID is the conversation the observation belongs to. | [optional] 
**StartTime** | Pointer to **time.Time** | StartTime is when the observation started. | [optional] 
**StatusCode** | Pointer to **string** | StatusCode is the observation&#39;s status, e.g. OK, ERROR. | [optional] 
**TotalCost** | Pointer to **float64** | TotalCost is the observation&#39;s cost. | [optional] 
**TotalTokens** | Pointer to **int64** | TotalTokens is the sum of prompt and completion tokens. | [optional] 
**TraceId** | Pointer to **string** | TraceID is the trace the observation belongs to. | [optional] 
**Type** | Pointer to **string** | Type is the observation kind, e.g. chat, embeddings, tool. | [optional] 
**UserId** | Pointer to **string** | UserID is the end user the observation is attributed to. | [optional] 

## Methods

### NewO11yO11yLLMObservation

`func NewO11yO11yLLMObservation() *O11yO11yLLMObservation`

NewO11yO11yLLMObservation instantiates a new O11yO11yLLMObservation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yLLMObservationWithDefaults

`func NewO11yO11yLLMObservationWithDefaults() *O11yO11yLLMObservation`

NewO11yO11yLLMObservationWithDefaults instantiates a new O11yO11yLLMObservation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletionTokens

`func (o *O11yO11yLLMObservation) GetCompletionTokens() int64`

GetCompletionTokens returns the CompletionTokens field if non-nil, zero value otherwise.

### GetCompletionTokensOk

`func (o *O11yO11yLLMObservation) GetCompletionTokensOk() (*int64, bool)`

GetCompletionTokensOk returns a tuple with the CompletionTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletionTokens

`func (o *O11yO11yLLMObservation) SetCompletionTokens(v int64)`

SetCompletionTokens sets CompletionTokens field to given value.

### HasCompletionTokens

`func (o *O11yO11yLLMObservation) HasCompletionTokens() bool`

HasCompletionTokens returns a boolean if a field has been set.

### GetId

`func (o *O11yO11yLLMObservation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *O11yO11yLLMObservation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *O11yO11yLLMObservation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *O11yO11yLLMObservation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLatencyMs

`func (o *O11yO11yLLMObservation) GetLatencyMs() float64`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *O11yO11yLLMObservation) GetLatencyMsOk() (*float64, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *O11yO11yLLMObservation) SetLatencyMs(v float64)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *O11yO11yLLMObservation) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### GetModel

`func (o *O11yO11yLLMObservation) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *O11yO11yLLMObservation) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *O11yO11yLLMObservation) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *O11yO11yLLMObservation) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yLLMObservation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yLLMObservation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yLLMObservation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yLLMObservation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetParentObservationId

`func (o *O11yO11yLLMObservation) GetParentObservationId() string`

GetParentObservationId returns the ParentObservationId field if non-nil, zero value otherwise.

### GetParentObservationIdOk

`func (o *O11yO11yLLMObservation) GetParentObservationIdOk() (*string, bool)`

GetParentObservationIdOk returns a tuple with the ParentObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentObservationId

`func (o *O11yO11yLLMObservation) SetParentObservationId(v string)`

SetParentObservationId sets ParentObservationId field to given value.

### HasParentObservationId

`func (o *O11yO11yLLMObservation) HasParentObservationId() bool`

HasParentObservationId returns a boolean if a field has been set.

### GetPromptTokens

`func (o *O11yO11yLLMObservation) GetPromptTokens() int64`

GetPromptTokens returns the PromptTokens field if non-nil, zero value otherwise.

### GetPromptTokensOk

`func (o *O11yO11yLLMObservation) GetPromptTokensOk() (*int64, bool)`

GetPromptTokensOk returns a tuple with the PromptTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPromptTokens

`func (o *O11yO11yLLMObservation) SetPromptTokens(v int64)`

SetPromptTokens sets PromptTokens field to given value.

### HasPromptTokens

`func (o *O11yO11yLLMObservation) HasPromptTokens() bool`

HasPromptTokens returns a boolean if a field has been set.

### GetProvider

`func (o *O11yO11yLLMObservation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *O11yO11yLLMObservation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *O11yO11yLLMObservation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *O11yO11yLLMObservation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetServiceName

`func (o *O11yO11yLLMObservation) GetServiceName() string`

GetServiceName returns the ServiceName field if non-nil, zero value otherwise.

### GetServiceNameOk

`func (o *O11yO11yLLMObservation) GetServiceNameOk() (*string, bool)`

GetServiceNameOk returns a tuple with the ServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceName

`func (o *O11yO11yLLMObservation) SetServiceName(v string)`

SetServiceName sets ServiceName field to given value.

### HasServiceName

`func (o *O11yO11yLLMObservation) HasServiceName() bool`

HasServiceName returns a boolean if a field has been set.

### GetSessionId

`func (o *O11yO11yLLMObservation) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *O11yO11yLLMObservation) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *O11yO11yLLMObservation) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *O11yO11yLLMObservation) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetStartTime

`func (o *O11yO11yLLMObservation) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *O11yO11yLLMObservation) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *O11yO11yLLMObservation) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *O11yO11yLLMObservation) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetStatusCode

`func (o *O11yO11yLLMObservation) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *O11yO11yLLMObservation) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *O11yO11yLLMObservation) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *O11yO11yLLMObservation) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTotalCost

`func (o *O11yO11yLLMObservation) GetTotalCost() float64`

GetTotalCost returns the TotalCost field if non-nil, zero value otherwise.

### GetTotalCostOk

`func (o *O11yO11yLLMObservation) GetTotalCostOk() (*float64, bool)`

GetTotalCostOk returns a tuple with the TotalCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCost

`func (o *O11yO11yLLMObservation) SetTotalCost(v float64)`

SetTotalCost sets TotalCost field to given value.

### HasTotalCost

`func (o *O11yO11yLLMObservation) HasTotalCost() bool`

HasTotalCost returns a boolean if a field has been set.

### GetTotalTokens

`func (o *O11yO11yLLMObservation) GetTotalTokens() int64`

GetTotalTokens returns the TotalTokens field if non-nil, zero value otherwise.

### GetTotalTokensOk

`func (o *O11yO11yLLMObservation) GetTotalTokensOk() (*int64, bool)`

GetTotalTokensOk returns a tuple with the TotalTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTokens

`func (o *O11yO11yLLMObservation) SetTotalTokens(v int64)`

SetTotalTokens sets TotalTokens field to given value.

### HasTotalTokens

`func (o *O11yO11yLLMObservation) HasTotalTokens() bool`

HasTotalTokens returns a boolean if a field has been set.

### GetTraceId

`func (o *O11yO11yLLMObservation) GetTraceId() string`

GetTraceId returns the TraceId field if non-nil, zero value otherwise.

### GetTraceIdOk

`func (o *O11yO11yLLMObservation) GetTraceIdOk() (*string, bool)`

GetTraceIdOk returns a tuple with the TraceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraceId

`func (o *O11yO11yLLMObservation) SetTraceId(v string)`

SetTraceId sets TraceId field to given value.

### HasTraceId

`func (o *O11yO11yLLMObservation) HasTraceId() bool`

HasTraceId returns a boolean if a field has been set.

### GetType

`func (o *O11yO11yLLMObservation) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *O11yO11yLLMObservation) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *O11yO11yLLMObservation) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *O11yO11yLLMObservation) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserId

`func (o *O11yO11yLLMObservation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *O11yO11yLLMObservation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *O11yO11yLLMObservation) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *O11yO11yLLMObservation) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


