# PricingModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Model identifier | [optional] 
**FullName** | Pointer to **string** | Human-readable name | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Features** | Pointer to **[]string** |  | [optional] 
**Tier** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **int32** | Context window in tokens | [optional] 
**Specs** | Pointer to [**PricingModelSpecs**](PricingModelSpecs.md) |  | [optional] 
**Pricing** | Pointer to [**PricingModelPricing**](PricingModelPricing.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** | Provider model ID (third-party models) | [optional] 
**IsFree** | Pointer to **bool** |  | [optional] 
**Featured** | Pointer to **bool** |  | [optional] 
**PricingUnit** | Pointer to **string** | Non-token pricing unit (minute, image, step) | [optional] 

## Methods

### NewPricingModel

`func NewPricingModel() *PricingModel`

NewPricingModel instantiates a new PricingModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingModelWithDefaults

`func NewPricingModelWithDefaults() *PricingModel`

NewPricingModelWithDefaults instantiates a new PricingModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PricingModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullName

`func (o *PricingModel) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PricingModel) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PricingModel) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *PricingModel) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *PricingModel) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingModel) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingModel) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingModel) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *PricingModel) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PricingModel) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PricingModel) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PricingModel) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetTier

`func (o *PricingModel) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PricingModel) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PricingModel) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PricingModel) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetContext

`func (o *PricingModel) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PricingModel) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PricingModel) SetContext(v int32)`

SetContext sets Context field to given value.

### HasContext

`func (o *PricingModel) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSpecs

`func (o *PricingModel) GetSpecs() PricingModelSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *PricingModel) GetSpecsOk() (*PricingModelSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *PricingModel) SetSpecs(v PricingModelSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *PricingModel) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetPricing

`func (o *PricingModel) GetPricing() PricingModelPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *PricingModel) GetPricingOk() (*PricingModelPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *PricingModel) SetPricing(v PricingModelPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *PricingModel) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvider

`func (o *PricingModel) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PricingModel) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PricingModel) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PricingModel) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCategory

`func (o *PricingModel) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PricingModel) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PricingModel) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PricingModel) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *PricingModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFree

`func (o *PricingModel) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *PricingModel) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *PricingModel) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *PricingModel) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetFeatured

`func (o *PricingModel) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *PricingModel) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *PricingModel) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *PricingModel) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPricingUnit

`func (o *PricingModel) GetPricingUnit() string`

GetPricingUnit returns the PricingUnit field if non-nil, zero value otherwise.

### GetPricingUnitOk

`func (o *PricingModel) GetPricingUnitOk() (*string, bool)`

GetPricingUnitOk returns a tuple with the PricingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUnit

`func (o *PricingModel) SetPricingUnit(v string)`

SetPricingUnit sets PricingUnit field to given value.

### HasPricingUnit

`func (o *PricingModel) HasPricingUnit() bool`

HasPricingUnit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


