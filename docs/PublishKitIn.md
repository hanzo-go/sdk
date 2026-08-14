# PublishKitIn

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
**Variants** | Pointer to [**[]Variant**](Variant.md) | Variants are the shapes this kit ships in, at most 32; the fork picks one. | [optional] 

## Methods

### NewPublishKitIn

`func NewPublishKitIn() *PublishKitIn`

NewPublishKitIn instantiates a new PublishKitIn object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPublishKitInWithDefaults

`func NewPublishKitInWithDefaults() *PublishKitIn`

NewPublishKitInWithDefaults instantiates a new PublishKitIn object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *PublishKitIn) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PublishKitIn) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PublishKitIn) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PublishKitIn) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDemo

`func (o *PublishKitIn) GetDemo() string`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *PublishKitIn) GetDemoOk() (*string, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *PublishKitIn) SetDemo(v string)`

SetDemo sets Demo field to given value.

### HasDemo

`func (o *PublishKitIn) HasDemo() bool`

HasDemo returns a boolean if a field has been set.

### GetDescription

`func (o *PublishKitIn) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PublishKitIn) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PublishKitIn) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PublishKitIn) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *PublishKitIn) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PublishKitIn) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PublishKitIn) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PublishKitIn) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFramework

`func (o *PublishKitIn) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *PublishKitIn) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *PublishKitIn) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *PublishKitIn) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetPreview

`func (o *PublishKitIn) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *PublishKitIn) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *PublishKitIn) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *PublishKitIn) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetSlug

`func (o *PublishKitIn) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *PublishKitIn) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *PublishKitIn) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *PublishKitIn) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *PublishKitIn) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *PublishKitIn) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *PublishKitIn) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *PublishKitIn) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTitle

`func (o *PublishKitIn) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *PublishKitIn) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *PublishKitIn) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *PublishKitIn) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUseCase

`func (o *PublishKitIn) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *PublishKitIn) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *PublishKitIn) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *PublishKitIn) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### GetVariants

`func (o *PublishKitIn) GetVariants() []Variant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *PublishKitIn) GetVariantsOk() (*[]Variant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *PublishKitIn) SetVariants(v []Variant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *PublishKitIn) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


