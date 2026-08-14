# MoneyInfra

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DoAvgDailyBurnCents** | Pointer to **int32** |  | [optional] 
**DoCreditRemainingCents** | Pointer to **int32** |  | [optional] 
**DoMonthToDateCents** | Pointer to **int32** |  | [optional] 
**Period** | Pointer to **string** |  | [optional] 
**TreasuryReserveCents** | Pointer to **int32** |  | [optional] 
**VendorCogsCents** | Pointer to **int32** |  | [optional] 
**Vendors** | Pointer to [**[]Vendor**](Vendor.md) |  | [optional] 

## Methods

### NewMoneyInfra

`func NewMoneyInfra() *MoneyInfra`

NewMoneyInfra instantiates a new MoneyInfra object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoneyInfraWithDefaults

`func NewMoneyInfraWithDefaults() *MoneyInfra`

NewMoneyInfraWithDefaults instantiates a new MoneyInfra object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDoAvgDailyBurnCents

`func (o *MoneyInfra) GetDoAvgDailyBurnCents() int32`

GetDoAvgDailyBurnCents returns the DoAvgDailyBurnCents field if non-nil, zero value otherwise.

### GetDoAvgDailyBurnCentsOk

`func (o *MoneyInfra) GetDoAvgDailyBurnCentsOk() (*int32, bool)`

GetDoAvgDailyBurnCentsOk returns a tuple with the DoAvgDailyBurnCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoAvgDailyBurnCents

`func (o *MoneyInfra) SetDoAvgDailyBurnCents(v int32)`

SetDoAvgDailyBurnCents sets DoAvgDailyBurnCents field to given value.

### HasDoAvgDailyBurnCents

`func (o *MoneyInfra) HasDoAvgDailyBurnCents() bool`

HasDoAvgDailyBurnCents returns a boolean if a field has been set.

### GetDoCreditRemainingCents

`func (o *MoneyInfra) GetDoCreditRemainingCents() int32`

GetDoCreditRemainingCents returns the DoCreditRemainingCents field if non-nil, zero value otherwise.

### GetDoCreditRemainingCentsOk

`func (o *MoneyInfra) GetDoCreditRemainingCentsOk() (*int32, bool)`

GetDoCreditRemainingCentsOk returns a tuple with the DoCreditRemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoCreditRemainingCents

`func (o *MoneyInfra) SetDoCreditRemainingCents(v int32)`

SetDoCreditRemainingCents sets DoCreditRemainingCents field to given value.

### HasDoCreditRemainingCents

`func (o *MoneyInfra) HasDoCreditRemainingCents() bool`

HasDoCreditRemainingCents returns a boolean if a field has been set.

### GetDoMonthToDateCents

`func (o *MoneyInfra) GetDoMonthToDateCents() int32`

GetDoMonthToDateCents returns the DoMonthToDateCents field if non-nil, zero value otherwise.

### GetDoMonthToDateCentsOk

`func (o *MoneyInfra) GetDoMonthToDateCentsOk() (*int32, bool)`

GetDoMonthToDateCentsOk returns a tuple with the DoMonthToDateCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoMonthToDateCents

`func (o *MoneyInfra) SetDoMonthToDateCents(v int32)`

SetDoMonthToDateCents sets DoMonthToDateCents field to given value.

### HasDoMonthToDateCents

`func (o *MoneyInfra) HasDoMonthToDateCents() bool`

HasDoMonthToDateCents returns a boolean if a field has been set.

### GetPeriod

`func (o *MoneyInfra) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *MoneyInfra) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *MoneyInfra) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *MoneyInfra) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.

### GetTreasuryReserveCents

`func (o *MoneyInfra) GetTreasuryReserveCents() int32`

GetTreasuryReserveCents returns the TreasuryReserveCents field if non-nil, zero value otherwise.

### GetTreasuryReserveCentsOk

`func (o *MoneyInfra) GetTreasuryReserveCentsOk() (*int32, bool)`

GetTreasuryReserveCentsOk returns a tuple with the TreasuryReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasuryReserveCents

`func (o *MoneyInfra) SetTreasuryReserveCents(v int32)`

SetTreasuryReserveCents sets TreasuryReserveCents field to given value.

### HasTreasuryReserveCents

`func (o *MoneyInfra) HasTreasuryReserveCents() bool`

HasTreasuryReserveCents returns a boolean if a field has been set.

### GetVendorCogsCents

`func (o *MoneyInfra) GetVendorCogsCents() int32`

GetVendorCogsCents returns the VendorCogsCents field if non-nil, zero value otherwise.

### GetVendorCogsCentsOk

`func (o *MoneyInfra) GetVendorCogsCentsOk() (*int32, bool)`

GetVendorCogsCentsOk returns a tuple with the VendorCogsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendorCogsCents

`func (o *MoneyInfra) SetVendorCogsCents(v int32)`

SetVendorCogsCents sets VendorCogsCents field to given value.

### HasVendorCogsCents

`func (o *MoneyInfra) HasVendorCogsCents() bool`

HasVendorCogsCents returns a boolean if a field has been set.

### GetVendors

`func (o *MoneyInfra) GetVendors() []Vendor`

GetVendors returns the Vendors field if non-nil, zero value otherwise.

### GetVendorsOk

`func (o *MoneyInfra) GetVendorsOk() (*[]Vendor, bool)`

GetVendorsOk returns a tuple with the Vendors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendors

`func (o *MoneyInfra) SetVendors(v []Vendor)`

SetVendors sets Vendors field to given value.

### HasVendors

`func (o *MoneyInfra) HasVendors() bool`

HasVendors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


