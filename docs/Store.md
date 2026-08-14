# Store

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentProvider** | Pointer to **string** |  | [optional] 
**Avatar** | Pointer to **string** |  | [optional] 
**BuiltinTools** | Pointer to **[]string** |  | [optional] 
**ChatCount** | Pointer to **int32** |  | [optional] 
**ChildModelProviders** | Pointer to **[]string** |  | [optional] 
**ChildStores** | Pointer to **[]string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**DisableFileUpload** | Pointer to **bool** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**EmbeddingProvider** | Pointer to **string** |  | [optional] 
**EnableTtsStreaming** | Pointer to **bool** |  | [optional] 
**ExampleQuestions** | Pointer to [**[]ExampleQuestion**](ExampleQuestion.md) |  | [optional] 
**FaviconUrl** | Pointer to **string** |  | [optional] 
**FileTree** | Pointer to [**TreeFile**](TreeFile.md) |  | [optional] 
**FooterHtml** | Pointer to **string** |  | [optional] 
**ForbiddenWords** | Pointer to **[]string** |  | [optional] 
**Frequency** | Pointer to **int32** |  | [optional] 
**HideThinking** | Pointer to **bool** |  | [optional] 
**HtmlTitle** | Pointer to **string** |  | [optional] 
**ImageProvider** | Pointer to **string** |  | [optional] 
**IsDefault** | Pointer to **bool** |  | [optional] 
**KnowledgeCount** | Pointer to **int32** |  | [optional] 
**LimitMinutes** | Pointer to **int32** |  | [optional] 
**LogoUrl** | Pointer to **string** |  | [optional] 
**MemoryLimit** | Pointer to **int32** |  | [optional] 
**MessageCount** | Pointer to **int32** |  | [optional] 
**ModelProvider** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NavItems** | Pointer to **[]string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Prompt** | Pointer to **string** |  | [optional] 
**PropertiesMap** | Pointer to [**map[string]Properties**](Properties.md) |  | [optional] 
**SearchProvider** | Pointer to **string** |  | [optional] 
**ShowAutoRead** | Pointer to **bool** |  | [optional] 
**SpeechToTextProvider** | Pointer to **string** |  | [optional] 
**SplitProvider** | Pointer to **string** |  | [optional] 
**State** | Pointer to **string** |  | [optional] 
**StorageProvider** | Pointer to **string** |  | [optional] 
**StorageSubpath** | Pointer to **string** |  | [optional] 
**SuggestionCount** | Pointer to **int32** |  | [optional] 
**TextToSpeechProvider** | Pointer to **string** |  | [optional] 
**ThemeColor** | Pointer to **string** |  | [optional] 
**Title** | Pointer to **string** |  | [optional] 
**VectorStoreId** | Pointer to **string** |  | [optional] 
**VectorStores** | Pointer to **[]string** |  | [optional] 
**Welcome** | Pointer to **string** |  | [optional] 
**WelcomeText** | Pointer to **string** |  | [optional] 
**WelcomeTitle** | Pointer to **string** |  | [optional] 

## Methods

### NewStore

`func NewStore() *Store`

NewStore instantiates a new Store object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoreWithDefaults

`func NewStoreWithDefaults() *Store`

NewStoreWithDefaults instantiates a new Store object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentProvider

`func (o *Store) GetAgentProvider() string`

GetAgentProvider returns the AgentProvider field if non-nil, zero value otherwise.

### GetAgentProviderOk

`func (o *Store) GetAgentProviderOk() (*string, bool)`

GetAgentProviderOk returns a tuple with the AgentProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentProvider

`func (o *Store) SetAgentProvider(v string)`

SetAgentProvider sets AgentProvider field to given value.

### HasAgentProvider

`func (o *Store) HasAgentProvider() bool`

HasAgentProvider returns a boolean if a field has been set.

### GetAvatar

`func (o *Store) GetAvatar() string`

GetAvatar returns the Avatar field if non-nil, zero value otherwise.

### GetAvatarOk

`func (o *Store) GetAvatarOk() (*string, bool)`

GetAvatarOk returns a tuple with the Avatar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvatar

`func (o *Store) SetAvatar(v string)`

SetAvatar sets Avatar field to given value.

### HasAvatar

`func (o *Store) HasAvatar() bool`

HasAvatar returns a boolean if a field has been set.

### GetBuiltinTools

`func (o *Store) GetBuiltinTools() []string`

