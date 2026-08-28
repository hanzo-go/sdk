# Provider

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ApiVersion** | Pointer to **string** |  | [optional] 
**Asset** | Pointer to **string** |  | [optional] 
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
**ErrorText** | Pointer to **string** |  | [optional] 
**Flavor** | Pointer to **string** |  | [optional] 
**FrequencyPenalty** | Pointer to **float32** |  | [optional] 
**InputPricePerThousandTokens** | Pointer to **float32** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**IsRemote** | Pointer to **bool** |  | [optional] 
**McpTools** | Pointer to [**[]AgentMcpTools**](AgentMcpTools.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**OutputPricePerThousandTokens** | Pointer to **float32** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**PresencePenalty** | Pointer to **float32** |  | [optional] 
**ProviderKey** | Pointer to **string** |  | [optional] 
**ProviderUrl** | Pointer to **string** |  | [optional] 
**RawText** | Pointer to **string** |  | [optional] 
**Region** | Pointer to **string** |  | [optional] 
**ResultSummary** | Pointer to **string** |  | [optional] 
**Runner** | Pointer to **string** |  | [optional] 
**SignCert** | Pointer to **string** |  | [optional] 
**SignKey** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**SubType** | Pointer to **string** |  | [optional] 
**Target** | Pointer to **string** |  | [optional] 
**TargetMode** | Pointer to **string** |  | [optional] 
**Temperature** | Pointer to **float32** |  | [optional] 
**TestContent** | Pointer to **string** |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TopK** | Pointer to **int32** |  | [optional] 
**TopP** | Pointer to **float32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UserCert** | Pointer to **string** |  | [optional] 
**UserKey** | Pointer to **string** |  | [optional] 

## Methods

### NewProvider

`func NewProvider() *Provider`

NewProvider instantiates a new Provider object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderWithDefaults

`func NewProviderWithDefaults() *Provider`

NewProviderWithDefaults instantiates a new Provider object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApiVersion

`func (o *Provider) GetApiVersion() string`

GetApiVersion returns the ApiVersion field if non-nil, zero value otherwise.

### GetApiVersionOk

`func (o *Provider) GetApiVersionOk() (*string, bool)`

GetApiVersionOk returns a tuple with the ApiVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiVersion

`func (o *Provider) SetApiVersion(v string)`

SetApiVersion sets ApiVersion field to given value.

### HasApiVersion

`func (o *Provider) HasApiVersion() bool`

HasApiVersion returns a boolean if a field has been set.

### GetAsset

`func (o *Provider) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *Provider) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *Provider) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *Provider) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetBrowserUrl

`func (o *Provider) GetBrowserUrl() string`

GetBrowserUrl returns the BrowserUrl field if non-nil, zero value otherwise.

### GetBrowserUrlOk

`func (o *Provider) GetBrowserUrlOk() (*string, bool)`

GetBrowserUrlOk returns a tuple with the BrowserUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowserUrl

`func (o *Provider) SetBrowserUrl(v string)`

SetBrowserUrl sets BrowserUrl field to given value.

### HasBrowserUrl

`func (o *Provider) HasBrowserUrl() bool`

HasBrowserUrl returns a boolean if a field has been set.

### GetCategory

`func (o *Provider) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Provider) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Provider) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Provider) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetChain

`func (o *Provider) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Provider) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Provider) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Provider) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetClientId

`func (o *Provider) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *Provider) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *Provider) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *Provider) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetClientSecret

