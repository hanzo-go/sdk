# PatrolPatrolAlarmIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the alarm-panel account number that reported. It is what an activation matches a site on when no site is named. | [optional] 
**Raw** | Pointer to **interface{}** |  | [optional] 
**Ref** | Pointer to **string** | Ref is the receiving centre&#39;s own reference for the activation. | [optional] 
**Site** | Pointer to **string** | Site names the site directly, by document name, when the caller already knows it. | [optional] 
**Source** | Pointer to **string** | Source is which receiver or operator the activation came through. | [optional] 
**Trigger** | Pointer to **string** | Trigger is what the panel reported. | [optional] 
**Zone** | Pointer to **string** | Zone is the panel zone that fired. | [optional] 

## Methods

### NewPatrolPatrolAlarmIn

`func NewPatrolPatrolAlarmIn() *PatrolPatrolAlarmIn`

NewPatrolPatrolAlarmIn instantiates a new PatrolPatrolAlarmIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolAlarmInWithDefaults

`func NewPatrolPatrolAlarmInWithDefaults() *PatrolPatrolAlarmIn`

NewPatrolPatrolAlarmInWithDefaults instantiates a new PatrolPatrolAlarmIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PatrolPatrolAlarmIn) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PatrolPatrolAlarmIn) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PatrolPatrolAlarmIn) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PatrolPatrolAlarmIn) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetRaw

`func (o *PatrolPatrolAlarmIn) GetRaw() interface{}`

GetRaw returns the Raw field if non-nil, zero value otherwise.

### GetRawOk

`func (o *PatrolPatrolAlarmIn) GetRawOk() (*interface{}, bool)`

GetRawOk returns a tuple with the Raw field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRaw

`func (o *PatrolPatrolAlarmIn) SetRaw(v interface{})`

SetRaw sets Raw field to given value.

### HasRaw

`func (o *PatrolPatrolAlarmIn) HasRaw() bool`

HasRaw returns a boolean if a field has been set.

### SetRawNil

`func (o *PatrolPatrolAlarmIn) SetRawNil(b bool)`

 SetRawNil sets the value for Raw to be an explicit nil

### UnsetRaw
`func (o *PatrolPatrolAlarmIn) UnsetRaw()`

UnsetRaw ensures that no value is present for Raw, not even an explicit nil
### GetRef

`func (o *PatrolPatrolAlarmIn) GetRef() string`

GetRef returns the Ref field if non-nil, zero value otherwise.

### GetRefOk

`func (o *PatrolPatrolAlarmIn) GetRefOk() (*string, bool)`

GetRefOk returns a tuple with the Ref field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRef

`func (o *PatrolPatrolAlarmIn) SetRef(v string)`

SetRef sets Ref field to given value.

### HasRef

`func (o *PatrolPatrolAlarmIn) HasRef() bool`

HasRef returns a boolean if a field has been set.

### GetSite

`func (o *PatrolPatrolAlarmIn) GetSite() string`

GetSite returns the Site field if non-nil, zero value otherwise.

### GetSiteOk

`func (o *PatrolPatrolAlarmIn) GetSiteOk() (*string, bool)`

GetSiteOk returns a tuple with the Site field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSite

`func (o *PatrolPatrolAlarmIn) SetSite(v string)`

SetSite sets Site field to given value.

### HasSite

`func (o *PatrolPatrolAlarmIn) HasSite() bool`

HasSite returns a boolean if a field has been set.

### GetSource

`func (o *PatrolPatrolAlarmIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PatrolPatrolAlarmIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PatrolPatrolAlarmIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PatrolPatrolAlarmIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTrigger

`func (o *PatrolPatrolAlarmIn) GetTrigger() string`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *PatrolPatrolAlarmIn) GetTriggerOk() (*string, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *PatrolPatrolAlarmIn) SetTrigger(v string)`

SetTrigger sets Trigger field to given value.

### HasTrigger

`func (o *PatrolPatrolAlarmIn) HasTrigger() bool`

HasTrigger returns a boolean if a field has been set.

### GetZone

`func (o *PatrolPatrolAlarmIn) GetZone() string`

GetZone returns the Zone field if non-nil, zero value otherwise.

### GetZoneOk

`func (o *PatrolPatrolAlarmIn) GetZoneOk() (*string, bool)`

GetZoneOk returns a tuple with the Zone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZone

`func (o *PatrolPatrolAlarmIn) SetZone(v string)`

SetZone sets Zone field to given value.

### HasZone

`func (o *PatrolPatrolAlarmIn) HasZone() bool`

HasZone returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


