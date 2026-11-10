# RewardBadgeOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Badge** | Pointer to **string** | Badge is which one the README carries: built (a hanzo.ai badge) or deploy (a Deploy to Hanzo button). | [optional] 
**Claimed** | Pointer to **bool** | Claimed is whether this check made a claim for staff: false when the project had already been claimed. | [optional] 
**Repo** | Pointer to **string** | Repo is the project, as GitHub names it. | [optional] 

## Methods

### NewRewardBadgeOut

`func NewRewardBadgeOut() *RewardBadgeOut`

NewRewardBadgeOut instantiates a new RewardBadgeOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardBadgeOutWithDefaults

`func NewRewardBadgeOutWithDefaults() *RewardBadgeOut`

NewRewardBadgeOutWithDefaults instantiates a new RewardBadgeOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBadge

`func (o *RewardBadgeOut) GetBadge() string`

GetBadge returns the Badge field if non-nil, zero value otherwise.

### GetBadgeOk

`func (o *RewardBadgeOut) GetBadgeOk() (*string, bool)`

GetBadgeOk returns a tuple with the Badge field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBadge

`func (o *RewardBadgeOut) SetBadge(v string)`

SetBadge sets Badge field to given value.

### HasBadge

`func (o *RewardBadgeOut) HasBadge() bool`

HasBadge returns a boolean if a field has been set.

### GetClaimed

`func (o *RewardBadgeOut) GetClaimed() bool`

GetClaimed returns the Claimed field if non-nil, zero value otherwise.

### GetClaimedOk

`func (o *RewardBadgeOut) GetClaimedOk() (*bool, bool)`

GetClaimedOk returns a tuple with the Claimed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimed

`func (o *RewardBadgeOut) SetClaimed(v bool)`

SetClaimed sets Claimed field to given value.

### HasClaimed

`func (o *RewardBadgeOut) HasClaimed() bool`

HasClaimed returns a boolean if a field has been set.

### GetRepo

`func (o *RewardBadgeOut) GetRepo() string`

GetRepo returns the Repo field if non-nil, zero value otherwise.

### GetRepoOk

`func (o *RewardBadgeOut) GetRepoOk() (*string, bool)`

GetRepoOk returns a tuple with the Repo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepo

`func (o *RewardBadgeOut) SetRepo(v string)`

SetRepo sets Repo field to given value.

### HasRepo

`func (o *RewardBadgeOut) HasRepo() bool`

HasRepo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


