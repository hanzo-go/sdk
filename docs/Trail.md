# Trail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Broken** | Pointer to **int32** | Broken counts the chains whose hash chain fails; each names where. | [optional] 
**Chains** | Pointer to [**[]Integrity**](Integrity.md) | Chains is every chain of the family, in name order. | [optional] 
**Intact** | Pointer to **int32** | Intact counts the chains that verified end to end. | [optional] 
**Records** | Pointer to **int32** | Records is the total number of records walked across every chain that was read. It counts nothing for an unread chain. | [optional] 
**Unread** | Pointer to **int32** | Unread counts the chains that could not be read. NOT a pass: nothing is known about their contents. | [optional] 

## Methods

### NewTrail

`func NewTrail() *Trail`

NewTrail instantiates a new Trail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrailWithDefaults

`func NewTrailWithDefaults() *Trail`

NewTrailWithDefaults instantiates a new Trail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBroken

`func (o *Trail) GetBroken() int32`

GetBroken returns the Broken field if non-nil, zero value otherwise.

### GetBrokenOk

`func (o *Trail) GetBrokenOk() (*int32, bool)`

GetBrokenOk returns a tuple with the Broken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBroken

`func (o *Trail) SetBroken(v int32)`

SetBroken sets Broken field to given value.

### HasBroken

`func (o *Trail) HasBroken() bool`

HasBroken returns a boolean if a field has been set.

### GetChains

`func (o *Trail) GetChains() []Integrity`

GetChains returns the Chains field if non-nil, zero value otherwise.

### GetChainsOk

`func (o *Trail) GetChainsOk() (*[]Integrity, bool)`

GetChainsOk returns a tuple with the Chains field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChains

`func (o *Trail) SetChains(v []Integrity)`

SetChains sets Chains field to given value.

### HasChains

`func (o *Trail) HasChains() bool`

HasChains returns a boolean if a field has been set.

### GetIntact

`func (o *Trail) GetIntact() int32`

GetIntact returns the Intact field if non-nil, zero value otherwise.

### GetIntactOk

`func (o *Trail) GetIntactOk() (*int32, bool)`

GetIntactOk returns a tuple with the Intact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntact

`func (o *Trail) SetIntact(v int32)`

SetIntact sets Intact field to given value.

### HasIntact

`func (o *Trail) HasIntact() bool`

HasIntact returns a boolean if a field has been set.

### GetRecords

`func (o *Trail) GetRecords() int32`

GetRecords returns the Records field if non-nil, zero value otherwise.

### GetRecordsOk

`func (o *Trail) GetRecordsOk() (*int32, bool)`

GetRecordsOk returns a tuple with the Records field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecords

`func (o *Trail) SetRecords(v int32)`

SetRecords sets Records field to given value.

### HasRecords

`func (o *Trail) HasRecords() bool`

HasRecords returns a boolean if a field has been set.

### GetUnread

`func (o *Trail) GetUnread() int32`

GetUnread returns the Unread field if non-nil, zero value otherwise.

### GetUnreadOk

`func (o *Trail) GetUnreadOk() (*int32, bool)`

GetUnreadOk returns a tuple with the Unread field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnread

`func (o *Trail) SetUnread(v int32)`

SetUnread sets Unread field to given value.

### HasUnread

`func (o *Trail) HasUnread() bool`

HasUnread returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


