# ReferrerRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is lifetime commission accrued, in cents. The board is sorted by this, descending. | [optional] 
**Code** | Pointer to **string** | Code is that affiliate&#39;s minted referral code; empty if it is not approved. | [optional] 
**Org** | Pointer to **string** | Org is the partner&#39;s own org slug. Named only here, on the SuperAdmin board — the partner-facing leaderboard shows an opt-in handle and never an org. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is accrued minus paid, in cents — what is still owed to this affiliate. Never negative. | [optional] 
**ReferredCount** | Pointer to **int32** | ReferredCount is how many orgs this affiliate is the DIRECT referrer of — its level-1 downline, not the whole three-level chain. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;applied\&quot;, \&quot;approved\&quot; or \&quot;suspended\&quot;. | [optional] 

## Methods

### NewReferrerRow

`func NewReferrerRow() *ReferrerRow`

NewReferrerRow instantiates a new ReferrerRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReferrerRowWithDefaults

`func NewReferrerRowWithDefaults() *ReferrerRow`

NewReferrerRowWithDefaults instantiates a new ReferrerRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *ReferrerRow) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *ReferrerRow) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *ReferrerRow) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *ReferrerRow) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetCode

`func (o *ReferrerRow) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ReferrerRow) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ReferrerRow) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ReferrerRow) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetOrg

`func (o *ReferrerRow) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *ReferrerRow) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *ReferrerRow) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *ReferrerRow) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPendingCents

`func (o *ReferrerRow) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *ReferrerRow) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *ReferrerRow) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *ReferrerRow) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetReferredCount

`func (o *ReferrerRow) GetReferredCount() int32`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *ReferrerRow) GetReferredCountOk() (*int32, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *ReferrerRow) SetReferredCount(v int32)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *ReferrerRow) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetStatus

`func (o *ReferrerRow) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ReferrerRow) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ReferrerRow) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ReferrerRow) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