GetBuiltinTools returns the BuiltinTools field if non-nil, zero value otherwise.

### GetBuiltinToolsOk

`func (o *Store) GetBuiltinToolsOk() (*[]string, bool)`

GetBuiltinToolsOk returns a tuple with the BuiltinTools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuiltinTools

`func (o *Store) SetBuiltinTools(v []string)`

SetBuiltinTools sets BuiltinTools field to given value.

### HasBuiltinTools

`func (o *Store) HasBuiltinTools() bool`

HasBuiltinTools returns a boolean if a field has been set.

### GetChatCount

`func (o *Store) GetChatCount() int32`

GetChatCount returns the ChatCount field if non-nil, zero value otherwise.

### GetChatCountOk

`func (o *Store) GetChatCountOk() (*int32, bool)`

GetChatCountOk returns a tuple with the ChatCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChatCount

`func (o *Store) SetChatCount(v int32)`

SetChatCount sets ChatCount field to given value.

### HasChatCount

`func (o *Store) HasChatCount() bool`

HasChatCount returns a boolean if a field has been set.

### GetChildModelProviders

`func (o *Store) GetChildModelProviders() []string`

GetChildModelProviders returns the ChildModelProviders field if non-nil, zero value otherwise.

### GetChildModelProvidersOk

`func (o *Store) GetChildModelProvidersOk() (*[]string, bool)`

GetChildModelProvidersOk returns a tuple with the ChildModelProviders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildModelProviders

`func (o *Store) SetChildModelProviders(v []string)`

SetChildModelProviders sets ChildModelProviders field to given value.

### HasChildModelProviders

`func (o *Store) HasChildModelProviders() bool`

HasChildModelProviders returns a boolean if a field has been set.

### GetChildStores

`func (o *Store) GetChildStores() []string`

GetChildStores returns the ChildStores field if non-nil, zero value otherwise.

### GetChildStoresOk

`func (o *Store) GetChildStoresOk() (*[]string, bool)`

GetChildStoresOk returns a tuple with the ChildStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChildStores

`func (o *Store) SetChildStores(v []string)`

SetChildStores sets ChildStores field to given value.

### HasChildStores

`func (o *Store) HasChildStores() bool`

HasChildStores returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Store) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Store) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Store) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Store) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetDisableFileUpload

`func (o *Store) GetDisableFileUpload() bool`

GetDisableFileUpload returns the DisableFileUpload field if non-nil, zero value otherwise.

### GetDisableFileUploadOk

`func (o *Store) GetDisableFileUploadOk() (*bool, bool)`

GetDisableFileUploadOk returns a tuple with the DisableFileUpload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisableFileUpload

`func (o *Store) SetDisableFileUpload(v bool)`

SetDisableFileUpload sets DisableFileUpload field to given value.

### HasDisableFileUpload

`func (o *Store) HasDisableFileUpload() bool`

HasDisableFileUpload returns a boolean if a field has been set.

### GetDisplayName

