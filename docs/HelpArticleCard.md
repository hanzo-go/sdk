# HelpArticleCard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category is the name of the knowledge-base section the article sits in, or empty when it is filed under none. | [optional] 
**Excerpt** | Pointer to **string** | Excerpt is the short summary the author wrote for listings, or empty. | [optional] 
**Slug** | Pointer to **string** | Slug is the article&#39;s stable public identifier and the path segment /v1/help/articles/{slug} addresses it by. | [optional] 
**Title** | Pointer to **string** | Title is the article&#39;s headline. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the unix second the article was last written, in the help center&#39;s own store. | [optional] 

## Methods

### NewHelpArticleCard

`func NewHelpArticleCard() *HelpArticleCard`

NewHelpArticleCard instantiates a new HelpArticleCard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpArticleCardWithDefaults

`func NewHelpArticleCardWithDefaults() *HelpArticleCard`

NewHelpArticleCardWithDefaults instantiates a new HelpArticleCard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *HelpArticleCard) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HelpArticleCard) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HelpArticleCard) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HelpArticleCard) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExcerpt

`func (o *HelpArticleCard) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *HelpArticleCard) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *HelpArticleCard) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *HelpArticleCard) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetSlug

`func (o *HelpArticleCard) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *HelpArticleCard) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *HelpArticleCard) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *HelpArticleCard) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *HelpArticleCard) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HelpArticleCard) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HelpArticleCard) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HelpArticleCard) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HelpArticleCard) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HelpArticleCard) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HelpArticleCard) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HelpArticleCard) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


