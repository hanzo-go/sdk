# AiModelInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Access** | Pointer to [**AiModelAccessInfo**](AiModelAccessInfo.md) |  | [optional] 
**ContextWindow** | Pointer to **int32** |  | [optional] 
**Created** | Pointer to **int32** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**MaxOutputTokens** | Pointer to **int32** |  | [optional] 
**Object** | Pointer to **string** |  | [optional] 
**Outputs** | Pointer to **[]string** |  | [optional] 
**OwnedBy** | Pointer to **string** |  | [optional] 
**Premium** | Pointer to **bool** |  | [optional] 
**Pricing** | Pointer to [**AiModelPricingInfo**](AiModelPricingInfo.md) |  | [optional] 
**Provider** | Pointer to **string** |  | [optional] 
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


