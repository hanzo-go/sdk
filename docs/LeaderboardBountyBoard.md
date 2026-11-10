# LeaderboardBountyBoard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hunters** | Pointer to [**[]LeaderboardBountyHunter**](LeaderboardBountyHunter.md) | Hunters are the top-ranked bug hunters and bounty contributors, best first. | [optional] 
**RecentAwards** | Pointer to [**[]LeaderboardBountyAwardView**](LeaderboardBountyAwardView.md) | RecentAwards are the newest verified bug reports and bounty payouts. | [optional] 
**Rules** | Pointer to [**[]LeaderboardBountyRule**](LeaderboardBountyRule.md) | Rules lists how contributors earn bounty points, AI credits, and leaderboard ranks. | [optional] 
**Self** | Pointer to [**LeaderboardBountyHunter**](LeaderboardBountyHunter.md) | Self is the signed-in caller&#39;s own bounty standing when signed in. | [optional] 

## Methods

### NewLeaderboardBountyBoard

`func NewLeaderboardBountyBoard() *LeaderboardBountyBoard`

NewLeaderboardBountyBoard instantiates a new LeaderboardBountyBoard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBountyBoardWithDefaults

`func NewLeaderboardBountyBoardWithDefaults() *LeaderboardBountyBoard`

NewLeaderboardBountyBoardWithDefaults instantiates a new LeaderboardBountyBoard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHunters

`func (o *LeaderboardBountyBoard) GetHunters() []LeaderboardBountyHunter`

GetHunters returns the Hunters field if non-nil, zero value otherwise.

### GetHuntersOk

`func (o *LeaderboardBountyBoard) GetHuntersOk() (*[]LeaderboardBountyHunter, bool)`

GetHuntersOk returns a tuple with the Hunters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHunters

`func (o *LeaderboardBountyBoard) SetHunters(v []LeaderboardBountyHunter)`

SetHunters sets Hunters field to given value.

### HasHunters

`func (o *LeaderboardBountyBoard) HasHunters() bool`

HasHunters returns a boolean if a field has been set.

### GetRecentAwards

`func (o *LeaderboardBountyBoard) GetRecentAwards() []LeaderboardBountyAwardView`

GetRecentAwards returns the RecentAwards field if non-nil, zero value otherwise.

### GetRecentAwardsOk

`func (o *LeaderboardBountyBoard) GetRecentAwardsOk() (*[]LeaderboardBountyAwardView, bool)`

GetRecentAwardsOk returns a tuple with the RecentAwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentAwards

`func (o *LeaderboardBountyBoard) SetRecentAwards(v []LeaderboardBountyAwardView)`

SetRecentAwards sets RecentAwards field to given value.

### HasRecentAwards

`func (o *LeaderboardBountyBoard) HasRecentAwards() bool`

HasRecentAwards returns a boolean if a field has been set.

### GetRules

`func (o *LeaderboardBountyBoard) GetRules() []LeaderboardBountyRule`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *LeaderboardBountyBoard) GetRulesOk() (*[]LeaderboardBountyRule, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *LeaderboardBountyBoard) SetRules(v []LeaderboardBountyRule)`

SetRules sets Rules field to given value.

### HasRules

`func (o *LeaderboardBountyBoard) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetSelf

`func (o *LeaderboardBountyBoard) GetSelf() LeaderboardBountyHunter`

GetSelf returns the Self field if non-nil, zero value otherwise.

### GetSelfOk

`func (o *LeaderboardBountyBoard) GetSelfOk() (*LeaderboardBountyHunter, bool)`

GetSelfOk returns a tuple with the Self field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelf

`func (o *LeaderboardBountyBoard) SetSelf(v LeaderboardBountyHunter)`

SetSelf sets Self field to given value.

### HasSelf

`func (o *LeaderboardBountyBoard) HasSelf() bool`

HasSelf returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


