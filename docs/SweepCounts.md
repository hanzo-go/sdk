# SweepCounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accrued** | Pointer to **int32** | Accrued is how many new royalty accruals it latched. | [optional] 
**Swept** | Pointer to **int32** | Swept is how many (author, deploying org) pairs the sweep examined. | [optional] 

## Methods

### NewSweepCounts

`func NewSweepCounts() *SweepCounts`

NewSweepCounts instantiates a new SweepCounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepCountsWithDefaults

`func NewSweepCountsWithDefaults() *SweepCounts`

NewSweepCountsWithDefaults instantiates a new SweepCounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrued

`func (o *SweepCounts) GetAccrued() int32`

GetAccrued returns the Accrued field if non-nil, zero value otherwise.

### GetAccruedOk

`func (o *SweepCounts) GetAccruedOk() (*int32, bool)`

GetAccruedOk returns a tuple with the Accrued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrued

`func (o *SweepCounts) SetAccrued(v int32)`

SetAccrued sets Accrued field to given value.

### HasAccrued

`func (o *SweepCounts) HasAccrued() bool`

HasAccrued returns a boolean if a field has been set.

### GetSwept

`func (o *SweepCounts) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *SweepCounts) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *SweepCounts) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *SweepCounts) HasSwept() bool`

HasSwept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


