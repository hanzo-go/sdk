# ProviderGoogleUserOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Configured is whether this deployment holds the OAuth client. False means nobody can connect here yet, and connect answers 503. | [optional] 
**Connected** | Pointer to **bool** | Connected is whether the caller has connected their own Google account in this org. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when the connection was made, RFC 3339 UTC. | [optional] 
**Email** | Pointer to **string** | Email is the Google account they connected: the From of every message sent as them. Absent when not connected. | [optional] 
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the current access token expires, RFC 3339 UTC. It is rotated before it does. | [optional] 
**Send** | Pointer to **bool** | Send is whether the grant lets Hanzo send mail as them. A connection is only kept when it does, so this is Connected unless the person later narrowed the grant at Google. | [optional] 

## Methods

### NewProviderGoogleUserOut

`func NewProviderGoogleUserOut() *ProviderGoogleUserOut`

NewProviderGoogleUserOut instantiates a new ProviderGoogleUserOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGoogleUserOutWithDefaults

`func NewProviderGoogleUserOutWithDefaults() *ProviderGoogleUserOut`

NewProviderGoogleUserOutWithDefaults instantiates a new ProviderGoogleUserOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *ProviderGoogleUserOut) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *ProviderGoogleUserOut) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *ProviderGoogleUserOut) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *ProviderGoogleUserOut) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetConnected

`func (o *ProviderGoogleUserOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderGoogleUserOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderGoogleUserOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderGoogleUserOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnectedAt

`func (o *ProviderGoogleUserOut) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ProviderGoogleUserOut) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ProviderGoogleUserOut) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *ProviderGoogleUserOut) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetEmail

`func (o *ProviderGoogleUserOut) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *ProviderGoogleUserOut) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *ProviderGoogleUserOut) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *ProviderGoogleUserOut) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ProviderGoogleUserOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProviderGoogleUserOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProviderGoogleUserOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProviderGoogleUserOut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetSend

`func (o *ProviderGoogleUserOut) GetSend() bool`

GetSend returns the Send field if non-nil, zero value otherwise.

### GetSendOk

`func (o *ProviderGoogleUserOut) GetSendOk() (*bool, bool)`

GetSendOk returns a tuple with the Send field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSend

`func (o *ProviderGoogleUserOut) SetSend(v bool)`

SetSend sets Send field to given value.

### HasSend

`func (o *ProviderGoogleUserOut) HasSend() bool`

HasSend returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


