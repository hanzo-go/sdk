# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnsweredTime** | Pointer to **string** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**Chat** | Pointer to **string** |  | [optional] 
**ClaimedTime** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DislikeUsers** | Pointer to **[]string** |  | [optional] 
**EmbeddingProvider** | Pointer to **string** |  | [optional] 
**ErrorText** | Pointer to **string** |  | [optional] 
**FileName** | Pointer to **string** |  | [optional] 
**IsAlerted** | Pointer to **bool** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**IsRegenerated** | Pointer to **bool** |  | [optional] 
**LikeUsers** | Pointer to **[]string** |  | [optional] 
**ModelProvider** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedNotify** | Pointer to **bool** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**ReasonText** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**SearchResults** | Pointer to [**[]ModelSearchResult**](ModelSearchResult.md) |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**Suggestions** | Pointer to [**[]Suggestion**](Suggestion.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TextTokenCount** | Pointer to **int32** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**ToolCalls** | Pointer to [**[]ModelToolCall**](ModelToolCall.md) |  | [optional] 
**TransactionId** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**VectorScores** | Pointer to [**[]VectorScore**](VectorScore.md) |  | [optional] 
**WebSearchEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnsweredTime

`func (o *Message) GetAnsweredTime() string`

GetAnsweredTime returns the AnsweredTime field if non-nil, zero value otherwise.

### GetAnsweredTimeOk

`func (o *Message) GetAnsweredTimeOk() (*string, bool)`

GetAnsweredTimeOk returns a tuple with the AnsweredTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredTime

`func (o *Message) SetAnsweredTime(v string)`

SetAnsweredTime sets AnsweredTime field to given value.

### HasAnsweredTime

`func (o *Message) HasAnsweredTime() bool`

HasAnsweredTime returns a boolean if a field has been set.

### GetAuthor

`func (o *Message) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *Message) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *Message) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *Message) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetChat

`func (o *Message) GetChat() string`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *Message) GetChatOk() (*string, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *Message) SetChat(v string)`

SetChat sets Chat field to given value.

### HasChat

`func (o *Message) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetClaimedTime

`func (o *Message) GetClaimedTime() string`

GetClaimedTime returns the ClaimedTime field if non-nil, zero value otherwise.

### GetClaimedTimeOk

`func (o *Message) GetClaimedTimeOk() (*string, bool)`

GetClaimedTimeOk returns a tuple with the ClaimedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimedTime

`func (o *Message) SetClaimedTime(v string)`

SetClaimedTime sets ClaimedTime field to given value.

### HasClaimedTime

`func (o *Message) HasClaimedTime() bool`

HasClaimedTime returns a boolean if a field has been set.

### GetComment

`func (o *Message) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *Message) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *Message) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *Message) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Message) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Message) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Message) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Message) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *Message) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Message) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Message) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Message) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDislikeUsers

`func (o *Message) GetDislikeUsers() []string`

GetDislikeUsers returns the DislikeUsers field if non-nil, zero value otherwise.

### GetDislikeUsersOk

`func (o *Message) GetDislikeUsersOk() (*[]string, bool)`

GetDislikeUsersOk returns a tuple with the DislikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikeUsers

`func (o *Message) SetDislikeUsers(v []string)`

SetDislikeUsers sets DislikeUsers field to given value.

### HasDislikeUsers

`func (o *Message) HasDislikeUsers() bool`

HasDislikeUsers returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *Message) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *Message) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *Message) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *Message) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetErrorText

`func (o *Message) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *Message) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *Message) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *Message) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetFileName

`func (o *Message) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *Message) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *Message) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *Message) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIsAlerted

`func (o *Message) GetIsAlerted() bool`

GetIsAlerted returns the IsAlerted field if non-nil, zero value otherwise.

### GetIsAlertedOk

`func (o *Message) GetIsAlertedOk() (*bool, bool)`

GetIsAlertedOk returns a tuple with the IsAlerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlerted

`func (o *Message) SetIsAlerted(v bool)`

SetIsAlerted sets IsAlerted field to given value.

### HasIsAlerted

`func (o *Message) HasIsAlerted() bool`

HasIsAlerted returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Message) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Message) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Message) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Message) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsHidden

`func (o *Message) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *Message) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *Message) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *Message) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsRegenerated

`func (o *Message) GetIsRegenerated() bool`

GetIsRegenerated returns the IsRegenerated field if non-nil, zero value otherwise.

### GetIsRegeneratedOk

`func (o *Message) GetIsRegeneratedOk() (*bool, bool)`

GetIsRegeneratedOk returns a tuple with the IsRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegenerated

`func (o *Message) SetIsRegenerated(v bool)`

SetIsRegenerated sets IsRegenerated field to given value.

### HasIsRegenerated

`func (o *Message) HasIsRegenerated() bool`

HasIsRegenerated returns a boolean if a field has been set.

### GetLikeUsers

`func (o *Message) GetLikeUsers() []string`

GetLikeUsers returns the LikeUsers field if non-nil, zero value otherwise.

### GetLikeUsersOk

`func (o *Message) GetLikeUsersOk() (*[]string, bool)`

GetLikeUsersOk returns a tuple with the LikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeUsers

`func (o *Message) SetLikeUsers(v []string)`

SetLikeUsers sets LikeUsers field to given value.

### HasLikeUsers

`func (o *Message) HasLikeUsers() bool`

HasLikeUsers returns a boolean if a field has been set.

### GetModelProvider

`func (o *Message) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *Message) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *Message) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *Message) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *Message) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Message) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Message) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Message) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedNotify

