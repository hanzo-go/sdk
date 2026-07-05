# NexusProvider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**BrowserUrl** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Chain** | Pointer to **string** |  | [optional] 
**ClientId** | Pointer to **string** |  | [optional] 
**ClientSecret** | Pointer to **string** |  | [optional] 
**CompatibleProvider** | Pointer to **string** |  | [optional] 
**ConfigText** | Pointer to **string** |  | [optional] 
**ContractMethod** | Pointer to **string** |  | [optional] 
**ContractName** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EnableThinking** | Pointer to **bool** |  | [optional] 
**Flavor** | Pointer to **string** |  | [optional] 
**FrequencyPenalty** | Pointer to **float32** |  | [optional] 
**InputPricePerThousandTokens** | Pointer to **float64** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**McpTools** | Pointer to [**[]NexusMcpTools**](NexusMcpTools.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**OutputPricePerThousandTokens** | Pointer to **float64** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PresencePenalty** | Pointer to **float32** |  | [optional] 
**ProviderKey** | Pointer to **string** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**SignCert** | Pointer to **string** |  | [optional] 
**SignKey** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**TestContent** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TopK** | Pointer to **int64** |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserCert** | Pointer to **string** |  | [optional] 
**UserKey** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusProvider

`func NewNexusProvider() *NexusProvider`

NewNexusProvider instantiates a new NexusProvider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusProviderWithDefaults

`func NewNexusProviderWithDefaults() *NexusProvider`

NewNexusProviderWithDefaults instantiates a new NexusProvider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *NexusProvider) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *NexusProvider) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *NexusProvider) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *NexusProvider) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetBrowserUrl

`func (o *NexusProvider) GetBrowserUrl() string`

GetBrowserUrl returns the BrowserUrl field if non-nil, zero value otherwise.

### GetBrowserUrlOk

`func (o *NexusProvider) GetBrowserUrlOk() (*string, bool)`

GetBrowserUrlOk returns a tuple with the BrowserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUrl

`func (o *NexusProvider) SetBrowserUrl(v string)`

SetBrowserUrl sets BrowserUrl field to given value.

### HasBrowserUrl

`func (o *NexusProvider) HasBrowserUrl() bool`

HasBrowserUrl returns a boolean if a field has been set.

### GetCategory

`func (o *NexusProvider) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NexusProvider) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NexusProvider) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NexusProvider) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChain

`func (o *NexusProvider) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *NexusProvider) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *NexusProvider) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *NexusProvider) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetClientId

`func (o *NexusProvider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *NexusProvider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *NexusProvider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *NexusProvider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *NexusProvider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *NexusProvider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *NexusProvider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *NexusProvider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCompatibleProvider

`func (o *NexusProvider) GetCompatibleProvider() string`

GetCompatibleProvider returns the CompatibleProvider field if non-nil, zero value otherwise.

### GetCompatibleProviderOk

`func (o *NexusProvider) GetCompatibleProviderOk() (*string, bool)`

GetCompatibleProviderOk returns a tuple with the CompatibleProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleProvider

`func (o *NexusProvider) SetCompatibleProvider(v string)`

SetCompatibleProvider sets CompatibleProvider field to given value.

### HasCompatibleProvider

`func (o *NexusProvider) HasCompatibleProvider() bool`

HasCompatibleProvider returns a boolean if a field has been set.

### GetConfigText

`func (o *NexusProvider) GetConfigText() string`

GetConfigText returns the ConfigText field if non-nil, zero value otherwise.

### GetConfigTextOk

`func (o *NexusProvider) GetConfigTextOk() (*string, bool)`

GetConfigTextOk returns a tuple with the ConfigText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigText

`func (o *NexusProvider) SetConfigText(v string)`

SetConfigText sets ConfigText field to given value.

### HasConfigText

`func (o *NexusProvider) HasConfigText() bool`

HasConfigText returns a boolean if a field has been set.

### GetContractMethod

`func (o *NexusProvider) GetContractMethod() string`

GetContractMethod returns the ContractMethod field if non-nil, zero value otherwise.

### GetContractMethodOk

`func (o *NexusProvider) GetContractMethodOk() (*string, bool)`

GetContractMethodOk returns a tuple with the ContractMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractMethod

`func (o *NexusProvider) SetContractMethod(v string)`

SetContractMethod sets ContractMethod field to given value.

### HasContractMethod

`func (o *NexusProvider) HasContractMethod() bool`

HasContractMethod returns a boolean if a field has been set.

### GetContractName

`func (o *NexusProvider) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *NexusProvider) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *NexusProvider) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *NexusProvider) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusProvider) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusProvider) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusProvider) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusProvider) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *NexusProvider) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NexusProvider) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NexusProvider) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NexusProvider) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusProvider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusProvider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusProvider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusProvider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableThinking

