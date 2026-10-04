# GraphGraphDeclared

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cardinality** | Pointer to **string** | Cardinality is how many values the relation holds at once: one, or many. Absent is the default — one for a property, many for an edge. | [optional] 
**Domain** | Pointer to **string** | Domain is the type the entity asserting it must have. Absent is any. | [optional] 
**Range** | Pointer to **string** | Range is the type the entity it points at must have; a relation with a range is an edge. Absent is any. | [optional] 
**Relation** | Pointer to **string** | Relation is the relation declared, without its &#x60;relation:&#x60; prefix. | [optional] 

## Methods

### NewGraphGraphDeclared

`func NewGraphGraphDeclared() *GraphGraphDeclared`

NewGraphGraphDeclared instantiates a new GraphGraphDeclared object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphDeclaredWithDefaults

`func NewGraphGraphDeclaredWithDefaults() *GraphGraphDeclared`

NewGraphGraphDeclaredWithDefaults instantiates a new GraphGraphDeclared object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCardinality

`func (o *GraphGraphDeclared) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *GraphGraphDeclared) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *GraphGraphDeclared) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.

### HasCardinality

`func (o *GraphGraphDeclared) HasCardinality() bool`

HasCardinality returns a boolean if a field has been set.

### GetDomain

`func (o *GraphGraphDeclared) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *GraphGraphDeclared) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *GraphGraphDeclared) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *GraphGraphDeclared) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetRange

`func (o *GraphGraphDeclared) GetRange() string`

GetRange returns the Range field if non-nil, zero value otherwise.

### GetRangeOk

`func (o *GraphGraphDeclared) GetRangeOk() (*string, bool)`

GetRangeOk returns a tuple with the Range field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRange

`func (o *GraphGraphDeclared) SetRange(v string)`

SetRange sets Range field to given value.

### HasRange

`func (o *GraphGraphDeclared) HasRange() bool`

HasRange returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphDeclared) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphDeclared) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphDeclared) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphDeclared) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


