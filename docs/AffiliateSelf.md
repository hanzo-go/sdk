# AffiliateSelf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int32** | AccruedCents is lifetime commission accrued, in cents. It only grows — a payout is recorded against paidCents and never reduces this. | [optional] 
**Code** | Pointer to **string** | Code is the minted referral code, the slug the ?aff link carries. Absent until staff approve; codes live in ONE global namespace across all affiliates. | [optional] 
**DefaultRateBps** | Pointer to **int32** | DefaultRateBps is the direct rate a new affiliate starts at, in basis points of margin (2000 &#x3D; 20%). Answered ONLY to a caller that has not applied, as the quote beside &#x60;schedule&#x60;. | [optional] 
**DownlineTotal** | Pointer to **int32** | DownlineTotal counts every org in the caller&#39;s downline across the levels. | [optional] 
**Handle** | Pointer to **string** | Handle is the opt-in public leaderboard name. Empty means opted out: the caller keeps its rank and still sees its own row, it is just not listed. | [optional] 
**Id** | Pointer to **string** | ID is the affiliate&#39;s server-minted handle, \&quot;aff_\&quot;-prefixed. Absent until the org applies. | [optional] 
**IsAffiliate** | Pointer to **bool** | IsAffiliate says whether the caller org has an affiliate record. On false the answer carries the rate SCHEDULE and the default rate instead of a downline, so the console can show what the caller would earn. | [optional] 
**Levels** | Pointer to [**[]LevelView**](LevelView.md) | Levels is the caller&#39;s downline per upline level, with the rate paid there. | [optional] 
**Link** | Pointer to **string** | Link is the shareable ?aff URL built from the code. Empty until a code is minted, since there is nothing to share before approval. | [optional] 
**MarginBps** | Pointer to **int32** | MarginBps is the platform gross-margin fraction, in basis points, that every rate here is a rate OF. Read live per request, so it is the value in force now, not the one that applied to commission already accrued. | [optional] 
**PaidCents** | Pointer to **int32** | PaidCents is lifetime commission already paid out, in cents — credits grants and record-only cash disbursements alike. | [optional] 
**Payouts** | Pointer to [**[]Remittance**](Remittance.md) | Payouts is the payout history, newest first, bounded to the last 100 rows. | [optional] 
**PendingCents** | Pointer to **int32** | PendingCents is accrued minus paid, in cents — what the platform still owes and the ceiling on the next payout. Never negative. | [optional] 
**RateBps** | Pointer to **int32** | RateBps is the caller&#39;s OWN direct (level 1) commission rate, in basis points of margin. Levels 2 and 3 are platform-wide and appear in &#x60;levels&#x60;. | [optional] 
**Schedule** | Pointer to [**[]LevelView**](LevelView.md) | Schedule is the rate schedule quoted to a caller that has not applied. | [optional] 
**Status** | Pointer to **string** | Status is \&quot;applied\&quot;, \&quot;approved\&quot; or \&quot;suspended\&quot;; absent for a caller that never applied. Only \&quot;approved\&quot; mints links and accrues. | [optional] 

## Methods

### NewAffiliateSelf

`func NewAffiliateSelf() *AffiliateSelf`

NewAffiliateSelf instantiates a new AffiliateSelf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateSelfWithDefaults

`func NewAffiliateSelfWithDefaults() *AffiliateSelf`

NewAffiliateSelfWithDefaults instantiates a new AffiliateSelf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AffiliateSelf) GetAccruedCents() int32`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliateSelf) GetAccruedCentsOk() (*int32, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliateSelf) SetAccruedCents(v int32)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliateSelf) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetCode

`func (o *AffiliateSelf) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AffiliateSelf) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AffiliateSelf) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AffiliateSelf) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDefaultRateBps

`func (o *AffiliateSelf) GetDefaultRateBps() int32`

GetDefaultRateBps returns the DefaultRateBps field if non-nil, zero value otherwise.

### GetDefaultRateBpsOk

`func (o *AffiliateSelf) GetDefaultRateBpsOk() (*int32, bool)`

GetDefaultRateBpsOk returns a tuple with the DefaultRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultRateBps

`func (o *AffiliateSelf) SetDefaultRateBps(v int32)`

SetDefaultRateBps sets DefaultRateBps field to given value.

### HasDefaultRateBps

`func (o *AffiliateSelf) HasDefaultRateBps() bool`

HasDefaultRateBps returns a boolean if a field has been set.

### GetDownlineTotal

`func (o *AffiliateSelf) GetDownlineTotal() int32`

GetDownlineTotal returns the DownlineTotal field if non-nil, zero value otherwise.

### GetDownlineTotalOk

`func (o *AffiliateSelf) GetDownlineTotalOk() (*int32, bool)`

GetDownlineTotalOk returns a tuple with the DownlineTotal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownlineTotal

`func (o *AffiliateSelf) SetDownlineTotal(v int32)`

SetDownlineTotal sets DownlineTotal field to given value.