`func (o *NexusProvider) GetEnableThinking() bool`

GetEnableThinking returns the EnableThinking field if non-nil, zero value otherwise.

### GetEnableThinkingOk

`func (o *NexusProvider) GetEnableThinkingOk() (*bool, bool)`

GetEnableThinkingOk returns a tuple with the EnableThinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThinking

`func (o *NexusProvider) SetEnableThinking(v bool)`

SetEnableThinking sets EnableThinking field to given value.

### HasEnableThinking

`func (o *NexusProvider) HasEnableThinking() bool`

HasEnableThinking returns a boolean if a field has been set.

### GetFlavor

`func (o *NexusProvider) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *NexusProvider) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *NexusProvider) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *NexusProvider) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *NexusProvider) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *NexusProvider) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *NexusProvider) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *NexusProvider) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetInputPricePerThousandTokens

`func (o *NexusProvider) GetInputPricePerThousandTokens() float64`

GetInputPricePerThousandTokens returns the InputPricePerThousandTokens field if non-nil, zero value otherwise.

### GetInputPricePerThousandTokensOk

`func (o *NexusProvider) GetInputPricePerThousandTokensOk() (*float64, bool)`

GetInputPricePerThousandTokensOk returns a tuple with the InputPricePerThousandTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerThousandTokens

`func (o *NexusProvider) SetInputPricePerThousandTokens(v float64)`

SetInputPricePerThousandTokens sets InputPricePerThousandTokens field to given value.

### HasInputPricePerThousandTokens

`func (o *NexusProvider) HasInputPricePerThousandTokens() bool`

HasInputPricePerThousandTokens returns a boolean if a field has been set.

### GetIsDefault

`func (o *NexusProvider) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NexusProvider) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NexusProvider) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *NexusProvider) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetMcpTools

`func (o *NexusProvider) GetMcpTools() []NexusMcpTools`

GetMcpTools returns the McpTools field if non-nil, zero value otherwise.

### GetMcpToolsOk

`func (o *NexusProvider) GetMcpToolsOk() (*[]NexusMcpTools, bool)`

GetMcpToolsOk returns a tuple with the McpTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpTools

`func (o *NexusProvider) SetMcpTools(v []NexusMcpTools)`

SetMcpTools sets McpTools field to given value.

### HasMcpTools

`func (o *NexusProvider) HasMcpTools() bool`

HasMcpTools returns a boolean if a field has been set.

### GetName

`func (o *NexusProvider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusProvider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusProvider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusProvider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *NexusProvider) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *NexusProvider) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *NexusProvider) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *NexusProvider) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOutputPricePerThousandTokens

`func (o *NexusProvider) GetOutputPricePerThousandTokens() float64`

GetOutputPricePerThousandTokens returns the OutputPricePerThousandTokens field if non-nil, zero value otherwise.

### GetOutputPricePerThousandTokensOk

`func (o *NexusProvider) GetOutputPricePerThousandTokensOk() (*float64, bool)`

GetOutputPricePerThousandTokensOk returns a tuple with the OutputPricePerThousandTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerThousandTokens

`func (o *NexusProvider) SetOutputPricePerThousandTokens(v float64)`

SetOutputPricePerThousandTokens sets OutputPricePerThousandTokens field to given value.

### HasOutputPricePerThousandTokens

`func (o *NexusProvider) HasOutputPricePerThousandTokens() bool`

HasOutputPricePerThousandTokens returns a boolean if a field has been set.

### GetOwner

`func (o *NexusProvider) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusProvider) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusProvider) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusProvider) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *NexusProvider) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *NexusProvider) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *NexusProvider) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *NexusProvider) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### GetProviderKey

