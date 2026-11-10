# AiDecideIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle decides over a state this org observed earlier, in place of State and Questions. | [optional] 
**Model** | Pointer to **string** | Model is kai, Kai&#39;s versioned id kai-&lt;12 hex&gt;, or Jev by its vendor id (typesafe/jev-1.13, ~typesafe/jev-latest). | [optional] 
**Observe** | Pointer to **string** | Observe holds the state under this id (1 to 128 of A-Z a-z 0-9 . _ -), so a later decision can name it as its Handle. | [optional] 
**Questions** | Pointer to **map[string]interface{}** | Questions are 1 to 100 named questions, each {\&quot;type\&quot;: \&quot;choice\&quot;|\&quot;noul\&quot;|\&quot;score\&quot;, \&quot;instructions\&quot;?, \&quot;criteria\&quot;?}; a choice names at least 2 labels and a score at least 1 level. Not sent with a Handle. | [optional] 
**State** | Pointer to **interface{}** |  | [optional] 

## Methods

### NewAiDecideIn

`func NewAiDecideIn() *AiDecideIn`

NewAiDecideIn instantiates a new AiDecideIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiDecideInWithDefaults

`func NewAiDecideInWithDefaults() *AiDecideIn`

NewAiDecideInWithDefaults instantiates a new AiDecideIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *AiDecideIn) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AiDecideIn) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AiDecideIn) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AiDecideIn) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetModel

`func (o *AiDecideIn) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *AiDecideIn) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *AiDecideIn) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *AiDecideIn) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetObserve

`func (o *AiDecideIn) GetObserve() string`

GetObserve returns the Observe field if non-nil, zero value otherwise.

### GetObserveOk

`func (o *AiDecideIn) GetObserveOk() (*string, bool)`

GetObserveOk returns a tuple with the Observe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObserve

`func (o *AiDecideIn) SetObserve(v string)`

SetObserve sets Observe field to given value.

### HasObserve

`func (o *AiDecideIn) HasObserve() bool`

HasObserve returns a boolean if a field has been set.

### GetQuestions

`func (o *AiDecideIn) GetQuestions() map[string]interface{}`

GetQuestions returns the Questions field if non-nil, zero value otherwise.

### GetQuestionsOk

`func (o *AiDecideIn) GetQuestionsOk() (*map[string]interface{}, bool)`

GetQuestionsOk returns a tuple with the Questions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuestions

`func (o *AiDecideIn) SetQuestions(v map[string]interface{})`

SetQuestions sets Questions field to given value.

### HasQuestions

`func (o *AiDecideIn) HasQuestions() bool`

HasQuestions returns a boolean if a field has been set.

### GetState

`func (o *AiDecideIn) GetState() interface{}`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *AiDecideIn) GetStateOk() (*interface{}, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *AiDecideIn) SetState(v interface{})`

SetState sets State field to given value.

### HasState

`func (o *AiDecideIn) HasState() bool`

HasState returns a boolean if a field has been set.

### SetStateNil

`func (o *AiDecideIn) SetStateNil(b bool)`

 SetStateNil sets the value for State to be an explicit nil

### UnsetState
`func (o *AiDecideIn) UnsetState()`

UnsetState ensures that no value is present for State, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


