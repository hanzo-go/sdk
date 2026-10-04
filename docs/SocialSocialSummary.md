# SocialSocialSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int64** | Accounts is how many accounts the org has connected, in any status. | [optional] 
**Posts** | Pointer to **int64** | Posts is how many posts the org has, in any state. | [optional] 
**Published** | Pointer to **int64** | Published is how many of them have published. | [optional] 
**Scheduled** | Pointer to **int64** | Scheduled is how many of them are waiting for their scheduled time. | [optional] 

## Methods

### NewSocialSocialSummary

`func NewSocialSocialSummary() *SocialSocialSummary`

NewSocialSocialSummary instantiates a new SocialSocialSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialSummaryWithDefaults

`func NewSocialSocialSummaryWithDefaults() *SocialSocialSummary`

NewSocialSocialSummaryWithDefaults instantiates a new SocialSocialSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SocialSocialSummary) GetAccounts() int64`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SocialSocialSummary) GetAccountsOk() (*int64, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SocialSocialSummary) SetAccounts(v int64)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SocialSocialSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetPosts

`func (o *SocialSocialSummary) GetPosts() int64`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *SocialSocialSummary) GetPostsOk() (*int64, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *SocialSocialSummary) SetPosts(v int64)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *SocialSocialSummary) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPublished

`func (o *SocialSocialSummary) GetPublished() int64`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SocialSocialSummary) GetPublishedOk() (*int64, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SocialSocialSummary) SetPublished(v int64)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SocialSocialSummary) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetScheduled

`func (o *SocialSocialSummary) GetScheduled() int64`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *SocialSocialSummary) GetScheduledOk() (*int64, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *SocialSocialSummary) SetScheduled(v int64)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *SocialSocialSummary) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


