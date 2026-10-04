# GraphGraphPathIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is how much this plane had heard, RFC 3339: the answer as it would have been given then, which is what makes a past answer reproducible. Absent is now. | [optional] 
**AsOf** | Pointer to **string** | AsOf finds the path through the graph as it stood at an instant of the world, RFC 3339: an edge that did not hold then is not crossed. Absent is now. | [optional] 
**Direction** | Pointer to **string** | Direction is out, in or both. Out crosses each edge from its entity to its value; in crosses it backwards; both crosses it either way. Absent is out. | [optional] 
**From** | **string** | From is the entity the path starts at. Required. | 
**Relations** | Pointer to **[]string** | Relations is the edge relations a path may cross. Absent crosses every relation. A causal chain is this list naming the relations that mean cause — &#x60;caused_by&#x60; and &#x60;influenced&#x60; under the decision convention. At most 256. | [optional] 
**To** | **string** | To is the entity the path ends at. Required. | 

## Methods

### NewGraphGraphPathIn

`func NewGraphGraphPathIn(from string, to string, ) *GraphGraphPathIn`

NewGraphGraphPathIn instantiates a new GraphGraphPathIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphPathInWithDefaults

`func NewGraphGraphPathInWithDefaults() *GraphGraphPathIn`

NewGraphGraphPathInWithDefaults instantiates a new GraphGraphPathIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphPathIn) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphPathIn) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphPathIn) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphPathIn) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphPathIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphPathIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphPathIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphPathIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetDirection

`func (o *GraphGraphPathIn) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *GraphGraphPathIn) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *GraphGraphPathIn) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *GraphGraphPathIn) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetFrom

`func (o *GraphGraphPathIn) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *GraphGraphPathIn) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *GraphGraphPathIn) SetFrom(v string)`

SetFrom sets From field to given value.


### GetRelations

`func (o *GraphGraphPathIn) GetRelations() []string`

GetRelations returns the Relations field if non-nil, zero value otherwise.

### GetRelationsOk

`func (o *GraphGraphPathIn) GetRelationsOk() (*[]string, bool)`

GetRelationsOk returns a tuple with the Relations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelations

`func (o *GraphGraphPathIn) SetRelations(v []string)`

SetRelations sets Relations field to given value.

### HasRelations

`func (o *GraphGraphPathIn) HasRelations() bool`

HasRelations returns a boolean if a field has been set.

### GetTo

`func (o *GraphGraphPathIn) GetTo() string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *GraphGraphPathIn) GetToOk() (*string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *GraphGraphPathIn) SetTo(v string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


