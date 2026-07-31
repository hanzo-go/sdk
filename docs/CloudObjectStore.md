# CloudObjectStore

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
**FileTree** | Pointer to [**CloudObjectFile**](CloudObjectFile.md) |  | [optional] 
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
**Prompts** | Pointer to [**[]CloudObjectPrompt**](CloudObjectPrompt.md) |  | [optional] 
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

### NewCloudObjectStore

`func NewCloudObjectStore() *CloudObjectStore`

NewCloudObjectStore instantiates a new CloudObjectStore object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectStoreWithDefaults

`func NewCloudObjectStoreWithDefaults() *CloudObjectStore`

NewCloudObjectStoreWithDefaults instantiates a new CloudObjectStore object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentProvider

`func (o *CloudObjectStore) GetAgentProvider() string`

GetAgentProvider returns the AgentProvider field if non-nil, zero value otherwise.

### GetAgentProviderOk

`func (o *CloudObjectStore) GetAgentProviderOk() (*string, bool)`

GetAgentProviderOk returns a tuple with the AgentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProvider

`func (o *CloudObjectStore) SetAgentProvider(v string)`

SetAgentProvider sets AgentProvider field to given value.

### HasAgentProvider

`func (o *CloudObjectStore) HasAgentProvider() bool`

HasAgentProvider returns a boolean if a field has been set.

### GetAvatar

`func (o *CloudObjectStore) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *CloudObjectStore) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *CloudObjectStore) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *CloudObjectStore) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetChatCount

`func (o *CloudObjectStore) GetChatCount() int64`

GetChatCount returns the ChatCount field if non-nil, zero value otherwise.

### GetChatCountOk

`func (o *CloudObjectStore) GetChatCountOk() (*int64, bool)`

GetChatCountOk returns a tuple with the ChatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatCount

`func (o *CloudObjectStore) SetChatCount(v int64)`

SetChatCount sets ChatCount field to given value.

### HasChatCount

`func (o *CloudObjectStore) HasChatCount() bool`

HasChatCount returns a boolean if a field has been set.

### GetChildModelProviders

`func (o *CloudObjectStore) GetChildModelProviders() []string`

GetChildModelProviders returns the ChildModelProviders field if non-nil, zero value otherwise.

### GetChildModelProvidersOk

`func (o *CloudObjectStore) GetChildModelProvidersOk() (*[]string, bool)`

GetChildModelProvidersOk returns a tuple with the ChildModelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildModelProviders

`func (o *CloudObjectStore) SetChildModelProviders(v []string)`

SetChildModelProviders sets ChildModelProviders field to given value.

### HasChildModelProviders

`func (o *CloudObjectStore) HasChildModelProviders() bool`

HasChildModelProviders returns a boolean if a field has been set.

### GetChildStores

`func (o *CloudObjectStore) GetChildStores() []string`

GetChildStores returns the ChildStores field if non-nil, zero value otherwise.

### GetChildStoresOk

`func (o *CloudObjectStore) GetChildStoresOk() (*[]string, bool)`

GetChildStoresOk returns a tuple with the ChildStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildStores

`func (o *CloudObjectStore) SetChildStores(v []string)`

SetChildStores sets ChildStores field to given value.

### HasChildStores

`func (o *CloudObjectStore) HasChildStores() bool`

HasChildStores returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectStore) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectStore) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectStore) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectStore) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisableFileUpload

`func (o *CloudObjectStore) GetDisableFileUpload() bool`

GetDisableFileUpload returns the DisableFileUpload field if non-nil, zero value otherwise.

### GetDisableFileUploadOk

`func (o *CloudObjectStore) GetDisableFileUploadOk() (*bool, bool)`

GetDisableFileUploadOk returns a tuple with the DisableFileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFileUpload

`func (o *CloudObjectStore) SetDisableFileUpload(v bool)`

SetDisableFileUpload sets DisableFileUpload field to given value.

### HasDisableFileUpload

`func (o *CloudObjectStore) HasDisableFileUpload() bool`

HasDisableFileUpload returns a boolean if a field has been set.

### GetDisplayName

`func (o *CloudObjectStore) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *CloudObjectStore) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *CloudObjectStore) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *CloudObjectStore) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *CloudObjectStore) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *CloudObjectStore) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *CloudObjectStore) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *CloudObjectStore) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetEnableTtsStreaming

`func (o *CloudObjectStore) GetEnableTtsStreaming() bool`

GetEnableTtsStreaming returns the EnableTtsStreaming field if non-nil, zero value otherwise.

### GetEnableTtsStreamingOk

`func (o *CloudObjectStore) GetEnableTtsStreamingOk() (*bool, bool)`

GetEnableTtsStreamingOk returns a tuple with the EnableTtsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTtsStreaming

`func (o *CloudObjectStore) SetEnableTtsStreaming(v bool)`

SetEnableTtsStreaming sets EnableTtsStreaming field to given value.

### HasEnableTtsStreaming

`func (o *CloudObjectStore) HasEnableTtsStreaming() bool`

HasEnableTtsStreaming returns a boolean if a field has been set.

### GetFileTree

`func (o *CloudObjectStore) GetFileTree() CloudObjectFile`

GetFileTree returns the FileTree field if non-nil, zero value otherwise.

### GetFileTreeOk

`func (o *CloudObjectStore) GetFileTreeOk() (*CloudObjectFile, bool)`

GetFileTreeOk returns a tuple with the FileTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTree

`func (o *CloudObjectStore) SetFileTree(v CloudObjectFile)`

SetFileTree sets FileTree field to given value.

### HasFileTree

`func (o *CloudObjectStore) HasFileTree() bool`

HasFileTree returns a boolean if a field has been set.

### GetFrequency

`func (o *CloudObjectStore) GetFrequency() int64`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *CloudObjectStore) GetFrequencyOk() (*int64, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *CloudObjectStore) SetFrequency(v int64)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *CloudObjectStore) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetImageProvider

`func (o *CloudObjectStore) GetImageProvider() string`

GetImageProvider returns the ImageProvider field if non-nil, zero value otherwise.

### GetImageProviderOk

`func (o *CloudObjectStore) GetImageProviderOk() (*string, bool)`

GetImageProviderOk returns a tuple with the ImageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProvider

`func (o *CloudObjectStore) SetImageProvider(v string)`

SetImageProvider sets ImageProvider field to given value.

### HasImageProvider

`func (o *CloudObjectStore) HasImageProvider() bool`

HasImageProvider returns a boolean if a field has been set.

### GetIsDefault

`func (o *CloudObjectStore) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CloudObjectStore) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CloudObjectStore) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *CloudObjectStore) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetKnowledgeCount

`func (o *CloudObjectStore) GetKnowledgeCount() int64`

GetKnowledgeCount returns the KnowledgeCount field if non-nil, zero value otherwise.

### GetKnowledgeCountOk

`func (o *CloudObjectStore) GetKnowledgeCountOk() (*int64, bool)`

GetKnowledgeCountOk returns a tuple with the KnowledgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCount

`func (o *CloudObjectStore) SetKnowledgeCount(v int64)`

SetKnowledgeCount sets KnowledgeCount field to given value.

### HasKnowledgeCount

`func (o *CloudObjectStore) HasKnowledgeCount() bool`

HasKnowledgeCount returns a boolean if a field has been set.

### GetLimitMinutes

`func (o *CloudObjectStore) GetLimitMinutes() int64`

GetLimitMinutes returns the LimitMinutes field if non-nil, zero value otherwise.

### GetLimitMinutesOk

`func (o *CloudObjectStore) GetLimitMinutesOk() (*int64, bool)`

GetLimitMinutesOk returns a tuple with the LimitMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMinutes

`func (o *CloudObjectStore) SetLimitMinutes(v int64)`

SetLimitMinutes sets LimitMinutes field to given value.

### HasLimitMinutes

`func (o *CloudObjectStore) HasLimitMinutes() bool`

HasLimitMinutes returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *CloudObjectStore) GetMemoryLimit() int64`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *CloudObjectStore) GetMemoryLimitOk() (*int64, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *CloudObjectStore) SetMemoryLimit(v int64)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *CloudObjectStore) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMessageCount

`func (o *CloudObjectStore) GetMessageCount() int64`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *CloudObjectStore) GetMessageCountOk() (*int64, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *CloudObjectStore) SetMessageCount(v int64)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *CloudObjectStore) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetModelProvider

