# GraphGraphCommunitiesIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339: the answer as it would have been given then, which is what makes a past answer reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf partitions the graph as it stood at an instant of the world, RFC 3339: the edges that held then. Absent is now. | [optional] 
**Level** | Pointer to **int64** | Level is how fine a partition to answer with. 0, the default, is the coarsest — the partition detection stops at, where modularity is highest — and each level above it is the pass before, finer. &#x60;levels&#x60; in the answer says how many there are. | [optional] 
**Relation** | Pointer to **string** | Relation narrows the graph to one edge relation, by its key. Absent partitions every edge. | [optional] 

## Methods

### NewGraphGraphCommunitiesIn

`func NewGraphGraphCommunitiesIn() *GraphGraphCommunitiesIn`

NewGraphGraphCommunitiesIn instantiates a new GraphGraphCommunitiesIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphCommunitiesInWithDefaults

`func NewGraphGraphCommunitiesInWithDefaults() *GraphGraphCommunitiesIn`

NewGraphGraphCommunitiesInWithDefaults instantiates a new GraphGraphCommunitiesIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphCommunitiesIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphCommunitiesIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphCommunitiesIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphCommunitiesIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphCommunitiesIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphCommunitiesIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphCommunitiesIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphCommunitiesIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetLevel

`func (o *GraphGraphCommunitiesIn) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GraphGraphCommunitiesIn) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GraphGraphCommunitiesIn) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GraphGraphCommunitiesIn) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphCommunitiesIn) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphCommunitiesIn) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphCommunitiesIn) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphCommunitiesIn) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


