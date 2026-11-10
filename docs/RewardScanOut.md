# RewardScanOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Held** | Pointer to **int64** |  | [optional] 
**Looked** | Pointer to **int64** | Looked is how many of the recent posts were read, Held how many carried your link and are now held, and Skipped how many did not. | [optional] 
**Posts** | Pointer to [**[]RewardPostView**](RewardPostView.md) |  | [optional] 
**Skipped** | Pointer to **int64** |  | [optional] 

## Methods

### NewRewardScanOut

`func NewRewardScanOut() *RewardScanOut`

NewRewardScanOut instantiates a new RewardScanOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardScanOutWithDefaults

`func NewRewardScanOutWithDefaults() *RewardScanOut`

NewRewardScanOutWithDefaults instantiates a new RewardScanOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeld

`func (o *RewardScanOut) GetHeld() int64`

GetHeld returns the Held field if non-nil, zero value otherwise.

### GetHeldOk

`func (o *RewardScanOut) GetHeldOk() (*int64, bool)`

GetHeldOk returns a tuple with the Held field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeld

`func (o *RewardScanOut) SetHeld(v int64)`

SetHeld sets Held field to given value.

### HasHeld

`func (o *RewardScanOut) HasHeld() bool`

HasHeld returns a boolean if a field has been set.

### GetLooked

`func (o *RewardScanOut) GetLooked() int64`

GetLooked returns the Looked field if non-nil, zero value otherwise.

### GetLookedOk

`func (o *RewardScanOut) GetLookedOk() (*int64, bool)`

GetLookedOk returns a tuple with the Looked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLooked

`func (o *RewardScanOut) SetLooked(v int64)`

SetLooked sets Looked field to given value.

### HasLooked

`func (o *RewardScanOut) HasLooked() bool`

HasLooked returns a boolean if a field has been set.

### GetPosts

`func (o *RewardScanOut) GetPosts() []RewardPostView`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *RewardScanOut) GetPostsOk() (*[]RewardPostView, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *RewardScanOut) SetPosts(v []RewardPostView)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *RewardScanOut) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetSkipped

`func (o *RewardScanOut) GetSkipped() int64`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *RewardScanOut) GetSkippedOk() (*int64, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *RewardScanOut) SetSkipped(v int64)`

SetSkipped sets Skipped field to given value.

### HasSkipped

`func (o *RewardScanOut) HasSkipped() bool`

HasSkipped returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


