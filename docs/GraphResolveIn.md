# GraphResolveIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is the instant to answer at, RFC 3339. Absent means now. | [optional] 
**Entity** | Pointer to **string** | Entity is the thing to answer about. Required. | [optional] 
**Relation** | Pointer to **string** | Relation is the one relation to settle. Required: this answers a single (entity, relation) pair, never a whole entity at once. | [optional] 

## Methods

### NewGraphResolveIn

`func NewGraphResolveIn() *GraphResolveIn`

NewGraphResolveIn instantiates a new GraphResolveIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphResolveInWithDefaults

`func NewGraphResolveInWithDefaults() *GraphResolveIn`

NewGraphResolveInWithDefaults instantiates a new GraphResolveIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *GraphResolveIn) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphResolveIn) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphResolveIn) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphResolveIn) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetEntity

`func (o *GraphResolveIn) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphResolveIn) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphResolveIn) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphResolveIn) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetRelation

`func (o *GraphResolveIn) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphResolveIn) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphResolveIn) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphResolveIn) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