`func (o *Store) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Store) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Store) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Store) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *Store) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *Store) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *Store) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *Store) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetEnableTtsStreaming

`func (o *Store) GetEnableTtsStreaming() bool`

GetEnableTtsStreaming returns the EnableTtsStreaming field if non-nil, zero value otherwise.

### GetEnableTtsStreamingOk

`func (o *Store) GetEnableTtsStreamingOk() (*bool, bool)`

GetEnableTtsStreamingOk returns a tuple with the EnableTtsStreaming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnableTtsStreaming

`func (o *Store) SetEnableTtsStreaming(v bool)`

SetEnableTtsStreaming sets EnableTtsStreaming field to given value.

### HasEnableTtsStreaming

`func (o *Store) HasEnableTtsStreaming() bool`

HasEnableTtsStreaming returns a boolean if a field has been set.

### GetExampleQuestions

`func (o *Store) GetExampleQuestions() []ExampleQuestion`

GetExampleQuestions returns the ExampleQuestions field if non-nil, zero value otherwise.

### GetExampleQuestionsOk

`func (o *Store) GetExampleQuestionsOk() (*[]ExampleQuestion, bool)`

GetExampleQuestionsOk returns a tuple with the ExampleQuestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExampleQuestions

`func (o *Store) SetExampleQuestions(v []ExampleQuestion)`

SetExampleQuestions sets ExampleQuestions field to given value.

### HasExampleQuestions

`func (o *Store) HasExampleQuestions() bool`

HasExampleQuestions returns a boolean if a field has been set.

### GetFaviconUrl

`func (o *Store) GetFaviconUrl() string`

GetFaviconUrl returns the FaviconUrl field if non-nil, zero value otherwise.

### GetFaviconUrlOk

`func (o *Store) GetFaviconUrlOk() (*string, bool)`

GetFaviconUrlOk returns a tuple with the FaviconUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFaviconUrl

`func (o *Store) SetFaviconUrl(v string)`

SetFaviconUrl sets FaviconUrl field to given value.

### HasFaviconUrl

`func (o *Store) HasFaviconUrl() bool`

HasFaviconUrl returns a boolean if a field has been set.

### GetFileTree

`func (o *Store) GetFileTree() TreeFile`

GetFileTree returns the FileTree field if non-nil, zero value otherwise.

### GetFileTreeOk

`func (o *Store) GetFileTreeOk() (*TreeFile, bool)`

GetFileTreeOk returns a tuple with the FileTree field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileTree

`func (o *Store) SetFileTree(v TreeFile)`

SetFileTree sets FileTree field to given value.

### HasFileTree

`func (o *Store) HasFileTree() bool`

HasFileTree returns a boolean if a field has been set.

### GetFooterHtml

`func (o *Store) GetFooterHtml() string`

GetFooterHtml returns the FooterHtml field if non-nil, zero value otherwise.

### GetFooterHtmlOk

`func (o *Store) GetFooterHtmlOk() (*string, bool)`

GetFooterHtmlOk returns a tuple with the FooterHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFooterHtml

`func (o *Store) SetFooterHtml(v string)`

SetFooterHtml sets FooterHtml field to given value.

### HasFooterHtml

`func (o *Store) HasFooterHtml() bool`

HasFooterHtml returns a boolean if a field has been set.

### GetForbiddenWords

`func (o *Store) GetForbiddenWords() []string`

GetForbiddenWords returns the ForbiddenWords field if non-nil, zero value otherwise.

### GetForbiddenWordsOk

`func (o *Store) GetForbiddenWordsOk() (*[]string, bool)`

GetForbiddenWordsOk returns a tuple with the ForbiddenWords field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForbiddenWords

`func (o *Store) SetForbiddenWords(v []string)`

SetForbiddenWords sets ForbiddenWords field to given value.

### HasForbiddenWords

`func (o *Store) HasForbiddenWords() bool`

HasForbiddenWords returns a boolean if a field has been set.

### GetFrequency

`func (o *Store) GetFrequency() int32`

GetFrequency returns the Frequency field if non-nil, zero value otherwise.

### GetFrequencyOk

`func (o *Store) GetFrequencyOk() (*int32, bool)`

GetFrequencyOk returns a tuple with the Frequency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrequency

`func (o *Store) SetFrequency(v int32)`

SetFrequency sets Frequency field to given value.

### HasFrequency

`func (o *Store) HasFrequency() bool`

HasFrequency returns a boolean if a field has been set.

### GetHideThinking

`func (o *Store) GetHideThinking() bool`

GetHideThinking returns the HideThinking field if non-nil, zero value otherwise.

### GetHideThinkingOk

`func (o *Store) GetHideThinkingOk() (*bool, bool)`

GetHideThinkingOk returns a tuple with the HideThinking field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHideThinking

`func (o *Store) SetHideThinking(v bool)`

SetHideThinking sets HideThinking field to given value.

### HasHideThinking

`func (o *Store) HasHideThinking() bool`

HasHideThinking returns a boolean if a field has been set.

### GetHtmlTitle

`func (o *Store) GetHtmlTitle() string`

GetHtmlTitle returns the HtmlTitle field if non-nil, zero value otherwise.

### GetHtmlTitleOk

`func (o *Store) GetHtmlTitleOk() (*string, bool)`

GetHtmlTitleOk returns a tuple with the HtmlTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHtmlTitle

`func (o *Store) SetHtmlTitle(v string)`

SetHtmlTitle sets HtmlTitle field to given value.

### HasHtmlTitle

`func (o *Store) HasHtmlTitle() bool`

HasHtmlTitle returns a boolean if a field has been set.

### GetImageProvider

`func (o *Store) GetImageProvider() string`

GetImageProvider returns the ImageProvider field if non-nil, zero value otherwise.

### GetImageProviderOk

`func (o *Store) GetImageProviderOk() (*string, bool)`

GetImageProviderOk returns a tuple with the ImageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageProvider

`func (o *Store) SetImageProvider(v string)`

SetImageProvider sets ImageProvider field to given value.

### HasImageProvider

`func (o *Store) HasImageProvider() bool`

HasImageProvider returns a boolean if a field has been set.

### GetIsDefault

`func (o *Store) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *Store) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *Store) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.

