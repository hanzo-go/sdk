# GraphGraphConclusion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decay** | Pointer to **float64** | Decay is the product of the confidences of the distinct statements the proof rests on: how much certainty survives the derivation. Absent unless every one states a confidence above zero and the proof has at most 1024 atoms. | [optional] 
**From** | Pointer to **string** | From is the latest start among the statements the proof rests on, RFC 3339. What a rule negates has no statement to date it, so a conclusion drawn through a negation may not have held from then. | [optional] 
**Object** | Pointer to **string** | Object is its second argument. Absent for a predicate the rules use with one argument. | [optional] 
**Predicate** | Pointer to **string** | Predicate is the relation concluded. | [optional] 
**Rule** | Pointer to **string** | Rule is the rule that concluded it, &#x60;rule:&lt;name&gt;&#x60;. | [optional] 
**RuleId** | Pointer to **string** | RuleID is the ID of the assertion stating that rule&#39;s text as it was in force. | [optional] 
**Subject** | Pointer to **string** | Subject is the tuple&#39;s first argument. | [optional] 
**Supports** | Pointer to [**[]GraphGraphSupport**](GraphGraphSupport.md) | Supports are the rule&#39;s positive atoms as they matched, in the order the rule writes them. What the rule negates is absent at the point and has no support to cite. | [optional] 
**Until** | Pointer to **string** | Until is the earliest until any of those statements states, RFC 3339. Absent is open. A statement a later assertion ends is a later question, asked at a later as_of. | [optional] 
**Weakest** | Pointer to **string** | Weakest is the ID of the statement with the lowest confidence among them, the first on a tie: the link the conclusion is only as strong as. Absent exactly when Decay is. | [optional] 

## Methods

### NewGraphGraphConclusion

`func NewGraphGraphConclusion() *GraphGraphConclusion`

NewGraphGraphConclusion instantiates a new GraphGraphConclusion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphConclusionWithDefaults

`func NewGraphGraphConclusionWithDefaults() *GraphGraphConclusion`

NewGraphGraphConclusionWithDefaults instantiates a new GraphGraphConclusion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecay

`func (o *GraphGraphConclusion) GetDecay() float64`

GetDecay returns the Decay field if non-nil, zero value otherwise.

### GetDecayOk

`func (o *GraphGraphConclusion) GetDecayOk() (*float64, bool)`

GetDecayOk returns a tuple with the Decay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecay

`func (o *GraphGraphConclusion) SetDecay(v float64)`

SetDecay sets Decay field to given value.

### HasDecay

`func (o *GraphGraphConclusion) HasDecay() bool`

HasDecay returns a boolean if a field has been set.

### GetFrom

`func (o *GraphGraphConclusion) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphGraphConclusion) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphGraphConclusion) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GraphGraphConclusion) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetObject

`func (o *GraphGraphConclusion) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *GraphGraphConclusion) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *GraphGraphConclusion) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *GraphGraphConclusion) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetPredicate

`func (o *GraphGraphConclusion) GetPredicate() string`

GetPredicate returns the Predicate field if non-nil, zero value otherwise.

### GetPredicateOk

`func (o *GraphGraphConclusion) GetPredicateOk() (*string, bool)`

GetPredicateOk returns a tuple with the Predicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPredicate

`func (o *GraphGraphConclusion) SetPredicate(v string)`

SetPredicate sets Predicate field to given value.

### HasPredicate

`func (o *GraphGraphConclusion) HasPredicate() bool`

HasPredicate returns a boolean if a field has been set.

### GetRule

`func (o *GraphGraphConclusion) GetRule() string`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *GraphGraphConclusion) GetRuleOk() (*string, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *GraphGraphConclusion) SetRule(v string)`

SetRule sets Rule field to given value.

### HasRule

`func (o *GraphGraphConclusion) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetRuleId

`func (o *GraphGraphConclusion) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *GraphGraphConclusion) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *GraphGraphConclusion) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *GraphGraphConclusion) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### GetSubject

`func (o *GraphGraphConclusion) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GraphGraphConclusion) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GraphGraphConclusion) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GraphGraphConclusion) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### GetSupports

`func (o *GraphGraphConclusion) GetSupports() []GraphGraphSupport`

GetSupports returns the Supports field if non-nil, zero value otherwise.

### GetSupportsOk

`func (o *GraphGraphConclusion) GetSupportsOk() (*[]GraphGraphSupport, bool)`

GetSupportsOk returns a tuple with the Supports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupports

`func (o *GraphGraphConclusion) SetSupports(v []GraphGraphSupport)`

SetSupports sets Supports field to given value.

### HasSupports

`func (o *GraphGraphConclusion) HasSupports() bool`

HasSupports returns a boolean if a field has been set.

### GetUntil

`func (o *GraphGraphConclusion) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *GraphGraphConclusion) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *GraphGraphConclusion) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *GraphGraphConclusion) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetWeakest

`func (o *GraphGraphConclusion) GetWeakest() string`

GetWeakest returns the Weakest field if non-nil, zero value otherwise.

### GetWeakestOk

`func (o *GraphGraphConclusion) GetWeakestOk() (*string, bool)`

GetWeakestOk returns a tuple with the Weakest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeakest

`func (o *GraphGraphConclusion) SetWeakest(v string)`

SetWeakest sets Weakest field to given value.

### HasWeakest

`func (o *GraphGraphConclusion) HasWeakest() bool`

HasWeakest returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