`func (o *CloudObjectStore) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *CloudObjectStore) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *CloudObjectStore) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *CloudObjectStore) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectStore) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectStore) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectStore) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectStore) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectStore) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectStore) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectStore) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectStore) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrompt

`func (o *CloudObjectStore) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *CloudObjectStore) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *CloudObjectStore) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *CloudObjectStore) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetPrompts

`func (o *CloudObjectStore) GetPrompts() []CloudObjectPrompt`

GetPrompts returns the Prompts field if non-nil, zero value otherwise.

### GetPromptsOk

`func (o *CloudObjectStore) GetPromptsOk() (*[]CloudObjectPrompt, bool)`

GetPromptsOk returns a tuple with the Prompts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompts

`func (o *CloudObjectStore) SetPrompts(v []CloudObjectPrompt)`

SetPrompts sets Prompts field to given value.

### HasPrompts

`func (o *CloudObjectStore) HasPrompts() bool`

HasPrompts returns a boolean if a field has been set.

### GetPropertiesMap

`func (o *CloudObjectStore) GetPropertiesMap() map[string]interface{}`

GetPropertiesMap returns the PropertiesMap field if non-nil, zero value otherwise.

### GetPropertiesMapOk

`func (o *CloudObjectStore) GetPropertiesMapOk() (*map[string]interface{}, bool)`

GetPropertiesMapOk returns a tuple with the PropertiesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesMap

`func (o *CloudObjectStore) SetPropertiesMap(v map[string]interface{})`

SetPropertiesMap sets PropertiesMap field to given value.

### HasPropertiesMap

`func (o *CloudObjectStore) HasPropertiesMap() bool`

HasPropertiesMap returns a boolean if a field has been set.

### GetSearchProvider

`func (o *CloudObjectStore) GetSearchProvider() string`

GetSearchProvider returns the SearchProvider field if non-nil, zero value otherwise.

### GetSearchProviderOk

`func (o *CloudObjectStore) GetSearchProviderOk() (*string, bool)`

GetSearchProviderOk returns a tuple with the SearchProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProvider

`func (o *CloudObjectStore) SetSearchProvider(v string)`

SetSearchProvider sets SearchProvider field to given value.

### HasSearchProvider

`func (o *CloudObjectStore) HasSearchProvider() bool`

HasSearchProvider returns a boolean if a field has been set.

### GetShowAutoRead

`func (o *CloudObjectStore) GetShowAutoRead() bool`

GetShowAutoRead returns the ShowAutoRead field if non-nil, zero value otherwise.

### GetShowAutoReadOk

`func (o *CloudObjectStore) GetShowAutoReadOk() (*bool, bool)`

GetShowAutoReadOk returns a tuple with the ShowAutoRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutoRead

`func (o *CloudObjectStore) SetShowAutoRead(v bool)`

SetShowAutoRead sets ShowAutoRead field to given value.

### HasShowAutoRead

`func (o *CloudObjectStore) HasShowAutoRead() bool`

HasShowAutoRead returns a boolean if a field has been set.

### GetSpeechToTextProvider

`func (o *CloudObjectStore) GetSpeechToTextProvider() string`

GetSpeechToTextProvider returns the SpeechToTextProvider field if non-nil, zero value otherwise.

### GetSpeechToTextProviderOk

`func (o *CloudObjectStore) GetSpeechToTextProviderOk() (*string, bool)`

GetSpeechToTextProviderOk returns a tuple with the SpeechToTextProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechToTextProvider

`func (o *CloudObjectStore) SetSpeechToTextProvider(v string)`

SetSpeechToTextProvider sets SpeechToTextProvider field to given value.

### HasSpeechToTextProvider

`func (o *CloudObjectStore) HasSpeechToTextProvider() bool`

HasSpeechToTextProvider returns a boolean if a field has been set.

### GetSplitProvider

`func (o *CloudObjectStore) GetSplitProvider() string`

GetSplitProvider returns the SplitProvider field if non-nil, zero value otherwise.

### GetSplitProviderOk

`func (o *CloudObjectStore) GetSplitProviderOk() (*string, bool)`

GetSplitProviderOk returns a tuple with the SplitProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitProvider

`func (o *CloudObjectStore) SetSplitProvider(v string)`

SetSplitProvider sets SplitProvider field to given value.

### HasSplitProvider

`func (o *CloudObjectStore) HasSplitProvider() bool`

HasSplitProvider returns a boolean if a field has been set.

### GetState

