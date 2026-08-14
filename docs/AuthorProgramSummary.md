# AuthorProgramSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is the page&#39;s lifetime royalty accrued, in integer USD cents. | [optional] 
**Approved** | Pointer to **int32** | Approved is how many are admitted and accruing. | [optional] 
**Connected** | Pointer to **int32** | Connected is how many of those are enrolled but not yet admitted to earning. | [optional] 
**PaidCents** | Pointer to **int32** | PaidCents is what has been RECORDED as paid across the page, in integer USD cents. Recorded, not settled: the money leaves in a human&#39;s hands. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is what the platform still owes across the page, in integer USD cents — the sum of each author&#39;s own accrued − paid, each floored at zero. | [optional] 
**Suspended** | Pointer to **int32** | Suspended is how many have been stopped from accruing further. An author holds exactly one status, so the three buckets never overlap and connected + approved + suspended &#x3D; total. | [optional] 
**Total** | Pointer to **int32** | Total is how many author records this response actually carried. The roll-up is folded over the SAME page as authors — newest first, bounded by limit (default 500, ceiling 1000) — so on a program larger than the page it summarizes that page, not the fleet. | [optional] 

## Methods

### NewAuthorProgramSummary

`func NewAuthorProgramSummary() *AuthorProgramSummary`

NewAuthorProgramSummary instantiates a new AuthorProgramSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthorProgramSummaryWithDefaults

`func NewAuthorProgramSummaryWithDefaults() *AuthorProgramSummary`

NewAuthorProgramSummaryWithDefaults instantiates a new AuthorProgramSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AuthorProgramSummary) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AuthorProgramSummary) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AuthorProgramSummary) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AuthorProgramSummary) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetApproved

`func (o *AuthorProgramSummary) GetApproved() int32`

GetApproved returns the Approved field if non-nil, zero value otherwise.

### GetApprovedOk

`func (o *AuthorProgramSummary) GetApprovedOk() (*int32, bool)`

GetApprovedOk returns a tuple with the Approved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApproved

`func (o *AuthorProgramSummary) SetApproved(v int32)`

SetApproved sets Approved field to given value.

### HasApproved

`func (o *AuthorProgramSummary) HasApproved() bool`

HasApproved returns a boolean if a field has been set.

### GetConnected

`func (o *AuthorProgramSummary) GetConnected() int32`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *AuthorProgramSummary) GetConnectedOk() (*int32, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *AuthorProgramSummary) SetConnected(v int32)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *AuthorProgramSummary) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetPaidCents

`func (o *AuthorProgramSummary) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AuthorProgramSummary) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AuthorProgramSummary) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AuthorProgramSummary) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AuthorProgramSummary) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AuthorProgramSummary) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AuthorProgramSummary) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AuthorProgramSummary) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetSuspended

`func (o *AuthorProgramSummary) GetSuspended() int32`

GetSuspended returns the Suspended field if non-nil, zero value otherwise.

### GetSuspendedOk

`func (o *AuthorProgramSummary) GetSuspendedOk() (*int32, bool)`

GetSuspendedOk returns a tuple with the Suspended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspended

`func (o *AuthorProgramSummary) SetSuspended(v int32)`

SetSuspended sets Suspended field to given value.

### HasSuspended

`func (o *AuthorProgramSummary) HasSuspended() bool`

HasSuspended returns a boolean if a field has been set.

### GetTotal

`func (o *AuthorProgramSummary) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AuthorProgramSummary) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AuthorProgramSummary) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AuthorProgramSummary) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


