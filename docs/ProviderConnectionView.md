# ProviderConnectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the human label of the connected third-party account (the Slack team name, the GitHub org login). Provider-supplied and sanitized on ingest. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when the connection was last (re)established, RFC 3339 UTC. | [optional] 
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the access token expires, RFC 3339 UTC; empty for a credential that does not expire. Reading the token rotates it inside the window, so a reader never sees an expired one. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s own id for the account (Slack team.id, GitHub installation_id) — the value inbound webhooks are mapped back to this org by. | [optional] 
**Id** | Pointer to **string** | ID is provider + \&quot;:\&quot; + label, the address every connection route takes. Empty on the org plane, where a provider&#39;s connection is addressed by the provider alone. | [optional] 
**Label** | Pointer to **string** | Label is the caller&#39;s own name for this connection (\&quot;default\&quot;, \&quot;work\&quot;). | [optional] 
**Provider** | Pointer to **string** | Provider is the connected provider&#39;s registry id. Empty where the reader already knows it, which is every org-plane card. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions the provider granted. Never null; [] when none. | [optional] 

## Methods

### NewProviderConnectionView

`func NewProviderConnectionView() *ProviderConnectionView`

NewProviderConnectionView instantiates a new ProviderConnectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderConnectionViewWithDefaults

`func NewProviderConnectionViewWithDefaults() *ProviderConnectionView`

NewProviderConnectionViewWithDefaults instantiates a new ProviderConnectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ProviderConnectionView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ProviderConnectionView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ProviderConnectionView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ProviderConnectionView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConnectedAt

`func (o *ProviderConnectionView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ProviderConnectionView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ProviderConnectionView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *ProviderConnectionView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *ProviderConnectionView) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ProviderConnectionView) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ProviderConnectionView) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *ProviderConnectionView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExternalId

`func (o *ProviderConnectionView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ProviderConnectionView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ProviderConnectionView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ProviderConnectionView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *ProviderConnectionView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ProviderConnectionView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ProviderConnectionView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ProviderConnectionView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ProviderConnectionView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ProviderConnectionView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ProviderConnectionView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ProviderConnectionView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *ProviderConnectionView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ProviderConnectionView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ProviderConnectionView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ProviderConnectionView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScopes

`func (o *ProviderConnectionView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ProviderConnectionView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ProviderConnectionView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ProviderConnectionView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


