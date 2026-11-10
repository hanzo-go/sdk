# LeaderboardBountyAwardView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is the unix timestamp (seconds) when the award was recorded. | [optional] 
**CreditsCents** | Pointer to **int64** | CreditsCents is the AI credit grant in US cents awarded with this bounty. | [optional] 
**Handle** | Pointer to **string** | Handle is the recipient&#39;s display name. | [optional] 
**Id** | Pointer to **string** | ID is the unique award identifier. | [optional] 
**Kind** | Pointer to **string** | Kind is bug_report or bounty. | [optional] 
**Points** | Pointer to **int64** | Points is the leaderboard points awarded. | [optional] 
**Summary** | Pointer to **string** | Summary describes what was reported or fixed. | [optional] 
**Target** | Pointer to **string** | Target is the repository or issue reference (e.g. hanzo-inc/cloud#index-lock). | [optional] 

## Methods

### NewLeaderboardBountyAwardView

`func NewLeaderboardBountyAwardView() *LeaderboardBountyAwardView`

NewLeaderboardBountyAwardView instantiates a new LeaderboardBountyAwardView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLeaderboardBountyAwardViewWithDefaults

`func NewLeaderboardBountyAwardViewWithDefaults() *LeaderboardBountyAwardView`

NewLeaderboardBountyAwardViewWithDefaults instantiates a new LeaderboardBountyAwardView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *LeaderboardBountyAwardView) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *LeaderboardBountyAwardView) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *LeaderboardBountyAwardView) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *LeaderboardBountyAwardView) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetCreditsCents

`func (o *LeaderboardBountyAwardView) GetCreditsCents() int64`

GetCreditsCents returns the CreditsCents field if non-nil, zero value otherwise.

### GetCreditsCentsOk

`func (o *LeaderboardBountyAwardView) GetCreditsCentsOk() (*int64, bool)`

GetCreditsCentsOk returns a tuple with the CreditsCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreditsCents

`func (o *LeaderboardBountyAwardView) SetCreditsCents(v int64)`

SetCreditsCents sets CreditsCents field to given value.

### HasCreditsCents

`func (o *LeaderboardBountyAwardView) HasCreditsCents() bool`

HasCreditsCents returns a boolean if a field has been set.

### GetHandle

`func (o *LeaderboardBountyAwardView) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *LeaderboardBountyAwardView) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *LeaderboardBountyAwardView) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *LeaderboardBountyAwardView) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *LeaderboardBountyAwardView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LeaderboardBountyAwardView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LeaderboardBountyAwardView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LeaderboardBountyAwardView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKind

`func (o *LeaderboardBountyAwardView) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *LeaderboardBountyAwardView) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *LeaderboardBountyAwardView) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *LeaderboardBountyAwardView) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetPoints

`func (o *LeaderboardBountyAwardView) GetPoints() int64`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *LeaderboardBountyAwardView) GetPointsOk() (*int64, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *LeaderboardBountyAwardView) SetPoints(v int64)`

SetPoints sets Points field to given value.

### HasPoints

`func (o *LeaderboardBountyAwardView) HasPoints() bool`

HasPoints returns a boolean if a field has been set.

### GetSummary

`func (o *LeaderboardBountyAwardView) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *LeaderboardBountyAwardView) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *LeaderboardBountyAwardView) SetSummary(v string)`

SetSummary sets Summary field to given value.

### HasSummary

`func (o *LeaderboardBountyAwardView) HasSummary() bool`

HasSummary returns a boolean if a field has been set.

### GetTarget

`func (o *LeaderboardBountyAwardView) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *LeaderboardBountyAwardView) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *LeaderboardBountyAwardView) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *LeaderboardBountyAwardView) HasTarget() bool`

HasTarget returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


