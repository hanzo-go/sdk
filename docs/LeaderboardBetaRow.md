# LeaderboardBetaRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Applied** | Pointer to **string** | Applied is the day the application began, 2006-01-02: the record&#39;s own, or the day of the lead it claims when that was earlier. | [optional] 
**Name** | Pointer to **string** | Name is the handle its owner listed under on the public board; null for a person who did not opt in, whom the page calls a builder. The caller&#39;s own row is always named. | [optional] 
**Qualified** | Pointer to **int64** | Qualified is how many of them have made metered spend. | [optional] 
**Rank** | Pointer to **int64** | Rank is the record&#39;s place, 1 first. | [optional] 
**Referrals** | Pointer to **int64** | Referrals is how many orgs signed up with the record&#39;s referral code. | [optional] 
**Score** | Pointer to **int64** | Score is what the row is ranked by: points, referrals, qualified referrals and tokens, as GET /v1/leaderboard/beta&#39;s description states. | [optional] 
**Tokens** | Pointer to **string** | Tokens is the person&#39;s all-time token usage as a range: 0, &lt;1M, 1M-10M, 10M-100M, 100M-1B or 1B+. | [optional] 
**TokensUsed** | Pointer to **int64** | TokensUsed is the exact count, on a named row and the caller&#39;s; absent on every other. | [optional] 
**You** | Pointer to **bool** | You marks the caller&#39;s own row. | [optional] 

## Methods

### NewLeaderboardBetaRow

`func NewLeaderboardBetaRow() *LeaderboardBetaRow`

NewLeaderboardBetaRow instantiates a new LeaderboardBetaRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBetaRowWithDefaults

`func NewLeaderboardBetaRowWithDefaults() *LeaderboardBetaRow`

NewLeaderboardBetaRowWithDefaults instantiates a new LeaderboardBetaRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplied

`func (o *LeaderboardBetaRow) GetApplied() string`

GetApplied returns the Applied field if non-nil, zero value otherwise.

### GetAppliedOk

`func (o *LeaderboardBetaRow) GetAppliedOk() (*string, bool)`

GetAppliedOk returns a tuple with the Applied field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplied

`func (o *LeaderboardBetaRow) SetApplied(v string)`

SetApplied sets Applied field to given value.

### HasApplied

`func (o *LeaderboardBetaRow) HasApplied() bool`

HasApplied returns a boolean if a field has been set.

### GetName

`func (o *LeaderboardBetaRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LeaderboardBetaRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LeaderboardBetaRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LeaderboardBetaRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQualified

`func (o *LeaderboardBetaRow) GetQualified() int64`

GetQualified returns the Qualified field if non-nil, zero value otherwise.

### GetQualifiedOk

`func (o *LeaderboardBetaRow) GetQualifiedOk() (*int64, bool)`

GetQualifiedOk returns a tuple with the Qualified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualified

`func (o *LeaderboardBetaRow) SetQualified(v int64)`

SetQualified sets Qualified field to given value.

### HasQualified

`func (o *LeaderboardBetaRow) HasQualified() bool`

HasQualified returns a boolean if a field has been set.

### GetRank

`func (o *LeaderboardBetaRow) GetRank() int64`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *LeaderboardBetaRow) GetRankOk() (*int64, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *LeaderboardBetaRow) SetRank(v int64)`

SetRank sets Rank field to given value.

### HasRank

`func (o *LeaderboardBetaRow) HasRank() bool`

HasRank returns a boolean if a field has been set.

### GetReferrals

`func (o *LeaderboardBetaRow) GetReferrals() int64`

GetReferrals returns the Referrals field if non-nil, zero value otherwise.

### GetReferralsOk

`func (o *LeaderboardBetaRow) GetReferralsOk() (*int64, bool)`

GetReferralsOk returns a tuple with the Referrals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferrals

`func (o *LeaderboardBetaRow) SetReferrals(v int64)`

SetReferrals sets Referrals field to given value.

### HasReferrals

`func (o *LeaderboardBetaRow) HasReferrals() bool`

HasReferrals returns a boolean if a field has been set.

### GetScore

`func (o *LeaderboardBetaRow) GetScore() int64`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *LeaderboardBetaRow) GetScoreOk() (*int64, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *LeaderboardBetaRow) SetScore(v int64)`

SetScore sets Score field to given value.

### HasScore

`func (o *LeaderboardBetaRow) HasScore() bool`

HasScore returns a boolean if a field has been set.

### GetTokens

`func (o *LeaderboardBetaRow) GetTokens() string`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *LeaderboardBetaRow) GetTokensOk() (*string, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *LeaderboardBetaRow) SetTokens(v string)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *LeaderboardBetaRow) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetTokensUsed

`func (o *LeaderboardBetaRow) GetTokensUsed() int64`

GetTokensUsed returns the TokensUsed field if non-nil, zero value otherwise.

### GetTokensUsedOk

`func (o *LeaderboardBetaRow) GetTokensUsedOk() (*int64, bool)`

GetTokensUsedOk returns a tuple with the TokensUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokensUsed

`func (o *LeaderboardBetaRow) SetTokensUsed(v int64)`

SetTokensUsed sets TokensUsed field to given value.

### HasTokensUsed

`func (o *LeaderboardBetaRow) HasTokensUsed() bool`

HasTokensUsed returns a boolean if a field has been set.

### GetYou

`func (o *LeaderboardBetaRow) GetYou() bool`

GetYou returns the You field if non-nil, zero value otherwise.

### GetYouOk

`func (o *LeaderboardBetaRow) GetYouOk() (*bool, bool)`

GetYouOk returns a tuple with the You field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYou

`func (o *LeaderboardBetaRow) SetYou(v bool)`

SetYou sets You field to given value.

### HasYou

`func (o *LeaderboardBetaRow) HasYou() bool`

HasYou returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


