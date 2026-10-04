# PatrolPatrolClock

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Elapsed** | Pointer to **int64** | Elapsed is the seconds run so far, or the seconds it took when settled. | [optional] 
**Label** | Pointer to **string** | Label names the clock: Verify, Dispatch or On site. | [optional] 
**Settled** | Pointer to **int64** | Settled is the seconds it took, or -1 while the clock is still running. | [optional] 
**Target** | Pointer to **int64** | Target is the contracted seconds. | [optional] 

## Methods

### NewPatrolPatrolClock

`func NewPatrolPatrolClock() *PatrolPatrolClock`

NewPatrolPatrolClock instantiates a new PatrolPatrolClock object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPatrolPatrolClockWithDefaults

`func NewPatrolPatrolClockWithDefaults() *PatrolPatrolClock`

NewPatrolPatrolClockWithDefaults instantiates a new PatrolPatrolClock object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetElapsed

`func (o *PatrolPatrolClock) GetElapsed() int64`

GetElapsed returns the Elapsed field if non-nil, zero value otherwise.

### GetElapsedOk

`func (o *PatrolPatrolClock) GetElapsedOk() (*int64, bool)`

GetElapsedOk returns a tuple with the Elapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElapsed

`func (o *PatrolPatrolClock) SetElapsed(v int64)`

SetElapsed sets Elapsed field to given value.

### HasElapsed

`func (o *PatrolPatrolClock) HasElapsed() bool`

HasElapsed returns a boolean if a field has been set.

### GetLabel

`func (o *PatrolPatrolClock) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *PatrolPatrolClock) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *PatrolPatrolClock) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *PatrolPatrolClock) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetSettled

`func (o *PatrolPatrolClock) GetSettled() int64`

GetSettled returns the Settled field if non-nil, zero value otherwise.

### GetSettledOk

`func (o *PatrolPatrolClock) GetSettledOk() (*int64, bool)`

GetSettledOk returns a tuple with the Settled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettled

`func (o *PatrolPatrolClock) SetSettled(v int64)`

SetSettled sets Settled field to given value.

### HasSettled

`func (o *PatrolPatrolClock) HasSettled() bool`

HasSettled returns a boolean if a field has been set.

### GetTarget

`func (o *PatrolPatrolClock) GetTarget() int64`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *PatrolPatrolClock) GetTargetOk() (*int64, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *PatrolPatrolClock) SetTarget(v int64)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *PatrolPatrolClock) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


