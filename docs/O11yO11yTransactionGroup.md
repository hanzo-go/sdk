# O11yO11yTransactionGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObjectGroup** | Pointer to [**O11yO11yObjectGroup**](O11yO11yObjectGroup.md) | ObjectGroup is the set of objects it allows the verb on. | [optional] 
**Relation** | Pointer to **string** | Relation is the verb the grant allows. | [optional] 

## Methods

### NewO11yO11yTransactionGroup

`func NewO11yO11yTransactionGroup() *O11yO11yTransactionGroup`

NewO11yO11yTransactionGroup instantiates a new O11yO11yTransactionGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTransactionGroupWithDefaults

`func NewO11yO11yTransactionGroupWithDefaults() *O11yO11yTransactionGroup`

NewO11yO11yTransactionGroupWithDefaults instantiates a new O11yO11yTransactionGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObjectGroup

`func (o *O11yO11yTransactionGroup) GetObjectGroup() O11yO11yObjectGroup`

GetObjectGroup returns the ObjectGroup field if non-nil, zero value otherwise.

### GetObjectGroupOk

`func (o *O11yO11yTransactionGroup) GetObjectGroupOk() (*O11yO11yObjectGroup, bool)`

GetObjectGroupOk returns a tuple with the ObjectGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectGroup

`func (o *O11yO11yTransactionGroup) SetObjectGroup(v O11yO11yObjectGroup)`

SetObjectGroup sets ObjectGroup field to given value.

### HasObjectGroup

`func (o *O11yO11yTransactionGroup) HasObjectGroup() bool`

HasObjectGroup returns a boolean if a field has been set.

### GetRelation

`func (o *O11yO11yTransactionGroup) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *O11yO11yTransactionGroup) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *O11yO11yTransactionGroup) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *O11yO11yTransactionGroup) HasRelation() bool`

HasRelation returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