### HasIsDefault

`func (o *Store) HasIsDefault() bool`

HasIsDefault returns a boolean if a field has been set.

### GetKnowledgeCount

`func (o *Store) GetKnowledgeCount() int32`

GetKnowledgeCount returns the KnowledgeCount field if non-nil, zero value otherwise.

### GetKnowledgeCountOk

`func (o *Store) GetKnowledgeCountOk() (*int32, bool)`

GetKnowledgeCountOk returns a tuple with the KnowledgeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKnowledgeCount

`func (o *Store) SetKnowledgeCount(v int32)`

SetKnowledgeCount sets KnowledgeCount field to given value.

### HasKnowledgeCount

`func (o *Store) HasKnowledgeCount() bool`

HasKnowledgeCount returns a boolean if a field has been set.

### GetLimitMinutes

`func (o *Store) GetLimitMinutes() int32`

GetLimitMinutes returns the LimitMinutes field if non-nil, zero value otherwise.

### GetLimitMinutesOk

`func (o *Store) GetLimitMinutesOk() (*int32, bool)`

GetLimitMinutesOk returns a tuple with the LimitMinutes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimitMinutes

`func (o *Store) SetLimitMinutes(v int32)`

SetLimitMinutes sets LimitMinutes field to given value.

### HasLimitMinutes

`func (o *Store) HasLimitMinutes() bool`

HasLimitMinutes returns a boolean if a field has been set.

### GetLogoUrl

`func (o *Store) GetLogoUrl() string`

GetLogoUrl returns the LogoUrl field if non-nil, zero value otherwise.

### GetLogoUrlOk

`func (o *Store) GetLogoUrlOk() (*string, bool)`

GetLogoUrlOk returns a tuple with the LogoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoUrl

`func (o *Store) SetLogoUrl(v string)`

SetLogoUrl sets LogoUrl field to given value.

### HasLogoUrl

`func (o *Store) HasLogoUrl() bool`

HasLogoUrl returns a boolean if a field has been set.

### GetMemoryLimit

`func (o *Store) GetMemoryLimit() int32`

GetMemoryLimit returns the MemoryLimit field if non-nil, zero value otherwise.

### GetMemoryLimitOk

`func (o *Store) GetMemoryLimitOk() (*int32, bool)`

GetMemoryLimitOk returns a tuple with the MemoryLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryLimit

`func (o *Store) SetMemoryLimit(v int32)`

SetMemoryLimit sets MemoryLimit field to given value.

### HasMemoryLimit

`func (o *Store) HasMemoryLimit() bool`

HasMemoryLimit returns a boolean if a field has been set.

### GetMessageCount

`func (o *Store) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *Store) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *Store) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *Store) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetModelProvider

`func (o *Store) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *Store) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *Store) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *Store) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *Store) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Store) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Store) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Store) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNavItems

`func (o *Store) GetNavItems() []string`

GetNavItems returns the NavItems field if non-nil, zero value otherwise.

### GetNavItemsOk

`func (o *Store) GetNavItemsOk() (*[]string, bool)`

GetNavItemsOk returns a tuple with the NavItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNavItems

`func (o *Store) SetNavItems(v []string)`

SetNavItems sets NavItems field to given value.

### HasNavItems

`func (o *Store) HasNavItems() bool`

HasNavItems returns a boolean if a field has been set.

### GetOwner

