# AiModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiModelAccessInfo**](AiModelAccessInfo.md) |  | [optional] 
**CanonicalSlug** | Pointer to **string** |  | [optional] 
**Class** | Pointer to **string** |  | [optional] 
**ContextWindow** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Family** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Inputs** | Pointer to **[]string** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to **[]string** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Premium** | Pointer to **bool** |  | [optional] 
**Pricing** | Pointer to [**AiModelPricingInfo**](AiModelPricingInfo.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
**SupportsReasoning** | Pointer to **bool** |  | [optional] 
**SupportsTools** | Pointer to **bool** |  | [optional] 
**SupportsVision** | Pointer to **bool** |  | [optional] 

## Methods

### NewAiModelInfo

`func NewAiModelInfo() *AiModelInfo`

NewAiModelInfo instantiates a new AiModelInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelInfoWithDefaults

`func NewAiModelInfoWithDefaults() *AiModelInfo`

NewAiModelInfoWithDefaults instantiates a new AiModelInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccess

`func (o *AiModelInfo) GetAccess() AiModelAccessInfo`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *AiModelInfo) GetAccessOk() (*AiModelAccessInfo, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *AiModelInfo) SetAccess(v AiModelAccessInfo)`

SetAccess sets Access field to given value.

### HasAccess

`func (o *AiModelInfo) HasAccess() bool`

HasAccess returns a boolean if a field has been set.

### GetCanonicalSlug

`func (o *AiModelInfo) GetCanonicalSlug() string`

GetCanonicalSlug returns the CanonicalSlug field if non-nil, zero value otherwise.

### GetCanonicalSlugOk

`func (o *AiModelInfo) GetCanonicalSlugOk() (*string, bool)`

GetCanonicalSlugOk returns a tuple with the CanonicalSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanonicalSlug

`func (o *AiModelInfo) SetCanonicalSlug(v string)`

SetCanonicalSlug sets CanonicalSlug field to given value.

### HasCanonicalSlug

`func (o *AiModelInfo) HasCanonicalSlug() bool`

HasCanonicalSlug returns a boolean if a field has been set.

### GetClass

`func (o *AiModelInfo) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AiModelInfo) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AiModelInfo) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *AiModelInfo) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetContextWindow

`func (o *AiModelInfo) GetContextWindow() int32`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *AiModelInfo) GetContextWindowOk() (*int32, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *AiModelInfo) SetContextWindow(v int32)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *AiModelInfo) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### GetCreated

`func (o *AiModelInfo) GetCreated() int32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AiModelInfo) GetCreatedOk() (*int32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AiModelInfo) SetCreated(v int32)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *AiModelInfo) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *AiModelInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AiModelInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AiModelInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AiModelInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFamily

`func (o *AiModelInfo) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AiModelInfo) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AiModelInfo) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AiModelInfo) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetId

`func (o *AiModelInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiModelInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiModelInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiModelInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *AiModelInfo) GetInputs() []string`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AiModelInfo) GetInputsOk() (*[]string, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AiModelInfo) SetInputs(v []string)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AiModelInfo) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AiModelInfo) GetMaxOutputTokens() int32`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AiModelInfo) GetMaxOutputTokensOk() (*int32, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AiModelInfo) SetMaxOutputTokens(v int32)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AiModelInfo) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetName

`func (o *AiModelInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiModelInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiModelInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiModelInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetObject

`func (o *AiModelInfo) GetObject() string`

GetObject returns the Object field if non-nil, zero value otherwise.

### GetObjectOk

`func (o *AiModelInfo) GetObjectOk() (*string, bool)`

GetObjectOk returns a tuple with the Object field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObject

`func (o *AiModelInfo) SetObject(v string)`

SetObject sets Object field to given value.

### HasObject

`func (o *AiModelInfo) HasObject() bool`

HasObject returns a boolean if a field has been set.

### GetOutputs

`func (o *AiModelInfo) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AiModelInfo) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AiModelInfo) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AiModelInfo) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetOwnedBy

`func (o *AiModelInfo) GetOwnedBy() string`

GetOwnedBy returns the OwnedBy field if non-nil, zero value otherwise.

### GetOwnedByOk

`func (o *AiModelInfo) GetOwnedByOk() (*string, bool)`

GetOwnedByOk returns a tuple with the OwnedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwnedBy

`func (o *AiModelInfo) SetOwnedBy(v string)`

SetOwnedBy sets OwnedBy field to given value.

### HasOwnedBy

`func (o *AiModelInfo) HasOwnedBy() bool`

HasOwnedBy returns a boolean if a field has been set.

### GetPremium

`func (o *AiModelInfo) GetPremium() bool`

GetPremium returns the Premium field if non-nil, zero value otherwise.

### GetPremiumOk

`func (o *AiModelInfo) GetPremiumOk() (*bool, bool)`

GetPremiumOk returns a tuple with the Premium field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPremium

`func (o *AiModelInfo) SetPremium(v bool)`

SetPremium sets Premium field to given value.

### HasPremium

`func (o *AiModelInfo) HasPremium() bool`

HasPremium returns a boolean if a field has been set.

### GetPricing

`func (o *AiModelInfo) GetPricing() AiModelPricingInfo`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AiModelInfo) GetPricingOk() (*AiModelPricingInfo, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AiModelInfo) SetPricing(v AiModelPricingInfo)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AiModelInfo) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetProvider

`func (o *AiModelInfo) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *AiModelInfo) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *AiModelInfo) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *AiModelInfo) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetSupportsReasoning

`func (o *AiModelInfo) GetSupportsReasoning() bool`

GetSupportsReasoning returns the SupportsReasoning field if non-nil, zero value otherwise.

### GetSupportsReasoningOk

`func (o *AiModelInfo) GetSupportsReasoningOk() (*bool, bool)`

GetSupportsReasoningOk returns a tuple with the SupportsReasoning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsReasoning

`func (o *AiModelInfo) SetSupportsReasoning(v bool)`

SetSupportsReasoning sets SupportsReasoning field to given value.

### HasSupportsReasoning

`func (o *AiModelInfo) HasSupportsReasoning() bool`

HasSupportsReasoning returns a boolean if a field has been set.

### GetSupportsTools

`func (o *AiModelInfo) GetSupportsTools() bool`

GetSupportsTools returns the SupportsTools field if non-nil, zero value otherwise.

### GetSupportsToolsOk

`func (o *AiModelInfo) GetSupportsToolsOk() (*bool, bool)`

GetSupportsToolsOk returns a tuple with the SupportsTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsTools

`func (o *AiModelInfo) SetSupportsTools(v bool)`

SetSupportsTools sets SupportsTools field to given value.

### HasSupportsTools

`func (o *AiModelInfo) HasSupportsTools() bool`

HasSupportsTools returns a boolean if a field has been set.

### GetSupportsVision

`func (o *AiModelInfo) GetSupportsVision() bool`

GetSupportsVision returns the SupportsVision field if non-nil, zero value otherwise.

### GetSupportsVisionOk

`func (o *AiModelInfo) GetSupportsVisionOk() (*bool, bool)`

GetSupportsVisionOk returns a tuple with the SupportsVision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVision

`func (o *AiModelInfo) SetSupportsVision(v bool)`

SetSupportsVision sets SupportsVision field to given value.

### HasSupportsVision

`func (o *AiModelInfo) HasSupportsVision() bool`

HasSupportsVision returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


