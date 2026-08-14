# Accruals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accrued** | Pointer to **int32** | Accrued is how many NEW commission accruals this run created, counted across every upline level. The accrual is latched at most once per (affiliate, source org, period), so a re-run inside the same month reports 0 having changed nothing — 0 means \&quot;already accrued\&quot;, not \&quot;failed\&quot;. | [optional] 
**RoyaltiesAccrued** | Pointer to **int32** | RoyaltiesAccrued is how many OSS-author royalty accruals the SAME spend read produced in the sibling authors program. One read drives both. | [optional] 
**RoyaltyFailures** | Pointer to **int32** | RoyaltyFailures is reported, not swallowed: a sweep that could not reach the royalty store must not read as one that found nothing owed. The count was already computed and then dropped on the floor, which is the same silence the typed leg was added to end. | [optional] 
**Swept** | Pointer to **int32** | Swept is how many source (referred) orgs the run visited, bounded at 500 per run. A source with no spend this period, or one whose spend could not be read, still counts as swept. | [optional] 

## Methods

### NewAccruals

`func NewAccruals() *Accruals`

NewAccruals instantiates a new Accruals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccrualsWithDefaults

`func NewAccrualsWithDefaults() *Accruals`

NewAccrualsWithDefaults instantiates a new Accruals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrued

`func (o *Accruals) GetAccrued() int32`

GetAccrued returns the Accrued field if non-nil, zero value otherwise.

### GetAccruedOk

`func (o *Accruals) GetAccruedOk() (*int32, bool)`

GetAccruedOk returns a tuple with the Accrued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrued

`func (o *Accruals) SetAccrued(v int32)`

SetAccrued sets Accrued field to given value.

### HasAccrued

`func (o *Accruals) HasAccrued() bool`

HasAccrued returns a boolean if a field has been set.

### GetRoyaltiesAccrued

`func (o *Accruals) GetRoyaltiesAccrued() int32`

GetRoyaltiesAccrued returns the RoyaltiesAccrued field if non-nil, zero value otherwise.

### GetRoyaltiesAccruedOk

`func (o *Accruals) GetRoyaltiesAccruedOk() (*int32, bool)`

GetRoyaltiesAccruedOk returns a tuple with the RoyaltiesAccrued field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltiesAccrued

`func (o *Accruals) SetRoyaltiesAccrued(v int32)`

SetRoyaltiesAccrued sets RoyaltiesAccrued field to given value.

### HasRoyaltiesAccrued

`func (o *Accruals) HasRoyaltiesAccrued() bool`

HasRoyaltiesAccrued returns a boolean if a field has been set.

### GetRoyaltyFailures

`func (o *Accruals) GetRoyaltyFailures() int32`

GetRoyaltyFailures returns the RoyaltyFailures field if non-nil, zero value otherwise.

### GetRoyaltyFailuresOk

`func (o *Accruals) GetRoyaltyFailuresOk() (*int32, bool)`

GetRoyaltyFailuresOk returns a tuple with the RoyaltyFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoyaltyFailures

`func (o *Accruals) SetRoyaltyFailures(v int32)`

SetRoyaltyFailures sets RoyaltyFailures field to given value.

### HasRoyaltyFailures

`func (o *Accruals) HasRoyaltyFailures() bool`

HasRoyaltyFailures returns a boolean if a field has been set.

### GetSwept

`func (o *Accruals) GetSwept() int32`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *Accruals) GetSweptOk() (*int32, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *Accruals) SetSwept(v int32)`

SetSwept sets Swept field to given value.

### HasSwept

`func (o *Accruals) HasSwept() bool`

HasSwept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