### HasDownlineTotal

`func (o *AffiliateSelf) HasDownlineTotal() bool`

HasDownlineTotal returns a boolean if a field has been set.

### GetHandle

`func (o *AffiliateSelf) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AffiliateSelf) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AffiliateSelf) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AffiliateSelf) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *AffiliateSelf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AffiliateSelf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AffiliateSelf) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AffiliateSelf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsAffiliate

`func (o *AffiliateSelf) GetIsAffiliate() bool`

GetIsAffiliate returns the IsAffiliate field if non-nil, zero value otherwise.

### GetIsAffiliateOk

`func (o *AffiliateSelf) GetIsAffiliateOk() (*bool, bool)`

GetIsAffiliateOk returns a tuple with the IsAffiliate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAffiliate

`func (o *AffiliateSelf) SetIsAffiliate(v bool)`

SetIsAffiliate sets IsAffiliate field to given value.

### HasIsAffiliate

`func (o *AffiliateSelf) HasIsAffiliate() bool`

HasIsAffiliate returns a boolean if a field has been set.

### GetLevels

`func (o *AffiliateSelf) GetLevels() []LevelView`

GetLevels returns the Levels field if non-nil, zero value otherwise.

### GetLevelsOk

`func (o *AffiliateSelf) GetLevelsOk() (*[]LevelView, bool)`

GetLevelsOk returns a tuple with the Levels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevels

`func (o *AffiliateSelf) SetLevels(v []LevelView)`

SetLevels sets Levels field to given value.

### HasLevels

`func (o *AffiliateSelf) HasLevels() bool`

HasLevels returns a boolean if a field has been set.

### GetLink

`func (o *AffiliateSelf) GetLink() string`

GetLink returns the Link field if non-nil, zero value otherwise.

### GetLinkOk

`func (o *AffiliateSelf) GetLinkOk() (*string, bool)`

GetLinkOk returns a tuple with the Link field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLink

`func (o *AffiliateSelf) SetLink(v string)`

SetLink sets Link field to given value.

### HasLink

`func (o *AffiliateSelf) HasLink() bool`

HasLink returns a boolean if a field has been set.

### GetMarginBps

`func (o *AffiliateSelf) GetMarginBps() int32`

GetMarginBps returns the MarginBps field if non-nil, zero value otherwise.

### GetMarginBpsOk

`func (o *AffiliateSelf) GetMarginBpsOk() (*int32, bool)`

GetMarginBpsOk returns a tuple with the MarginBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginBps

`func (o *AffiliateSelf) SetMarginBps(v int32)`

SetMarginBps sets MarginBps field to given value.

### HasMarginBps

`func (o *AffiliateSelf) HasMarginBps() bool`

HasMarginBps returns a boolean if a field has been set.

### GetPaidCents

`func (o *AffiliateSelf) GetPaidCents() int32`

GetPaidCents returns the PaidCents field if non-nil, zero value otherwise.

### GetPaidCentsOk

`func (o *AffiliateSelf) GetPaidCentsOk() (*int32, bool)`

GetPaidCentsOk returns a tuple with the PaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidCents

`func (o *AffiliateSelf) SetPaidCents(v int32)`

SetPaidCents sets PaidCents field to given value.

### HasPaidCents

`func (o *AffiliateSelf) HasPaidCents() bool`

HasPaidCents returns a boolean if a field has been set.

### GetPayouts

`func (o *AffiliateSelf) GetPayouts() []Remittance`

GetPayouts returns the Payouts field if non-nil, zero value otherwise.

### GetPayoutsOk

`func (o *AffiliateSelf) GetPayoutsOk() (*[]Remittance, bool)`

GetPayoutsOk returns a tuple with the Payouts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayouts

`func (o *AffiliateSelf) SetPayouts(v []Remittance)`

SetPayouts sets Payouts field to given value.

### HasPayouts

`func (o *AffiliateSelf) HasPayouts() bool`

HasPayouts returns a boolean if a field has been set.

### GetPendingCents

`func (o *AffiliateSelf) GetPendingCents() int32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *AffiliateSelf) GetPendingCentsOk() (*int32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *AffiliateSelf) SetPendingCents(v int32)`

SetPendingCents sets PendingCents field to given value.

### HasPendingCents

`func (o *AffiliateSelf) HasPendingCents() bool`

HasPendingCents returns a boolean if a field has been set.

### GetRateBps

`func (o *AffiliateSelf) GetRateBps() int32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *AffiliateSelf) GetRateBpsOk() (*int32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *AffiliateSelf) SetRateBps(v int32)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *AffiliateSelf) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### GetSchedule

`func (o *AffiliateSelf) GetSchedule() []LevelView`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AffiliateSelf) GetScheduleOk() (*[]LevelView, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AffiliateSelf) SetSchedule(v []LevelView)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AffiliateSelf) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetStatus

`func (o *AffiliateSelf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AffiliateSelf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AffiliateSelf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AffiliateSelf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


