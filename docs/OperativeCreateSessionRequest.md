# OperativeCreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Provider** | **string** | API provider to use | 
**Model** | Pointer to **string** | Model name (defaults to provider default) | [optional] 
**ApiKey** | Pointer to **string** | API key for the provider (if not set via env) | [optional] 
**SystemPromptSuffix** | Pointer to **string** | Extra instructions appended to the system prompt | [optional] 
**ToolVersion** | Pointer to **string** | Tool version to use | [optional] [default to "computer_use_20250124"]
**MaxTokens** | Pointer to **int32** | Maximum output tokens | [optional] [default to 128000]
**ThinkingBudget** | Pointer to **int32** | Token budget for thinking (null to disable) | [optional] 
**OnlyNMostRecentImages** | Pointer to **int32** | Number of recent screenshots to send to model | [optional] [default to 1]
**TokenEfficientToolsBeta** | Pointer to **bool** | Enable token-efficient tools beta flag | [optional] [default to false]
**DisplayWidth** | Pointer to **int32** | Virtual display width in pixels | [optional] [default to 1024]
**DisplayHeight** | Pointer to **int32** | Virtual display height in pixels | [optional] [default to 768]

## Methods

### NewOperativeCreateSessionRequest

`func NewOperativeCreateSessionRequest(provider string, ) *OperativeCreateSessionRequest`

NewOperativeCreateSessionRequest instantiates a new OperativeCreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperativeCreateSessionRequestWithDefaults

`func NewOperativeCreateSessionRequestWithDefaults() *OperativeCreateSessionRequest`

NewOperativeCreateSessionRequestWithDefaults instantiates a new OperativeCreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProvider

`func (o *OperativeCreateSessionRequest) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *OperativeCreateSessionRequest) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *OperativeCreateSessionRequest) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetModel

`func (o *OperativeCreateSessionRequest) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *OperativeCreateSessionRequest) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *OperativeCreateSessionRequest) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *OperativeCreateSessionRequest) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetApiKey

