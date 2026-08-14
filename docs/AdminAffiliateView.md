# AdminAffiliateView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is lifetime commission accrued, in cents. It only grows — a payout moves paidCents, never this. | [optional] 
**ApprovedAt** | Pointer to **int32** | ApprovedAt is when staff approved, Unix seconds UTC. 0 means never approved. | [optional] 
**Code** | Pointer to **string** | Code is the minted referral code, the slug the ?aff link carries. Empty until approval mints it. Codes are one global namespace across all affiliates. | [optional] 
**CreatedAt** | Pointer to **int32** | CreatedAt is when the org applied, Unix seconds UTC. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate&#39;s server-minted handle, \&quot;aff_\&quot;-prefixed — the id the approve, suspend, rate and payout routes address. | [optional] 
**Org** | Pointer to **string** | Org is the partner&#39;s own org slug. It appears ONLY on this cross-tenant admin view; no partner-facing read ever names another org. | [optional] 
**PaidCents** | Pointer to **int32** | PaidCents is lifetime commission already paid out, in cents — credits grants and record-only cash disbursements alike. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is accrued minus paid, in cents: what is still owed, and the hard ceiling the next payout is reserved against. Never negative. | [optional] 
**RateBps** | Pointer to **int32** | RateBps is this affiliate&#39;s DIRECT (level 1) commission rate in basis points OF Hanzo&#39;s margin (2000 &#x3D; 20% of margin, never of the customer&#39;s bill). Levels 2 and 3 are platform-wide switches and are not carried per affiliate. | [optional] 
**ReferredCount** | Pointer to **int32** | ReferredCount is how many orgs this affiliate is the DIRECT referrer of, counted from the attribution edges. It is 0 on the single-affiliate answers (approve, suspend, rate, payout), which do not run the count. | [optional] 
**RequestedCode** | Pointer to **string** | RequestedCode is the vanity code the applicant asked for. A request, not an allocation: approval mints a different slug if this one was taken. Absent when none was asked for. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;applied\&quot;, \&quot;approved\&quot; or \&quot;suspended\&quot;. Only \&quot;approved\&quot; resolves for attribution and accrues; \&quot;suspended\&quot; stops future earning and claws nothing back. | [optional] 
**SuspendedAt** | Pointer to **int32** | SuspendedAt is when staff suspended, Unix seconds UTC. 0 means never suspended; it is not cleared by a later re-approval. | [optional] 

## Methods

### NewAdminAffiliateView

`func NewAdminAffiliateView() *AdminAffiliateView`

NewAdminAffiliateView instantiates a new AdminAffiliateView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminAffiliateViewWithDefaults

`func NewAdminAffiliateViewWithDefaults() *AdminAffiliateView`

NewAdminAffiliateViewWithDefaults instantiates a new AdminAffiliateView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AdminAffiliateView) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AdminAffiliateView) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AdminAffiliateView) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AdminAffiliateView) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetApprovedAt

`func (o *AdminAffiliateView) GetApprovedAt() int32`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *AdminAffiliateView) GetApprovedAtOk() (*int32, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *AdminAffiliateView) SetApprovedAt(v int32)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *AdminAffiliateView) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### GetCode

`func (o *AdminAffiliateView) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AdminAffiliateView) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AdminAffiliateView) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AdminAffiliateView) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *AdminAffiliateView) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *AdminAffiliateView) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *AdminAffiliateView) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *AdminAffiliateView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetId

`func (o *AdminAffiliateView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AdminAffiliateView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AdminAffiliateView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AdminAffiliateView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetOrg

`func (o *AdminAffiliateView) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *AdminAffiliateView) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *AdminAffiliateView) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *AdminAffiliateView) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPaidCents

`func (o *AdminAffiliateView) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AdminAffiliateView) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AdminAffiliateView) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AdminAffiliateView) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPendingCents

`func (o *AdminAffiliateView) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AdminAffiliateView) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AdminAffiliateView) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AdminAffiliateView) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetRateBps

`func (o *AdminAffiliateView) GetRateBps() int32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AdminAffiliateView) GetRateBpsOk() (*int32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AdminAffiliateView) SetRateBps(v int32)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AdminAffiliateView) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetReferredCount

`func (o *AdminAffiliateView) GetReferredCount() int32`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AdminAffiliateView) GetReferredCountOk() (*int32, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AdminAffiliateView) SetReferredCount(v int32)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AdminAffiliateView) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AdminAffiliateView) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AdminAffiliateView) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AdminAffiliateView) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AdminAffiliateView) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetStatus

`func (o *AdminAffiliateView) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AdminAffiliateView) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AdminAffiliateView) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AdminAffiliateView) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSuspendedAt

`func (o *AdminAffiliateView) GetSuspendedAt() int32`

GetSuspendedAt returns the SuspendedAt field if non-nil, zero value otherwise.

### GetSuspendedAtOk

`func (o *AdminAffiliateView) GetSuspendedAtOk() (*int32, bool)`

GetSuspendedAtOk returns a tuple with the SuspendedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuspendedAt

`func (o *AdminAffiliateView) SetSuspendedAt(v int32)`

SetSuspendedAt sets SuspendedAt field to given value.

### HasSuspendedAt

`func (o *AdminAffiliateView) HasSuspendedAt() bool`

HasSuspendedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


