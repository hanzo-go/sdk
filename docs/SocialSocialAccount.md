# SocialSocialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int64** | CreatedAt is when the account was connected, as a unix timestamp in seconds. | [optional] 
**Handle** | Pointer to **string** | Handle is the account&#39;s public name on the network, as the customer knows it. Trimmed and bounded at 1024 characters.  Example: \&quot;@acme\&quot; | [optional] 
**Id** | Pointer to **string** | ID is the account&#39;s identifier, minted on connect and the id every later call addresses it by.  Example: \&quot;acct_7f3c1a\&quot; | [optional] 
**Provider** | Pointer to **string** | Provider is the network this account is on: x, facebook, instagram, linkedin, tiktok, youtube or threads.  Example: \&quot;x\&quot; | [optional] 
**Status** | Pointer to **string** | Status is the connection lifecycle: connected, disconnected or error. Only a connected account is a publish target. | [optional] 
**UpdatedAt** | Pointer to **int64** | UpdatedAt is when the account row last changed, as a unix timestamp in seconds. The listing is ordered by it, newest first. | [optional] 

## Methods

### NewSocialSocialAccount

`func NewSocialSocialAccount() *SocialSocialAccount`

NewSocialSocialAccount instantiates a new SocialSocialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialSocialAccountWithDefaults

`func NewSocialSocialAccountWithDefaults() *SocialSocialAccount`

NewSocialSocialAccountWithDefaults instantiates a new SocialSocialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SocialSocialAccount) GetCreatedAt() int64`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SocialSocialAccount) GetCreatedAtOk() (*int64, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SocialSocialAccount) SetCreatedAt(v int64)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SocialSocialAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHandle

`func (o *SocialSocialAccount) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SocialSocialAccount) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SocialSocialAccount) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SocialSocialAccount) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *SocialSocialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialSocialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialSocialAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialSocialAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *SocialSocialAccount) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialSocialAccount) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialSocialAccount) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialSocialAccount) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *SocialSocialAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialSocialAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialSocialAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialSocialAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SocialSocialAccount) GetUpdatedAt() int64`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SocialSocialAccount) GetUpdatedAtOk() (*int64, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SocialSocialAccount) SetUpdatedAt(v int64)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SocialSocialAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


