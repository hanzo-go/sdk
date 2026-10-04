# AiDecisionsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Answers** | [**map[string]AiDecisionsAnswer**](AiDecisionsAnswer.md) |  | 
**Id** | **string** |  | 
**LatencyMs** | **float32** |  | 
**Model** | **string** |  | 
**Provider** | **string** |  | 
**Routing** | [**AiDecisionsRouting**](AiDecisionsRouting.md) |  | 
**StateHash** | **string** |  | 
**Usage** | [**AiDecisionsUsage**](AiDecisionsUsage.md) |  | 

## Methods

### NewAiDecisionsResponse

`func NewAiDecisionsResponse(answers map[string]AiDecisionsAnswer, id string, latencyMs float32, model string, provider string, routing AiDecisionsRouting, stateHash string, usage AiDecisionsUsage, ) *AiDecisionsResponse`

NewAiDecisionsResponse instantiates a new AiDecisionsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecisionsResponseWithDefaults

`func NewAiDecisionsResponseWithDefaults() *AiDecisionsResponse`

NewAiDecisionsResponseWithDefaults instantiates a new AiDecisionsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswers

`func (o *AiDecisionsResponse) GetAnswers() map[string]AiDecisionsAnswer`

GetAnswers returns the Answers field if non-nil, zero value otherwise.

### GetAnswersOk

`func (o *AiDecisionsResponse) GetAnswersOk() (*map[string]AiDecisionsAnswer, bool)`

GetAnswersOk returns a tuple with the Answers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswers

`func (o *AiDecisionsResponse) SetAnswers(v map[string]AiDecisionsAnswer)`

SetAnswers sets Answers field to given value.


### GetId

`func (o *AiDecisionsResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiDecisionsResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiDecisionsResponse) SetId(v string)`

SetId sets Id field to given value.


### GetLatencyMs

`func (o *AiDecisionsResponse) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *AiDecisionsResponse) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *AiDecisionsResponse) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.


### GetModel

`func (o *AiDecisionsResponse) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiDecisionsResponse) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiDecisionsResponse) SetModel(v string)`

SetModel sets Model field to given value.


### GetProvider

`func (o *AiDecisionsResponse) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiDecisionsResponse) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiDecisionsResponse) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetRouting

`func (o *AiDecisionsResponse) GetRouting() AiDecisionsRouting`

GetRouting returns the Routing field if non-nil, zero value otherwise.

### GetRoutingOk

`func (o *AiDecisionsResponse) GetRoutingOk() (*AiDecisionsRouting, bool)`

GetRoutingOk returns a tuple with the Routing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRouting

`func (o *AiDecisionsResponse) SetRouting(v AiDecisionsRouting)`

SetRouting sets Routing field to given value.


### GetStateHash

`func (o *AiDecisionsResponse) GetStateHash() string`

GetStateHash returns the StateHash field if non-nil, zero value otherwise.

### GetStateHashOk

`func (o *AiDecisionsResponse) GetStateHashOk() (*string, bool)`

GetStateHashOk returns a tuple with the StateHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateHash

`func (o *AiDecisionsResponse) SetStateHash(v string)`

SetStateHash sets StateHash field to given value.


### GetUsage

`func (o *AiDecisionsResponse) GetUsage() AiDecisionsUsage`

GetUsage returns the Usage field if non-nil, zero value otherwise.

### GetUsageOk

`func (o *AiDecisionsResponse) GetUsageOk() (*AiDecisionsUsage, bool)`

GetUsageOk returns a tuple with the Usage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsage

`func (o *AiDecisionsResponse) SetUsage(v AiDecisionsUsage)`

SetUsage sets Usage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


