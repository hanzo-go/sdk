# GraphGraphOntologyPredicate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID is the raw predicate code stored in graph.db, such as PS1, PS500 or nm_isCurrentUser. | [optional] 
**Label** | Pointer to **string** | Label is the plain-language reading of what the predicate records. | [optional] 
**Mechanism** | Pointer to **string** | Mechanism records whether the daemon mechanism is implemented, unverified or unlocated. | [optional] 
**Method** | Pointer to **string** | Method names how the daemon obtains the value: Copied, Joined, Rule, Apple-shipped, Embedding model, XGBoost, Face clustering or Local LLM. | [optional] 
**Names** | Pointer to **bool** | Names is true when the predicate links its subject to another entity rather than a scalar literal. | [optional] 
**Relation** | Pointer to **string** | Relation is the canonical snake_case relation key this plane files the predicate under when ingesting an Apple graph. | [optional] 
**Source** | Pointer to **string** | Source names the OS store or subsystem the daemon reads the predicate from. | [optional] 

## Methods

### NewGraphGraphOntologyPredicate

`func NewGraphGraphOntologyPredicate() *GraphGraphOntologyPredicate`

NewGraphGraphOntologyPredicate instantiates a new GraphGraphOntologyPredicate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphOntologyPredicateWithDefaults

`func NewGraphGraphOntologyPredicateWithDefaults() *GraphGraphOntologyPredicate`

NewGraphGraphOntologyPredicateWithDefaults instantiates a new GraphGraphOntologyPredicate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GraphGraphOntologyPredicate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GraphGraphOntologyPredicate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GraphGraphOntologyPredicate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GraphGraphOntologyPredicate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *GraphGraphOntologyPredicate) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GraphGraphOntologyPredicate) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GraphGraphOntologyPredicate) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GraphGraphOntologyPredicate) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMechanism

`func (o *GraphGraphOntologyPredicate) GetMechanism() string`

GetMechanism returns the Mechanism field if non-nil, zero value otherwise.

### GetMechanismOk

`func (o *GraphGraphOntologyPredicate) GetMechanismOk() (*string, bool)`

GetMechanismOk returns a tuple with the Mechanism field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMechanism

`func (o *GraphGraphOntologyPredicate) SetMechanism(v string)`

SetMechanism sets Mechanism field to given value.

### HasMechanism

`func (o *GraphGraphOntologyPredicate) HasMechanism() bool`

HasMechanism returns a boolean if a field has been set.

### GetMethod

`func (o *GraphGraphOntologyPredicate) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *GraphGraphOntologyPredicate) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *GraphGraphOntologyPredicate) SetMethod(v string)`

SetMethod sets Method field to given value.

### HasMethod

`func (o *GraphGraphOntologyPredicate) HasMethod() bool`

HasMethod returns a boolean if a field has been set.

### GetNames

`func (o *GraphGraphOntologyPredicate) GetNames() bool`

GetNames returns the Names field if non-nil, zero value otherwise.

### GetNamesOk

`func (o *GraphGraphOntologyPredicate) GetNamesOk() (*bool, bool)`

GetNamesOk returns a tuple with the Names field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNames

`func (o *GraphGraphOntologyPredicate) SetNames(v bool)`

SetNames sets Names field to given value.

### HasNames

`func (o *GraphGraphOntologyPredicate) HasNames() bool`

HasNames returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphOntologyPredicate) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphOntologyPredicate) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphOntologyPredicate) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphOntologyPredicate) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSource

`func (o *GraphGraphOntologyPredicate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GraphGraphOntologyPredicate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GraphGraphOntologyPredicate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *GraphGraphOntologyPredicate) HasSource() bool`

HasSource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


