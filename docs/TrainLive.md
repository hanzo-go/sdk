# TrainLive

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Damage** | Pointer to **int64** |  | [optional] 
**DamageRate** | Pointer to **float64** |  | [optional] 
**Fires** | Pointer to **int64** | Fires counts decisions it answered otherwise than the base would have. | [optional] 
**LastFired** | Pointer to **int64** | LastFired is when it last fired, unix seconds; 0 never. | [optional] 
**Neutral** | Pointer to **int64** |  | [optional] 
**RepairRate** | Pointer to **float64** |  | [optional] 
**Repairs** | Pointer to **int64** | Repairs are fires that changed the label and no correction contradicts; Damage, fires whose label a correction put back to the base&#39;s; Neutral, fires that kept the base&#39;s label. | [optional] 
**Version** | Pointer to **int64** |  | [optional] 

## Methods

### NewTrainLive

`func NewTrainLive() *TrainLive`

NewTrainLive instantiates a new TrainLive object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrainLiveWithDefaults

`func NewTrainLiveWithDefaults() *TrainLive`

NewTrainLiveWithDefaults instantiates a new TrainLive object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDamage

`func (o *TrainLive) GetDamage() int64`

GetDamage returns the Damage field if non-nil, zero value otherwise.

### GetDamageOk

`func (o *TrainLive) GetDamageOk() (*int64, bool)`

GetDamageOk returns a tuple with the Damage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamage

`func (o *TrainLive) SetDamage(v int64)`

SetDamage sets Damage field to given value.

### HasDamage

`func (o *TrainLive) HasDamage() bool`

HasDamage returns a boolean if a field has been set.

### GetDamageRate

`func (o *TrainLive) GetDamageRate() float64`

GetDamageRate returns the DamageRate field if non-nil, zero value otherwise.

### GetDamageRateOk

`func (o *TrainLive) GetDamageRateOk() (*float64, bool)`

GetDamageRateOk returns a tuple with the DamageRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDamageRate

`func (o *TrainLive) SetDamageRate(v float64)`

SetDamageRate sets DamageRate field to given value.

### HasDamageRate

`func (o *TrainLive) HasDamageRate() bool`

HasDamageRate returns a boolean if a field has been set.

### GetFires

`func (o *TrainLive) GetFires() int64`

GetFires returns the Fires field if non-nil, zero value otherwise.

### GetFiresOk

`func (o *TrainLive) GetFiresOk() (*int64, bool)`

GetFiresOk returns a tuple with the Fires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFires

`func (o *TrainLive) SetFires(v int64)`

SetFires sets Fires field to given value.

### HasFires

`func (o *TrainLive) HasFires() bool`

HasFires returns a boolean if a field has been set.

### GetLastFired

`func (o *TrainLive) GetLastFired() int64`

GetLastFired returns the LastFired field if non-nil, zero value otherwise.

### GetLastFiredOk

`func (o *TrainLive) GetLastFiredOk() (*int64, bool)`

GetLastFiredOk returns a tuple with the LastFired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFired

`func (o *TrainLive) SetLastFired(v int64)`

SetLastFired sets LastFired field to given value.

### HasLastFired

`func (o *TrainLive) HasLastFired() bool`

HasLastFired returns a boolean if a field has been set.

### GetNeutral

`func (o *TrainLive) GetNeutral() int64`

GetNeutral returns the Neutral field if non-nil, zero value otherwise.

### GetNeutralOk

`func (o *TrainLive) GetNeutralOk() (*int64, bool)`

GetNeutralOk returns a tuple with the Neutral field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeutral

`func (o *TrainLive) SetNeutral(v int64)`

SetNeutral sets Neutral field to given value.

### HasNeutral

`func (o *TrainLive) HasNeutral() bool`

HasNeutral returns a boolean if a field has been set.

### GetRepairRate

`func (o *TrainLive) GetRepairRate() float64`

GetRepairRate returns the RepairRate field if non-nil, zero value otherwise.

### GetRepairRateOk

`func (o *TrainLive) GetRepairRateOk() (*float64, bool)`

GetRepairRateOk returns a tuple with the RepairRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairRate

`func (o *TrainLive) SetRepairRate(v float64)`

SetRepairRate sets RepairRate field to given value.

### HasRepairRate

`func (o *TrainLive) HasRepairRate() bool`

HasRepairRate returns a boolean if a field has been set.

### GetRepairs

`func (o *TrainLive) GetRepairs() int64`

GetRepairs returns the Repairs field if non-nil, zero value otherwise.

### GetRepairsOk

`func (o *TrainLive) GetRepairsOk() (*int64, bool)`

GetRepairsOk returns a tuple with the Repairs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepairs

`func (o *TrainLive) SetRepairs(v int64)`

SetRepairs sets Repairs field to given value.

### HasRepairs

`func (o *TrainLive) HasRepairs() bool`

HasRepairs returns a boolean if a field has been set.

### GetVersion

`func (o *TrainLive) GetVersion() int64`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *TrainLive) GetVersionOk() (*int64, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *TrainLive) SetVersion(v int64)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *TrainLive) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


