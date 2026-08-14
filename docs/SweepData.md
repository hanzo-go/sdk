# SweepData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is the amount moved into the reserve fund. | [optional] 
**Created** | Pointer to **bool** | Created is false when this period had already been swept — the accrual is idempotent. | [optional] 
**Period** | Pointer to **string** | Period is the period actually accrued. | [optional] 
**ReserveCents** | Pointer to **int32** | ReserveCents is the fund balance after the accrual. | [optional] 
**RevenueCents** | Pointer to **int32** | RevenueCents is the revenue the share was computed from. | [optional] 

## Methods

### NewSweepData

`func NewSweepData() *SweepData`

NewSweepData instantiates a new SweepData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSweepDataWithDefaults

`func NewSweepDataWithDefaults() *SweepData`

NewSweepDataWithDefaults instantiates a new SweepData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *SweepData) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *SweepData) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *SweepData) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *SweepData) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetCreated

`func (o *SweepData) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SweepData) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SweepData) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SweepData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPeriod

`func (o *SweepData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *SweepData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *SweepData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *SweepData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetReserveCents

`func (o *SweepData) GetReserveCents() int32`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *SweepData) GetReserveCentsOk() (*int32, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *SweepData) SetReserveCents(v int32)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *SweepData) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.

### GetRevenueCents

`func (o *SweepData) GetRevenueCents() int32`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *SweepData) GetRevenueCentsOk() (*int32, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *SweepData) SetRevenueCents(v int32)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *SweepData) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


