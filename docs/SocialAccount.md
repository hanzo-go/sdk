# SocialAccount

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreatedAt** | Pointer to **int32** | CreatedAt is when the account was connected, as a unix timestamp in seconds. | [optional] 
**Handle** | Pointer to **string** | Handle is the account&#39;s public name on the network, as the customer knows it. Trimmed and bounded at 1024 characters.  Example: \&quot;@acme\&quot; | [optional] 
**Id** | Pointer to **string** | ID is the account&#39;s identifier, minted on connect and the id every later call addresses it by.  Example: \&quot;acct_7f3c1a\&quot; | [optional] 
**Provider** | Pointer to **string** | Provider is the network this account is on: x, facebook, instagram, linkedin, tiktok, youtube or threads.  Example: \&quot;x\&quot; | [optional] 
**Status** | Pointer to **string** | Status is the connection lifecycle: connected, disconnected or error. Only a connected account is a publish target. | [optional] 
**UpdatedAt** | Pointer to **int32** | UpdatedAt is when the account row last changed, as a unix timestamp in seconds. The listing is ordered by it, newest first. | [optional] 

## Methods

### NewSocialAccount

`func NewSocialAccount() *SocialAccount`

NewSocialAccount instantiates a new SocialAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAccountWithDefaults

`func NewSocialAccountWithDefaults() *SocialAccount`

NewSocialAccountWithDefaults instantiates a new SocialAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreatedAt

`func (o *SocialAccount) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SocialAccount) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SocialAccount) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SocialAccount) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetHandle

`func (o *SocialAccount) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SocialAccount) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SocialAccount) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SocialAccount) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetId

`func (o *SocialAccount) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SocialAccount) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SocialAccount) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SocialAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetProvider

`func (o *SocialAccount) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialAccount) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialAccount) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialAccount) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *SocialAccount) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialAccount) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialAccount) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialAccount) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SocialAccount) GetUpdatedAt() int32`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SocialAccount) GetUpdatedAtOk() (*int32, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SocialAccount) SetUpdatedAt(v int32)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SocialAccount) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


