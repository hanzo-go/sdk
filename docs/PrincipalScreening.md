# PrincipalScreening

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **int64** | At is when it ran, unix seconds. | [optional] 
**Lists** | Pointer to [**[]PrincipalList**](PrincipalList.md) | Lists is each publisher&#39;s readiness at the moment of screening. | [optional] 
**Matches** | Pointer to [**[]PrincipalMatch**](PrincipalMatch.md) | Matches are what matched — answered to the org&#39;s own admins and to a reviewer, never to another org. | [optional] 
**Reason** | Pointer to **string** | Reason says why, in words. | [optional] 
**Status** | Pointer to **string** | Status is clear, review (a possible match awaits a platform reviewer), blocked (a confirmed match), unscreened (the org states nothing to screen), or unavailable (the lists are not fit to screen against). | [optional] 

## Methods

### NewPrincipalScreening

`func NewPrincipalScreening() *PrincipalScreening`

NewPrincipalScreening instantiates a new PrincipalScreening object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalScreeningWithDefaults

`func NewPrincipalScreeningWithDefaults() *PrincipalScreening`

NewPrincipalScreeningWithDefaults instantiates a new PrincipalScreening object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *PrincipalScreening) GetAt() int64`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *PrincipalScreening) GetAtOk() (*int64, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *PrincipalScreening) SetAt(v int64)`

SetAt sets At field to given value.

### HasAt

`func (o *PrincipalScreening) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetLists

`func (o *PrincipalScreening) GetLists() []PrincipalList`

GetLists returns the Lists field if non-nil, zero value otherwise.

### GetListsOk

`func (o *PrincipalScreening) GetListsOk() (*[]PrincipalList, bool)`

GetListsOk returns a tuple with the Lists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLists

`func (o *PrincipalScreening) SetLists(v []PrincipalList)`

SetLists sets Lists field to given value.

### HasLists

`func (o *PrincipalScreening) HasLists() bool`

HasLists returns a boolean if a field has been set.

### GetMatches

`func (o *PrincipalScreening) GetMatches() []PrincipalMatch`

GetMatches returns the Matches field if non-nil, zero value otherwise.

### GetMatchesOk

`func (o *PrincipalScreening) GetMatchesOk() (*[]PrincipalMatch, bool)`

GetMatchesOk returns a tuple with the Matches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatches

`func (o *PrincipalScreening) SetMatches(v []PrincipalMatch)`

SetMatches sets Matches field to given value.

### HasMatches

`func (o *PrincipalScreening) HasMatches() bool`

HasMatches returns a boolean if a field has been set.

### GetReason

`func (o *PrincipalScreening) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PrincipalScreening) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PrincipalScreening) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PrincipalScreening) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalScreening) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalScreening) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalScreening) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalScreening) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


