# StarterKit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | groups the kit in the gallery browser (\&quot;Portfolio\&quot;, \&quot;SaaS\&quot;) | [optional] 
**Demo** | Pointer to **string** | live demo (&lt;slug&gt;.hanzo.app), when deployed | [optional] 
**Description** | Pointer to **string** | the browse-card blurb | [optional] 
**Features** | Pointer to **[]string** | the highlights the card lists, at most 32 | [optional] 
**Framework** | Pointer to **string** | the stack the kit is built on (\&quot;Next.js 14.2 + TS\&quot;) | [optional] 
**Org** | Pointer to **string** | owner of a PRIVATE template; empty in the public catalog | [optional] 
**Preview** | Pointer to **string** | the still image the browse card renders | [optional] 
**Rating** | Pointer to **float64** | Rating is public-gallery curation, on the same terms as Tier: catalog-only, never accepted from a request, absent on a customer&#39;s own kit. | [optional] 
**Slug** | Pointer to **string** | the kit&#39;s identity — lowercase alphanumeric with dashes, max 40 | [optional] 
**Source** | Pointer to **string** | the repository the kit is forked from | [optional] 
**Tier** | Pointer to **int64** | Tier is public-gallery curation, carried verbatim from the embedded catalog. No request can set it — neither write body has the field and neither builds a kit carrying one — so it is absent on every customer-published kit. | [optional] 
**Title** | Pointer to **string** | display name | [optional] 
**UseCase** | Pointer to **string** | what the kit is for, in a phrase | [optional] 
**Variants** | Pointer to [**[]Variant**](Variant.md) | the shapes this template ships in | [optional] 

## Methods

### NewStarterKit

`func NewStarterKit() *StarterKit`

NewStarterKit instantiates a new StarterKit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStarterKitWithDefaults

`func NewStarterKitWithDefaults() *StarterKit`

NewStarterKitWithDefaults instantiates a new StarterKit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *StarterKit) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *StarterKit) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *StarterKit) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *StarterKit) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetDemo

`func (o *StarterKit) GetDemo() string`

GetDemo returns the Demo field if non-nil, zero value otherwise.

### GetDemoOk

`func (o *StarterKit) GetDemoOk() (*string, bool)`

GetDemoOk returns a tuple with the Demo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDemo

`func (o *StarterKit) SetDemo(v string)`

SetDemo sets Demo field to given value.

### HasDemo

`func (o *StarterKit) HasDemo() bool`

HasDemo returns a boolean if a field has been set.

### GetDescription

`func (o *StarterKit) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *StarterKit) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *StarterKit) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *StarterKit) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *StarterKit) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *StarterKit) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *StarterKit) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *StarterKit) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetFramework

`func (o *StarterKit) GetFramework() string`

GetFramework returns the Framework field if non-nil, zero value otherwise.

### GetFrameworkOk

`func (o *StarterKit) GetFrameworkOk() (*string, bool)`

GetFrameworkOk returns a tuple with the Framework field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFramework

`func (o *StarterKit) SetFramework(v string)`

SetFramework sets Framework field to given value.

### HasFramework

`func (o *StarterKit) HasFramework() bool`

HasFramework returns a boolean if a field has been set.

### GetOrg

`func (o *StarterKit) GetOrg() string`

GetOrg returns the Org field if non-nil, zero value otherwise.

### GetOrgOk

`func (o *StarterKit) GetOrgOk() (*string, bool)`

GetOrgOk returns a tuple with the Org field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrg

`func (o *StarterKit) SetOrg(v string)`

SetOrg sets Org field to given value.

### HasOrg

`func (o *StarterKit) HasOrg() bool`

HasOrg returns a boolean if a field has been set.

### GetPreview

`func (o *StarterKit) GetPreview() string`

GetPreview returns the Preview field if non-nil, zero value otherwise.

### GetPreviewOk

`func (o *StarterKit) GetPreviewOk() (*string, bool)`

GetPreviewOk returns a tuple with the Preview field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreview

`func (o *StarterKit) SetPreview(v string)`

SetPreview sets Preview field to given value.

### HasPreview

`func (o *StarterKit) HasPreview() bool`

HasPreview returns a boolean if a field has been set.

### GetRating

`func (o *StarterKit) GetRating() float64`

GetRating returns the Rating field if non-nil, zero value otherwise.

### GetRatingOk

`func (o *StarterKit) GetRatingOk() (*float64, bool)`

GetRatingOk returns a tuple with the Rating field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRating

`func (o *StarterKit) SetRating(v float64)`

SetRating sets Rating field to given value.

### HasRating

`func (o *StarterKit) HasRating() bool`

HasRating returns a boolean if a field has been set.

### GetSlug

`func (o *StarterKit) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *StarterKit) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *StarterKit) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *StarterKit) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### GetSource

`func (o *StarterKit) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *StarterKit) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *StarterKit) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *StarterKit) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTier

`func (o *StarterKit) GetTier() int64`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *StarterKit) GetTierOk() (*int64, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *StarterKit) SetTier(v int64)`

SetTier sets Tier field to given value.

### HasTier

`func (o *StarterKit) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetTitle

`func (o *StarterKit) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *StarterKit) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *StarterKit) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *StarterKit) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetUseCase

`func (o *StarterKit) GetUseCase() string`

GetUseCase returns the UseCase field if non-nil, zero value otherwise.

### GetUseCaseOk

`func (o *StarterKit) GetUseCaseOk() (*string, bool)`

GetUseCaseOk returns a tuple with the UseCase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseCase

`func (o *StarterKit) SetUseCase(v string)`

SetUseCase sets UseCase field to given value.

### HasUseCase

`func (o *StarterKit) HasUseCase() bool`

HasUseCase returns a boolean if a field has been set.

### GetVariants

`func (o *StarterKit) GetVariants() []Variant`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *StarterKit) GetVariantsOk() (*[]Variant, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *StarterKit) SetVariants(v []Variant)`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *StarterKit) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


