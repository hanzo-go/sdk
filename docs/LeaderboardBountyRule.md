# LeaderboardBountyRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreditsCents** | Pointer to **int64** | CreditsCents is the default AI credit grant in US cents awarded on verification (0 when variable or points-only). | [optional] 
**Description** | Pointer to **string** | Description explains how the action is verified and credited. | [optional] 
**Kind** | Pointer to **string** | Kind is the action identifier: bug_report, bounty, star, qualified_referral, or tokens. | [optional] 
**Label** | Pointer to **string** | Label is the human-readable title of the reward tier. | [optional] 
**Points** | Pointer to **int64** | Points is how many leaderboard points the action awards. | [optional] 

## Methods

### NewLeaderboardBountyRule

`func NewLeaderboardBountyRule() *LeaderboardBountyRule`

NewLeaderboardBountyRule instantiates a new LeaderboardBountyRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBountyRuleWithDefaults

`func NewLeaderboardBountyRuleWithDefaults() *LeaderboardBountyRule`

NewLeaderboardBountyRuleWithDefaults instantiates a new LeaderboardBountyRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreditsCents

`func (o *LeaderboardBountyRule) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *LeaderboardBountyRule) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *LeaderboardBountyRule) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *LeaderboardBountyRule) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetDescription

`func (o *LeaderboardBountyRule) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LeaderboardBountyRule) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LeaderboardBountyRule) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LeaderboardBountyRule) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *LeaderboardBountyRule) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LeaderboardBountyRule) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LeaderboardBountyRule) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LeaderboardBountyRule) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetLabel

`func (o *LeaderboardBountyRule) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *LeaderboardBountyRule) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *LeaderboardBountyRule) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *LeaderboardBountyRule) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetPoints

`func (o *LeaderboardBountyRule) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *LeaderboardBountyRule) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *LeaderboardBountyRule) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *LeaderboardBountyRule) HasPoints() bool`

HasPoints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


