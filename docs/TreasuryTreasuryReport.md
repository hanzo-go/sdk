# TreasuryTreasuryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int64** | lifetime revenue-share into the fund | [optional] 
**ByProgramCents** | Pointer to **map[string]int64** | program → lifetime paid | [optional] 
**PaidCents** | Pointer to **int64** | lifetime backed payouts out of the fund | [optional] 
**Policy** | Pointer to [**TreasurySharePolicy**](TreasurySharePolicy.md) | current revenue-share policy | [optional] 
**ReserveCents** | Pointer to **int64** | fund:reserve balance (available now) | [optional] 
**SolventForPayout** | Pointer to **bool** | reserve &gt; 0: at least some payout is backable | [optional] 

## Methods

### NewTreasuryTreasuryReport

`func NewTreasuryTreasuryReport() *TreasuryTreasuryReport`

NewTreasuryTreasuryReport instantiates a new TreasuryTreasuryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreasuryTreasuryReportWithDefaults

`func NewTreasuryTreasuryReportWithDefaults() *TreasuryTreasuryReport`

NewTreasuryTreasuryReportWithDefaults instantiates a new TreasuryTreasuryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *TreasuryTreasuryReport) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *TreasuryTreasuryReport) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *TreasuryTreasuryReport) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *TreasuryTreasuryReport) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetByProgramCents

`func (o *TreasuryTreasuryReport) GetByProgramCents() map[string]int64`

GetByProgramCents returns the ByProgramCents field if non-nil, zero value otherwise.

### GetByProgramCentsOk

`func (o *TreasuryTreasuryReport) GetByProgramCentsOk() (*map[string]int64, bool)`

GetByProgramCentsOk returns a tuple with the ByProgramCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProgramCents

`func (o *TreasuryTreasuryReport) SetByProgramCents(v map[string]int64)`

SetByProgramCents sets ByProgramCents field to given value.

### HasByProgramCents

`func (o *TreasuryTreasuryReport) HasByProgramCents() bool`

HasByProgramCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *TreasuryTreasuryReport) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *TreasuryTreasuryReport) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *TreasuryTreasuryReport) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *TreasuryTreasuryReport) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPolicy

`func (o *TreasuryTreasuryReport) GetPolicy() TreasurySharePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *TreasuryTreasuryReport) GetPolicyOk() (*TreasurySharePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *TreasuryTreasuryReport) SetPolicy(v TreasurySharePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *TreasuryTreasuryReport) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetReserveCents

`func (o *TreasuryTreasuryReport) GetReserveCents() int64`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *TreasuryTreasuryReport) GetReserveCentsOk() (*int64, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *TreasuryTreasuryReport) SetReserveCents(v int64)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *TreasuryTreasuryReport) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.

### GetSolventForPayout

`func (o *TreasuryTreasuryReport) GetSolventForPayout() bool`

GetSolventForPayout returns the SolventForPayout field if non-nil, zero value otherwise.

### GetSolventForPayoutOk

`func (o *TreasuryTreasuryReport) GetSolventForPayoutOk() (*bool, bool)`

GetSolventForPayoutOk returns a tuple with the SolventForPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolventForPayout

`func (o *TreasuryTreasuryReport) SetSolventForPayout(v bool)`

SetSolventForPayout sets SolventForPayout field to given value.

### HasSolventForPayout

`func (o *TreasuryTreasuryReport) HasSolventForPayout() bool`

HasSolventForPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


