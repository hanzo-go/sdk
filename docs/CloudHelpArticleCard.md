# CloudHelpArticleCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the name of the knowledge-base section the article sits in, or empty when it is filed under none. | [optional] 
**Excerpt** | Pointer to **string** | Excerpt is the short summary the author wrote for listings, or empty. | [optional] 
**Slug** | Pointer to **string** | Slug is the article&#39;s stable public identifier and the path segment /v1/help/articles/{slug} addresses it by. | [optional] 
**Title** | Pointer to **string** | Title is the article&#39;s headline. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is the unix second the article was last written, in the help center&#39;s own store. | [optional] 

## Methods

### NewCloudHelpArticleCard

`func NewCloudHelpArticleCard() *CloudHelpArticleCard`

NewCloudHelpArticleCard instantiates a new CloudHelpArticleCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudHelpArticleCardWithDefaults

`func NewCloudHelpArticleCardWithDefaults() *CloudHelpArticleCard`

NewCloudHelpArticleCardWithDefaults instantiates a new CloudHelpArticleCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CloudHelpArticleCard) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CloudHelpArticleCard) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CloudHelpArticleCard) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CloudHelpArticleCard) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExcerpt

`func (o *CloudHelpArticleCard) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *CloudHelpArticleCard) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *CloudHelpArticleCard) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *CloudHelpArticleCard) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetSlug

`func (o *CloudHelpArticleCard) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CloudHelpArticleCard) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CloudHelpArticleCard) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CloudHelpArticleCard) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *CloudHelpArticleCard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudHelpArticleCard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudHelpArticleCard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudHelpArticleCard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *CloudHelpArticleCard) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CloudHelpArticleCard) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CloudHelpArticleCard) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *CloudHelpArticleCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


