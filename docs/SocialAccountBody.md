# SocialAccountBody

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Handle** | Pointer to **string** | Handle is the account&#39;s public name on the network, as the customer knows it (&#x60;@acme&#x60;). Trimmed and bounded at 1024 characters.  Example: \&quot;@acme\&quot; | [optional] 
**Provider** | Pointer to **string** | Provider is the network this account is on: x, facebook, instagram, linkedin, tiktok, youtube or threads. Omitted means x. Anything else is refused rather than coerced, because a stored account on a network that does not exist can never be published through.  Example: \&quot;x\&quot; | [optional] 
**Status** | Pointer to **string** | Status is the connection lifecycle: connected, disconnected or error. Omitted means connected. Only a connected account is a publish target. | [optional] 

## Methods

### NewSocialAccountBody

`func NewSocialAccountBody() *SocialAccountBody`

NewSocialAccountBody instantiates a new SocialAccountBody object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSocialAccountBodyWithDefaults

`func NewSocialAccountBodyWithDefaults() *SocialAccountBody`

NewSocialAccountBodyWithDefaults instantiates a new SocialAccountBody object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHandle

`func (o *SocialAccountBody) GetHandle() string`

GetHandle returns the Handle field if non-nil, zero value otherwise.

### GetHandleOk

`func (o *SocialAccountBody) GetHandleOk() (*string, bool)`

GetHandleOk returns a tuple with the Handle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandle

`func (o *SocialAccountBody) SetHandle(v string)`

SetHandle sets Handle field to given value.

### HasHandle

`func (o *SocialAccountBody) HasHandle() bool`

HasHandle returns a boolean if a field has been set.

### GetProvider

`func (o *SocialAccountBody) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *SocialAccountBody) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *SocialAccountBody) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *SocialAccountBody) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetStatus

`func (o *SocialAccountBody) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SocialAccountBody) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SocialAccountBody) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SocialAccountBody) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