`func (o *Provider) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *Provider) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *Provider) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *Provider) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetCompatibleProvider

`func (o *Provider) GetCompatibleProvider() string`

GetCompatibleProvider returns the CompatibleProvider field if non-nil, zero value otherwise.

### GetCompatibleProviderOk

`func (o *Provider) GetCompatibleProviderOk() (*string, bool)`

GetCompatibleProviderOk returns a tuple with the CompatibleProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleProvider

`func (o *Provider) SetCompatibleProvider(v string)`

SetCompatibleProvider sets CompatibleProvider field to given value.

### HasCompatibleProvider

`func (o *Provider) HasCompatibleProvider() bool`

HasCompatibleProvider returns a boolean if a field has been set.

### GetConfigText

`func (o *Provider) GetConfigText() string`

GetConfigText returns the ConfigText field if non-nil, zero value otherwise.

### GetConfigTextOk

`func (o *Provider) GetConfigTextOk() (*string, bool)`

GetConfigTextOk returns a tuple with the ConfigText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigText

`func (o *Provider) SetConfigText(v string)`

SetConfigText sets ConfigText field to given value.

### HasConfigText

`func (o *Provider) HasConfigText() bool`

HasConfigText returns a boolean if a field has been set.

### GetContractMethod

`func (o *Provider) GetContractMethod() string`

GetContractMethod returns the ContractMethod field if non-nil, zero value otherwise.

### GetContractMethodOk

`func (o *Provider) GetContractMethodOk() (*string, bool)`

GetContractMethodOk returns a tuple with the ContractMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractMethod

`func (o *Provider) SetContractMethod(v string)`

SetContractMethod sets ContractMethod field to given value.

### HasContractMethod

`func (o *Provider) HasContractMethod() bool`

HasContractMethod returns a boolean if a field has been set.

### GetContractName

`func (o *Provider) GetContractName() string`

GetContractName returns the ContractName field if non-nil, zero value otherwise.

### GetContractNameOk

`func (o *Provider) GetContractNameOk() (*string, bool)`

GetContractNameOk returns a tuple with the ContractName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractName

`func (o *Provider) SetContractName(v string)`

SetContractName sets ContractName field to given value.

### HasContractName

`func (o *Provider) HasContractName() bool`

HasContractName returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Provider) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Provider) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Provider) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Provider) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *Provider) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Provider) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Provider) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Provider) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *Provider) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Provider) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Provider) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Provider) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEnableThinking

`func (o *Provider) GetEnableThinking() bool`

GetEnableThinking returns the EnableThinking field if non-nil, zero value otherwise.

### GetEnableThinkingOk

`func (o *Provider) GetEnableThinkingOk() (*bool, bool)`

GetEnableThinkingOk returns a tuple with the EnableThinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableThinking

`func (o *Provider) SetEnableThinking(v bool)`

SetEnableThinking sets EnableThinking field to given value.

### HasEnableThinking

`func (o *Provider) HasEnableThinking() bool`

HasEnableThinking returns a boolean if a field has been set.

### GetErrorText

`func (o *Provider) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *Provider) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *Provider) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *Provider) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetFlavor

`func (o *Provider) GetFlavor() string`

GetFlavor returns the Flavor field if non-nil, zero value otherwise.

### GetFlavorOk

`func (o *Provider) GetFlavorOk() (*string, bool)`

GetFlavorOk returns a tuple with the Flavor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlavor

`func (o *Provider) SetFlavor(v string)`

SetFlavor sets Flavor field to given value.

### HasFlavor

`func (o *Provider) HasFlavor() bool`

HasFlavor returns a boolean if a field has been set.

### GetFrequencyPenalty

`func (o *Provider) GetFrequencyPenalty() float32`

GetFrequencyPenalty returns the FrequencyPenalty field if non-nil, zero value otherwise.

### GetFrequencyPenaltyOk

`func (o *Provider) GetFrequencyPenaltyOk() (*float32, bool)`

GetFrequencyPenaltyOk returns a tuple with the FrequencyPenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequencyPenalty

`func (o *Provider) SetFrequencyPenalty(v float32)`

SetFrequencyPenalty sets FrequencyPenalty field to given value.

### HasFrequencyPenalty

`func (o *Provider) HasFrequencyPenalty() bool`

HasFrequencyPenalty returns a boolean if a field has been set.

### GetInputPricePerThousandTokens

`func (o *Provider) GetInputPricePerThousandTokens() float32`

GetInputPricePerThousandTokens returns the InputPricePerThousandTokens field if non-nil, zero value otherwise.

