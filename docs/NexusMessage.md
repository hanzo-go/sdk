# NexusMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Author** | Pointer to **string** |  | [optional] 
**Chat** | Pointer to **string** |  | [optional] 
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
**Price** | Pointer to **float64** |  | [optional] 
**ReasonText** | Pointer to **string** |  | [optional] 
**ReplyTo** | Pointer to **string** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**Suggestions** | Pointer to [**[]NexusSuggestion**](NexusSuggestion.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TextTokenCount** | Pointer to **int64** |  | [optional] 
**TokenCount** | Pointer to **int64** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**VectorScores** | Pointer to [**[]NexusVectorScore**](NexusVectorScore.md) |  | [optional] 

## Methods

### NewNexusMessage

`func NewNexusMessage() *NexusMessage`

NewNexusMessage instantiates a new NexusMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusMessageWithDefaults

`func NewNexusMessageWithDefaults() *NexusMessage`

NewNexusMessageWithDefaults instantiates a new NexusMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *NexusMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *NexusMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *NexusMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *NexusMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetChat

`func (o *NexusMessage) GetChat() string`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *NexusMessage) GetChatOk() (*string, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *NexusMessage) SetChat(v string)`

SetChat sets Chat field to given value.

### HasChat

`func (o *NexusMessage) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetComment

`func (o *NexusMessage) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *NexusMessage) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *NexusMessage) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *NexusMessage) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusMessage) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusMessage) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusMessage) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusMessage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *NexusMessage) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NexusMessage) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NexusMessage) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NexusMessage) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDislikeUsers

`func (o *NexusMessage) GetDislikeUsers() []string`

GetDislikeUsers returns the DislikeUsers field if non-nil, zero value otherwise.

### GetDislikeUsersOk

`func (o *NexusMessage) GetDislikeUsersOk() (*[]string, bool)`

GetDislikeUsersOk returns a tuple with the DislikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikeUsers

`func (o *NexusMessage) SetDislikeUsers(v []string)`

SetDislikeUsers sets DislikeUsers field to given value.

### HasDislikeUsers

`func (o *NexusMessage) HasDislikeUsers() bool`

HasDislikeUsers returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *NexusMessage) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *NexusMessage) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *NexusMessage) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *NexusMessage) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetErrorText

`func (o *NexusMessage) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *NexusMessage) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *NexusMessage) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *NexusMessage) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetFileName

`func (o *NexusMessage) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *NexusMessage) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *NexusMessage) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *NexusMessage) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIsAlerted

`func (o *NexusMessage) GetIsAlerted() bool`

GetIsAlerted returns the IsAlerted field if non-nil, zero value otherwise.

### GetIsAlertedOk

`func (o *NexusMessage) GetIsAlertedOk() (*bool, bool)`

GetIsAlertedOk returns a tuple with the IsAlerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlerted

`func (o *NexusMessage) SetIsAlerted(v bool)`

SetIsAlerted sets IsAlerted field to given value.

### HasIsAlerted

`func (o *NexusMessage) HasIsAlerted() bool`

HasIsAlerted returns a boolean if a field has been set.

### GetIsDeleted

`func (o *NexusMessage) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *NexusMessage) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *NexusMessage) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *NexusMessage) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsHidden

