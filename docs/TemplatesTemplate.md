# TemplatesTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Framework** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**UseCase** | Pointer to **string** |  | [optional] 
**Tier** | Pointer to **int32** | Optional pricing tier | [optional] 
**Rating** | Pointer to **float64** | Optional gallery rating | [optional] 
**Source** | Pointer to **string** | Live gallery fork/deploy URL | [optional] 
**Preview** | Pointer to **string** | Live gallery screenshot/preview URL | [optional] 

## Methods

### NewTemplatesTemplate

`func NewTemplatesTemplate() *TemplatesTemplate`

NewTemplatesTemplate instantiates a new TemplatesTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatesTemplateWithDefaults

`func NewTemplatesTemplateWithDefaults() *TemplatesTemplate`

NewTemplatesTemplateWithDefaults instantiates a new TemplatesTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *TemplatesTemplate) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TemplatesTemplate) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TemplatesTemplate) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TemplatesTemplate) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetTitle

`func (o *TemplatesTemplate) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplatesTemplate) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplatesTemplate) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TemplatesTemplate) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetCategory

`func (o *TemplatesTemplate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TemplatesTemplate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TemplatesTemplate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TemplatesTemplate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDescription

`func (o *TemplatesTemplate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplatesTemplate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplatesTemplate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplatesTemplate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFramework

`func (o *TemplatesTemplate) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TemplatesTemplate) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TemplatesTemplate) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TemplatesTemplate) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetFeatures

`func (o *TemplatesTemplate) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TemplatesTemplate) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TemplatesTemplate) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TemplatesTemplate) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetUseCase

`func (o *TemplatesTemplate) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *TemplatesTemplate) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *TemplatesTemplate) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *TemplatesTemplate) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### GetTier

`func (o *TemplatesTemplate) GetTier() int32`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *TemplatesTemplate) GetTierOk() (*int32, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *TemplatesTemplate) SetTier(v int32)`

SetTier sets Tier field to given value.

### HasTier

`func (o *TemplatesTemplate) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetRating

`func (o *TemplatesTemplate) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *TemplatesTemplate) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *TemplatesTemplate) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *TemplatesTemplate) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetSource

`func (o *TemplatesTemplate) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TemplatesTemplate) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TemplatesTemplate) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TemplatesTemplate) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetPreview

`func (o *TemplatesTemplate) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *TemplatesTemplate) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *TemplatesTemplate) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *TemplatesTemplate) HasPreview() bool`

HasPreview returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


