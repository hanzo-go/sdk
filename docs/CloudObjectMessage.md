# CloudObjectMessage

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
**Suggestions** | Pointer to [**[]CloudObjectSuggestion**](CloudObjectSuggestion.md) |  | [optional] 
**Text** | Pointer to **string** |  | [optional] 
**TextTokenCount** | Pointer to **int64** |  | [optional] 
**TokenCount** | Pointer to **int64** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**VectorScores** | Pointer to [**[]CloudObjectVectorScore**](CloudObjectVectorScore.md) |  | [optional] 

## Methods

### NewCloudObjectMessage

`func NewCloudObjectMessage() *CloudObjectMessage`

NewCloudObjectMessage instantiates a new CloudObjectMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudObjectMessageWithDefaults

`func NewCloudObjectMessageWithDefaults() *CloudObjectMessage`

NewCloudObjectMessageWithDefaults instantiates a new CloudObjectMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthor

`func (o *CloudObjectMessage) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *CloudObjectMessage) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *CloudObjectMessage) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *CloudObjectMessage) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetChat

`func (o *CloudObjectMessage) GetChat() string`

GetChat returns the Chat field if non-nil, zero value otherwise.

### GetChatOk

`func (o *CloudObjectMessage) GetChatOk() (*string, bool)`

GetChatOk returns a tuple with the Chat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChat

`func (o *CloudObjectMessage) SetChat(v string)`

SetChat sets Chat field to given value.

### HasChat

`func (o *CloudObjectMessage) HasChat() bool`

HasChat returns a boolean if a field has been set.

### GetComment

`func (o *CloudObjectMessage) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CloudObjectMessage) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CloudObjectMessage) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CloudObjectMessage) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetCreatedTime

`func (o *CloudObjectMessage) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *CloudObjectMessage) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *CloudObjectMessage) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *CloudObjectMessage) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *CloudObjectMessage) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CloudObjectMessage) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CloudObjectMessage) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CloudObjectMessage) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDislikeUsers

`func (o *CloudObjectMessage) GetDislikeUsers() []string`

GetDislikeUsers returns the DislikeUsers field if non-nil, zero value otherwise.

### GetDislikeUsersOk

`func (o *CloudObjectMessage) GetDislikeUsersOk() (*[]string, bool)`

GetDislikeUsersOk returns a tuple with the DislikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDislikeUsers

`func (o *CloudObjectMessage) SetDislikeUsers(v []string)`

SetDislikeUsers sets DislikeUsers field to given value.

### HasDislikeUsers

`func (o *CloudObjectMessage) HasDislikeUsers() bool`

HasDislikeUsers returns a boolean if a field has been set.

### GetEmbeddingProvider

`func (o *CloudObjectMessage) GetEmbeddingProvider() string`

GetEmbeddingProvider returns the EmbeddingProvider field if non-nil, zero value otherwise.

### GetEmbeddingProviderOk

`func (o *CloudObjectMessage) GetEmbeddingProviderOk() (*string, bool)`

GetEmbeddingProviderOk returns a tuple with the EmbeddingProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmbeddingProvider

`func (o *CloudObjectMessage) SetEmbeddingProvider(v string)`

SetEmbeddingProvider sets EmbeddingProvider field to given value.

### HasEmbeddingProvider

`func (o *CloudObjectMessage) HasEmbeddingProvider() bool`

HasEmbeddingProvider returns a boolean if a field has been set.

### GetErrorText

`func (o *CloudObjectMessage) GetErrorText() string`

GetErrorText returns the ErrorText field if non-nil, zero value otherwise.

### GetErrorTextOk

`func (o *CloudObjectMessage) GetErrorTextOk() (*string, bool)`

GetErrorTextOk returns a tuple with the ErrorText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorText

`func (o *CloudObjectMessage) SetErrorText(v string)`

SetErrorText sets ErrorText field to given value.

### HasErrorText

`func (o *CloudObjectMessage) HasErrorText() bool`

HasErrorText returns a boolean if a field has been set.

### GetFileName

`func (o *CloudObjectMessage) GetFileName() string`

GetFileName returns the FileName field if non-nil, zero value otherwise.

### GetFileNameOk

`func (o *CloudObjectMessage) GetFileNameOk() (*string, bool)`

GetFileNameOk returns a tuple with the FileName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileName

`func (o *CloudObjectMessage) SetFileName(v string)`

SetFileName sets FileName field to given value.

### HasFileName

`func (o *CloudObjectMessage) HasFileName() bool`

HasFileName returns a boolean if a field has been set.

### GetIsAlerted

`func (o *CloudObjectMessage) GetIsAlerted() bool`

GetIsAlerted returns the IsAlerted field if non-nil, zero value otherwise.

### GetIsAlertedOk

`func (o *CloudObjectMessage) GetIsAlertedOk() (*bool, bool)`

GetIsAlertedOk returns a tuple with the IsAlerted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsAlerted

`func (o *CloudObjectMessage) SetIsAlerted(v bool)`

SetIsAlerted sets IsAlerted field to given value.

### HasIsAlerted

`func (o *CloudObjectMessage) HasIsAlerted() bool`

HasIsAlerted returns a boolean if a field has been set.

### GetIsDeleted

`func (o *CloudObjectMessage) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *CloudObjectMessage) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *CloudObjectMessage) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *CloudObjectMessage) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsHidden

