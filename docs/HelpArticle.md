# HelpArticle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **string** | Body is the article&#39;s rich-text content as the author saved it. | [optional] 
**Category** | Pointer to **string** | Category is the name of the knowledge-base section the article sits in, or empty when it is filed under none. | [optional] 
**Excerpt** | Pointer to **string** | Excerpt is the short summary the author wrote for listings, or empty. | [optional] 
**Slug** | Pointer to **string** | Slug is the article&#39;s stable public identifier — the path segment it was addressed by. | [optional] 
**Title** | Pointer to **string** | Title is the article&#39;s headline. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is the unix second the article was last written. | [optional] 

## Methods

### NewHelpArticle

`func NewHelpArticle() *HelpArticle`

NewHelpArticle instantiates a new HelpArticle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHelpArticleWithDefaults

`func NewHelpArticleWithDefaults() *HelpArticle`

NewHelpArticleWithDefaults instantiates a new HelpArticle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *HelpArticle) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *HelpArticle) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *HelpArticle) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *HelpArticle) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetCategory

`func (o *HelpArticle) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *HelpArticle) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *HelpArticle) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *HelpArticle) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetExcerpt

`func (o *HelpArticle) GetExcerpt() string`

GetExcerpt returns the Excerpt field if non-nil, zero value otherwise.

### GetExcerptOk

`func (o *HelpArticle) GetExcerptOk() (*string, bool)`

GetExcerptOk returns a tuple with the Excerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcerpt

`func (o *HelpArticle) SetExcerpt(v string)`

SetExcerpt sets Excerpt field to given value.

### HasExcerpt

`func (o *HelpArticle) HasExcerpt() bool`

HasExcerpt returns a boolean if a field has been set.

### GetSlug

`func (o *HelpArticle) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *HelpArticle) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *HelpArticle) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *HelpArticle) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *HelpArticle) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *HelpArticle) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *HelpArticle) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *HelpArticle) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *HelpArticle) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *HelpArticle) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *HelpArticle) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *HelpArticle) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


