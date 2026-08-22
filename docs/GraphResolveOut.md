# GraphResolveOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AsOf** | Pointer to **string** | AsOf is the instant this answer was taken at, RFC 3339: the one asked for, or the server&#39;s clock when none was. | [optional] 
**Conflicts** | Pointer to [**[]WireFact**](WireFact.md) | Conflicts is every OTHER assertion knowable at AsOf, strongest first. They are not all disagreements: one that repeats the winner&#39;s value ranks below it and is listed here too. | [optional] 
**Contested** | Pointer to **bool** | Contested is true when at least one conflict claims a value different from the winner&#39;s. Any number of conflicts that all agree leaves it false. | [optional] 
**Entity** | Pointer to **string** | Entity is the entity the question named, echoed so a stored answer still says what it is about. | [optional] 
**Known** | Pointer to **bool** | Known is false when this plane held nothing knowable at AsOf. That is an answer, not an error. | [optional] 
**Relation** | Pointer to **string** | Relation is the relation the question named, echoed for the same reason. | [optional] 
**Truncated** | Pointer to **bool** | Truncated says this pair holds more assertions than one read returns, so the winner was decided from the most recent ceiling-full of them. It is reported because a provenance plane that trims silently is a plane that answers confidently and wrongly; narrow the question with as_of to see what it dropped. | [optional] 
**Winner** | Pointer to [**WireFact**](WireFact.md) | Winner is the assertion in force — the strongest of those knowable at AsOf under the order &#x60;rule&#x60; names. Absent exactly when Known is false. | [optional] 

## Methods

### NewGraphResolveOut

`func NewGraphResolveOut() *GraphResolveOut`

NewGraphResolveOut instantiates a new GraphResolveOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGraphResolveOutWithDefaults

`func NewGraphResolveOutWithDefaults() *GraphResolveOut`

NewGraphResolveOutWithDefaults instantiates a new GraphResolveOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsOf

`func (o *GraphResolveOut) GetAsOf() string`

GetAsOf returns the AsOf field if non-nil, zero value otherwise.

### GetAsOfOk

`func (o *GraphResolveOut) GetAsOfOk() (*string, bool)`

GetAsOfOk returns a tuple with the AsOf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsOf

`func (o *GraphResolveOut) SetAsOf(v string)`

SetAsOf sets AsOf field to given value.

### HasAsOf

`func (o *GraphResolveOut) HasAsOf() bool`

HasAsOf returns a boolean if a field has been set.

### GetConflicts

`func (o *GraphResolveOut) GetConflicts() []WireFact`

GetConflicts returns the Conflicts field if non-nil, zero value otherwise.

### GetConflictsOk

`func (o *GraphResolveOut) GetConflictsOk() (*[]WireFact, bool)`

GetConflictsOk returns a tuple with the Conflicts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConflicts

`func (o *GraphResolveOut) SetConflicts(v []WireFact)`

SetConflicts sets Conflicts field to given value.

### HasConflicts

`func (o *GraphResolveOut) HasConflicts() bool`

HasConflicts returns a boolean if a field has been set.

### GetContested

`func (o *GraphResolveOut) GetContested() bool`

GetContested returns the Contested field if non-nil, zero value otherwise.

### GetContestedOk

`func (o *GraphResolveOut) GetContestedOk() (*bool, bool)`

GetContestedOk returns a tuple with the Contested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContested

`func (o *GraphResolveOut) SetContested(v bool)`

SetContested sets Contested field to given value.

### HasContested

`func (o *GraphResolveOut) HasContested() bool`

HasContested returns a boolean if a field has been set.

### GetEntity

`func (o *GraphResolveOut) GetEntity() string`

GetEntity returns the Entity field if non-nil, zero value otherwise.

### GetEntityOk

`func (o *GraphResolveOut) GetEntityOk() (*string, bool)`

GetEntityOk returns a tuple with the Entity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntity

`func (o *GraphResolveOut) SetEntity(v string)`

SetEntity sets Entity field to given value.

### HasEntity

`func (o *GraphResolveOut) HasEntity() bool`

HasEntity returns a boolean if a field has been set.

### GetKnown

`func (o *GraphResolveOut) GetKnown() bool`

GetKnown returns the Known field if non-nil, zero value otherwise.

### GetKnownOk

`func (o *GraphResolveOut) GetKnownOk() (*bool, bool)`

GetKnownOk returns a tuple with the Known field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnown

`func (o *GraphResolveOut) SetKnown(v bool)`

SetKnown sets Known field to given value.

### HasKnown

`func (o *GraphResolveOut) HasKnown() bool`

HasKnown returns a boolean if a field has been set.

### GetRelation

`func (o *GraphResolveOut) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *GraphResolveOut) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *GraphResolveOut) SetRelation(v string)`

SetRelation sets Relation field to given value.

### HasRelation

`func (o *GraphResolveOut) HasRelation() bool`

HasRelation returns a boolean if a field has been set.

### GetTruncated

`func (o *GraphResolveOut) GetTruncated() bool`

GetTruncated returns the Truncated field if non-nil, zero value otherwise.

### GetTruncatedOk

`func (o *GraphResolveOut) GetTruncatedOk() (*bool, bool)`

GetTruncatedOk returns a tuple with the Truncated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTruncated

`func (o *GraphResolveOut) SetTruncated(v bool)`

SetTruncated sets Truncated field to given value.

### HasTruncated

`func (o *GraphResolveOut) HasTruncated() bool`

HasTruncated returns a boolean if a field has been set.

### GetWinner

`func (o *GraphResolveOut) GetWinner() WireFact`

GetWinner returns the Winner field if non-nil, zero value otherwise.

### GetWinnerOk

`func (o *GraphResolveOut) GetWinnerOk() (*WireFact, bool)`

GetWinnerOk returns a tuple with the Winner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWinner

`func (o *GraphResolveOut) SetWinner(v WireFact)`

SetWinner sets Winner field to given value.

### HasWinner

`func (o *GraphResolveOut) HasWinner() bool`

HasWinner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


