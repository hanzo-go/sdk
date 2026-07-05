# NexusChat

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
**MessageCount** | Pointer to **int64** |  | [optional] 
**ModelProvider** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**NeedTitle** | Pointer to **bool** |  | [optional] 
**Organization** | Pointer to **string** |  | [optional] 
**Owner** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float64** |  | [optional] 
**Store** | Pointer to **string** |  | [optional] 
**TokenCount** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**UpdatedTime** | Pointer to **string** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**User1** | Pointer to **string** |  | [optional] 
**User2** | Pointer to **string** |  | [optional] 
**UserAgent** | Pointer to **string** |  | [optional] 
**UserAgentDesc** | Pointer to **string** |  | [optional] 
**Users** | Pointer to **[]string** |  | [optional] 

## Methods

### NewNexusChat

`func NewNexusChat() *NexusChat`

NewNexusChat instantiates a new NexusChat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNexusChatWithDefaults

`func NewNexusChatWithDefaults() *NexusChat`

NewNexusChatWithDefaults instantiates a new NexusChat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *NexusChat) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *NexusChat) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *NexusChat) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *NexusChat) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetClientIp

`func (o *NexusChat) GetClientIp() string`

GetClientIp returns the ClientIp field if non-nil, zero value otherwise.

### GetClientIpOk

`func (o *NexusChat) GetClientIpOk() (*string, bool)`

GetClientIpOk returns a tuple with the ClientIp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIp

`func (o *NexusChat) SetClientIp(v string)`

SetClientIp sets ClientIp field to given value.

### HasClientIp

`func (o *NexusChat) HasClientIp() bool`

HasClientIp returns a boolean if a field has been set.

### GetClientIpDesc

`func (o *NexusChat) GetClientIpDesc() string`

GetClientIpDesc returns the ClientIpDesc field if non-nil, zero value otherwise.

### GetClientIpDescOk

`func (o *NexusChat) GetClientIpDescOk() (*string, bool)`

GetClientIpDescOk returns a tuple with the ClientIpDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientIpDesc

`func (o *NexusChat) SetClientIpDesc(v string)`

SetClientIpDesc sets ClientIpDesc field to given value.

### HasClientIpDesc

`func (o *NexusChat) HasClientIpDesc() bool`

HasClientIpDesc returns a boolean if a field has been set.

### GetCreatedTime

`func (o *NexusChat) GetCreatedTime() string`

GetCreatedTime returns the CreatedTime field if non-nil, zero value otherwise.

### GetCreatedTimeOk

`func (o *NexusChat) GetCreatedTimeOk() (*string, bool)`

GetCreatedTimeOk returns a tuple with the CreatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedTime

`func (o *NexusChat) SetCreatedTime(v string)`

SetCreatedTime sets CreatedTime field to given value.

### HasCreatedTime

`func (o *NexusChat) HasCreatedTime() bool`

HasCreatedTime returns a boolean if a field has been set.

### GetCurrency

`func (o *NexusChat) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *NexusChat) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *NexusChat) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *NexusChat) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetDisplayName

`func (o *NexusChat) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *NexusChat) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *NexusChat) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *NexusChat) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetIsDeleted

`func (o *NexusChat) GetIsDeleted() bool`

GetIsDeleted returns the IsDeleted field if non-nil, zero value otherwise.

### GetIsDeletedOk

`func (o *NexusChat) GetIsDeletedOk() (*bool, bool)`

GetIsDeletedOk returns a tuple with the IsDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDeleted

`func (o *NexusChat) SetIsDeleted(v bool)`

SetIsDeleted sets IsDeleted field to given value.

### HasIsDeleted

`func (o *NexusChat) HasIsDeleted() bool`

HasIsDeleted returns a boolean if a field has been set.

### GetIsHidden