`func (o *Message) GetNeedNotify() bool`

GetNeedNotify returns the NeedNotify field if non-nil, zero value otherwise.

### GetNeedNotifyOk

`func (o *Message) GetNeedNotifyOk() (*bool, bool)`

GetNeedNotifyOk returns a tuple with the NeedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNotify

`func (o *Message) SetNeedNotify(v bool)`

SetNeedNotify sets NeedNotify field to given value.

### HasNeedNotify

`func (o *Message) HasNeedNotify() bool`

HasNeedNotify returns a boolean if a field has been set.

### GetOrganization

`func (o *Message) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Message) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Message) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Message) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *Message) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Message) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Message) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Message) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *Message) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Message) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Message) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Message) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReasonText

`func (o *Message) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *Message) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *Message) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *Message) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetReplyTo

`func (o *Message) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *Message) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *Message) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *Message) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetSearchResults

`func (o *Message) GetSearchResults() []ModelSearchResult`

GetSearchResults returns the SearchResults field if non-nil, zero value otherwise.

### GetSearchResultsOk

`func (o *Message) GetSearchResultsOk() (*[]ModelSearchResult, bool)`

GetSearchResultsOk returns a tuple with the SearchResults field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchResults

`func (o *Message) SetSearchResults(v []ModelSearchResult)`

SetSearchResults sets SearchResults field to given value.

### HasSearchResults

`func (o *Message) HasSearchResults() bool`

HasSearchResults returns a boolean if a field has been set.

### GetStore

`func (o *Message) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *Message) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *Message) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *Message) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSuggestions

`func (o *Message) GetSuggestions() []Suggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *Message) GetSuggestionsOk() (*[]Suggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *Message) SetSuggestions(v []Suggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *Message) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetText

`func (o *Message) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *Message) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *Message) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *Message) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextTokenCount

`func (o *Message) GetTextTokenCount() int32`

GetTextTokenCount returns the TextTokenCount field if non-nil, zero value otherwise.

### GetTextTokenCountOk

`func (o *Message) GetTextTokenCountOk() (*int32, bool)`

GetTextTokenCountOk returns a tuple with the TextTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTokenCount

`func (o *Message) SetTextTokenCount(v int32)`

SetTextTokenCount sets TextTokenCount field to given value.

### HasTextTokenCount

`func (o *Message) HasTextTokenCount() bool`

HasTextTokenCount returns a boolean if a field has been set.

### GetTokenCount

`func (o *Message) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *Message) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *Message) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *Message) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetToolCalls

`func (o *Message) GetToolCalls() []ModelToolCall`

GetToolCalls returns the ToolCalls field if non-nil, zero value otherwise.

### GetToolCallsOk

`func (o *Message) GetToolCallsOk() (*[]ModelToolCall, bool)`

GetToolCallsOk returns a tuple with the ToolCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToolCalls

`func (o *Message) SetToolCalls(v []ModelToolCall)`

SetToolCalls sets ToolCalls field to given value.

### HasToolCalls

`func (o *Message) HasToolCalls() bool`

HasToolCalls returns a boolean if a field has been set.

### GetTransactionId

`func (o *Message) GetTransactionId() string`

GetTransactionId returns the TransactionId field if non-nil, zero value otherwise.

### GetTransactionIdOk

`func (o *Message) GetTransactionIdOk() (*string, bool)`

GetTransactionIdOk returns a tuple with the TransactionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionId

`func (o *Message) SetTransactionId(v string)`

SetTransactionId sets TransactionId field to given value.

### HasTransactionId

`func (o *Message) HasTransactionId() bool`

HasTransactionId returns a boolean if a field has been set.

### GetUser

`func (o *Message) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Message) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Message) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Message) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVectorScores

`func (o *Message) GetVectorScores() []VectorScore`

GetVectorScores returns the VectorScores field if non-nil, zero value otherwise.

### GetVectorScoresOk

`func (o *Message) GetVectorScoresOk() (*[]VectorScore, bool)`

GetVectorScoresOk returns a tuple with the VectorScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorScores

`func (o *Message) SetVectorScores(v []VectorScore)`

SetVectorScores sets VectorScores field to given value.

### HasVectorScores

`func (o *Message) HasVectorScores() bool`

HasVectorScores returns a boolean if a field has been set.

### GetWebSearchEnabled

`func (o *Message) GetWebSearchEnabled() bool`

GetWebSearchEnabled returns the WebSearchEnabled field if non-nil, zero value otherwise.

### GetWebSearchEnabledOk

`func (o *Message) GetWebSearchEnabledOk() (*bool, bool)`

GetWebSearchEnabledOk returns a tuple with the WebSearchEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebSearchEnabled

`func (o *Message) SetWebSearchEnabled(v bool)`

SetWebSearchEnabled sets WebSearchEnabled field to given value.

### HasWebSearchEnabled

`func (o *Message) HasWebSearchEnabled() bool`

HasWebSearchEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


