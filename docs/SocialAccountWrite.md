# SocialAccountWrite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the account&#39;s public name on the network. Omitting it BLANKS the stored handle: this is a replacement, not a merge.  Example: \&quot;@acme\&quot; | [optional] 
**Provider** | Pointer to **string** | Provider is the network this account is on: x, facebook, instagram, linkedin, tiktok, youtube or threads. Omitted means x.  Example: \&quot;x\&quot; | [optional] 
**Status** | Pointer to **string** | Status is the connection lifecycle: connected, disconnected or error. Omitting it RESETS the account to connected. | [optional] 

## Methods

### NewSocialAccountWrite

`func NewSocialAccountWrite() *SocialAccountWrite`

NewSocialAccountWrite instantiates a new SocialAccountWrite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAccountWriteWithDefaults

`func NewSocialAccountWriteWithDefaults() *SocialAccountWrite`

NewSocialAccountWriteWithDefaults instantiates a new SocialAccountWrite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *SocialAccountWrite) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SocialAccountWrite) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SocialAccountWrite) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SocialAccountWrite) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetProvider

`func (o *SocialAccountWrite) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialAccountWrite) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialAccountWrite) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialAccountWrite) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *SocialAccountWrite) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialAccountWrite) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialAccountWrite) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialAccountWrite) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


