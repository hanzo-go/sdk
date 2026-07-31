# CloudSweepData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is the amount moved into the reserve fund. | [optional] 
**Created** | Pointer to **bool** | Created is false when this period had already been swept — the accrual is idempotent. | [optional] 
**Period** | Pointer to **string** | Period is the period actually accrued. | [optional] 
**ReserveCents** | Pointer to **int32** | ReserveCents is the fund balance after the accrual. | [optional] 
**RevenueCents** | Pointer to **int32** | RevenueCents is the revenue the share was computed from. | [optional] 

## Methods

### NewCloudSweepData

`func NewCloudSweepData() *CloudSweepData`

NewCloudSweepData instantiates a new CloudSweepData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudSweepDataWithDefaults

`func NewCloudSweepDataWithDefaults() *CloudSweepData`

NewCloudSweepDataWithDefaults instantiates a new CloudSweepData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *CloudSweepData) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *CloudSweepData) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *CloudSweepData) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *CloudSweepData) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetCreated

`func (o *CloudSweepData) GetCreated() bool`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CloudSweepData) GetCreatedOk() (*bool, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CloudSweepData) SetCreated(v bool)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *CloudSweepData) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetPeriod

`func (o *CloudSweepData) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *CloudSweepData) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *CloudSweepData) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *CloudSweepData) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetReserveCents

`func (o *CloudSweepData) GetReserveCents() int32`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *CloudSweepData) GetReserveCentsOk() (*int32, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *CloudSweepData) SetReserveCents(v int32)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *CloudSweepData) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.

### GetRevenueCents

`func (o *CloudSweepData) GetRevenueCents() int32`

GetRevenueCents returns the RevenueCents field if non-nil, zero value otherwise.

### GetRevenueCentsOk

`func (o *CloudSweepData) GetRevenueCentsOk() (*int32, bool)`

GetRevenueCentsOk returns a tuple with the RevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueCents

`func (o *CloudSweepData) SetRevenueCents(v int32)`

SetRevenueCents sets RevenueCents field to given value.

### HasRevenueCents

`func (o *CloudSweepData) HasRevenueCents() bool`

HasRevenueCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


