# AffiliateAffiliateStanding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int64** | AccruedCents is lifetime commission accrued, in cents. | [optional] 
**Code** | Pointer to **string** | Code is the minted referral code; empty until staff approve. | [optional] 
**DefaultRateBps** | Pointer to **int64** | DefaultRateBps is the direct rate a new affiliate would get, answered only to a caller that has not applied. | [optional] 
**Handle** | Pointer to **string** | Handle is the opt-in public leaderboard name; empty means opted out. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate&#39;s server-minted handle, \&quot;aff_\&quot;-prefixed — what staff approve, suspend, re-rate and pay against. Absent until the org applies. | [optional] 
**IsAffiliate** | Pointer to **bool** | IsAffiliate says whether the caller org has an affiliate record at all. It is the ONE field an org that never applied gets besides defaultRateBps: on false, read nothing else here — every other field is absent, not zero. | [optional] 
**Link** | Pointer to **string** | Link is the shareable ?aff URL; empty until a code is minted. | [optional] 
**MarginBps** | Pointer to **int64** | MarginBps is the platform gross-margin fraction commission is a rate OF. | [optional] 
**PaidCents** | Pointer to **int64** | PaidCents is lifetime commission already paid out, in cents. | [optional] 
**Payouts** | Pointer to [**[]AffiliateRemittance**](AffiliateRemittance.md) | Payouts is the payout history, newest rows bounded. | [optional] 
**PendingCents** | Pointer to **int64** | PendingCents is accrued minus paid — what the platform still owes. | [optional] 
**RateBps** | Pointer to **int64** | RateBps is the affiliate&#39;s own direct commission rate, in basis points. | [optional] 
**ReferredCount** | Pointer to **int64** | ReferredCount is how many orgs this affiliate has referred. | [optional] 
**RequestedCode** | Pointer to **string** | RequestedCode is the vanity code asked for at apply time — a request, not an allocation. Approval mints &#x60;code&#x60;, which may be a different slug if this one was already taken. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;applied\&quot;, \&quot;approved\&quot; or \&quot;suspended\&quot;. Only an approved affiliate has a code that resolves for attribution and accrues commission; suspended keeps what it already earned but stops earning more. | [optional] 

## Methods

### NewAffiliateAffiliateStanding

`func NewAffiliateAffiliateStanding() *AffiliateAffiliateStanding`

NewAffiliateAffiliateStanding instantiates a new AffiliateAffiliateStanding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateAffiliateStandingWithDefaults

`func NewAffiliateAffiliateStandingWithDefaults() *AffiliateAffiliateStanding`

NewAffiliateAffiliateStandingWithDefaults instantiates a new AffiliateAffiliateStanding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AffiliateAffiliateStanding) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliateAffiliateStanding) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliateAffiliateStanding) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliateAffiliateStanding) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetCode

`func (o *AffiliateAffiliateStanding) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateAffiliateStanding) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateAffiliateStanding) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateAffiliateStanding) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefaultRateBps

`func (o *AffiliateAffiliateStanding) GetDefaultRateBps() int64`

GetDefaultRateBps returns the DefaultRateBps field if non-nil, zero value otherwise.

### GetDefaultRateBpsOk

`func (o *AffiliateAffiliateStanding) GetDefaultRateBpsOk() (*int64, bool)`

GetDefaultRateBpsOk returns a tuple with the DefaultRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRateBps

`func (o *AffiliateAffiliateStanding) SetDefaultRateBps(v int64)`

SetDefaultRateBps sets DefaultRateBps field to given value.

### HasDefaultRateBps

`func (o *AffiliateAffiliateStanding) HasDefaultRateBps() bool`

HasDefaultRateBps returns a boolean if a field has been set.

### GetHandle

`func (o *AffiliateAffiliateStanding) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AffiliateAffiliateStanding) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AffiliateAffiliateStanding) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AffiliateAffiliateStanding) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *AffiliateAffiliateStanding) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliateAffiliateStanding) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliateAffiliateStanding) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliateAffiliateStanding) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAffiliate

`func (o *AffiliateAffiliateStanding) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliateAffiliateStanding) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliateAffiliateStanding) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliateAffiliateStanding) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetLink

`func (o *AffiliateAffiliateStanding) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AffiliateAffiliateStanding) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AffiliateAffiliateStanding) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AffiliateAffiliateStanding) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMarginBps

`func (o *AffiliateAffiliateStanding) GetMarginBps() int64`

GetMarginBps returns the MarginBps field if non-nil, zero value otherwise.

### GetMarginBpsOk

`func (o *AffiliateAffiliateStanding) GetMarginBpsOk() (*int64, bool)`

GetMarginBpsOk returns a tuple with the MarginBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBps

`func (o *AffiliateAffiliateStanding) SetMarginBps(v int64)`

SetMarginBps sets MarginBps field to given value.

### HasMarginBps

`func (o *AffiliateAffiliateStanding) HasMarginBps() bool`

HasMarginBps returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliateAffiliateStanding) GetPaidCents() int64`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliateAffiliateStanding) GetPaidCentsOk() (*int64, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliateAffiliateStanding) SetPaidCents(v int64)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliateAffiliateStanding) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AffiliateAffiliateStanding) GetPayouts() []AffiliateRemittance`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AffiliateAffiliateStanding) GetPayoutsOk() (*[]AffiliateRemittance, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AffiliateAffiliateStanding) SetPayouts(v []AffiliateRemittance)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AffiliateAffiliateStanding) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliateAffiliateStanding) GetPendingCents() int64`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliateAffiliateStanding) GetPendingCentsOk() (*int64, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliateAffiliateStanding) SetPendingCents(v int64)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliateAffiliateStanding) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliateAffiliateStanding) GetRateBps() int64`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliateAffiliateStanding) GetRateBpsOk() (*int64, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliateAffiliateStanding) SetRateBps(v int64)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliateAffiliateStanding) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetReferredCount

`func (o *AffiliateAffiliateStanding) GetReferredCount() int64`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AffiliateAffiliateStanding) GetReferredCountOk() (*int64, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AffiliateAffiliateStanding) SetReferredCount(v int64)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AffiliateAffiliateStanding) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.

### GetRequestedCode

`func (o *AffiliateAffiliateStanding) GetRequestedCode() string`

GetRequestedCode returns the RequestedCode field if non-nil, zero value otherwise.

### GetRequestedCodeOk

`func (o *AffiliateAffiliateStanding) GetRequestedCodeOk() (*string, bool)`

GetRequestedCodeOk returns a tuple with the RequestedCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCode

`func (o *AffiliateAffiliateStanding) SetRequestedCode(v string)`

SetRequestedCode sets RequestedCode field to given value.

### HasRequestedCode

`func (o *AffiliateAffiliateStanding) HasRequestedCode() bool`

HasRequestedCode returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliateAffiliateStanding) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliateAffiliateStanding) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliateAffiliateStanding) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliateAffiliateStanding) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