`func (o *OperativeCreateSessionRequest) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *OperativeCreateSessionRequest) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *OperativeCreateSessionRequest) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *OperativeCreateSessionRequest) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetSystemPromptSuffix

`func (o *OperativeCreateSessionRequest) GetSystemPromptSuffix() string`

GetSystemPromptSuffix returns the SystemPromptSuffix field if non-nil, zero value otherwise.

### GetSystemPromptSuffixOk

`func (o *OperativeCreateSessionRequest) GetSystemPromptSuffixOk() (*string, bool)`

GetSystemPromptSuffixOk returns a tuple with the SystemPromptSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystemPromptSuffix

`func (o *OperativeCreateSessionRequest) SetSystemPromptSuffix(v string)`

SetSystemPromptSuffix sets SystemPromptSuffix field to given value.

### HasSystemPromptSuffix

`func (o *OperativeCreateSessionRequest) HasSystemPromptSuffix() bool`

HasSystemPromptSuffix returns a boolean if a field has been set.

### GetToolVersion

`func (o *OperativeCreateSessionRequest) GetToolVersion() string`

GetToolVersion returns the ToolVersion field if non-nil, zero value otherwise.

### GetToolVersionOk

`func (o *OperativeCreateSessionRequest) GetToolVersionOk() (*string, bool)`

GetToolVersionOk returns a tuple with the ToolVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolVersion

`func (o *OperativeCreateSessionRequest) SetToolVersion(v string)`

SetToolVersion sets ToolVersion field to given value.

### HasToolVersion

`func (o *OperativeCreateSessionRequest) HasToolVersion() bool`

HasToolVersion returns a boolean if a field has been set.

### GetMaxTokens

`func (o *OperativeCreateSessionRequest) GetMaxTokens() int32`

GetMaxTokens returns the MaxTokens field if non-nil, zero value otherwise.

### GetMaxTokensOk

`func (o *OperativeCreateSessionRequest) GetMaxTokensOk() (*int32, bool)`

GetMaxTokensOk returns a tuple with the MaxTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTokens

`func (o *OperativeCreateSessionRequest) SetMaxTokens(v int32)`

SetMaxTokens sets MaxTokens field to given value.

### HasMaxTokens

`func (o *OperativeCreateSessionRequest) HasMaxTokens() bool`

HasMaxTokens returns a boolean if a field has been set.

### GetThinkingBudget

`func (o *OperativeCreateSessionRequest) GetThinkingBudget() int32`

GetThinkingBudget returns the ThinkingBudget field if non-nil, zero value otherwise.

### GetThinkingBudgetOk

`func (o *OperativeCreateSessionRequest) GetThinkingBudgetOk() (*int32, bool)`

GetThinkingBudgetOk returns a tuple with the ThinkingBudget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThinkingBudget

`func (o *OperativeCreateSessionRequest) SetThinkingBudget(v int32)`

SetThinkingBudget sets ThinkingBudget field to given value.

### HasThinkingBudget

`func (o *OperativeCreateSessionRequest) HasThinkingBudget() bool`

HasThinkingBudget returns a boolean if a field has been set.

### GetOnlyNMostRecentImages

`func (o *OperativeCreateSessionRequest) GetOnlyNMostRecentImages() int32`

GetOnlyNMostRecentImages returns the OnlyNMostRecentImages field if non-nil, zero value otherwise.

### GetOnlyNMostRecentImagesOk

`func (o *OperativeCreateSessionRequest) GetOnlyNMostRecentImagesOk() (*int32, bool)`

GetOnlyNMostRecentImagesOk returns a tuple with the OnlyNMostRecentImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlyNMostRecentImages

`func (o *OperativeCreateSessionRequest) SetOnlyNMostRecentImages(v int32)`

SetOnlyNMostRecentImages sets OnlyNMostRecentImages field to given value.

### HasOnlyNMostRecentImages

`func (o *OperativeCreateSessionRequest) HasOnlyNMostRecentImages() bool`

HasOnlyNMostRecentImages returns a boolean if a field has been set.

### GetTokenEfficientToolsBeta

`func (o *OperativeCreateSessionRequest) GetTokenEfficientToolsBeta() bool`

GetTokenEfficientToolsBeta returns the TokenEfficientToolsBeta field if non-nil, zero value otherwise.

### GetTokenEfficientToolsBetaOk

`func (o *OperativeCreateSessionRequest) GetTokenEfficientToolsBetaOk() (*bool, bool)`

GetTokenEfficientToolsBetaOk returns a tuple with the TokenEfficientToolsBeta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenEfficientToolsBeta

`func (o *OperativeCreateSessionRequest) SetTokenEfficientToolsBeta(v bool)`

SetTokenEfficientToolsBeta sets TokenEfficientToolsBeta field to given value.

### HasTokenEfficientToolsBeta

`func (o *OperativeCreateSessionRequest) HasTokenEfficientToolsBeta() bool`

HasTokenEfficientToolsBeta returns a boolean if a field has been set.

### GetDisplayWidth

`func (o *OperativeCreateSessionRequest) GetDisplayWidth() int32`

GetDisplayWidth returns the DisplayWidth field if non-nil, zero value otherwise.

### GetDisplayWidthOk

`func (o *OperativeCreateSessionRequest) GetDisplayWidthOk() (*int32, bool)`

GetDisplayWidthOk returns a tuple with the DisplayWidth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayWidth

`func (o *OperativeCreateSessionRequest) SetDisplayWidth(v int32)`

SetDisplayWidth sets DisplayWidth field to given value.

### HasDisplayWidth

`func (o *OperativeCreateSessionRequest) HasDisplayWidth() bool`

HasDisplayWidth returns a boolean if a field has been set.

### GetDisplayHeight

`func (o *OperativeCreateSessionRequest) GetDisplayHeight() int32`

GetDisplayHeight returns the DisplayHeight field if non-nil, zero value otherwise.

### GetDisplayHeightOk

`func (o *OperativeCreateSessionRequest) GetDisplayHeightOk() (*int32, bool)`

GetDisplayHeightOk returns a tuple with the DisplayHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayHeight

`func (o *OperativeCreateSessionRequest) SetDisplayHeight(v int32)`

SetDisplayHeight sets DisplayHeight field to given value.

### HasDisplayHeight

`func (o *OperativeCreateSessionRequest) HasDisplayHeight() bool`

HasDisplayHeight returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