### GetInputPricePerThousandTokensOk

`func (o *Provider) GetInputPricePerThousandTokensOk() (*float32, bool)`

GetInputPricePerThousandTokensOk returns a tuple with the InputPricePerThousandTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInputPricePerThousandTokens

`func (o *Provider) SetInputPricePerThousandTokens(v float32)`

SetInputPricePerThousandTokens sets InputPricePerThousandTokens field to given value.

### HasInputPricePerThousandTokens

`func (o *Provider) HasInputPricePerThousandTokens() bool`

HasInputPricePerThousandTokens returns a boolean if a field has been set.

### GetIsDefault

`func (o *Provider) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Provider) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Provider) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Provider) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetIsRemote

`func (o *Provider) GetIsRemote() bool`

GetIsRemote returns the IsRemote field if non-nil, zero value otherwise.

### GetIsRemoteOk

`func (o *Provider) GetIsRemoteOk() (*bool, bool)`

GetIsRemoteOk returns a tuple with the IsRemote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRemote

`func (o *Provider) SetIsRemote(v bool)`

SetIsRemote sets IsRemote field to given value.

### HasIsRemote

`func (o *Provider) HasIsRemote() bool`

HasIsRemote returns a boolean if a field has been set.

### GetMcpTools

`func (o *Provider) GetMcpTools() []AgentMcpTools`

GetMcpTools returns the McpTools field if non-nil, zero value otherwise.

### GetMcpToolsOk

`func (o *Provider) GetMcpToolsOk() (*[]AgentMcpTools, bool)`

GetMcpToolsOk returns a tuple with the McpTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMcpTools

`func (o *Provider) SetMcpTools(v []AgentMcpTools)`

SetMcpTools sets McpTools field to given value.

### HasMcpTools

`func (o *Provider) HasMcpTools() bool`

HasMcpTools returns a boolean if a field has been set.

### GetName

`func (o *Provider) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Provider) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Provider) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Provider) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNetwork

`func (o *Provider) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *Provider) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *Provider) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *Provider) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetOutputPricePerThousandTokens

`func (o *Provider) GetOutputPricePerThousandTokens() float32`

GetOutputPricePerThousandTokens returns the OutputPricePerThousandTokens field if non-nil, zero value otherwise.

### GetOutputPricePerThousandTokensOk

`func (o *Provider) GetOutputPricePerThousandTokensOk() (*float32, bool)`

GetOutputPricePerThousandTokensOk returns a tuple with the OutputPricePerThousandTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutputPricePerThousandTokens

`func (o *Provider) SetOutputPricePerThousandTokens(v float32)`

SetOutputPricePerThousandTokens sets OutputPricePerThousandTokens field to given value.

### HasOutputPricePerThousandTokens

`func (o *Provider) HasOutputPricePerThousandTokens() bool`

HasOutputPricePerThousandTokens returns a boolean if a field has been set.

### GetOwner

`func (o *Provider) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Provider) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Provider) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Provider) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPresencePenalty

`func (o *Provider) GetPresencePenalty() float32`

GetPresencePenalty returns the PresencePenalty field if non-nil, zero value otherwise.

### GetPresencePenaltyOk

`func (o *Provider) GetPresencePenaltyOk() (*float32, bool)`

GetPresencePenaltyOk returns a tuple with the PresencePenalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresencePenalty

`func (o *Provider) SetPresencePenalty(v float32)`

SetPresencePenalty sets PresencePenalty field to given value.

### HasPresencePenalty

`func (o *Provider) HasPresencePenalty() bool`

HasPresencePenalty returns a boolean if a field has been set.

### GetProviderKey

`func (o *Provider) GetProviderKey() string`

GetProviderKey returns the ProviderKey field if non-nil, zero value otherwise.

### GetProviderKeyOk

`func (o *Provider) GetProviderKeyOk() (*string, bool)`

GetProviderKeyOk returns a tuple with the ProviderKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderKey

`func (o *Provider) SetProviderKey(v string)`

SetProviderKey sets ProviderKey field to given value.

