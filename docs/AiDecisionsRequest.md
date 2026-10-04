# AiDecisionsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** |  | [optional] 
**Model** | **string** |  | 
**Observe** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to **interface{}** |  | [optional] 
**Questions** | Pointer to [**map[string]AiDecisionsQuestion**](AiDecisionsQuestion.md) |  | [optional] 
**SessionId** | Pointer to **string** |  | [optional] 
**State** | Pointer to [**AiDecisionSidesFalse**](AiDecisionSidesFalse.md) |  | [optional] 
**Trace** | Pointer to **interface{}** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 

## Methods

### NewAiDecisionsRequest

`func NewAiDecisionsRequest(model string, ) *AiDecisionsRequest`

NewAiDecisionsRequest instantiates a new AiDecisionsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsRequestWithDefaults

`func NewAiDecisionsRequestWithDefaults() *AiDecisionsRequest`

NewAiDecisionsRequestWithDefaults instantiates a new AiDecisionsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *AiDecisionsRequest) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AiDecisionsRequest) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AiDecisionsRequest) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AiDecisionsRequest) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetModel

`func (o *AiDecisionsRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiDecisionsRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiDecisionsRequest) SetModel(v string)`

SetModel sets Model field to given value.


### GetObserve

`func (o *AiDecisionsRequest) GetObserve() string`

GetObserve returns the Observe field if non-nil, zero value otherwise.

### GetObserveOk

`func (o *AiDecisionsRequest) GetObserveOk() (*string, bool)`

GetObserveOk returns a tuple with the Observe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserve

`func (o *AiDecisionsRequest) SetObserve(v string)`

SetObserve sets Observe field to given value.

### HasObserve

`func (o *AiDecisionsRequest) HasObserve() bool`

HasObserve returns a boolean if a field has been set.

### GetProvider

`func (o *AiDecisionsRequest) GetProvider() interface{}`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiDecisionsRequest) GetProviderOk() (*interface{}, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiDecisionsRequest) SetProvider(v interface{})`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiDecisionsRequest) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### SetProviderNil

`func (o *AiDecisionsRequest) SetProviderNil(b bool)`

 SetProviderNil sets the value for Provider to be an explicit nil

### UnsetProvider
`func (o *AiDecisionsRequest) UnsetProvider()`

UnsetProvider ensures that no value is present for Provider, not even an explicit nil
### GetQuestions

`func (o *AiDecisionsRequest) GetQuestions() map[string]AiDecisionsQuestion`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *AiDecisionsRequest) GetQuestionsOk() (*map[string]AiDecisionsQuestion, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *AiDecisionsRequest) SetQuestions(v map[string]AiDecisionsQuestion)`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *AiDecisionsRequest) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetSessionId

`func (o *AiDecisionsRequest) GetSessionId() string`

GetSessionId returns the SessionId field if non-nil, zero value otherwise.

### GetSessionIdOk

`func (o *AiDecisionsRequest) GetSessionIdOk() (*string, bool)`

GetSessionIdOk returns a tuple with the SessionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionId

`func (o *AiDecisionsRequest) SetSessionId(v string)`

SetSessionId sets SessionId field to given value.

### HasSessionId

`func (o *AiDecisionsRequest) HasSessionId() bool`

HasSessionId returns a boolean if a field has been set.

### GetState

`func (o *AiDecisionsRequest) GetState() AiDecisionSidesFalse`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiDecisionsRequest) GetStateOk() (*AiDecisionSidesFalse, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiDecisionsRequest) SetState(v AiDecisionSidesFalse)`

SetState sets State field to given value.

### HasState

`func (o *AiDecisionsRequest) HasState() bool`

HasState returns a boolean if a field has been set.

### GetTrace

`func (o *AiDecisionsRequest) GetTrace() interface{}`

GetTrace returns the Trace field if non-nil, zero value otherwise.

### GetTraceOk

`func (o *AiDecisionsRequest) GetTraceOk() (*interface{}, bool)`

GetTraceOk returns a tuple with the Trace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrace

`func (o *AiDecisionsRequest) SetTrace(v interface{})`

SetTrace sets Trace field to given value.

### HasTrace

`func (o *AiDecisionsRequest) HasTrace() bool`

HasTrace returns a boolean if a field has been set.

### SetTraceNil

`func (o *AiDecisionsRequest) SetTraceNil(b bool)`

 SetTraceNil sets the value for Trace to be an explicit nil

### UnsetTrace
`func (o *AiDecisionsRequest) UnsetTrace()`

UnsetTrace ensures that no value is present for Trace, not even an explicit nil
### GetUser

`func (o *AiDecisionsRequest) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AiDecisionsRequest) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AiDecisionsRequest) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *AiDecisionsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


