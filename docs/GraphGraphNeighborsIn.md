# GraphGraphNeighborsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339: the answer as it would have been given then, which is what makes a past answer reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf walks the graph as it stood at an instant of the world, RFC 3339: the edges that held then. Absent is now. | [optional] 
**Depth** | Pointer to **int64** | Depth is how many hops. Absent is one. | [optional] 
**Direction** | Pointer to **string** | Direction is out, in or both. Out follows an edge from its entity to its value — what the node points at; in follows it the other way — what points at the node; both is the union of the two, not a third rule. Absent is out. | [optional] 
**Relation** | Pointer to **string** | Relation narrows the walk to one edge relation. Absent follows all. Only edges IN FORCE are ever followed — the assertion that wins its (entity, relation) at the instant — so a superseded or retracted edge is not a hop, and neither is a property. | [optional] 
**Seeds** | **[]string** | Seeds is where the walk starts. At least one. | 

## Methods

### NewGraphGraphNeighborsIn

`func NewGraphGraphNeighborsIn(seeds []string, ) *GraphGraphNeighborsIn`

NewGraphGraphNeighborsIn instantiates a new GraphGraphNeighborsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphNeighborsInWithDefaults

`func NewGraphGraphNeighborsInWithDefaults() *GraphGraphNeighborsIn`

NewGraphGraphNeighborsInWithDefaults instantiates a new GraphGraphNeighborsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphNeighborsIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphNeighborsIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphNeighborsIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphNeighborsIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphNeighborsIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphNeighborsIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphNeighborsIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphNeighborsIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetDepth

`func (o *GraphGraphNeighborsIn) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GraphGraphNeighborsIn) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GraphGraphNeighborsIn) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GraphGraphNeighborsIn) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetDirection

`func (o *GraphGraphNeighborsIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GraphGraphNeighborsIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GraphGraphNeighborsIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GraphGraphNeighborsIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphNeighborsIn) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphNeighborsIn) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphNeighborsIn) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphNeighborsIn) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSeeds

`func (o *GraphGraphNeighborsIn) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *GraphGraphNeighborsIn) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *GraphGraphNeighborsIn) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


