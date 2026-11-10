# AiModel

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Capabilities** | Pointer to **[]string** | Capabilities are tools, vision and reasoning, those the model supports. | [optional] 
**Class** | Pointer to **string** | Class is premium, ours or free. | [optional] 
**ContextWindow** | Pointer to **int64** | ContextWindow is the most tokens the model reads at once, absent when the catalog does not say. | [optional] 
**Expires** | Pointer to **string** | Expires is the day the model&#39;s vendor stops serving it (2006-01-02), absent when it states none. | [optional] 
**Family** | Pointer to **string** | Family is the Hanzo family (enso, zen, kai, zoo), absent for a third-party model. | [optional] 
**Free** | Pointer to **bool** | Free is whether a call bills nothing. | [optional] 
**Id** | Pointer to **string** | ID is the model&#39;s id, the one a completion or a decision names. | [optional] 
**Inputs** | Pointer to **[]string** | Inputs are the modalities the model reads. | [optional] 
**MaxOutputTokens** | Pointer to **int64** | MaxOutputTokens is the most one answer may hold, absent when the catalog does not say. | [optional] 
**Name** | Pointer to **string** | Name is the model&#39;s display name. | [optional] 
**Outputs** | Pointer to **[]string** | Outputs are the modalities the model writes. | [optional] 
**Pricing** | Pointer to [**AiPrice**](AiPrice.md) | Pricing is the model&#39;s list price. | [optional] 
**Zdr** | Pointer to **bool** | ZDR is whether the model is eligible for Zero Data Retention (non-free, non-data-shared commercial route). | [optional] 

## Methods

### NewAiModel

`func NewAiModel() *AiModel`

NewAiModel instantiates a new AiModel object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAiModelWithDefaults

`func NewAiModelWithDefaults() *AiModel`

NewAiModelWithDefaults instantiates a new AiModel object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCapabilities

`func (o *AiModel) GetCapabilities() []string`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *AiModel) GetCapabilitiesOk() (*[]string, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *AiModel) SetCapabilities(v []string)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *AiModel) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetClass

`func (o *AiModel) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *AiModel) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *AiModel) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *AiModel) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetContextWindow

`func (o *AiModel) GetContextWindow() int64`

GetContextWindow returns the ContextWindow field if non-nil, zero value otherwise.

### GetContextWindowOk

`func (o *AiModel) GetContextWindowOk() (*int64, bool)`

GetContextWindowOk returns a tuple with the ContextWindow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextWindow

`func (o *AiModel) SetContextWindow(v int64)`

SetContextWindow sets ContextWindow field to given value.

### HasContextWindow

`func (o *AiModel) HasContextWindow() bool`

HasContextWindow returns a boolean if a field has been set.

### GetExpires

`func (o *AiModel) GetExpires() string`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *AiModel) GetExpiresOk() (*string, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *AiModel) SetExpires(v string)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *AiModel) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetFamily

`func (o *AiModel) GetFamily() string`

GetFamily returns the Family field if non-nil, zero value otherwise.

### GetFamilyOk

`func (o *AiModel) GetFamilyOk() (*string, bool)`

GetFamilyOk returns a tuple with the Family field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFamily

`func (o *AiModel) SetFamily(v string)`

SetFamily sets Family field to given value.

### HasFamily

`func (o *AiModel) HasFamily() bool`

HasFamily returns a boolean if a field has been set.

### GetFree

`func (o *AiModel) GetFree() bool`

GetFree returns the Free field if non-nil, zero value otherwise.

### GetFreeOk

`func (o *AiModel) GetFreeOk() (*bool, bool)`

GetFreeOk returns a tuple with the Free field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFree

`func (o *AiModel) SetFree(v bool)`

SetFree sets Free field to given value.

### HasFree

`func (o *AiModel) HasFree() bool`

HasFree returns a boolean if a field has been set.

### GetId

`func (o *AiModel) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AiModel) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AiModel) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *AiModel) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInputs

`func (o *AiModel) GetInputs() []string`

GetInputs returns the Inputs field if non-nil, zero value otherwise.

### GetInputsOk

`func (o *AiModel) GetInputsOk() (*[]string, bool)`

GetInputsOk returns a tuple with the Inputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputs

`func (o *AiModel) SetInputs(v []string)`

SetInputs sets Inputs field to given value.

### HasInputs

`func (o *AiModel) HasInputs() bool`

HasInputs returns a boolean if a field has been set.

### GetMaxOutputTokens

`func (o *AiModel) GetMaxOutputTokens() int64`

GetMaxOutputTokens returns the MaxOutputTokens field if non-nil, zero value otherwise.

### GetMaxOutputTokensOk

`func (o *AiModel) GetMaxOutputTokensOk() (*int64, bool)`

GetMaxOutputTokensOk returns a tuple with the MaxOutputTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOutputTokens

`func (o *AiModel) SetMaxOutputTokens(v int64)`

SetMaxOutputTokens sets MaxOutputTokens field to given value.

### HasMaxOutputTokens

`func (o *AiModel) HasMaxOutputTokens() bool`

HasMaxOutputTokens returns a boolean if a field has been set.

### GetName

`func (o *AiModel) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AiModel) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AiModel) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AiModel) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutputs

`func (o *AiModel) GetOutputs() []string`

GetOutputs returns the Outputs field if non-nil, zero value otherwise.

### GetOutputsOk

`func (o *AiModel) GetOutputsOk() (*[]string, bool)`

GetOutputsOk returns a tuple with the Outputs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputs

`func (o *AiModel) SetOutputs(v []string)`

SetOutputs sets Outputs field to given value.

### HasOutputs

`func (o *AiModel) HasOutputs() bool`

HasOutputs returns a boolean if a field has been set.

### GetPricing

`func (o *AiModel) GetPricing() AiPrice`

GetPricing returns the Pricing field if non-nil, zero value otherwise.

### GetPricingOk

`func (o *AiModel) GetPricingOk() (*AiPrice, bool)`

GetPricingOk returns a tuple with the Pricing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPricing

`func (o *AiModel) SetPricing(v AiPrice)`

SetPricing sets Pricing field to given value.

### HasPricing

`func (o *AiModel) HasPricing() bool`

HasPricing returns a boolean if a field has been set.

### GetZdr

`func (o *AiModel) GetZdr() bool`

GetZdr returns the Zdr field if non-nil, zero value otherwise.

### GetZdrOk

`func (o *AiModel) GetZdrOk() (*bool, bool)`

GetZdrOk returns a tuple with the Zdr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZdr

`func (o *AiModel) SetZdr(v bool)`

SetZdr sets Zdr field to given value.

### HasZdr

`func (o *AiModel) HasZdr() bool`

HasZdr returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


