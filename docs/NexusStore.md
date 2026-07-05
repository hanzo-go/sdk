# NexusStore

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentProvider** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**ChatCount** | Pointer to **int64** |  | [optional] 
**ChildModelProviders** | Pointer to **[]string** |  | [optional] 
**ChildStores** | Pointer to **[]string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisableFileUpload** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EmbeddingProvider** | Pointer to **string** |  | [optional] 
**EnableTtsStreaming** | Pointer to **bool** |  | [optional] 
**FileTree** | Pointer to [**NexusFile**](NexusFile.md) |  | [optional] 
**Frequency** | Pointer to **int64** |  | [optional] 
**ImageProvider** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**KnowledgeCount** | Pointer to **int64** |  | [optional] 
**LimitMinutes** | Pointer to **int64** |  | [optional] 
**MemoryLimit** | Pointer to **int64** |  | [optional] 
**MessageCount** | Pointer to **int64** |  | [optional] 
**ModelProvider** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**Prompts** | Pointer to [**[]NexusPrompt**](NexusPrompt.md) |  | [optional] 
**PropertiesMap** | Pointer to **map[string]interface{}** |  | [optional] 
**SearchProvider** | Pointer to **string** |  | [optional] 
**ShowAutoRead** | Pointer to **bool** |  | [optional] 
**SpeechToTextProvider** | Pointer to **string** |  | [optional] 
**SplitProvider** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**StorageSubpath** | Pointer to **string** |  | [optional] 
**SuggestionCount** | Pointer to **int64** |  | [optional] 
**TextToSpeechProvider** | Pointer to **string** |  | [optional] 
**ThemeColor** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**Welcome** | Pointer to **string** |  | [optional] 
**WelcomeText** | Pointer to **string** |  | [optional] 
**WelcomeTitle** | Pointer to **string** |  | [optional] 

## Methods

### NewNexusStore

`func NewNexusStore() *NexusStore`

NewNexusStore instantiates a new NexusStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusStoreWithDefaults

`func NewNexusStoreWithDefaults() *NexusStore`

NewNexusStoreWithDefaults instantiates a new NexusStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentProvider

`func (o *NexusStore) GetAgentProvider() string`

GetAgentProvider returns the AgentProvider field if non-nil, zero value otherwise.

### GetAgentProviderOk

`func (o *NexusStore) GetAgentProviderOk() (*string, bool)`

GetAgentProviderOk returns a tuple with the AgentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProvider

`func (o *NexusStore) SetAgentProvider(v string)`

SetAgentProvider sets AgentProvider field to given value.

### HasAgentProvider

`func (o *NexusStore) HasAgentProvider() bool`

HasAgentProvider returns a boolean if a field has been set.

### GetAvatar

