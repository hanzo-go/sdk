# GraphGraphCommunitiesOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant it was taken at, RFC 3339, the same way. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the instant the graph was read at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Bound** | Pointer to **int64** | Bound is the most members listed in one answer. | [optional] 
**Ceiling** | Pointer to **int64** | Ceiling is the most edge assertions one partition reads — every row, so an edge three sources asserted counts three times. A graph over it is refused rather than partitioned from part of itself. | [optional] 
**Communities** | Pointer to [**[]GraphGraphCommunity**](GraphGraphCommunity.md) | Communities are the communities, largest first. | [optional] 
**Edges** | Pointer to **int64** | Edges is how many distinct undirected edges were partitioned. | [optional] 
**Entities** | Pointer to **int64** | Entities is how many entities the edge graph holds. An entity with no edge in force at AsOf is in no community. | [optional] 
**Level** | Pointer to **int64** | Level is the level answered. | [optional] 
**Levels** | Pointer to **int64** | Levels is how many levels detection found. A graph where no two entities belong together more than chance has one, of single entities. | [optional] 
**Modularity** | Pointer to **float64** | Modularity is Newman&#39;s Q of this level&#39;s partition, from -0.5 to 1: the share of edge weight inside communities, less what chance would put there. Above about 0.3 is usually read as real structure rather than chance. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says the member lists stop at Bound entities in all, filled in community order. Every community is still listed with its full Size. | [optional] 

## Methods

### NewGraphGraphCommunitiesOut

`func NewGraphGraphCommunitiesOut() *GraphGraphCommunitiesOut`

NewGraphGraphCommunitiesOut instantiates a new GraphGraphCommunitiesOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphCommunitiesOutWithDefaults

`func NewGraphGraphCommunitiesOutWithDefaults() *GraphGraphCommunitiesOut`

NewGraphGraphCommunitiesOutWithDefaults instantiates a new GraphGraphCommunitiesOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphCommunitiesOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphCommunitiesOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphCommunitiesOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphCommunitiesOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphCommunitiesOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphCommunitiesOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphCommunitiesOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphCommunitiesOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetBound

`func (o *GraphGraphCommunitiesOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphGraphCommunitiesOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphGraphCommunitiesOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphGraphCommunitiesOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetCeiling

`func (o *GraphGraphCommunitiesOut) GetCeiling() int64`

GetCeiling returns the Ceiling field if non-nil, zero value otherwise.

### GetCeilingOk

`func (o *GraphGraphCommunitiesOut) GetCeilingOk() (*int64, bool)`

GetCeilingOk returns a tuple with the Ceiling field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCeiling

`func (o *GraphGraphCommunitiesOut) SetCeiling(v int64)`

SetCeiling sets Ceiling field to given value.

### HasCeiling

`func (o *GraphGraphCommunitiesOut) HasCeiling() bool`

HasCeiling returns a boolean if a field has been set.

### GetCommunities

`func (o *GraphGraphCommunitiesOut) GetCommunities() []GraphGraphCommunity`

GetCommunities returns the Communities field if non-nil, zero value otherwise.

### GetCommunitiesOk

`func (o *GraphGraphCommunitiesOut) GetCommunitiesOk() (*[]GraphGraphCommunity, bool)`

GetCommunitiesOk returns a tuple with the Communities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommunities

`func (o *GraphGraphCommunitiesOut) SetCommunities(v []GraphGraphCommunity)`

SetCommunities sets Communities field to given value.

### HasCommunities

`func (o *GraphGraphCommunitiesOut) HasCommunities() bool`

HasCommunities returns a boolean if a field has been set.

### GetEdges

`func (o *GraphGraphCommunitiesOut) GetEdges() int64`

GetEdges returns the Edges field if non-nil, zero value otherwise.

### GetEdgesOk

`func (o *GraphGraphCommunitiesOut) GetEdgesOk() (*int64, bool)`

GetEdgesOk returns a tuple with the Edges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEdges

`func (o *GraphGraphCommunitiesOut) SetEdges(v int64)`

SetEdges sets Edges field to given value.

### HasEdges

`func (o *GraphGraphCommunitiesOut) HasEdges() bool`

HasEdges returns a boolean if a field has been set.

### GetEntities

`func (o *GraphGraphCommunitiesOut) GetEntities() int64`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *GraphGraphCommunitiesOut) GetEntitiesOk() (*int64, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *GraphGraphCommunitiesOut) SetEntities(v int64)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *GraphGraphCommunitiesOut) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetLevel

`func (o *GraphGraphCommunitiesOut) GetLevel() int64`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *GraphGraphCommunitiesOut) GetLevelOk() (*int64, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *GraphGraphCommunitiesOut) SetLevel(v int64)`

SetLevel sets Level field to given value.

### HasLevel

`func (o *GraphGraphCommunitiesOut) HasLevel() bool`

HasLevel returns a boolean if a field has been set.

### GetLevels

`func (o *GraphGraphCommunitiesOut) GetLevels() int64`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *GraphGraphCommunitiesOut) GetLevelsOk() (*int64, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *GraphGraphCommunitiesOut) SetLevels(v int64)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *GraphGraphCommunitiesOut) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetModularity

`func (o *GraphGraphCommunitiesOut) GetModularity() float64`

GetModularity returns the Modularity field if non-nil, zero value otherwise.

### GetModularityOk

`func (o *GraphGraphCommunitiesOut) GetModularityOk() (*float64, bool)`

GetModularityOk returns a tuple with the Modularity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModularity

`func (o *GraphGraphCommunitiesOut) SetModularity(v float64)`

SetModularity sets Modularity field to given value.

### HasModularity

`func (o *GraphGraphCommunitiesOut) HasModularity() bool`

HasModularity returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphCommunitiesOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphCommunitiesOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphCommunitiesOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphCommunitiesOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