`func (o *CloudObjectStore) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CloudObjectStore) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CloudObjectStore) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *CloudObjectStore) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageProvider

`func (o *CloudObjectStore) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *CloudObjectStore) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *CloudObjectStore) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *CloudObjectStore) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetStorageSubpath

`func (o *CloudObjectStore) GetStorageSubpath() string`

GetStorageSubpath returns the StorageSubpath field if non-nil, zero value otherwise.

### GetStorageSubpathOk

`func (o *CloudObjectStore) GetStorageSubpathOk() (*string, bool)`

GetStorageSubpathOk returns a tuple with the StorageSubpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSubpath

`func (o *CloudObjectStore) SetStorageSubpath(v string)`

SetStorageSubpath sets StorageSubpath field to given value.

### HasStorageSubpath

`func (o *CloudObjectStore) HasStorageSubpath() bool`

HasStorageSubpath returns a boolean if a field has been set.

### GetSuggestionCount

`func (o *CloudObjectStore) GetSuggestionCount() int64`

GetSuggestionCount returns the SuggestionCount field if non-nil, zero value otherwise.

### GetSuggestionCountOk

`func (o *CloudObjectStore) GetSuggestionCountOk() (*int64, bool)`

GetSuggestionCountOk returns a tuple with the SuggestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionCount

`func (o *CloudObjectStore) SetSuggestionCount(v int64)`

SetSuggestionCount sets SuggestionCount field to given value.

### HasSuggestionCount

`func (o *CloudObjectStore) HasSuggestionCount() bool`

HasSuggestionCount returns a boolean if a field has been set.

### GetTextToSpeechProvider

`func (o *CloudObjectStore) GetTextToSpeechProvider() string`

GetTextToSpeechProvider returns the TextToSpeechProvider field if non-nil, zero value otherwise.

### GetTextToSpeechProviderOk

`func (o *CloudObjectStore) GetTextToSpeechProviderOk() (*string, bool)`

GetTextToSpeechProviderOk returns a tuple with the TextToSpeechProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextToSpeechProvider

`func (o *CloudObjectStore) SetTextToSpeechProvider(v string)`

SetTextToSpeechProvider sets TextToSpeechProvider field to given value.

### HasTextToSpeechProvider

`func (o *CloudObjectStore) HasTextToSpeechProvider() bool`

HasTextToSpeechProvider returns a boolean if a field has been set.

### GetThemeColor

`func (o *CloudObjectStore) GetThemeColor() string`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *CloudObjectStore) GetThemeColorOk() (*string, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColor

`func (o *CloudObjectStore) SetThemeColor(v string)`

SetThemeColor sets ThemeColor field to given value.

### HasThemeColor

`func (o *CloudObjectStore) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.

### GetTitle

`func (o *CloudObjectStore) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CloudObjectStore) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CloudObjectStore) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CloudObjectStore) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetWelcome

`func (o *CloudObjectStore) GetWelcome() string`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *CloudObjectStore) GetWelcomeOk() (*string, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *CloudObjectStore) SetWelcome(v string)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *CloudObjectStore) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.

### GetWelcomeText

`func (o *CloudObjectStore) GetWelcomeText() string`

GetWelcomeText returns the WelcomeText field if non-nil, zero value otherwise.

### GetWelcomeTextOk

`func (o *CloudObjectStore) GetWelcomeTextOk() (*string, bool)`

GetWelcomeTextOk returns a tuple with the WelcomeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeText

`func (o *CloudObjectStore) SetWelcomeText(v string)`

SetWelcomeText sets WelcomeText field to given value.

### HasWelcomeText

`func (o *CloudObjectStore) HasWelcomeText() bool`

HasWelcomeText returns a boolean if a field has been set.

### GetWelcomeTitle

`func (o *CloudObjectStore) GetWelcomeTitle() string`

GetWelcomeTitle returns the WelcomeTitle field if non-nil, zero value otherwise.

### GetWelcomeTitleOk

`func (o *CloudObjectStore) GetWelcomeTitleOk() (*string, bool)`

GetWelcomeTitleOk returns a tuple with the WelcomeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTitle

`func (o *CloudObjectStore) SetWelcomeTitle(v string)`

SetWelcomeTitle sets WelcomeTitle field to given value.

### HasWelcomeTitle

`func (o *CloudObjectStore) HasWelcomeTitle() bool`

HasWelcomeTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


