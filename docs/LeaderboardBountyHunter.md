# LeaderboardBountyHunter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badge** | Pointer to **string** | Badge is the gamified tier title earned by the contributor. | [optional] 
**BountiesMerged** | Pointer to **int64** | BountiesMerged is how many bounty pull requests this contributor has landed. | [optional] 
**BountyPoints** | Pointer to **int64** | BountyPoints is the sum of points earned from bug reports and bounties. | [optional] 
**BugsReported** | Pointer to **int64** | BugsReported is how many verified bug reports this contributor has filed. | [optional] 
**CreditsEarnedCents** | Pointer to **int64** | CreditsEarnedCents is the total AI credits in US cents granted for bounties and bug reports. | [optional] 
**GlobalTokenRank** | Pointer to **int64** | GlobalTokenRank is the contributor&#39;s global organization rank by token volume (0 if unranked). | [optional] 
**Handle** | Pointer to **string** | Handle is the contributor&#39;s public display handle. | [optional] 
**Org** | Pointer to **string** | Org is the contributor&#39;s organization slug when listed. | [optional] 
**Rank** | Pointer to **int64** | Rank is the contributor&#39;s 1-based standing on the bounty board. | [optional] 
**Self** | Pointer to **bool** | Self marks the signed-in caller&#39;s own row. | [optional] 
**TokensUsed** | Pointer to **int64** | TokensUsed is the contributor&#39;s all-time token volume in the warehouse. | [optional] 
**TotalScore** | Pointer to **int64** | TotalScore is the org&#39;s points by client.Points, the one score every board shares. | [optional] 

## Methods

### NewLeaderboardBountyHunter

`func NewLeaderboardBountyHunter() *LeaderboardBountyHunter`

NewLeaderboardBountyHunter instantiates a new LeaderboardBountyHunter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBountyHunterWithDefaults

`func NewLeaderboardBountyHunterWithDefaults() *LeaderboardBountyHunter`

NewLeaderboardBountyHunterWithDefaults instantiates a new LeaderboardBountyHunter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadge

`func (o *LeaderboardBountyHunter) GetBadge() string`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *LeaderboardBountyHunter) GetBadgeOk() (*string, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *LeaderboardBountyHunter) SetBadge(v string)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *LeaderboardBountyHunter) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetBountiesMerged

`func (o *LeaderboardBountyHunter) GetBountiesMerged() int64`

GetBountiesMerged returns the BountiesMerged field if non-nil, zero value otherwise.

### GetBountiesMergedOk

`func (o *LeaderboardBountyHunter) GetBountiesMergedOk() (*int64, bool)`

GetBountiesMergedOk returns a tuple with the BountiesMerged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountiesMerged

`func (o *LeaderboardBountyHunter) SetBountiesMerged(v int64)`

SetBountiesMerged sets BountiesMerged field to given value.

### HasBountiesMerged

`func (o *LeaderboardBountyHunter) HasBountiesMerged() bool`

HasBountiesMerged returns a boolean if a field has been set.

### GetBountyPoints

`func (o *LeaderboardBountyHunter) GetBountyPoints() int64`

GetBountyPoints returns the BountyPoints field if non-nil, zero value otherwise.

### GetBountyPointsOk

`func (o *LeaderboardBountyHunter) GetBountyPointsOk() (*int64, bool)`

GetBountyPointsOk returns a tuple with the BountyPoints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBountyPoints

`func (o *LeaderboardBountyHunter) SetBountyPoints(v int64)`

SetBountyPoints sets BountyPoints field to given value.

### HasBountyPoints

`func (o *LeaderboardBountyHunter) HasBountyPoints() bool`

HasBountyPoints returns a boolean if a field has been set.

### GetBugsReported

`func (o *LeaderboardBountyHunter) GetBugsReported() int64`

GetBugsReported returns the BugsReported field if non-nil, zero value otherwise.

### GetBugsReportedOk

`func (o *LeaderboardBountyHunter) GetBugsReportedOk() (*int64, bool)`

