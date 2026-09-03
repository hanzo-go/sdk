# GraphNeighborsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Bound** | Pointer to **int64** | Bound is the ceiling this walk was held to, the same for every caller, so Truncated can be read against a number rather than guessed at. | [optional] 
**Depth** | Pointer to **int64** | Depth is the deepest hop count actually reached. It is at most the depth asked for, and smaller when the walk ran out of edges first. | [optional] 
**Entities** | Pointer to **[]string** | Entities is everything reached, the seeds included, ordered by the fewest hops that reach each one and then by key. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says the bound stopped the walk. The bound is part of the answer rather than a silent short read. | [optional] 

## Methods

### NewGraphNeighborsOut

`func NewGraphNeighborsOut() *GraphNeighborsOut`

NewGraphNeighborsOut instantiates a new GraphNeighborsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphNeighborsOutWithDefaults

`func NewGraphNeighborsOutWithDefaults() *GraphNeighborsOut`

NewGraphNeighborsOutWithDefaults instantiates a new GraphNeighborsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBound

`func (o *GraphNeighborsOut) GetBound() int64`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *GraphNeighborsOut) GetBoundOk() (*int64, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *GraphNeighborsOut) SetBound(v int64)`

SetBound sets Bound field to given value.

### HasBound

`func (o *GraphNeighborsOut) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetDepth

`func (o *GraphNeighborsOut) GetDepth() int64`

GetDepth returns the Depth field if non-nil, zero value otherwise.

### GetDepthOk

`func (o *GraphNeighborsOut) GetDepthOk() (*int64, bool)`

GetDepthOk returns a tuple with the Depth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepth

`func (o *GraphNeighborsOut) SetDepth(v int64)`

SetDepth sets Depth field to given value.

### HasDepth

`func (o *GraphNeighborsOut) HasDepth() bool`

HasDepth returns a boolean if a field has been set.

### GetEntities

`func (o *GraphNeighborsOut) GetEntities() []string`

GetEntities returns the Entities field if non-nil, zero value otherwise.

### GetEntitiesOk

`func (o *GraphNeighborsOut) GetEntitiesOk() (*[]string, bool)`

GetEntitiesOk returns a tuple with the Entities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntities

`func (o *GraphNeighborsOut) SetEntities(v []string)`

SetEntities sets Entities field to given value.

### HasEntities

`func (o *GraphNeighborsOut) HasEntities() bool`

HasEntities returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphNeighborsOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphNeighborsOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphNeighborsOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphNeighborsOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


