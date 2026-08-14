# ReferenceSetsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Refused** | Pointer to **[]string** | Refused names the sets that cannot be consulted at all. A key checked against one of these is UNKNOWN, not clean. | [optional] 
**Sets** | Pointer to [**[]ReferenceSet**](ReferenceSet.md) | Sets is the whole catalog, in a stable order. | [optional] 
**Stale** | Pointer to **[]string** | Stale names the sets past their freshness bound — the list to alarm on. | [optional] 

## Methods

### NewReferenceSetsOut

`func NewReferenceSetsOut() *ReferenceSetsOut`

NewReferenceSetsOut instantiates a new ReferenceSetsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferenceSetsOutWithDefaults

`func NewReferenceSetsOutWithDefaults() *ReferenceSetsOut`

NewReferenceSetsOutWithDefaults instantiates a new ReferenceSetsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefused

`func (o *ReferenceSetsOut) GetRefused() []string`

GetRefused returns the Refused field if non-nil, zero value otherwise.

### GetRefusedOk

`func (o *ReferenceSetsOut) GetRefusedOk() (*[]string, bool)`

GetRefusedOk returns a tuple with the Refused field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefused

`func (o *ReferenceSetsOut) SetRefused(v []string)`

SetRefused sets Refused field to given value.

### HasRefused

`func (o *ReferenceSetsOut) HasRefused() bool`

HasRefused returns a boolean if a field has been set.

### GetSets

`func (o *ReferenceSetsOut) GetSets() []ReferenceSet`

GetSets returns the Sets field if non-nil, zero value otherwise.

### GetSetsOk

`func (o *ReferenceSetsOut) GetSetsOk() (*[]ReferenceSet, bool)`

GetSetsOk returns a tuple with the Sets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSets

`func (o *ReferenceSetsOut) SetSets(v []ReferenceSet)`

SetSets sets Sets field to given value.

### HasSets

`func (o *ReferenceSetsOut) HasSets() bool`

HasSets returns a boolean if a field has been set.

### GetStale

`func (o *ReferenceSetsOut) GetStale() []string`

GetStale returns the Stale field if non-nil, zero value otherwise.

### GetStaleOk

`func (o *ReferenceSetsOut) GetStaleOk() (*[]string, bool)`

GetStaleOk returns a tuple with the Stale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStale

`func (o *ReferenceSetsOut) SetStale(v []string)`

SetStale sets Stale field to given value.

### HasStale

`func (o *ReferenceSetsOut) HasStale() bool`

HasStale returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


