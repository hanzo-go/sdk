# O11yO11yFeature

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultVariant** | Pointer to **string** | DefaultVariant is the variant used when nothing overrides it. | [optional] 
**Description** | Pointer to **string** | Description says what the flag gates. | [optional] 
**Kind** | Pointer to **string** | Kind is the flag&#39;s value kind, e.g. boolean. | [optional] 
**Name** | Pointer to **string** | Name is the flag&#39;s name. | [optional] 
**ResolvedValue** | Pointer to **map[string]interface{}** | ResolvedValue is the value resolved for the caller&#39;s org. | [optional] 
**Stage** | Pointer to **string** | Stage is the flag&#39;s lifecycle stage, e.g. stable. | [optional] 
**Variants** | Pointer to **map[string]map[string]interface{}** | Variants are the flag&#39;s possible values, by variant name. | [optional] 

## Methods

### NewO11yO11yFeature

`func NewO11yO11yFeature() *O11yO11yFeature`

NewO11yO11yFeature instantiates a new O11yO11yFeature object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yFeatureWithDefaults

`func NewO11yO11yFeatureWithDefaults() *O11yO11yFeature`

NewO11yO11yFeatureWithDefaults instantiates a new O11yO11yFeature object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultVariant

`func (o *O11yO11yFeature) GetDefaultVariant() string`

GetDefaultVariant returns the DefaultVariant field if non-nil, zero value otherwise.

### GetDefaultVariantOk

`func (o *O11yO11yFeature) GetDefaultVariantOk() (*string, bool)`

GetDefaultVariantOk returns a tuple with the DefaultVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVariant

`func (o *O11yO11yFeature) SetDefaultVariant(v string)`

SetDefaultVariant sets DefaultVariant field to given value.

### HasDefaultVariant

`func (o *O11yO11yFeature) HasDefaultVariant() bool`

HasDefaultVariant returns a boolean if a field has been set.

### GetDescription

`func (o *O11yO11yFeature) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *O11yO11yFeature) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *O11yO11yFeature) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *O11yO11yFeature) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetKind

`func (o *O11yO11yFeature) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *O11yO11yFeature) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *O11yO11yFeature) SetKind(v string)`

SetKind sets Kind field to given value.

### HasKind

`func (o *O11yO11yFeature) HasKind() bool`

HasKind returns a boolean if a field has been set.

### GetName

`func (o *O11yO11yFeature) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *O11yO11yFeature) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *O11yO11yFeature) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *O11yO11yFeature) HasName() bool`

HasName returns a boolean if a field has been set.

### GetResolvedValue

`func (o *O11yO11yFeature) GetResolvedValue() map[string]interface{}`

GetResolvedValue returns the ResolvedValue field if non-nil, zero value otherwise.

### GetResolvedValueOk

`func (o *O11yO11yFeature) GetResolvedValueOk() (*map[string]interface{}, bool)`

GetResolvedValueOk returns a tuple with the ResolvedValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResolvedValue

`func (o *O11yO11yFeature) SetResolvedValue(v map[string]interface{})`

SetResolvedValue sets ResolvedValue field to given value.

### HasResolvedValue

`func (o *O11yO11yFeature) HasResolvedValue() bool`

HasResolvedValue returns a boolean if a field has been set.

### GetStage

`func (o *O11yO11yFeature) GetStage() string`

GetStage returns the Stage field if non-nil, zero value otherwise.

### GetStageOk

`func (o *O11yO11yFeature) GetStageOk() (*string, bool)`

GetStageOk returns a tuple with the Stage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStage

`func (o *O11yO11yFeature) SetStage(v string)`

SetStage sets Stage field to given value.

### HasStage

`func (o *O11yO11yFeature) HasStage() bool`

HasStage returns a boolean if a field has been set.

### GetVariants

`func (o *O11yO11yFeature) GetVariants() map[string]map[string]interface{}`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *O11yO11yFeature) GetVariantsOk() (*map[string]map[string]interface{}, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *O11yO11yFeature) SetVariants(v map[string]map[string]interface{})`

SetVariants sets Variants field to given value.

### HasVariants

`func (o *O11yO11yFeature) HasVariants() bool`

HasVariants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


