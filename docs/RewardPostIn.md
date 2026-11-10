# RewardPostIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is who you are on that network, so the post&#39;s author can be matched. Optional; without it the author match is unknown. | [optional] 
**Url** | Pointer to **string** | URL is the address of the post, on X, Bluesky, Reddit, Hacker News, Mastodon or LinkedIn. | [optional] 

## Methods

### NewRewardPostIn

`func NewRewardPostIn() *RewardPostIn`

NewRewardPostIn instantiates a new RewardPostIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRewardPostInWithDefaults

`func NewRewardPostInWithDefaults() *RewardPostIn`

NewRewardPostInWithDefaults instantiates a new RewardPostIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *RewardPostIn) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *RewardPostIn) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *RewardPostIn) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *RewardPostIn) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetUrl

`func (o *RewardPostIn) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *RewardPostIn) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *RewardPostIn) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *RewardPostIn) HasUrl() bool`

HasUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