`func (o *NexusStore) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *NexusStore) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *NexusStore) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *NexusStore) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChatCount

`func (o *NexusStore) GetChatCount() int64`

GetChatCount returns the ChatCount field if non-nil, zero value otherwise.

### GetChatCountOk

`func (o *NexusStore) GetChatCountOk() (*int64, bool)`

GetChatCountOk returns a tuple with the ChatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatCount

`func (o *NexusStore) SetChatCount(v int64)`

SetChatCount sets ChatCount field to given value.

### HasChatCount

`func (o *NexusStore) HasChatCount() bool`

HasChatCount returns a boolean if a field has been set.

### GetChildModelProviders

`func (o *NexusStore) GetChildModelProviders() []string`

GetChildModelProviders returns the ChildModelProviders field if non-nil, zero value otherwise.

### GetChildModelProvidersOk

`func (o *NexusStore) GetChildModelProvidersOk() (*[]string, bool)`

GetChildModelProvidersOk returns a tuple with the ChildModelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildModelProviders

`func (o *NexusStore) SetChildModelProviders(v []string)`

SetChildModelProviders sets ChildModelProviders field to given value.

### HasChildModelProviders

`func (o *NexusStore) HasChildModelProviders() bool`

HasChildModelProviders returns a boolean if a field has been set.

### GetChildStores

`func (o *NexusStore) GetChildStores() []string`

GetChildStores returns the ChildStores field if non-nil, zero value otherwise.

### GetChildStoresOk

`func (o *NexusStore) GetChildStoresOk() (*[]string, bool)`

GetChildStoresOk returns a tuple with the ChildStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildStores

`func (o *NexusStore) SetChildStores(v []string)`

SetChildStores sets ChildStores field to given value.

### HasChildStores

`func (o *NexusStore) HasChildStores() bool`

HasChildStores returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusStore) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusStore) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusStore) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusStore) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisableFileUpload

`func (o *NexusStore) GetDisableFileUpload() bool`

GetDisableFileUpload returns the DisableFileUpload field if non-nil, zero value otherwise.

### GetDisableFileUploadOk

`func (o *NexusStore) GetDisableFileUploadOk() (*bool, bool)`

GetDisableFileUploadOk returns a tuple with the DisableFileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFileUpload

`func (o *NexusStore) SetDisableFileUpload(v bool)`

SetDisableFileUpload sets DisableFileUpload field to given value.

### HasDisableFileUpload

`func (o *NexusStore) HasDisableFileUpload() bool`

HasDisableFileUpload returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusStore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusStore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusStore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusStore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *NexusStore) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *NexusStore) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *NexusStore) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *NexusStore) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetEnableTtsStreaming

`func (o *NexusStore) GetEnableTtsStreaming() bool`

GetEnableTtsStreaming returns the EnableTtsStreaming field if non-nil, zero value otherwise.

### GetEnableTtsStreamingOk

`func (o *NexusStore) GetEnableTtsStreamingOk() (*bool, bool)`

GetEnableTtsStreamingOk returns a tuple with the EnableTtsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTtsStreaming

`func (o *NexusStore) SetEnableTtsStreaming(v bool)`

SetEnableTtsStreaming sets EnableTtsStreaming field to given value.

### HasEnableTtsStreaming

`func (o *NexusStore) HasEnableTtsStreaming() bool`

HasEnableTtsStreaming returns a boolean if a field has been set.

### GetFileTree

`func (o *NexusStore) GetFileTree() NexusFile`

GetFileTree returns the FileTree field if non-nil, zero value otherwise.

### GetFileTreeOk

`func (o *NexusStore) GetFileTreeOk() (*NexusFile, bool)`

GetFileTreeOk returns a tuple with the FileTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTree

`func (o *NexusStore) SetFileTree(v NexusFile)`

SetFileTree sets FileTree field to given value.

### HasFileTree

`func (o *NexusStore) HasFileTree() bool`

HasFileTree returns a boolean if a field has been set.

### GetFrequency

`func (o *NexusStore) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *NexusStore) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *NexusStore) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *NexusStore) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetImageProvider

`func (o *NexusStore) GetImageProvider() string`

GetImageProvider returns the ImageProvider field if non-nil, zero value otherwise.

### GetImageProviderOk

`func (o *NexusStore) GetImageProviderOk() (*string, bool)`

GetImageProviderOk returns a tuple with the ImageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProvider

`func (o *NexusStore) SetImageProvider(v string)`

SetImageProvider sets ImageProvider field to given value.

### HasImageProvider

`func (o *NexusStore) HasImageProvider() bool`

HasImageProvider returns a boolean if a field has been set.

### GetIsDefault

`func (o *NexusStore) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *NexusStore) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *NexusStore) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *NexusStore) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetKnowledgeCount

`func (o *NexusStore) GetKnowledgeCount() int64`

GetKnowledgeCount returns the KnowledgeCount field if non-nil, zero value otherwise.

### GetKnowledgeCountOk

`func (o *NexusStore) GetKnowledgeCountOk() (*int64, bool)`

GetKnowledgeCountOk returns a tuple with the KnowledgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCount

`func (o *NexusStore) SetKnowledgeCount(v int64)`

SetKnowledgeCount sets KnowledgeCount field to given value.

### HasKnowledgeCount

`func (o *NexusStore) HasKnowledgeCount() bool`

HasKnowledgeCount returns a boolean if a field has been set.

### GetLimitMinutes

`func (o *NexusStore) GetLimitMinutes() int64`

GetLimitMinutes returns the LimitMinutes field if non-nil, zero value otherwise.

### GetLimitMinutesOk

`func (o *NexusStore) GetLimitMinutesOk() (*int64, bool)`

GetLimitMinutesOk returns a tuple with the LimitMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMinutes

`func (o *NexusStore) SetLimitMinutes(v int64)`

SetLimitMinutes sets LimitMinutes field to given value.

### HasLimitMinutes

`func (o *NexusStore) HasLimitMinutes() bool`

HasLimitMinutes returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *NexusStore) GetMemoryLimit() int64`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *NexusStore) GetMemoryLimitOk() (*int64, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *NexusStore) SetMemoryLimit(v int64)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *NexusStore) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMessageCount

