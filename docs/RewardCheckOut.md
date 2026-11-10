# RewardCheckOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Follows** | Pointer to **int64** |  | [optional] 
**Forks** | Pointer to **int64** |  | [optional] 
**Login** | Pointer to **string** | Login is the GitHub account the caller linked. | [optional] 
**Stars** | Pointer to **int64** | Stars, Follows and Forks are the claims this check made, each for one repository or account of ours that GitHub shows the caller has starred, followed or forked, and nobody had claimed before. | [optional] 

## Methods

### NewRewardCheckOut

`func NewRewardCheckOut() *RewardCheckOut`

NewRewardCheckOut instantiates a new RewardCheckOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardCheckOutWithDefaults

`func NewRewardCheckOutWithDefaults() *RewardCheckOut`

NewRewardCheckOutWithDefaults instantiates a new RewardCheckOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFollows

`func (o *RewardCheckOut) GetFollows() int64`

GetFollows returns the Follows field if non-nil, zero value otherwise.

### GetFollowsOk

`func (o *RewardCheckOut) GetFollowsOk() (*int64, bool)`

GetFollowsOk returns a tuple with the Follows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollows

`func (o *RewardCheckOut) SetFollows(v int64)`

SetFollows sets Follows field to given value.

### HasFollows

`func (o *RewardCheckOut) HasFollows() bool`

HasFollows returns a boolean if a field has been set.

### GetForks

`func (o *RewardCheckOut) GetForks() int64`

GetForks returns the Forks field if non-nil, zero value otherwise.

### GetForksOk

`func (o *RewardCheckOut) GetForksOk() (*int64, bool)`

GetForksOk returns a tuple with the Forks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForks

`func (o *RewardCheckOut) SetForks(v int64)`

SetForks sets Forks field to given value.

### HasForks

`func (o *RewardCheckOut) HasForks() bool`

HasForks returns a boolean if a field has been set.

### GetLogin

`func (o *RewardCheckOut) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *RewardCheckOut) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *RewardCheckOut) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *RewardCheckOut) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetStars

`func (o *RewardCheckOut) GetStars() int64`

GetStars returns the Stars field if non-nil, zero value otherwise.

### GetStarsOk

`func (o *RewardCheckOut) GetStarsOk() (*int64, bool)`

GetStarsOk returns a tuple with the Stars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStars

`func (o *RewardCheckOut) SetStars(v int64)`

SetStars sets Stars field to given value.

### HasStars

`func (o *RewardCheckOut) HasStars() bool`

HasStars returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


