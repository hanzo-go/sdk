# Chat

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** |  | [optional] 
**ClientIp** | Pointer to **string** |  | [optional] 
**ClientIpDesc** | Pointer to **string** |  | [optional] 
**CreatedTime** | Pointer to **string** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**DisplayName** | Pointer to **string** |  | [optional] 
**IsDeleted** | Pointer to **bool** |  | [optional] 
**IsHidden** | Pointer to **bool** |  | [optional] 
**MessageCount** | Pointer to **int32** |  | [optional] 
**ModelProvider** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedTitle** | Pointer to **bool** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int32** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**User1** | Pointer to **string** |  | [optional] 
**User2** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserAgentDesc** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewChat

`func NewChat() *Chat`

NewChat instantiates a new Chat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewChatWithDefaults

`func NewChatWithDefaults() *Chat`

NewChatWithDefaults instantiates a new Chat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *Chat) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *Chat) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *Chat) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *Chat) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientIp

`func (o *Chat) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *Chat) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *Chat) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *Chat) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientIpDesc

`func (o *Chat) GetClientIpDesc() string`

GetClientIpDesc returns the ClientIpDesc field if non-nil, zero value otherwise.

### GetClientIpDescOk

`func (o *Chat) GetClientIpDescOk() (*string, bool)`

GetClientIpDescOk returns a tuple with the ClientIpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpDesc

`func (o *Chat) SetClientIpDesc(v string)`

SetClientIpDesc sets ClientIpDesc field to given value.

### HasClientIpDesc

`func (o *Chat) HasClientIpDesc() bool`

HasClientIpDesc returns a boolean if a field has been set.

### GetCreatedTime

`func (o *Chat) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *Chat) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *Chat) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *Chat) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *Chat) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Chat) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Chat) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Chat) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *Chat) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *Chat) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *Chat) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *Chat) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *Chat) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *Chat) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *Chat) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *Chat) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsHidden

`func (o *Chat) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *Chat) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *Chat) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *Chat) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetMessageCount

`func (o *Chat) GetMessageCount() int32`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *Chat) GetMessageCountOk() (*int32, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *Chat) SetMessageCount(v int32)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *Chat) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetModelProvider

`func (o *Chat) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *Chat) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *Chat) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *Chat) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *Chat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Chat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Chat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Chat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedTitle

`func (o *Chat) GetNeedTitle() bool`

GetNeedTitle returns the NeedTitle field if non-nil, zero value otherwise.

### GetNeedTitleOk

`func (o *Chat) GetNeedTitleOk() (*bool, bool)`

GetNeedTitleOk returns a tuple with the NeedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTitle

`func (o *Chat) SetNeedTitle(v bool)`

SetNeedTitle sets NeedTitle field to given value.

### HasNeedTitle

`func (o *Chat) HasNeedTitle() bool`

HasNeedTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *Chat) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *Chat) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *Chat) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *Chat) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *Chat) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *Chat) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *Chat) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *Chat) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *Chat) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Chat) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Chat) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Chat) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStore

`func (o *Chat) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *Chat) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *Chat) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *Chat) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTokenCount

`func (o *Chat) GetTokenCount() int32`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *Chat) GetTokenCountOk() (*int32, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *Chat) SetTokenCount(v int32)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *Chat) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetType

`func (o *Chat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Chat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Chat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Chat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *Chat) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *Chat) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *Chat) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *Chat) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUser

`func (o *Chat) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *Chat) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *Chat) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *Chat) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUser1

`func (o *Chat) GetUser1() string`

GetUser1 returns the User1 field if non-nil, zero value otherwise.

### GetUser1Ok

`func (o *Chat) GetUser1Ok() (*string, bool)`

GetUser1Ok returns a tuple with the User1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser1

`func (o *Chat) SetUser1(v string)`

SetUser1 sets User1 field to given value.

### HasUser1

`func (o *Chat) HasUser1() bool`

HasUser1 returns a boolean if a field has been set.

### GetUser2

`func (o *Chat) GetUser2() string`

GetUser2 returns the User2 field if non-nil, zero value otherwise.

### GetUser2Ok

`func (o *Chat) GetUser2Ok() (*string, bool)`

GetUser2Ok returns a tuple with the User2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser2

`func (o *Chat) SetUser2(v string)`

SetUser2 sets User2 field to given value.

### HasUser2

`func (o *Chat) HasUser2() bool`

HasUser2 returns a boolean if a field has been set.

### GetUserAgent

`func (o *Chat) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *Chat) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *Chat) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *Chat) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentDesc

`func (o *Chat) GetUserAgentDesc() string`

GetUserAgentDesc returns the UserAgentDesc field if non-nil, zero value otherwise.

### GetUserAgentDescOk

`func (o *Chat) GetUserAgentDescOk() (*string, bool)`

GetUserAgentDescOk returns a tuple with the UserAgentDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentDesc

`func (o *Chat) SetUserAgentDesc(v string)`

SetUserAgentDesc sets UserAgentDesc field to given value.

### HasUserAgentDesc

`func (o *Chat) HasUserAgentDesc() bool`

HasUserAgentDesc returns a boolean if a field has been set.

### GetUsers

`func (o *Chat) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Chat) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Chat) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *Chat) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