`func (o *Store) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Store) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Store) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Store) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrompt

`func (o *Store) GetPrompt() string`

GetPrompt returns the Prompt field if non-nil, zero value otherwise.

### GetPromptOk

`func (o *Store) GetPromptOk() (*string, bool)`

GetPromptOk returns a tuple with the Prompt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrompt

`func (o *Store) SetPrompt(v string)`

SetPrompt sets Prompt field to given value.

### HasPrompt

`func (o *Store) HasPrompt() bool`

HasPrompt returns a boolean if a field has been set.

### GetPropertiesMap

`func (o *Store) GetPropertiesMap() map[string]Properties`

GetPropertiesMap returns the PropertiesMap field if non-nil, zero value otherwise.

### GetPropertiesMapOk

`func (o *Store) GetPropertiesMapOk() (*map[string]Properties, bool)`

GetPropertiesMapOk returns a tuple with the PropertiesMap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPropertiesMap

`func (o *Store) SetPropertiesMap(v map[string]Properties)`

SetPropertiesMap sets PropertiesMap field to given value.

### HasPropertiesMap

`func (o *Store) HasPropertiesMap() bool`

HasPropertiesMap returns a boolean if a field has been set.

### GetSearchProvider

`func (o *Store) GetSearchProvider() string`

GetSearchProvider returns the SearchProvider field if non-nil, zero value otherwise.

### GetSearchProviderOk

`func (o *Store) GetSearchProviderOk() (*string, bool)`

GetSearchProviderOk returns a tuple with the SearchProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchProvider

`func (o *Store) SetSearchProvider(v string)`

SetSearchProvider sets SearchProvider field to given value.

### HasSearchProvider

`func (o *Store) HasSearchProvider() bool`

HasSearchProvider returns a boolean if a field has been set.

### GetShowAutoRead

`func (o *Store) GetShowAutoRead() bool`

GetShowAutoRead returns the ShowAutoRead field if non-nil, zero value otherwise.

### GetShowAutoReadOk

`func (o *Store) GetShowAutoReadOk() (*bool, bool)`

GetShowAutoReadOk returns a tuple with the ShowAutoRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShowAutoRead

`func (o *Store) SetShowAutoRead(v bool)`

SetShowAutoRead sets ShowAutoRead field to given value.

### HasShowAutoRead

`func (o *Store) HasShowAutoRead() bool`

HasShowAutoRead returns a boolean if a field has been set.

### GetSpeechToTextProvider

`func (o *Store) GetSpeechToTextProvider() string`

GetSpeechToTextProvider returns the SpeechToTextProvider field if non-nil, zero value otherwise.

### GetSpeechToTextProviderOk

`func (o *Store) GetSpeechToTextProviderOk() (*string, bool)`

GetSpeechToTextProviderOk returns a tuple with the SpeechToTextProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpeechToTextProvider

`func (o *Store) SetSpeechToTextProvider(v string)`

SetSpeechToTextProvider sets SpeechToTextProvider field to given value.

### HasSpeechToTextProvider

`func (o *Store) HasSpeechToTextProvider() bool`

HasSpeechToTextProvider returns a boolean if a field has been set.

### GetSplitProvider

`func (o *Store) GetSplitProvider() string`

GetSplitProvider returns the SplitProvider field if non-nil, zero value otherwise.

### GetSplitProviderOk

`func (o *Store) GetSplitProviderOk() (*string, bool)`

GetSplitProviderOk returns a tuple with the SplitProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSplitProvider

`func (o *Store) SetSplitProvider(v string)`

SetSplitProvider sets SplitProvider field to given value.

### HasSplitProvider

`func (o *Store) HasSplitProvider() bool`

HasSplitProvider returns a boolean if a field has been set.

### GetState

`func (o *Store) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Store) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Store) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Store) HasState() bool`

HasState returns a boolean if a field has been set.

### GetStorageProvider

`func (o *Store) GetStorageProvider() string`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *Store) GetStorageProviderOk() (*string, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *Store) SetStorageProvider(v string)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *Store) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetStorageSubpath

`func (o *Store) GetStorageSubpath() string`

GetStorageSubpath returns the StorageSubpath field if non-nil, zero value otherwise.

### GetStorageSubpathOk

`func (o *Store) GetStorageSubpathOk() (*string, bool)`

GetStorageSubpathOk returns a tuple with the StorageSubpath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSubpath

`func (o *Store) SetStorageSubpath(v string)`

SetStorageSubpath sets StorageSubpath field to given value.

### HasStorageSubpath

`func (o *Store) HasStorageSubpath() bool`

HasStorageSubpath returns a boolean if a field has been set.

### GetSuggestionCount

`func (o *Store) GetSuggestionCount() int32`

GetSuggestionCount returns the SuggestionCount field if non-nil, zero value otherwise.

### GetSuggestionCountOk

`func (o *Store) GetSuggestionCountOk() (*int32, bool)`

GetSuggestionCountOk returns a tuple with the SuggestionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestionCount

`func (o *Store) SetSuggestionCount(v int32)`

SetSuggestionCount sets SuggestionCount field to given value.

### HasSuggestionCount

`func (o *Store) HasSuggestionCount() bool`

HasSuggestionCount returns a boolean if a field has been set.

### GetTextToSpeechProvider

`func (o *Store) GetTextToSpeechProvider() string`

GetTextToSpeechProvider returns the TextToSpeechProvider field if non-nil, zero value otherwise.

### GetTextToSpeechProviderOk

`func (o *Store) GetTextToSpeechProviderOk() (*string, bool)`

GetTextToSpeechProviderOk returns a tuple with the TextToSpeechProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextToSpeechProvider

`func (o *Store) SetTextToSpeechProvider(v string)`

SetTextToSpeechProvider sets TextToSpeechProvider field to given value.

### HasTextToSpeechProvider

`func (o *Store) HasTextToSpeechProvider() bool`

HasTextToSpeechProvider returns a boolean if a field has been set.

### GetThemeColor

`func (o *Store) GetThemeColor() string`

GetThemeColor returns the ThemeColor field if non-nil, zero value otherwise.

### GetThemeColorOk

`func (o *Store) GetThemeColorOk() (*string, bool)`

GetThemeColorOk returns a tuple with the ThemeColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThemeColor

`func (o *Store) SetThemeColor(v string)`

SetThemeColor sets ThemeColor field to given value.

### HasThemeColor

`func (o *Store) HasThemeColor() bool`

HasThemeColor returns a boolean if a field has been set.

### GetTitle

`func (o *Store) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *Store) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *Store) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *Store) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### GetVectorStoreId