`func (o *NexusStore) GetMessageCount() int64`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *NexusStore) GetMessageCountOk() (*int64, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *NexusStore) SetMessageCount(v int64)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *NexusStore) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetModelProvider

`func (o *NexusStore) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *NexusStore) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *NexusStore) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *NexusStore) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *NexusStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *NexusStore) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusStore) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusStore) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusStore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrompt

`func (o *NexusStore) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *NexusStore) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *NexusStore) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *NexusStore) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetPrompts

`func (o *NexusStore) GetPrompts() []NexusPrompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *NexusStore) GetPromptsOk() (*[]NexusPrompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *NexusStore) SetPrompts(v []NexusPrompt)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *NexusStore) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetPropertiesMap

`func (o *NexusStore) GetPropertiesMap() map[string]interface{}`

GetPropertiesMap returns the PropertiesMap field if non-nil, zero value otherwise.

### GetPropertiesMapOk

`func (o *NexusStore) GetPropertiesMapOk() (*map[string]interface{}, bool)`

GetPropertiesMapOk returns a tuple with the PropertiesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesMap

`func (o *NexusStore) SetPropertiesMap(v map[string]interface{})`

SetPropertiesMap sets PropertiesMap field to given value.

### HasPropertiesMap

`func (o *NexusStore) HasPropertiesMap() bool`

HasPropertiesMap returns a boolean if a field has been set.

### GetSearchProvider

`func (o *NexusStore) GetSearchProvider() string`

GetSearchProvider returns the SearchProvider field if non-nil, zero value otherwise.

### GetSearchProviderOk

`func (o *NexusStore) GetSearchProviderOk() (*string, bool)`

GetSearchProviderOk returns a tuple with the SearchProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProvider

`func (o *NexusStore) SetSearchProvider(v string)`

SetSearchProvider sets SearchProvider field to given value.

### HasSearchProvider

`func (o *NexusStore) HasSearchProvider() bool`

HasSearchProvider returns a boolean if a field has been set.

### GetShowAutoRead

`func (o *NexusStore) GetShowAutoRead() bool`

GetShowAutoRead returns the ShowAutoRead field if non-nil, zero value otherwise.

### GetShowAutoReadOk

`func (o *NexusStore) GetShowAutoReadOk() (*bool, bool)`

GetShowAutoReadOk returns a tuple with the ShowAutoRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutoRead

`func (o *NexusStore) SetShowAutoRead(v bool)`

SetShowAutoRead sets ShowAutoRead field to given value.

### HasShowAutoRead

`func (o *NexusStore) HasShowAutoRead() bool`

HasShowAutoRead returns a boolean if a field has been set.

### GetSpeechToTextProvider

`func (o *NexusStore) GetSpeechToTextProvider() string`

GetSpeechToTextProvider returns the SpeechToTextProvider field if non-nil, zero value otherwise.

### GetSpeechToTextProviderOk

`func (o *NexusStore) GetSpeechToTextProviderOk() (*string, bool)`

GetSpeechToTextProviderOk returns a tuple with the SpeechToTextProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechToTextProvider

`func (o *NexusStore) SetSpeechToTextProvider(v string)`

SetSpeechToTextProvider sets SpeechToTextProvider field to given value.

### HasSpeechToTextProvider

`func (o *NexusStore) HasSpeechToTextProvider() bool`

HasSpeechToTextProvider returns a boolean if a field has been set.

### GetSplitProvider

`func (o *NexusStore) GetSplitProvider() string`

GetSplitProvider returns the SplitProvider field if non-nil, zero value otherwise.

### GetSplitProviderOk

`func (o *NexusStore) GetSplitProviderOk() (*string, bool)`

GetSplitProviderOk returns a tuple with the SplitProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitProvider

`func (o *NexusStore) SetSplitProvider(v string)`

SetSplitProvider sets SplitProvider field to given value.

### HasSplitProvider

`func (o *NexusStore) HasSplitProvider() bool`

HasSplitProvider returns a boolean if a field has been set.

### GetState

`func (o *NexusStore) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *NexusStore) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *NexusStore) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *NexusStore) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageProvider

