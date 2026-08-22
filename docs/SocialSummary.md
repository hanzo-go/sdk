# SocialSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **int32** | Accounts is how many accounts the org has connected, in any status. | [optional] 
**Posts** | Pointer to **int32** | Posts is how many posts the org has, in any state. | [optional] 
**Published** | Pointer to **int32** | Published is how many of them have published. | [optional] 
**Scheduled** | Pointer to **int32** | Scheduled is how many of them are waiting for their scheduled time. | [optional] 

## Methods

### NewSocialSummary

`func NewSocialSummary() *SocialSummary`

NewSocialSummary instantiates a new SocialSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSummaryWithDefaults

`func NewSocialSummaryWithDefaults() *SocialSummary`

NewSocialSummaryWithDefaults instantiates a new SocialSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *SocialSummary) GetAccounts() int32`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SocialSummary) GetAccountsOk() (*int32, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SocialSummary) SetAccounts(v int32)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *SocialSummary) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetPosts

`func (o *SocialSummary) GetPosts() int32`

GetPosts returns the Posts field if non-nil, zero value otherwise.

### GetPostsOk

`func (o *SocialSummary) GetPostsOk() (*int32, bool)`

GetPostsOk returns a tuple with the Posts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosts

`func (o *SocialSummary) SetPosts(v int32)`

SetPosts sets Posts field to given value.

### HasPosts

`func (o *SocialSummary) HasPosts() bool`

HasPosts returns a boolean if a field has been set.

### GetPublished

`func (o *SocialSummary) GetPublished() int32`

GetPublished returns the Published field if non-nil, zero value otherwise.

### GetPublishedOk

`func (o *SocialSummary) GetPublishedOk() (*int32, bool)`

GetPublishedOk returns a tuple with the Published field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublished

`func (o *SocialSummary) SetPublished(v int32)`

SetPublished sets Published field to given value.

### HasPublished

`func (o *SocialSummary) HasPublished() bool`

HasPublished returns a boolean if a field has been set.

### GetScheduled

`func (o *SocialSummary) GetScheduled() int32`

GetScheduled returns the Scheduled field if non-nil, zero value otherwise.

### GetScheduledOk

`func (o *SocialSummary) GetScheduledOk() (*int32, bool)`

GetScheduledOk returns a tuple with the Scheduled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduled

`func (o *SocialSummary) SetScheduled(v int32)`

SetScheduled sets Scheduled field to given value.

### HasScheduled

`func (o *SocialSummary) HasScheduled() bool`

HasScheduled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