`func (o *Store) GetVectorStoreId() string`

GetVectorStoreId returns the VectorStoreId field if non-nil, zero value otherwise.

### GetVectorStoreIdOk

`func (o *Store) GetVectorStoreIdOk() (*string, bool)`

GetVectorStoreIdOk returns a tuple with the VectorStoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorStoreId

`func (o *Store) SetVectorStoreId(v string)`

SetVectorStoreId sets VectorStoreId field to given value.

### HasVectorStoreId

`func (o *Store) HasVectorStoreId() bool`

HasVectorStoreId returns a boolean if a field has been set.

### GetVectorStores

`func (o *Store) GetVectorStores() []string`

GetVectorStores returns the VectorStores field if non-nil, zero value otherwise.

### GetVectorStoresOk

`func (o *Store) GetVectorStoresOk() (*[]string, bool)`

GetVectorStoresOk returns a tuple with the VectorStores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorStores

`func (o *Store) SetVectorStores(v []string)`

SetVectorStores sets VectorStores field to given value.

### HasVectorStores

`func (o *Store) HasVectorStores() bool`

HasVectorStores returns a boolean if a field has been set.

### GetWelcome

`func (o *Store) GetWelcome() string`

GetWelcome returns the Welcome field if non-nil, zero value otherwise.

### GetWelcomeOk

`func (o *Store) GetWelcomeOk() (*string, bool)`

GetWelcomeOk returns a tuple with the Welcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcome

`func (o *Store) SetWelcome(v string)`

SetWelcome sets Welcome field to given value.

### HasWelcome

`func (o *Store) HasWelcome() bool`

HasWelcome returns a boolean if a field has been set.

### GetWelcomeText

`func (o *Store) GetWelcomeText() string`

GetWelcomeText returns the WelcomeText field if non-nil, zero value otherwise.

### GetWelcomeTextOk

`func (o *Store) GetWelcomeTextOk() (*string, bool)`

GetWelcomeTextOk returns a tuple with the WelcomeText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeText

`func (o *Store) SetWelcomeText(v string)`

SetWelcomeText sets WelcomeText field to given value.

### HasWelcomeText

`func (o *Store) HasWelcomeText() bool`

HasWelcomeText returns a boolean if a field has been set.

### GetWelcomeTitle

`func (o *Store) GetWelcomeTitle() string`

GetWelcomeTitle returns the WelcomeTitle field if non-nil, zero value otherwise.

### GetWelcomeTitleOk

`func (o *Store) GetWelcomeTitleOk() (*string, bool)`

GetWelcomeTitleOk returns a tuple with the WelcomeTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWelcomeTitle

`func (o *Store) SetWelcomeTitle(v string)`

SetWelcomeTitle sets WelcomeTitle field to given value.

### HasWelcomeTitle

`func (o *Store) HasWelcomeTitle() bool`

HasWelcomeTitle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


