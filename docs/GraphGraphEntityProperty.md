# GraphGraphEntityProperty

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Conflicts** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Conflicts is every weaker account of the winner&#39;s moment that names a different value, strongest first. | [optional] 
**Contested** | Pointer to **bool** | Contested is true exactly when Conflicts is not empty. | [optional] 
**Relation** | Pointer to **string** | Relation is the scalar relation&#39;s key. | [optional] 
**Value** | Pointer to **string** | Value is the winning scalar value holding at the point. | [optional] 
**Winner** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Winner is the assertion that speaks for this relation at the point. | [optional] 

## Methods

### NewGraphGraphEntityProperty

`func NewGraphGraphEntityProperty() *GraphGraphEntityProperty`

NewGraphGraphEntityProperty instantiates a new GraphGraphEntityProperty object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphEntityPropertyWithDefaults

`func NewGraphGraphEntityPropertyWithDefaults() *GraphGraphEntityProperty`

NewGraphGraphEntityPropertyWithDefaults instantiates a new GraphGraphEntityProperty object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConflicts

`func (o *GraphGraphEntityProperty) GetConflicts() []GraphWireFact`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *GraphGraphEntityProperty) GetConflictsOk() (*[]GraphWireFact, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *GraphGraphEntityProperty) SetConflicts(v []GraphWireFact)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *GraphGraphEntityProperty) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetContested

`func (o *GraphGraphEntityProperty) GetContested() bool`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *GraphGraphEntityProperty) GetContestedOk() (*bool, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *GraphGraphEntityProperty) SetContested(v bool)`

SetContested sets Contested field to given value.

### HasContested

`func (o *GraphGraphEntityProperty) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphEntityProperty) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphEntityProperty) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphEntityProperty) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphEntityProperty) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetValue

`func (o *GraphGraphEntityProperty) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GraphGraphEntityProperty) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GraphGraphEntityProperty) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GraphGraphEntityProperty) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetWinner

`func (o *GraphGraphEntityProperty) GetWinner() GraphWireFact`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GraphGraphEntityProperty) GetWinnerOk() (*GraphWireFact, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GraphGraphEntityProperty) SetWinner(v GraphWireFact)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GraphGraphEntityProperty) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


