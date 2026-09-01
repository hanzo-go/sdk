# Roster

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chains** | Pointer to [**[]Market**](Market.md) | Chains is &#x60;[]&#x60; where the registry answered and named none, and &#x60;null&#x60; where it did not answer — never absent, because a missing key and an empty list read alike and only one of them means \&quot;there are none\&quot;. Each row carries its OWN reach for its figures, so a chain whose indexer is down is one row saying so. | [optional] 
**Reach** | Pointer to [**Reach**](Reach.md) | Reach is how far the read of the REGISTRY got. It governs the list: a registry that did not answer yields no rows, and the reason it did not is here rather than in an empty array a caller would read as \&quot;no chains exist\&quot;. | [optional] 

## Methods

### NewRoster

`func NewRoster() *Roster`

NewRoster instantiates a new Roster object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRosterWithDefaults

`func NewRosterWithDefaults() *Roster`

NewRosterWithDefaults instantiates a new Roster object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChains

`func (o *Roster) GetChains() []Market`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *Roster) GetChainsOk() (*[]Market, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *Roster) SetChains(v []Market)`

SetChains sets Chains field to given value.

### HasChains

`func (o *Roster) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetReach

`func (o *Roster) GetReach() Reach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *Roster) GetReachOk() (*Reach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *Roster) SetReach(v Reach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *Roster) HasReach() bool`

HasReach returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