### HasProviderKey

`func (o *Provider) HasProviderKey() bool`

HasProviderKey returns a boolean if a field has been set.

### GetProviderUrl

`func (o *Provider) GetProviderUrl() string`

GetProviderUrl returns the ProviderUrl field if non-nil, zero value otherwise.

### GetProviderUrlOk

`func (o *Provider) GetProviderUrlOk() (*string, bool)`

GetProviderUrlOk returns a tuple with the ProviderUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderUrl

`func (o *Provider) SetProviderUrl(v string)`

SetProviderUrl sets ProviderUrl field to given value.

### HasProviderUrl

`func (o *Provider) HasProviderUrl() bool`

HasProviderUrl returns a boolean if a field has been set.

### GetRawText

`func (o *Provider) GetRawText() string`

GetRawText returns the RawText field if non-nil, zero value otherwise.

### GetRawTextOk

`func (o *Provider) GetRawTextOk() (*string, bool)`

GetRawTextOk returns a tuple with the RawText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawText

`func (o *Provider) SetRawText(v string)`

SetRawText sets RawText field to given value.

### HasRawText

`func (o *Provider) HasRawText() bool`

HasRawText returns a boolean if a field has been set.

### GetRegion

`func (o *Provider) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Provider) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Provider) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *Provider) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetResultSummary

`func (o *Provider) GetResultSummary() string`

GetResultSummary returns the ResultSummary field if non-nil, zero value otherwise.

### GetResultSummaryOk

`func (o *Provider) GetResultSummaryOk() (*string, bool)`

GetResultSummaryOk returns a tuple with the ResultSummary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResultSummary

`func (o *Provider) SetResultSummary(v string)`

SetResultSummary sets ResultSummary field to given value.

### HasResultSummary

`func (o *Provider) HasResultSummary() bool`

HasResultSummary returns a boolean if a field has been set.

### GetRunner

`func (o *Provider) GetRunner() string`

GetRunner returns the Runner field if non-nil, zero value otherwise.

### GetRunnerOk

`func (o *Provider) GetRunnerOk() (*string, bool)`

GetRunnerOk returns a tuple with the Runner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunner

`func (o *Provider) SetRunner(v string)`

SetRunner sets Runner field to given value.

### HasRunner

`func (o *Provider) HasRunner() bool`

HasRunner returns a boolean if a field has been set.

### GetSignCert

`func (o *Provider) GetSignCert() string`

GetSignCert returns the SignCert field if non-nil, zero value otherwise.

### GetSignCertOk

`func (o *Provider) GetSignCertOk() (*string, bool)`

GetSignCertOk returns a tuple with the SignCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignCert

`func (o *Provider) SetSignCert(v string)`

SetSignCert sets SignCert field to given value.

### HasSignCert

`func (o *Provider) HasSignCert() bool`

HasSignCert returns a boolean if a field has been set.

### GetSignKey

`func (o *Provider) GetSignKey() string`

GetSignKey returns the SignKey field if non-nil, zero value otherwise.

### GetSignKeyOk

`func (o *Provider) GetSignKeyOk() (*string, bool)`

GetSignKeyOk returns a tuple with the SignKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignKey

`func (o *Provider) SetSignKey(v string)`

SetSignKey sets SignKey field to given value.

### HasSignKey

`func (o *Provider) HasSignKey() bool`

HasSignKey returns a boolean if a field has been set.

### GetState

`func (o *Provider) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Provider) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Provider) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Provider) HasState() bool`

HasState returns a boolean if a field has been set.

### GetSubType

`func (o *Provider) GetSubType() string`

GetSubType returns the SubType field if non-nil, zero value otherwise.

### GetSubTypeOk

`func (o *Provider) GetSubTypeOk() (*string, bool)`

GetSubTypeOk returns a tuple with the SubType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubType

`func (o *Provider) SetSubType(v string)`

SetSubType sets SubType field to given value.

### HasSubType

`func (o *Provider) HasSubType() bool`

HasSubType returns a boolean if a field has been set.

### GetTarget

`func (o *Provider) GetTarget() string`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Provider) GetTargetOk() (*string, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Provider) SetTarget(v string)`