`func (o *CloudObjectMessage) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *CloudObjectMessage) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *CloudObjectMessage) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *CloudObjectMessage) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetIsRegenerated

`func (o *CloudObjectMessage) GetIsRegenerated() bool`

GetIsRegenerated returns the IsRegenerated field if non-nil, zero value otherwise.

### GetIsRegeneratedOk

`func (o *CloudObjectMessage) GetIsRegeneratedOk() (*bool, bool)`

GetIsRegeneratedOk returns a tuple with the IsRegenerated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRegenerated

`func (o *CloudObjectMessage) SetIsRegenerated(v bool)`

SetIsRegenerated sets IsRegenerated field to given value.

### HasIsRegenerated

`func (o *CloudObjectMessage) HasIsRegenerated() bool`

HasIsRegenerated returns a boolean if a field has been set.

### GetLikeUsers

`func (o *CloudObjectMessage) GetLikeUsers() []string`

GetLikeUsers returns the LikeUsers field if non-nil, zero value otherwise.

### GetLikeUsersOk

`func (o *CloudObjectMessage) GetLikeUsersOk() (*[]string, bool)`

GetLikeUsersOk returns a tuple with the LikeUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeUsers

`func (o *CloudObjectMessage) SetLikeUsers(v []string)`

SetLikeUsers sets LikeUsers field to given value.

### HasLikeUsers

`func (o *CloudObjectMessage) HasLikeUsers() bool`

HasLikeUsers returns a boolean if a field has been set.

### GetModelProvider

`func (o *CloudObjectMessage) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *CloudObjectMessage) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *CloudObjectMessage) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *CloudObjectMessage) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *CloudObjectMessage) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CloudObjectMessage) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CloudObjectMessage) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CloudObjectMessage) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedNotify

`func (o *CloudObjectMessage) GetNeedNotify() bool`

GetNeedNotify returns the NeedNotify field if non-nil, zero value otherwise.

### GetNeedNotifyOk

`func (o *CloudObjectMessage) GetNeedNotifyOk() (*bool, bool)`

GetNeedNotifyOk returns a tuple with the NeedNotify field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedNotify

`func (o *CloudObjectMessage) SetNeedNotify(v bool)`

SetNeedNotify sets NeedNotify field to given value.

### HasNeedNotify

`func (o *CloudObjectMessage) HasNeedNotify() bool`

HasNeedNotify returns a boolean if a field has been set.

### GetOrganization

