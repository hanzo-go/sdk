# TemplatePublishKitIn

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | Category groups the kit in the gallery browser. | [optional] 
**Demo** | Pointer to **string** | Demo is the deployed site itself, when there is one. | [optional] 
**Description** | Pointer to **string** | Description is the browse-card blurb, max 4096 characters. | [optional] 
**Features** | Pointer to **[]string** | Features are the highlights the card lists, at most 32. | [optional] 
**Framework** | Pointer to **string** | Framework is the stack the kit is built on (\&quot;Next.js 14\&quot;). | [optional] 
**Preview** | Pointer to **string** | Preview is the still image the browse card renders, max 4096 characters. | [optional] 
**Slug** | Pointer to **string** | Slug is the kit&#39;s identity — lowercase alphanumeric with dashes, max 40. | [optional] 
**Source** | Pointer to **string** | Source is the repository the kit is forked from, max 4096 characters. | [optional] 
**Title** | Pointer to **string** | Title is the display name. Required, max 200 characters. | [optional] 
**UseCase** | Pointer to **string** | UseCase is what the kit is for, in a phrase. | [optional] 
**Variants** | Pointer to [**[]TemplateVariant**](TemplateVariant.md) | Variants are the shapes this kit ships in, at most 32; the fork picks one. | [optional] 

## Methods

### NewTemplatePublishKitIn

`func NewTemplatePublishKitIn() *TemplatePublishKitIn`

NewTemplatePublishKitIn instantiates a new TemplatePublishKitIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTemplatePublishKitInWithDefaults

`func NewTemplatePublishKitInWithDefaults() *TemplatePublishKitIn`

NewTemplatePublishKitInWithDefaults instantiates a new TemplatePublishKitIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *TemplatePublishKitIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *TemplatePublishKitIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *TemplatePublishKitIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *TemplatePublishKitIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDemo

`func (o *TemplatePublishKitIn) GetDemo() string`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *TemplatePublishKitIn) GetDemoOk() (*string, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *TemplatePublishKitIn) SetDemo(v string)`

SetDemo sets Demo field to given value.

### HasDemo

`func (o *TemplatePublishKitIn) HasDemo() bool`

HasDemo returns a boolean if a field has been set.

### GetDescription

`func (o *TemplatePublishKitIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TemplatePublishKitIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TemplatePublishKitIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TemplatePublishKitIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *TemplatePublishKitIn) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *TemplatePublishKitIn) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *TemplatePublishKitIn) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *TemplatePublishKitIn) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFramework

`func (o *TemplatePublishKitIn) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *TemplatePublishKitIn) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *TemplatePublishKitIn) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *TemplatePublishKitIn) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetPreview

`func (o *TemplatePublishKitIn) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *TemplatePublishKitIn) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *TemplatePublishKitIn) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *TemplatePublishKitIn) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSlug

`func (o *TemplatePublishKitIn) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *TemplatePublishKitIn) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *TemplatePublishKitIn) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *TemplatePublishKitIn) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *TemplatePublishKitIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *TemplatePublishKitIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *TemplatePublishKitIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *TemplatePublishKitIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *TemplatePublishKitIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *TemplatePublishKitIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *TemplatePublishKitIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *TemplatePublishKitIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUseCase

`func (o *TemplatePublishKitIn) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *TemplatePublishKitIn) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *TemplatePublishKitIn) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *TemplatePublishKitIn) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### GetVariants

`func (o *TemplatePublishKitIn) GetVariants() []TemplateVariant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *TemplatePublishKitIn) GetVariantsOk() (*[]TemplateVariant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *TemplatePublishKitIn) SetVariants(v []TemplateVariant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *TemplatePublishKitIn) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