SetTarget sets Target field to given value.

### HasTarget

`func (o *Provider) HasTarget() bool`

HasTarget returns a boolean if a field has been set.

### GetTargetMode

`func (o *Provider) GetTargetMode() string`

GetTargetMode returns the TargetMode field if non-nil, zero value otherwise.

### GetTargetModeOk

`func (o *Provider) GetTargetModeOk() (*string, bool)`

GetTargetModeOk returns a tuple with the TargetMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetMode

`func (o *Provider) SetTargetMode(v string)`

SetTargetMode sets TargetMode field to given value.

### HasTargetMode

`func (o *Provider) HasTargetMode() bool`

HasTargetMode returns a boolean if a field has been set.

### GetTemperature

`func (o *Provider) GetTemperature() float32`

GetTemperature returns the Temperature field if non-nil, zero value otherwise.

### GetTemperatureOk

`func (o *Provider) GetTemperatureOk() (*float32, bool)`

GetTemperatureOk returns a tuple with the Temperature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemperature

`func (o *Provider) SetTemperature(v float32)`

SetTemperature sets Temperature field to given value.

### HasTemperature

`func (o *Provider) HasTemperature() bool`

HasTemperature returns a boolean if a field has been set.

### GetTestContent

`func (o *Provider) GetTestContent() string`

GetTestContent returns the TestContent field if non-nil, zero value otherwise.

### GetTestContentOk

`func (o *Provider) GetTestContentOk() (*string, bool)`

GetTestContentOk returns a tuple with the TestContent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTestContent

`func (o *Provider) SetTestContent(v string)`

SetTestContent sets TestContent field to given value.

### HasTestContent

`func (o *Provider) HasTestContent() bool`

HasTestContent returns a boolean if a field has been set.

### GetText

`func (o *Provider) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Provider) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Provider) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Provider) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTopK

`func (o *Provider) GetTopK() int32`

GetTopK returns the TopK field if non-nil, zero value otherwise.

### GetTopKOk

`func (o *Provider) GetTopKOk() (*int32, bool)`

GetTopKOk returns a tuple with the TopK field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopK

`func (o *Provider) SetTopK(v int32)`

SetTopK sets TopK field to given value.

### HasTopK

`func (o *Provider) HasTopK() bool`

HasTopK returns a boolean if a field has been set.

### GetTopP

`func (o *Provider) GetTopP() float32`

GetTopP returns the TopP field if non-nil, zero value otherwise.

### GetTopPOk

`func (o *Provider) GetTopPOk() (*float32, bool)`

GetTopPOk returns a tuple with the TopP field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopP

`func (o *Provider) SetTopP(v float32)`

SetTopP sets TopP field to given value.

### HasTopP

`func (o *Provider) HasTopP() bool`

HasTopP returns a boolean if a field has been set.

### GetType

`func (o *Provider) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Provider) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Provider) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Provider) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserCert

`func (o *Provider) GetUserCert() string`

GetUserCert returns the UserCert field if non-nil, zero value otherwise.

### GetUserCertOk

`func (o *Provider) GetUserCertOk() (*string, bool)`

GetUserCertOk returns a tuple with the UserCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserCert

`func (o *Provider) SetUserCert(v string)`

SetUserCert sets UserCert field to given value.

### HasUserCert

`func (o *Provider) HasUserCert() bool`

HasUserCert returns a boolean if a field has been set.

### GetUserKey

`func (o *Provider) GetUserKey() string`

GetUserKey returns the UserKey field if non-nil, zero value otherwise.

### GetUserKeyOk

`func (o *Provider) GetUserKeyOk() (*string, bool)`

GetUserKeyOk returns a tuple with the UserKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserKey

`func (o *Provider) SetUserKey(v string)`

SetUserKey sets UserKey field to given value.

### HasUserKey

`func (o *Provider) HasUserKey() bool`

HasUserKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


