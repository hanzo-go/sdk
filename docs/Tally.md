# Tally

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedLifetimeCents** | Pointer to **int32** | AccruedLifetimeCents is all commission ever accrued, summed across every affiliate, in cents. It only grows; a payout does not reduce it. | [optional] 
**Affiliates** | Pointer to **int32** | Affiliates is how many affiliate rows the board read, at every status. The read is bounded at 1000 rows, so a larger fleet reports the bound. | [optional] 
**Approved** | Pointer to **int32** | Approved is how many of those rows are approved — the only ones whose code resolves for attribution and whose balance can grow. | [optional] 
**PaidLifetimeCents** | Pointer to **int32** | PaidLifetimeCents is all commission ever paid out, in cents: credits grants plus record-only cash disbursements. | [optional] 
**PendingLiabilityCents** | Pointer to **int32** | PendingLiabilityCents is accrued minus paid across every affiliate, in cents. Read it as money OWED and not yet disbursed — a liability, not spend. | [optional] 

## Methods

### NewTally

`func NewTally() *Tally`

NewTally instantiates a new Tally object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTallyWithDefaults

`func NewTallyWithDefaults() *Tally`

NewTallyWithDefaults instantiates a new Tally object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedLifetimeCents

`func (o *Tally) GetAccruedLifetimeCents() int32`

GetAccruedLifetimeCents returns the AccruedLifetimeCents field if non-nil, zero value otherwise.

### GetAccruedLifetimeCentsOk

`func (o *Tally) GetAccruedLifetimeCentsOk() (*int32, bool)`

GetAccruedLifetimeCentsOk returns a tuple with the AccruedLifetimeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedLifetimeCents

`func (o *Tally) SetAccruedLifetimeCents(v int32)`

SetAccruedLifetimeCents sets AccruedLifetimeCents field to given value.

### HasAccruedLifetimeCents

`func (o *Tally) HasAccruedLifetimeCents() bool`

HasAccruedLifetimeCents returns a boolean if a field has been set.

### GetAffiliates

`func (o *Tally) GetAffiliates() int32`

GetAffiliates returns the Affiliates field if non-nil, zero value otherwise.

### GetAffiliatesOk

`func (o *Tally) GetAffiliatesOk() (*int32, bool)`

GetAffiliatesOk returns a tuple with the Affiliates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffiliates

`func (o *Tally) SetAffiliates(v int32)`

SetAffiliates sets Affiliates field to given value.

### HasAffiliates

`func (o *Tally) HasAffiliates() bool`

HasAffiliates returns a boolean if a field has been set.

### GetApproved

`func (o *Tally) GetApproved() int32`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Tally) GetApprovedOk() (*int32, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Tally) SetApproved(v int32)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *Tally) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetPaidLifetimeCents

`func (o *Tally) GetPaidLifetimeCents() int32`

GetPaidLifetimeCents returns the PaidLifetimeCents field if non-nil, zero value otherwise.

### GetPaidLifetimeCentsOk

`func (o *Tally) GetPaidLifetimeCentsOk() (*int32, bool)`

GetPaidLifetimeCentsOk returns a tuple with the PaidLifetimeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidLifetimeCents

`func (o *Tally) SetPaidLifetimeCents(v int32)`

SetPaidLifetimeCents sets PaidLifetimeCents field to given value.

### HasPaidLifetimeCents

`func (o *Tally) HasPaidLifetimeCents() bool`

HasPaidLifetimeCents returns a boolean if a field has been set.

### GetPendingLiabilityCents

`func (o *Tally) GetPendingLiabilityCents() int32`

GetPendingLiabilityCents returns the PendingLiabilityCents field if non-nil, zero value otherwise.

### GetPendingLiabilityCentsOk

`func (o *Tally) GetPendingLiabilityCentsOk() (*int32, bool)`

GetPendingLiabilityCentsOk returns a tuple with the PendingLiabilityCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingLiabilityCents

`func (o *Tally) SetPendingLiabilityCents(v int32)`

SetPendingLiabilityCents sets PendingLiabilityCents field to given value.

### HasPendingLiabilityCents

`func (o *Tally) HasPendingLiabilityCents() bool`

HasPendingLiabilityCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


