# SocialSocialPosts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to [**[]SocialSocialPost**](SocialSocialPost.md) | Data is the posts, most-recently-updated first, bounded by the limit. | [optional] 

## Methods

### NewSocialSocialPosts

`func NewSocialSocialPosts() *SocialSocialPosts`

NewSocialSocialPosts instantiates a new SocialSocialPosts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialPostsWithDefaults

`func NewSocialSocialPostsWithDefaults() *SocialSocialPosts`

NewSocialSocialPostsWithDefaults instantiates a new SocialSocialPosts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *SocialSocialPosts) GetData() []SocialSocialPost`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *SocialSocialPosts) GetDataOk() (*[]SocialSocialPost, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *SocialSocialPosts) SetData(v []SocialSocialPost)`

SetData sets Data field to given value.

### HasData

`func (o *SocialSocialPosts) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