`func (o *NexusProvider) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *NexusProvider) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *NexusProvider) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *NexusProvider) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderUrl

`func (o *NexusProvider) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *NexusProvider) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *NexusProvider) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *NexusProvider) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetRegion

`func (o *NexusProvider) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *NexusProvider) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *NexusProvider) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *NexusProvider) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetSignCert

`func (o *NexusProvider) GetSignCert() string`

GetSignCert returns the SignCert field if non-nil, zero value otherwise.

### GetSignCertOk

`func (o *NexusProvider) GetSignCertOk() (*string, bool)`

GetSignCertOk returns a tuple with the SignCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignCert

`func (o *NexusProvider) SetSignCert(v string)`

SetSignCert sets SignCert field to given value.

### HasSignCert

`func (o *NexusProvider) HasSignCert() bool`

HasSignCert returns a boolean if a field has been set.

### GetSignKey

`func (o *NexusProvider) GetSignKey() string`

GetSignKey returns the SignKey field if non-nil, zero value otherwise.

### GetSignKeyOk

`func (o *NexusProvider) GetSignKeyOk() (*string, bool)`

GetSignKeyOk returns a tuple with the SignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignKey

`func (o *NexusProvider) SetSignKey(v string)`

SetSignKey sets SignKey field to given value.

### HasSignKey

`func (o *NexusProvider) HasSignKey() bool`

HasSignKey returns a boolean if a field has been set.

### GetState

`func (o *NexusProvider) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NexusProvider) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NexusProvider) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NexusProvider) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubType

`func (o *NexusProvider) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *NexusProvider) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *NexusProvider) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *NexusProvider) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTemperature

`func (o *NexusProvider) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *NexusProvider) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *NexusProvider) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *NexusProvider) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTestContent

`func (o *NexusProvider) GetTestContent() string`

GetTestContent returns the TestContent field if non-nil, zero value otherwise.

### GetTestContentOk

`func (o *NexusProvider) GetTestContentOk() (*string, bool)`

GetTestContentOk returns a tuple with the TestContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestContent

`func (o *NexusProvider) SetTestContent(v string)`

SetTestContent sets TestContent field to given value.

### HasTestContent

`func (o *NexusProvider) HasTestContent() bool`

HasTestContent returns a boolean if a field has been set.

### GetText

`func (o *NexusProvider) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NexusProvider) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NexusProvider) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *NexusProvider) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTopK

`func (o *NexusProvider) GetTopK() int64`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *NexusProvider) GetTopKOk() (*int64, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *NexusProvider) SetTopK(v int64)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *NexusProvider) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetTopP

`func (o *NexusProvider) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *NexusProvider) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *NexusProvider) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *NexusProvider) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetType

`func (o *NexusProvider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexusProvider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexusProvider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexusProvider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserCert

`func (o *NexusProvider) GetUserCert() string`

GetUserCert returns the UserCert field if non-nil, zero value otherwise.

### GetUserCertOk

`func (o *NexusProvider) GetUserCertOk() (*string, bool)`

GetUserCertOk returns a tuple with the UserCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCert

`func (o *NexusProvider) SetUserCert(v string)`

SetUserCert sets UserCert field to given value.

### HasUserCert

`func (o *NexusProvider) HasUserCert() bool`

HasUserCert returns a boolean if a field has been set.

### GetUserKey

`func (o *NexusProvider) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *NexusProvider) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *NexusProvider) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.

### HasUserKey

`func (o *NexusProvider) HasUserKey() bool`

HasUserKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


