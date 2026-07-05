# PricingModelListResponseDataInner

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
**Object** | Pointer to **string** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 

## Methods

### NewPricingModelListResponseDataInner

`func NewPricingModelListResponseDataInner() *PricingModelListResponseDataInner`

NewPricingModelListResponseDataInner instantiates a new PricingModelListResponseDataInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPricingModelListResponseDataInnerWithDefaults

`func NewPricingModelListResponseDataInnerWithDefaults() *PricingModelListResponseDataInner`

NewPricingModelListResponseDataInnerWithDefaults instantiates a new PricingModelListResponseDataInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PricingModelListResponseDataInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PricingModelListResponseDataInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PricingModelListResponseDataInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PricingModelListResponseDataInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFullName

`func (o *PricingModelListResponseDataInner) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *PricingModelListResponseDataInner) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *PricingModelListResponseDataInner) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *PricingModelListResponseDataInner) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetDescription

`func (o *PricingModelListResponseDataInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *PricingModelListResponseDataInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *PricingModelListResponseDataInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *PricingModelListResponseDataInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFeatures

`func (o *PricingModelListResponseDataInner) GetFeatures() []string`

GetFeatures returns the Features field if non-nil, zero value otherwise.

### GetFeaturesOk

`func (o *PricingModelListResponseDataInner) GetFeaturesOk() (*[]string, bool)`

GetFeaturesOk returns a tuple with the Features field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatures

`func (o *PricingModelListResponseDataInner) SetFeatures(v []string)`

SetFeatures sets Features field to given value.

### HasFeatures

`func (o *PricingModelListResponseDataInner) HasFeatures() bool`

HasFeatures returns a boolean if a field has been set.

### GetTier

`func (o *PricingModelListResponseDataInner) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *PricingModelListResponseDataInner) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *PricingModelListResponseDataInner) SetTier(v string)`

SetTier sets Tier field to given value.

### HasTier

`func (o *PricingModelListResponseDataInner) HasTier() bool`

HasTier returns a boolean if a field has been set.

### GetContext

`func (o *PricingModelListResponseDataInner) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *PricingModelListResponseDataInner) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *PricingModelListResponseDataInner) SetContext(v int32)`

SetContext sets Context field to given value.

### HasContext

`func (o *PricingModelListResponseDataInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetSpecs

`func (o *PricingModelListResponseDataInner) GetSpecs() PricingModelSpecs`

GetSpecs returns the Specs field if non-nil, zero value otherwise.

### GetSpecsOk

`func (o *PricingModelListResponseDataInner) GetSpecsOk() (*PricingModelSpecs, bool)`

GetSpecsOk returns a tuple with the Specs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecs

`func (o *PricingModelListResponseDataInner) SetSpecs(v PricingModelSpecs)`

SetSpecs sets Specs field to given value.

### HasSpecs

`func (o *PricingModelListResponseDataInner) HasSpecs() bool`

HasSpecs returns a boolean if a field has been set.

### GetPricing

`func (o *PricingModelListResponseDataInner) GetPricing() PricingModelPricing`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *PricingModelListResponseDataInner) GetPricingOk() (*PricingModelPricing, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *PricingModelListResponseDataInner) SetPricing(v PricingModelPricing)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *PricingModelListResponseDataInner) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvider

`func (o *PricingModelListResponseDataInner) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *PricingModelListResponseDataInner) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *PricingModelListResponseDataInner) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *PricingModelListResponseDataInner) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetCategory

`func (o *PricingModelListResponseDataInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *PricingModelListResponseDataInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *PricingModelListResponseDataInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *PricingModelListResponseDataInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetId

`func (o *PricingModelListResponseDataInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PricingModelListResponseDataInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PricingModelListResponseDataInner) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PricingModelListResponseDataInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFree

`func (o *PricingModelListResponseDataInner) GetIsFree() bool`

GetIsFree returns the IsFree field if non-nil, zero value otherwise.

### GetIsFreeOk

`func (o *PricingModelListResponseDataInner) GetIsFreeOk() (*bool, bool)`

GetIsFreeOk returns a tuple with the IsFree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFree

`func (o *PricingModelListResponseDataInner) SetIsFree(v bool)`

SetIsFree sets IsFree field to given value.

### HasIsFree

`func (o *PricingModelListResponseDataInner) HasIsFree() bool`

HasIsFree returns a boolean if a field has been set.

### GetFeatured

`func (o *PricingModelListResponseDataInner) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *PricingModelListResponseDataInner) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *PricingModelListResponseDataInner) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.

### HasFeatured

`func (o *PricingModelListResponseDataInner) HasFeatured() bool`

HasFeatured returns a boolean if a field has been set.

### GetPricingUnit

`func (o *PricingModelListResponseDataInner) GetPricingUnit() string`

GetPricingUnit returns the PricingUnit field if non-nil, zero value otherwise.

### GetPricingUnitOk

`func (o *PricingModelListResponseDataInner) GetPricingUnitOk() (*string, bool)`

GetPricingUnitOk returns a tuple with the PricingUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricingUnit

`func (o *PricingModelListResponseDataInner) SetPricingUnit(v string)`

SetPricingUnit sets PricingUnit field to given value.

### HasPricingUnit

`func (o *PricingModelListResponseDataInner) HasPricingUnit() bool`

HasPricingUnit returns a boolean if a field has been set.

### GetObject

`func (o *PricingModelListResponseDataInner) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *PricingModelListResponseDataInner) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *PricingModelListResponseDataInner) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *PricingModelListResponseDataInner) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOwnedBy

`func (o *PricingModelListResponseDataInner) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *PricingModelListResponseDataInner) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *PricingModelListResponseDataInner) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *PricingModelListResponseDataInner) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


