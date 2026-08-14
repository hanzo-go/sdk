# Totals

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is lifetime commission accrued summed over those rows, in cents. | [optional] 
**Applied** | Pointer to **int32** | Applied is how many of those rows are still awaiting approval — no code, no accrual yet. | [optional] 
**Approved** | Pointer to **int32** | Approved is how many are approved: the only rows whose code resolves for attribution and whose balance can still grow. | [optional] 
**PaidCents** | Pointer to **int32** | PaidCents is lifetime commission already paid out summed over those rows, in cents. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is accrued minus paid summed over those rows, in cents — the outstanding liability across the page. | [optional] 
**Suspended** | Pointer to **int32** | Suspended is how many were suspended. What they already accrued stays accrued and stays payable. | [optional] 
**Total** | Pointer to **int32** | Total is how many affiliate rows this page covered, at every status. It is the page, not the table: a limit that truncates truncates this too. | [optional] 

## Methods

### NewTotals

`func NewTotals() *Totals`

NewTotals instantiates a new Totals object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTotalsWithDefaults

`func NewTotalsWithDefaults() *Totals`

NewTotalsWithDefaults instantiates a new Totals object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *Totals) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *Totals) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *Totals) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *Totals) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetApplied

`func (o *Totals) GetApplied() int32`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *Totals) GetAppliedOk() (*int32, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *Totals) SetApplied(v int32)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *Totals) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### GetApproved

`func (o *Totals) GetApproved() int32`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *Totals) GetApprovedOk() (*int32, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *Totals) SetApproved(v int32)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *Totals) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetPaidCents

`func (o *Totals) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *Totals) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *Totals) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *Totals) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *Totals) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *Totals) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *Totals) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *Totals) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetSuspended

`func (o *Totals) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *Totals) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *Totals) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *Totals) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTotal

`func (o *Totals) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Totals) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Totals) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Totals) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


