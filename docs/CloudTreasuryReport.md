# CloudTreasuryReport

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | lifetime revenue-share into the fund | [optional] 
**ByProgramCents** | Pointer to **map[string]int32** | program → lifetime paid | [optional] 
**PaidCents** | Pointer to **int32** | lifetime backed payouts out of the fund | [optional] 
**Policy** | Pointer to [**CloudSharePolicy**](CloudSharePolicy.md) | current revenue-share policy | [optional] 
**ReserveCents** | Pointer to **int32** | fund:reserve balance (available now) | [optional] 
**SolventForPayout** | Pointer to **bool** | reserve &gt; 0: at least some payout is backable | [optional] 

## Methods

### NewCloudTreasuryReport

`func NewCloudTreasuryReport() *CloudTreasuryReport`

NewCloudTreasuryReport instantiates a new CloudTreasuryReport object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudTreasuryReportWithDefaults

`func NewCloudTreasuryReportWithDefaults() *CloudTreasuryReport`

NewCloudTreasuryReportWithDefaults instantiates a new CloudTreasuryReport object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *CloudTreasuryReport) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *CloudTreasuryReport) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *CloudTreasuryReport) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *CloudTreasuryReport) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetByProgramCents

`func (o *CloudTreasuryReport) GetByProgramCents() map[string]int32`

GetByProgramCents returns the ByProgramCents field if non-nil, zero value otherwise.

### GetByProgramCentsOk

`func (o *CloudTreasuryReport) GetByProgramCentsOk() (*map[string]int32, bool)`

GetByProgramCentsOk returns a tuple with the ByProgramCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetByProgramCents

`func (o *CloudTreasuryReport) SetByProgramCents(v map[string]int32)`

SetByProgramCents sets ByProgramCents field to given value.

### HasByProgramCents

`func (o *CloudTreasuryReport) HasByProgramCents() bool`

HasByProgramCents returns a boolean if a field has been set.

### GetPaidCents

`func (o *CloudTreasuryReport) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *CloudTreasuryReport) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *CloudTreasuryReport) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *CloudTreasuryReport) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPolicy

`func (o *CloudTreasuryReport) GetPolicy() CloudSharePolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CloudTreasuryReport) GetPolicyOk() (*CloudSharePolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CloudTreasuryReport) SetPolicy(v CloudSharePolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CloudTreasuryReport) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetReserveCents

`func (o *CloudTreasuryReport) GetReserveCents() int32`

GetReserveCents returns the ReserveCents field if non-nil, zero value otherwise.

### GetReserveCentsOk

`func (o *CloudTreasuryReport) GetReserveCentsOk() (*int32, bool)`

GetReserveCentsOk returns a tuple with the ReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveCents

`func (o *CloudTreasuryReport) SetReserveCents(v int32)`

SetReserveCents sets ReserveCents field to given value.

### HasReserveCents

`func (o *CloudTreasuryReport) HasReserveCents() bool`

HasReserveCents returns a boolean if a field has been set.

### GetSolventForPayout

`func (o *CloudTreasuryReport) GetSolventForPayout() bool`

GetSolventForPayout returns the SolventForPayout field if non-nil, zero value otherwise.

### GetSolventForPayoutOk

`func (o *CloudTreasuryReport) GetSolventForPayoutOk() (*bool, bool)`

GetSolventForPayoutOk returns a tuple with the SolventForPayout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolventForPayout

`func (o *CloudTreasuryReport) SetSolventForPayout(v bool)`

SetSolventForPayout sets SolventForPayout field to given value.

### HasSolventForPayout

`func (o *CloudTreasuryReport) HasSolventForPayout() bool`

HasSolventForPayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


