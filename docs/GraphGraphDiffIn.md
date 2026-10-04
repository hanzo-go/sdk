# GraphGraphDiffIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Entity** | Pointer to **string** | Entity narrows the diff to what was asserted about one entity. Absent diffs the whole graph, bounded by Limit. | [optional] 
**From** | Pointer to **string** | From is the earlier instant of the world, RFC 3339. What held here is the baseline, so a statement beginning exactly at From is part of it and not a change. Absent is To: the diff is then of what was heard, not of what happened. | [optional] 
**FromKnown** | Pointer to **string** | FromKnown is how much this plane had heard at the baseline, RFC 3339. Absent is ToKnown: the diff is then of the world, as known once. | [optional] 
**Limit** | Pointer to **int64** | Limit caps how many (entity, relation) pairs are examined, in key order. Absent, zero, or anything above the walk ceiling is the ceiling. | [optional] 
**To** | Pointer to **string** | To is the later instant of the world, RFC 3339, and may not precede From. A statement beginning or ending exactly at To is in the diff. Absent is now. | [optional] 
**ToKnown** | Pointer to **string** | ToKnown is how much it had heard at the later point, RFC 3339, and may not precede FromKnown. Absent is now. | [optional] 

## Methods

### NewGraphGraphDiffIn

`func NewGraphGraphDiffIn() *GraphGraphDiffIn`

NewGraphGraphDiffIn instantiates a new GraphGraphDiffIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphDiffInWithDefaults

`func NewGraphGraphDiffInWithDefaults() *GraphGraphDiffIn`

NewGraphGraphDiffInWithDefaults instantiates a new GraphGraphDiffIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEntity

`func (o *GraphGraphDiffIn) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphDiffIn) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphDiffIn) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphGraphDiffIn) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetFrom

`func (o *GraphGraphDiffIn) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphGraphDiffIn) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphGraphDiffIn) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *GraphGraphDiffIn) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetFromKnown

`func (o *GraphGraphDiffIn) GetFromKnown() string`

GetFromKnown returns the FromKnown field if non-nil, zero value otherwise.

### GetFromKnownOk

`func (o *GraphGraphDiffIn) GetFromKnownOk() (*string, bool)`

GetFromKnownOk returns a tuple with the FromKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromKnown

`func (o *GraphGraphDiffIn) SetFromKnown(v string)`

SetFromKnown sets FromKnown field to given value.

### HasFromKnown

`func (o *GraphGraphDiffIn) HasFromKnown() bool`

HasFromKnown returns a boolean if a field has been set.

### GetLimit

`func (o *GraphGraphDiffIn) GetLimit() int64`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GraphGraphDiffIn) GetLimitOk() (*int64, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GraphGraphDiffIn) SetLimit(v int64)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *GraphGraphDiffIn) HasLimit() bool`

HasLimit returns a boolean if a field has been set.

### GetTo

`func (o *GraphGraphDiffIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphGraphDiffIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphGraphDiffIn) SetTo(v string)`

SetTo sets To field to given value.

### HasTo

`func (o *GraphGraphDiffIn) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetToKnown

`func (o *GraphGraphDiffIn) GetToKnown() string`

GetToKnown returns the ToKnown field if non-nil, zero value otherwise.

### GetToKnownOk

`func (o *GraphGraphDiffIn) GetToKnownOk() (*string, bool)`

GetToKnownOk returns a tuple with the ToKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToKnown

`func (o *GraphGraphDiffIn) SetToKnown(v string)`

SetToKnown sets ToKnown field to given value.

### HasToKnown

`func (o *GraphGraphDiffIn) HasToKnown() bool`

HasToKnown returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


