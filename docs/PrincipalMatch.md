# PrincipalMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecidedAt** | Pointer to **int64** | DecidedAt is when, unix seconds. | [optional] 
**DecidedBy** | Pointer to **string** | DecidedBy is the reviewer, or \&quot;exact designated address\&quot;. | [optional] 
**Entry** | Pointer to **string** | Entry is the designated subject&#39;s primary name, or the jurisdiction&#39;s. | [optional] 
**FoundAt** | Pointer to **int64** | FoundAt is when a screening first found it, unix seconds. | [optional] 
**Hits** | Pointer to [**[]PrincipalHit**](PrincipalHit.md) | Hits are the org&#39;s names, addresses and countries that matched, in the order they were first seen. A decision covers the hits it was made on. | [optional] 
**Id** | Pointer to **string** | ID addresses the match for a reviewer&#39;s decision. It is the org and the designation together, so a decision is about the org whatever it is called. | [optional] 
**List** | Pointer to **string** | List is the publisher — ofac, un, eu or ofsi — or jurisdiction for an embargoed place. | [optional] 
**Note** | Pointer to **string** | Note is the reviewer&#39;s reason. | [optional] 
**Org** | Pointer to **string** | Org is the org the match names. | [optional] 
**Reason** | Pointer to **string** | Reason is what first matched, and which identifiers agreed or conflicted. | [optional] 
**Ref** | Pointer to **string** | Ref is the entry&#39;s reference on that list, or the country code. | [optional] 
**Score** | Pointer to **float64** | Score is the highest similarity seen, 0 to 1; an address or a jurisdiction is 1. | [optional] 
**Status** | Pointer to **string** | Status is open (a reviewer decides), cleared (a namesake, decided) or confirmed (the designated person, decided — or a designated address). | [optional] 

## Methods

### NewPrincipalMatch

`func NewPrincipalMatch() *PrincipalMatch`

NewPrincipalMatch instantiates a new PrincipalMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPrincipalMatchWithDefaults

`func NewPrincipalMatchWithDefaults() *PrincipalMatch`

NewPrincipalMatchWithDefaults instantiates a new PrincipalMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecidedAt

`func (o *PrincipalMatch) GetDecidedAt() int64`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *PrincipalMatch) GetDecidedAtOk() (*int64, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *PrincipalMatch) SetDecidedAt(v int64)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *PrincipalMatch) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### GetDecidedBy

`func (o *PrincipalMatch) GetDecidedBy() string`

GetDecidedBy returns the DecidedBy field if non-nil, zero value otherwise.

### GetDecidedByOk

`func (o *PrincipalMatch) GetDecidedByOk() (*string, bool)`

GetDecidedByOk returns a tuple with the DecidedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedBy

`func (o *PrincipalMatch) SetDecidedBy(v string)`

SetDecidedBy sets DecidedBy field to given value.

### HasDecidedBy

`func (o *PrincipalMatch) HasDecidedBy() bool`

HasDecidedBy returns a boolean if a field has been set.

### GetEntry

`func (o *PrincipalMatch) GetEntry() string`

GetEntry returns the Entry field if non-nil, zero value otherwise.

### GetEntryOk

`func (o *PrincipalMatch) GetEntryOk() (*string, bool)`

GetEntryOk returns a tuple with the Entry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntry

`func (o *PrincipalMatch) SetEntry(v string)`

SetEntry sets Entry field to given value.

### HasEntry

`func (o *PrincipalMatch) HasEntry() bool`

HasEntry returns a boolean if a field has been set.

### GetFoundAt

`func (o *PrincipalMatch) GetFoundAt() int64`

GetFoundAt returns the FoundAt field if non-nil, zero value otherwise.

### GetFoundAtOk

`func (o *PrincipalMatch) GetFoundAtOk() (*int64, bool)`

GetFoundAtOk returns a tuple with the FoundAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFoundAt

`func (o *PrincipalMatch) SetFoundAt(v int64)`

SetFoundAt sets FoundAt field to given value.

### HasFoundAt

`func (o *PrincipalMatch) HasFoundAt() bool`

HasFoundAt returns a boolean if a field has been set.

### GetHits

`func (o *PrincipalMatch) GetHits() []PrincipalHit`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *PrincipalMatch) GetHitsOk() (*[]PrincipalHit, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *PrincipalMatch) SetHits(v []PrincipalHit)`

SetHits sets Hits field to given value.

### HasHits

`func (o *PrincipalMatch) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetId

`func (o *PrincipalMatch) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PrincipalMatch) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PrincipalMatch) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PrincipalMatch) HasId() bool`

HasId returns a boolean if a field has been set.

### GetList

`func (o *PrincipalMatch) GetList() string`

GetList returns the List field if non-nil, zero value otherwise.

### GetListOk

`func (o *PrincipalMatch) GetListOk() (*string, bool)`

GetListOk returns a tuple with the List field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetList

`func (o *PrincipalMatch) SetList(v string)`

SetList sets List field to given value.

### HasList

`func (o *PrincipalMatch) HasList() bool`

HasList returns a boolean if a field has been set.

### GetNote

`func (o *PrincipalMatch) GetNote() string`

GetNote returns the Note field if non-nil, zero value otherwise.

### GetNoteOk

`func (o *PrincipalMatch) GetNoteOk() (*string, bool)`

GetNoteOk returns a tuple with the Note field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNote

`func (o *PrincipalMatch) SetNote(v string)`

SetNote sets Note field to given value.

### HasNote

`func (o *PrincipalMatch) HasNote() bool`

HasNote returns a boolean if a field has been set.

### GetOrg

`func (o *PrincipalMatch) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *PrincipalMatch) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *PrincipalMatch) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *PrincipalMatch) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetReason

`func (o *PrincipalMatch) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *PrincipalMatch) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *PrincipalMatch) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *PrincipalMatch) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetRef

`func (o *PrincipalMatch) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PrincipalMatch) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PrincipalMatch) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PrincipalMatch) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetScore

`func (o *PrincipalMatch) GetScore() float64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *PrincipalMatch) GetScoreOk() (*float64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *PrincipalMatch) SetScore(v float64)`

SetScore sets Score field to given value.

### HasScore

`func (o *PrincipalMatch) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetStatus

`func (o *PrincipalMatch) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PrincipalMatch) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PrincipalMatch) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PrincipalMatch) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


