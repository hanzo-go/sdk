# RewardPostList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | Pointer to **int64** | Pending is how many posts are held and count for nothing yet. | [optional] 
**Posts** | Pointer to [**[]RewardPostView**](RewardPostView.md) |  | [optional] 

## Methods

### NewRewardPostList

`func NewRewardPostList() *RewardPostList`

NewRewardPostList instantiates a new RewardPostList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardPostListWithDefaults

`func NewRewardPostListWithDefaults() *RewardPostList`

NewRewardPostListWithDefaults instantiates a new RewardPostList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *RewardPostList) GetPending() int64`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *RewardPostList) GetPendingOk() (*int64, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *RewardPostList) SetPending(v int64)`

SetPending sets Pending field to given value.

### HasPending

`func (o *RewardPostList) HasPending() bool`

HasPending returns a boolean if a field has been set.

### GetPosts

`func (o *RewardPostList) GetPosts() []RewardPostView`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *RewardPostList) GetPostsOk() (*[]RewardPostView, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *RewardPostList) SetPosts(v []RewardPostView)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *RewardPostList) HasPosts() bool`

HasPosts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


