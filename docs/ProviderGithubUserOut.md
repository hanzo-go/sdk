# ProviderGithubUserOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Configured** | Pointer to **bool** | Configured is whether this deployment holds the App&#39;s OAuth client. False means nobody can connect here yet, and connect answers 503. | [optional] 
**Connected** | Pointer to **bool** | Connected is whether the caller has connected their own GitHub account in this org. False means their repositories come from nowhere: an org admin still sees the org&#39;s installations, and everyone else is asked to connect. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when the connection was made, RFC 3339 UTC. | [optional] 
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the current access token expires, RFC 3339 UTC. It is rotated before it does; absent for a token GitHub does not expire. | [optional] 
**Login** | Pointer to **string** | Login is the GitHub account they connected. Absent when not connected. | [optional] 

## Methods

### NewProviderGithubUserOut

`func NewProviderGithubUserOut() *ProviderGithubUserOut`

NewProviderGithubUserOut instantiates a new ProviderGithubUserOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderGithubUserOutWithDefaults

`func NewProviderGithubUserOutWithDefaults() *ProviderGithubUserOut`

NewProviderGithubUserOutWithDefaults instantiates a new ProviderGithubUserOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigured

`func (o *ProviderGithubUserOut) GetConfigured() bool`

GetConfigured returns the Configured field if non-nil, zero value otherwise.

### GetConfiguredOk

`func (o *ProviderGithubUserOut) GetConfiguredOk() (*bool, bool)`

GetConfiguredOk returns a tuple with the Configured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigured

`func (o *ProviderGithubUserOut) SetConfigured(v bool)`

SetConfigured sets Configured field to given value.

### HasConfigured

`func (o *ProviderGithubUserOut) HasConfigured() bool`

HasConfigured returns a boolean if a field has been set.

### GetConnected

`func (o *ProviderGithubUserOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ProviderGithubUserOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ProviderGithubUserOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ProviderGithubUserOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetConnectedAt

`func (o *ProviderGithubUserOut) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ProviderGithubUserOut) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ProviderGithubUserOut) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *ProviderGithubUserOut) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ProviderGithubUserOut) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProviderGithubUserOut) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProviderGithubUserOut) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProviderGithubUserOut) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetLogin

`func (o *ProviderGithubUserOut) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ProviderGithubUserOut) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ProviderGithubUserOut) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ProviderGithubUserOut) HasLogin() bool`

HasLogin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


