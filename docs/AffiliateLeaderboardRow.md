# AffiliateLeaderboardRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccruedCents** | Pointer to **int64** | AccruedCents is that affiliate&#39;s lifetime commission accrued, in cents, and what the board is ordered by. An aggregate: no per-customer figure is exposed. | [optional] 
**Handle** | Pointer to **string** | Handle is the affiliate&#39;s self-chosen display name — the only identity the board ever carries. The org behind it is never disclosed. | [optional] 
**IsYou** | Pointer to **bool** | IsYou marks the caller&#39;s own row, so a client can highlight it without matching on a handle. Absent on every other row. | [optional] 
**Rank** | Pointer to **int64** | Rank is the position in the GLOBAL approved set ordered by lifetime accrued commission, 1-based. Affiliates that set no handle still occupy their rank and are simply not listed, so the visible ranks have gaps and the board is not a complete roster. On the caller&#39;s own row the rank is computed over the whole set, so it is exact well outside the top page. | [optional] 
**ReferredCount** | Pointer to **int64** | ReferredCount is how many orgs that affiliate directly referred — a count only, never which orgs. | [optional] 

## Methods

### NewAffiliateLeaderboardRow

`func NewAffiliateLeaderboardRow() *AffiliateLeaderboardRow`

NewAffiliateLeaderboardRow instantiates a new AffiliateLeaderboardRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAffiliateLeaderboardRowWithDefaults

`func NewAffiliateLeaderboardRowWithDefaults() *AffiliateLeaderboardRow`

NewAffiliateLeaderboardRowWithDefaults instantiates a new AffiliateLeaderboardRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccruedCents

`func (o *AffiliateLeaderboardRow) GetAccruedCents() int64`

GetAccruedCents returns the AccruedCents field if non-nil, zero value otherwise.

### GetAccruedCentsOk

`func (o *AffiliateLeaderboardRow) GetAccruedCentsOk() (*int64, bool)`

GetAccruedCentsOk returns a tuple with the AccruedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccruedCents

`func (o *AffiliateLeaderboardRow) SetAccruedCents(v int64)`

SetAccruedCents sets AccruedCents field to given value.

### HasAccruedCents

`func (o *AffiliateLeaderboardRow) HasAccruedCents() bool`

HasAccruedCents returns a boolean if a field has been set.

### GetHandle

`func (o *AffiliateLeaderboardRow) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *AffiliateLeaderboardRow) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *AffiliateLeaderboardRow) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *AffiliateLeaderboardRow) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetIsYou

`func (o *AffiliateLeaderboardRow) GetIsYou() bool`

GetIsYou returns the IsYou field if non-nil, zero value otherwise.

### GetIsYouOk

`func (o *AffiliateLeaderboardRow) GetIsYouOk() (*bool, bool)`

GetIsYouOk returns a tuple with the IsYou field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsYou

`func (o *AffiliateLeaderboardRow) SetIsYou(v bool)`

SetIsYou sets IsYou field to given value.

### HasIsYou

`func (o *AffiliateLeaderboardRow) HasIsYou() bool`

HasIsYou returns a boolean if a field has been set.

### GetRank

`func (o *AffiliateLeaderboardRow) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *AffiliateLeaderboardRow) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *AffiliateLeaderboardRow) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *AffiliateLeaderboardRow) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetReferredCount

`func (o *AffiliateLeaderboardRow) GetReferredCount() int64`

GetReferredCount returns the ReferredCount field if non-nil, zero value otherwise.

### GetReferredCountOk

`func (o *AffiliateLeaderboardRow) GetReferredCountOk() (*int64, bool)`

GetReferredCountOk returns a tuple with the ReferredCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferredCount

`func (o *AffiliateLeaderboardRow) SetReferredCount(v int64)`

SetReferredCount sets ReferredCount field to given value.

### HasReferredCount

`func (o *AffiliateLeaderboardRow) HasReferredCount() bool`

HasReferredCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