`func (o *NexusStore) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *NexusStore) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *NexusStore) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *NexusStore) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetStorageSubpath

`func (o *NexusStore) GetStorageSubpath() string`

GetStorageSubpath returns the StorageSubpath field if non-nil, zero value otherwise.

### GetStorageSubpathOk

`func (o *NexusStore) GetStorageSubpathOk() (*string, bool)`

GetStorageSubpathOk returns a tuple with the StorageSubpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSubpath

`func (o *NexusStore) SetStorageSubpath(v string)`

SetStorageSubpath sets StorageSubpath field to given value.

### HasStorageSubpath

`func (o *NexusStore) HasStorageSubpath() bool`

HasStorageSubpath returns a boolean if a field has been set.

### GetSuggestionCount

`func (o *NexusStore) GetSuggestionCount() int64`

GetSuggestionCount returns the SuggestionCount field if non-nil, zero value otherwise.

### GetSuggestionCountOk

`func (o *NexusStore) GetSuggestionCountOk() (*int64, bool)`

GetSuggestionCountOk returns a tuple with the SuggestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionCount

`func (o *NexusStore) SetSuggestionCount(v int64)`

SetSuggestionCount sets SuggestionCount field to given value.

### HasSuggestionCount

`func (o *NexusStore) HasSuggestionCount() bool`

HasSuggestionCount returns a boolean if a field has been set.

### GetTextToSpeechProvider

`func (o *NexusStore) GetTextToSpeechProvider() string`

GetTextToSpeechProvider returns the TextToSpeechProvider field if non-nil, zero value otherwise.

### GetTextToSpeechProviderOk

`func (o *NexusStore) GetTextToSpeechProviderOk() (*string, bool)`

GetTextToSpeechProviderOk returns a tuple with the TextToSpeechProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextToSpeechProvider

`func (o *NexusStore) SetTextToSpeechProvider(v string)`

SetTextToSpeechProvider sets TextToSpeechProvider field to given value.

### HasTextToSpeechProvider

`func (o *NexusStore) HasTextToSpeechProvider() bool`

HasTextToSpeechProvider returns a boolean if a field has been set.

### GetThemeColor

`func (o *NexusStore) GetThemeColor() string`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *NexusStore) GetThemeColorOk() (*string, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColor

`func (o *NexusStore) SetThemeColor(v string)`

SetThemeColor sets ThemeColor field to given value.

### HasThemeColor

`func (o *NexusStore) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.

### GetTitle

`func (o *NexusStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *NexusStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *NexusStore) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *NexusStore) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWelcome

`func (o *NexusStore) GetWelcome() string`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *NexusStore) GetWelcomeOk() (*string, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *NexusStore) SetWelcome(v string)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *NexusStore) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.

### GetWelcomeText

`func (o *NexusStore) GetWelcomeText() string`

GetWelcomeText returns the WelcomeText field if non-nil, zero value otherwise.

### GetWelcomeTextOk

`func (o *NexusStore) GetWelcomeTextOk() (*string, bool)`

GetWelcomeTextOk returns a tuple with the WelcomeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeText

`func (o *NexusStore) SetWelcomeText(v string)`

SetWelcomeText sets WelcomeText field to given value.

### HasWelcomeText

`func (o *NexusStore) HasWelcomeText() bool`

HasWelcomeText returns a boolean if a field has been set.

### GetWelcomeTitle

`func (o *NexusStore) GetWelcomeTitle() string`

GetWelcomeTitle returns the WelcomeTitle field if non-nil, zero value otherwise.

### GetWelcomeTitleOk

`func (o *NexusStore) GetWelcomeTitleOk() (*string, bool)`

GetWelcomeTitleOk returns a tuple with the WelcomeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTitle

`func (o *NexusStore) SetWelcomeTitle(v string)`

SetWelcomeTitle sets WelcomeTitle field to given value.

### HasWelcomeTitle

`func (o *NexusStore) HasWelcomeTitle() bool`

HasWelcomeTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