GetBugsReportedOk returns a tuple with the BugsReported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBugsReported

`func (o *LeaderboardBountyHunter) SetBugsReported(v int64)`

SetBugsReported sets BugsReported field to given value.

### HasBugsReported

`func (o *LeaderboardBountyHunter) HasBugsReported() bool`

HasBugsReported returns a boolean if a field has been set.

### GetCreditsEarnedCents

`func (o *LeaderboardBountyHunter) GetCreditsEarnedCents() int64`

GetCreditsEarnedCents returns the CreditsEarnedCents field if non-nil, zero value otherwise.

### GetCreditsEarnedCentsOk

`func (o *LeaderboardBountyHunter) GetCreditsEarnedCentsOk() (*int64, bool)`

GetCreditsEarnedCentsOk returns a tuple with the CreditsEarnedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsEarnedCents

`func (o *LeaderboardBountyHunter) SetCreditsEarnedCents(v int64)`

SetCreditsEarnedCents sets CreditsEarnedCents field to given value.

### HasCreditsEarnedCents

`func (o *LeaderboardBountyHunter) HasCreditsEarnedCents() bool`

HasCreditsEarnedCents returns a boolean if a field has been set.

### GetGlobalTokenRank

`func (o *LeaderboardBountyHunter) GetGlobalTokenRank() int64`

GetGlobalTokenRank returns the GlobalTokenRank field if non-nil, zero value otherwise.

### GetGlobalTokenRankOk

`func (o *LeaderboardBountyHunter) GetGlobalTokenRankOk() (*int64, bool)`

GetGlobalTokenRankOk returns a tuple with the GlobalTokenRank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGlobalTokenRank

`func (o *LeaderboardBountyHunter) SetGlobalTokenRank(v int64)`

SetGlobalTokenRank sets GlobalTokenRank field to given value.

### HasGlobalTokenRank

`func (o *LeaderboardBountyHunter) HasGlobalTokenRank() bool`

HasGlobalTokenRank returns a boolean if a field has been set.

### GetHandle

`func (o *LeaderboardBountyHunter) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardBountyHunter) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardBountyHunter) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardBountyHunter) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetOrg

`func (o *LeaderboardBountyHunter) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *LeaderboardBountyHunter) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *LeaderboardBountyHunter) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *LeaderboardBountyHunter) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetRank

`func (o *LeaderboardBountyHunter) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *LeaderboardBountyHunter) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *LeaderboardBountyHunter) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *LeaderboardBountyHunter) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardBountyHunter) GetSelf() bool`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardBountyHunter) GetSelfOk() (*bool, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardBountyHunter) SetSelf(v bool)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardBountyHunter) HasSelf() bool`

HasSelf returns a boolean if a field has been set.

### GetTokensUsed

`func (o *LeaderboardBountyHunter) GetTokensUsed() int64`

GetTokensUsed returns the TokensUsed field if non-nil, zero value otherwise.

### GetTokensUsedOk

`func (o *LeaderboardBountyHunter) GetTokensUsedOk() (*int64, bool)`

GetTokensUsedOk returns a tuple with the TokensUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensUsed

`func (o *LeaderboardBountyHunter) SetTokensUsed(v int64)`

SetTokensUsed sets TokensUsed field to given value.

### HasTokensUsed

`func (o *LeaderboardBountyHunter) HasTokensUsed() bool`

HasTokensUsed returns a boolean if a field has been set.

### GetTotalScore

`func (o *LeaderboardBountyHunter) GetTotalScore() int64`

GetTotalScore returns the TotalScore field if non-nil, zero value otherwise.

### GetTotalScoreOk

`func (o *LeaderboardBountyHunter) GetTotalScoreOk() (*int64, bool)`

GetTotalScoreOk returns a tuple with the TotalScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalScore

`func (o *LeaderboardBountyHunter) SetTotalScore(v int64)`

SetTotalScore sets TotalScore field to given value.

### HasTotalScore

`func (o *LeaderboardBountyHunter) HasTotalScore() bool`

HasTotalScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


