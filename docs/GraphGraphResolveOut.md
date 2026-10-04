# GraphGraphResolveOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsKnown** | Pointer to **string** | AsKnown is the knowledge instant it was taken at, RFC 3339, the same way. | [optional] 
**AsOf** | Pointer to **string** | AsOf is the valid instant this answer was taken at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Conflicts** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Conflicts is, for a relation holding one value, every weaker account of the winner&#39;s moment that names a different value, strongest first. They lost the order and stay visible, so a disagreement is never resolved into silence. | [optional] 
**Contested** | Pointer to **bool** | Contested is true exactly when Conflicts is not empty. | [optional] 
**Entity** | Pointer to **string** | Entity is the entity the question named, echoed so a stored answer still says what it is about. | [optional] 
**Held** | Pointer to [**[]GraphWireFact**](GraphWireFact.md) | Held is every statement holding at the point, Winner first, one per value: one for a relation that holds one value at a time, and every one the pair holds for a relation that holds many — the five holders of a position over five terms are five. | [optional] 
**Known** | Pointer to **bool** | Known is false when nothing held at the point. That is an answer, not an error. | [optional] 
**Relation** | Pointer to **string** | Relation is the relation the question named, echoed for the same reason. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says this pair holds more statements than one read returns, so the answer was decided from the ceiling-full begun most recently. It is reported because a provenance plane that trims silently is a plane that answers confidently and wrongly; narrow the question with as_of to see what it dropped. | [optional] 
**Winner** | Pointer to [**GraphWireFact**](GraphWireFact.md) | Winner is the statement holding at the point that began most recently, then the strongest under the order &#x60;rule&#x60; names. Absent exactly when Known is false. | [optional] 

## Methods

### NewGraphGraphResolveOut

`func NewGraphGraphResolveOut() *GraphGraphResolveOut`

NewGraphGraphResolveOut instantiates a new GraphGraphResolveOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphGraphResolveOutWithDefaults

`func NewGraphGraphResolveOutWithDefaults() *GraphGraphResolveOut`

NewGraphGraphResolveOutWithDefaults instantiates a new GraphGraphResolveOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsKnown

`func (o *GraphGraphResolveOut) GetAsKnown() string`

GetAsKnown returns the AsKnown field if non-nil, zero value otherwise.

### GetAsKnownOk

`func (o *GraphGraphResolveOut) GetAsKnownOk() (*string, bool)`

GetAsKnownOk returns a tuple with the AsKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsKnown

`func (o *GraphGraphResolveOut) SetAsKnown(v string)`

SetAsKnown sets AsKnown field to given value.

### HasAsKnown

`func (o *GraphGraphResolveOut) HasAsKnown() bool`

HasAsKnown returns a boolean if a field has been set.

### GetAsOf

`func (o *GraphGraphResolveOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphGraphResolveOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphGraphResolveOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphGraphResolveOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetConflicts

`func (o *GraphGraphResolveOut) GetConflicts() []GraphWireFact`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *GraphGraphResolveOut) GetConflictsOk() (*[]GraphWireFact, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *GraphGraphResolveOut) SetConflicts(v []GraphWireFact)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *GraphGraphResolveOut) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetContested

`func (o *GraphGraphResolveOut) GetContested() bool`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *GraphGraphResolveOut) GetContestedOk() (*bool, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *GraphGraphResolveOut) SetContested(v bool)`

SetContested sets Contested field to given value.

### HasContested

`func (o *GraphGraphResolveOut) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetEntity

`func (o *GraphGraphResolveOut) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphGraphResolveOut) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphGraphResolveOut) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphGraphResolveOut) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetHeld

`func (o *GraphGraphResolveOut) GetHeld() []GraphWireFact`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *GraphGraphResolveOut) GetHeldOk() (*[]GraphWireFact, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *GraphGraphResolveOut) SetHeld(v []GraphWireFact)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *GraphGraphResolveOut) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetKnown

`func (o *GraphGraphResolveOut) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *GraphGraphResolveOut) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *GraphGraphResolveOut) SetKnown(v bool)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *GraphGraphResolveOut) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetRelation

`func (o *GraphGraphResolveOut) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphGraphResolveOut) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphGraphResolveOut) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphGraphResolveOut) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphGraphResolveOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphGraphResolveOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphGraphResolveOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphGraphResolveOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetWinner

`func (o *GraphGraphResolveOut) GetWinner() GraphWireFact`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GraphGraphResolveOut) GetWinnerOk() (*GraphWireFact, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GraphGraphResolveOut) SetWinner(v GraphWireFact)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GraphGraphResolveOut) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