`func (o *CloudObjectMessage) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *CloudObjectMessage) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *CloudObjectMessage) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *CloudObjectMessage) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *CloudObjectMessage) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *CloudObjectMessage) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *CloudObjectMessage) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *CloudObjectMessage) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *CloudObjectMessage) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CloudObjectMessage) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CloudObjectMessage) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CloudObjectMessage) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetReasonText

`func (o *CloudObjectMessage) GetReasonText() string`

GetReasonText returns the ReasonText field if non-nil, zero value otherwise.

### GetReasonTextOk

`func (o *CloudObjectMessage) GetReasonTextOk() (*string, bool)`

GetReasonTextOk returns a tuple with the ReasonText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReasonText

`func (o *CloudObjectMessage) SetReasonText(v string)`

SetReasonText sets ReasonText field to given value.

### HasReasonText

`func (o *CloudObjectMessage) HasReasonText() bool`

HasReasonText returns a boolean if a field has been set.

### GetReplyTo

`func (o *CloudObjectMessage) GetReplyTo() string`

GetReplyTo returns the ReplyTo field if non-nil, zero value otherwise.

### GetReplyToOk

`func (o *CloudObjectMessage) GetReplyToOk() (*string, bool)`

GetReplyToOk returns a tuple with the ReplyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReplyTo

`func (o *CloudObjectMessage) SetReplyTo(v string)`

SetReplyTo sets ReplyTo field to given value.

### HasReplyTo

`func (o *CloudObjectMessage) HasReplyTo() bool`

HasReplyTo returns a boolean if a field has been set.

### GetStore

`func (o *CloudObjectMessage) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *CloudObjectMessage) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *CloudObjectMessage) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *CloudObjectMessage) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetSuggestions

`func (o *CloudObjectMessage) GetSuggestions() []CloudObjectSuggestion`

GetSuggestions returns the Suggestions field if non-nil, zero value otherwise.

### GetSuggestionsOk

`func (o *CloudObjectMessage) GetSuggestionsOk() (*[]CloudObjectSuggestion, bool)`

GetSuggestionsOk returns a tuple with the Suggestions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestions

`func (o *CloudObjectMessage) SetSuggestions(v []CloudObjectSuggestion)`

SetSuggestions sets Suggestions field to given value.

### HasSuggestions

`func (o *CloudObjectMessage) HasSuggestions() bool`

HasSuggestions returns a boolean if a field has been set.

### GetText

`func (o *CloudObjectMessage) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *CloudObjectMessage) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *CloudObjectMessage) SetText(v string)`

SetText sets Text field to given value.

### HasText

`func (o *CloudObjectMessage) HasText() bool`

HasText returns a boolean if a field has been set.

### GetTextTokenCount

`func (o *CloudObjectMessage) GetTextTokenCount() int64`

GetTextTokenCount returns the TextTokenCount field if non-nil, zero value otherwise.

### GetTextTokenCountOk

`func (o *CloudObjectMessage) GetTextTokenCountOk() (*int64, bool)`

GetTextTokenCountOk returns a tuple with the TextTokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextTokenCount

`func (o *CloudObjectMessage) SetTextTokenCount(v int64)`

SetTextTokenCount sets TextTokenCount field to given value.

### HasTextTokenCount

`func (o *CloudObjectMessage) HasTextTokenCount() bool`

HasTextTokenCount returns a boolean if a field has been set.

### GetTokenCount

`func (o *CloudObjectMessage) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *CloudObjectMessage) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *CloudObjectMessage) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *CloudObjectMessage) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetUser

`func (o *CloudObjectMessage) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CloudObjectMessage) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CloudObjectMessage) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *CloudObjectMessage) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVectorScores

`func (o *CloudObjectMessage) GetVectorScores() []CloudObjectVectorScore`

GetVectorScores returns the VectorScores field if non-nil, zero value otherwise.

### GetVectorScoresOk

`func (o *CloudObjectMessage) GetVectorScoresOk() (*[]CloudObjectVectorScore, bool)`

GetVectorScoresOk returns a tuple with the VectorScores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVectorScores

`func (o *CloudObjectMessage) SetVectorScores(v []CloudObjectVectorScore)`

SetVectorScores sets VectorScores field to given value.

### HasVectorScores

`func (o *CloudObjectMessage) HasVectorScores() bool`

HasVectorScores returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