`func (o *NexusChat) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *NexusChat) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *NexusChat) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.

### HasIsHidden

`func (o *NexusChat) HasIsHidden() bool`

HasIsHidden returns a boolean if a field has been set.

### GetMessageCount

`func (o *NexusChat) GetMessageCount() int64`

GetMessageCount returns the MessageCount field if non-nil, zero value otherwise.

### GetMessageCountOk

`func (o *NexusChat) GetMessageCountOk() (*int64, bool)`

GetMessageCountOk returns a tuple with the MessageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageCount

`func (o *NexusChat) SetMessageCount(v int64)`

SetMessageCount sets MessageCount field to given value.

### HasMessageCount

`func (o *NexusChat) HasMessageCount() bool`

HasMessageCount returns a boolean if a field has been set.

### GetModelProvider

`func (o *NexusChat) GetModelProvider() string`

GetModelProvider returns the ModelProvider field if non-nil, zero value otherwise.

### GetModelProviderOk

`func (o *NexusChat) GetModelProviderOk() (*string, bool)`

GetModelProviderOk returns a tuple with the ModelProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelProvider

`func (o *NexusChat) SetModelProvider(v string)`

SetModelProvider sets ModelProvider field to given value.

### HasModelProvider

`func (o *NexusChat) HasModelProvider() bool`

HasModelProvider returns a boolean if a field has been set.

### GetName

`func (o *NexusChat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NexusChat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NexusChat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NexusChat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNeedTitle

`func (o *NexusChat) GetNeedTitle() bool`

GetNeedTitle returns the NeedTitle field if non-nil, zero value otherwise.

### GetNeedTitleOk

`func (o *NexusChat) GetNeedTitleOk() (*bool, bool)`

GetNeedTitleOk returns a tuple with the NeedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNeedTitle

`func (o *NexusChat) SetNeedTitle(v bool)`

SetNeedTitle sets NeedTitle field to given value.

### HasNeedTitle

`func (o *NexusChat) HasNeedTitle() bool`

HasNeedTitle returns a boolean if a field has been set.

### GetOrganization

`func (o *NexusChat) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *NexusChat) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *NexusChat) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *NexusChat) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetOwner

`func (o *NexusChat) GetOwner() string`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *NexusChat) GetOwnerOk() (*string, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *NexusChat) SetOwner(v string)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *NexusChat) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetPrice

`func (o *NexusChat) GetPrice() float64`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *NexusChat) GetPriceOk() (*float64, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *NexusChat) SetPrice(v float64)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *NexusChat) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetStore

`func (o *NexusChat) GetStore() string`

GetStore returns the Store field if non-nil, zero value otherwise.

### GetStoreOk

`func (o *NexusChat) GetStoreOk() (*string, bool)`

GetStoreOk returns a tuple with the Store field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStore

`func (o *NexusChat) SetStore(v string)`

SetStore sets Store field to given value.

### HasStore

`func (o *NexusChat) HasStore() bool`

HasStore returns a boolean if a field has been set.

### GetTokenCount

`func (o *NexusChat) GetTokenCount() int64`

GetTokenCount returns the TokenCount field if non-nil, zero value otherwise.

### GetTokenCountOk

`func (o *NexusChat) GetTokenCountOk() (*int64, bool)`

GetTokenCountOk returns a tuple with the TokenCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenCount

`func (o *NexusChat) SetTokenCount(v int64)`

SetTokenCount sets TokenCount field to given value.

### HasTokenCount

`func (o *NexusChat) HasTokenCount() bool`

HasTokenCount returns a boolean if a field has been set.

### GetType

`func (o *NexusChat) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NexusChat) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NexusChat) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NexusChat) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdatedTime

`func (o *NexusChat) GetUpdatedTime() string`

GetUpdatedTime returns the UpdatedTime field if non-nil, zero value otherwise.

### GetUpdatedTimeOk

`func (o *NexusChat) GetUpdatedTimeOk() (*string, bool)`

GetUpdatedTimeOk returns a tuple with the UpdatedTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedTime

`func (o *NexusChat) SetUpdatedTime(v string)`

SetUpdatedTime sets UpdatedTime field to given value.

### HasUpdatedTime

`func (o *NexusChat) HasUpdatedTime() bool`

HasUpdatedTime returns a boolean if a field has been set.

### GetUser

`func (o *NexusChat) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *NexusChat) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *NexusChat) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *NexusChat) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetUser1

`func (o *NexusChat) GetUser1() string`

GetUser1 returns the User1 field if non-nil, zero value otherwise.

### GetUser1Ok

`func (o *NexusChat) GetUser1Ok() (*string, bool)`

GetUser1Ok returns a tuple with the User1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser1

`func (o *NexusChat) SetUser1(v string)`

SetUser1 sets User1 field to given value.

### HasUser1

`func (o *NexusChat) HasUser1() bool`

HasUser1 returns a boolean if a field has been set.

### GetUser2

`func (o *NexusChat) GetUser2() string`

GetUser2 returns the User2 field if non-nil, zero value otherwise.

### GetUser2Ok

`func (o *NexusChat) GetUser2Ok() (*string, bool)`

GetUser2Ok returns a tuple with the User2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser2

`func (o *NexusChat) SetUser2(v string)`

SetUser2 sets User2 field to given value.

### HasUser2

`func (o *NexusChat) HasUser2() bool`

HasUser2 returns a boolean if a field has been set.

### GetUserAgent

`func (o *NexusChat) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *NexusChat) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *NexusChat) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *NexusChat) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### GetUserAgentDesc

`func (o *NexusChat) GetUserAgentDesc() string`

GetUserAgentDesc returns the UserAgentDesc field if non-nil, zero value otherwise.

### GetUserAgentDescOk

`func (o *NexusChat) GetUserAgentDescOk() (*string, bool)`

GetUserAgentDescOk returns a tuple with the UserAgentDesc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgentDesc

`func (o *NexusChat) SetUserAgentDesc(v string)`

SetUserAgentDesc sets UserAgentDesc field to given value.

### HasUserAgentDesc

`func (o *NexusChat) HasUserAgentDesc() bool`

HasUserAgentDesc returns a boolean if a field has been set.

### GetUsers

`func (o *NexusChat) GetUsers() []string`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *NexusChat) GetUsersOk() (*[]string, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *NexusChat) SetUsers(v []string)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *NexusChat) HasUsers() bool`

HasUsers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


