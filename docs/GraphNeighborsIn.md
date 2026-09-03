# GraphNeighborsIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf walks the graph as it stood at an instant, RFC 3339. Absent walks it as it stands now. | [optional] 
**Depth** | Pointer to **int64** | Depth is how many hops. Absent is one. | [optional] 
**Direction** | Pointer to **string** | Direction is out, in or both. Out follows an edge from its entity to its value — what the node points at; in follows it the other way — what points at the node; both is the union of the two, not a third rule. Absent is out. | [optional] 
**Relation** | Pointer to **string** | Relation narrows the walk to one edge relation. Absent follows all. Only edges are ever followed: an assertion whose value is a scalar is a property and is never a hop. | [optional] 
**Seeds** | Pointer to **[]string** | Seeds is where the walk starts. At least one. | [optional] 

## Methods

### NewGraphNeighborsIn

`func NewGraphNeighborsIn() *GraphNeighborsIn`

NewGraphNeighborsIn instantiates a new GraphNeighborsIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNeighborsInWithDefaults

`func NewGraphNeighborsInWithDefaults() *GraphNeighborsIn`

NewGraphNeighborsInWithDefaults instantiates a new GraphNeighborsIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *GraphNeighborsIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphNeighborsIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphNeighborsIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphNeighborsIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetDepth

`func (o *GraphNeighborsIn) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GraphNeighborsIn) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GraphNeighborsIn) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GraphNeighborsIn) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetDirection

`func (o *GraphNeighborsIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GraphNeighborsIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GraphNeighborsIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GraphNeighborsIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetRelation

`func (o *GraphNeighborsIn) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphNeighborsIn) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphNeighborsIn) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphNeighborsIn) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetSeeds

`func (o *GraphNeighborsIn) GetSeeds() []string`

GetSeeds returns the Seeds field if non-nil, zero value otherwise.

### GetSeedsOk

`func (o *GraphNeighborsIn) GetSeedsOk() (*[]string, bool)`

GetSeedsOk returns a tuple with the Seeds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeeds

`func (o *GraphNeighborsIn) SetSeeds(v []string)`

SetSeeds sets Seeds field to given value.

### HasSeeds

`func (o *GraphNeighborsIn) HasSeeds() bool`

HasSeeds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


