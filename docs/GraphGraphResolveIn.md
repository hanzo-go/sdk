# GraphGraphResolveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339: the answer as it would have been given then, which is what makes a past answer reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the instant of the world to answer about, RFC 3339: what held then. Absent is now. | [optional] 
**Entity** | **string** | Entity is the thing to answer about. Required. | 
**Relation** | **string** | Relation is the one relation to settle. Required: this answers a single (entity, relation) pair, never a whole entity at once. | 

## Methods

### NewGraphGraphResolveIn

`func NewGraphGraphResolveIn(entity string, relation string, ) *GraphGraphResolveIn`

NewGraphGraphResolveIn instantiates a new GraphGraphResolveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphResolveInWithDefaults

`func NewGraphGraphResolveInWithDefaults() *GraphGraphResolveIn`

NewGraphGraphResolveInWithDefaults instantiates a new GraphGraphResolveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphResolveIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphResolveIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphResolveIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphResolveIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphResolveIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphResolveIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphResolveIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphResolveIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetEntity

`func (o *GraphGraphResolveIn) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphResolveIn) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphResolveIn) SetEntity(v string)`

SetEntity sets Entity field to given value.


### GetRelation

`func (o *GraphGraphResolveIn) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphResolveIn) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphResolveIn) SetRelation(v string)`

SetRelation sets Relation field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