`func (o *NexusMessage) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *NexusMessage) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *NexusMessage) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *NexusMessage) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsRegenerated

`func (o *NexusMessage) GetIsRegenerated() bool`

GetIsRegenerated returns the IsRegenerated field if non-nil, zero value otherwise.

### GetIsRegeneratedOk

`func (o *NexusMessage) GetIsRegeneratedOk() (*bool, bool)`

GetIsRegeneratedOk returns a tuple with the IsRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegenerated

`func (o *NexusMessage) SetIsRegenerated(v bool)`

SetIsRegenerated sets IsRegenerated field to given value.

### HasIsRegenerated

`func (o *NexusMessage) HasIsRegenerated() bool`

HasIsRegenerated returns a boolean if a field has been set.

### GetLikeUsers

`func (o *NexusMessage) GetLikeUsers() []string`

GetLikeUsers returns the LikeUsers field if non-nil, zero value otherwise.

### GetLikeUsersOk

`func (o *NexusMessage) GetLikeUsersOk() (*[]string, bool)`

GetLikeUsersOk returns a tuple with the LikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeUsers

`func (o *NexusMessage) SetLikeUsers(v []string)`

SetLikeUsers sets LikeUsers field to given value.

### HasLikeUsers

`func (o *NexusMessage) HasLikeUsers() bool`

HasLikeUsers returns a boolean if a field has been set.

### GetModelProvider

`func (o *NexusMessage) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *NexusMessage) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *NexusMessage) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *NexusMessage) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *NexusMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedNotify

`func (o *NexusMessage) GetNeedNotify() bool`

GetNeedNotify returns the NeedNotify field if non-nil, zero value otherwise.

### GetNeedNotifyOk

`func (o *NexusMessage) GetNeedNotifyOk() (*bool, bool)`

GetNeedNotifyOk returns a tuple with the NeedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNotify

`func (o *NexusMessage) SetNeedNotify(v bool)`

SetNeedNotify sets NeedNotify field to given value.

### HasNeedNotify

`func (o *NexusMessage) HasNeedNotify() bool`

HasNeedNotify returns a boolean if a field has been set.

### GetOrganization

`func (o *NexusMessage) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NexusMessage) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NexusMessage) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *NexusMessage) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *NexusMessage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusMessage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusMessage) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusMessage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *NexusMessage) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NexusMessage) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NexusMessage) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NexusMessage) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReasonText

`func (o *NexusMessage) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *NexusMessage) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *NexusMessage) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *NexusMessage) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetReplyTo

`func (o *NexusMessage) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *NexusMessage) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *NexusMessage) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *NexusMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetStore

`func (o *NexusMessage) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *NexusMessage) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *NexusMessage) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *NexusMessage) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSuggestions

`func (o *NexusMessage) GetSuggestions() []NexusSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *NexusMessage) GetSuggestionsOk() (*[]NexusSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *NexusMessage) SetSuggestions(v []NexusSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *NexusMessage) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetText

`func (o *NexusMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *NexusMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *NexusMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *NexusMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextTokenCount

`func (o *NexusMessage) GetTextTokenCount() int64`

GetTextTokenCount returns the TextTokenCount field if non-nil, zero value otherwise.

### GetTextTokenCountOk

`func (o *NexusMessage) GetTextTokenCountOk() (*int64, bool)`

GetTextTokenCountOk returns a tuple with the TextTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTokenCount

`func (o *NexusMessage) SetTextTokenCount(v int64)`

SetTextTokenCount sets TextTokenCount field to given value.

### HasTextTokenCount

`func (o *NexusMessage) HasTextTokenCount() bool`

HasTextTokenCount returns a boolean if a field has been set.

### GetTokenCount

`func (o *NexusMessage) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *NexusMessage) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *NexusMessage) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *NexusMessage) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetUser

`func (o *NexusMessage) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NexusMessage) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NexusMessage) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *NexusMessage) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVectorScores

`func (o *NexusMessage) GetVectorScores() []NexusVectorScore`

GetVectorScores returns the VectorScores field if non-nil, zero value otherwise.

### GetVectorScoresOk

`func (o *NexusMessage) GetVectorScoresOk() (*[]NexusVectorScore, bool)`

GetVectorScoresOk returns a tuple with the VectorScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorScores

`func (o *NexusMessage) SetVectorScores(v []NexusVectorScore)`

SetVectorScores sets VectorScores field to given value.

### HasVectorScores

`func (o *NexusMessage) HasVectorScores() bool`

HasVectorScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


