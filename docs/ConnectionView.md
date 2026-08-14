# ConnectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the human label of the connected third-party account (the Slack team name, the GitHub org login). Provider-supplied and sanitized on ingest. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when the connection was last (re)established, RFC 3339 UTC. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s own id for the account (Slack team.id, GitHub installation_id) — the value inbound webhooks are mapped back to this org by. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions the provider granted. Never null; [] when none. | [optional] 

## Methods

### NewConnectionView

`func NewConnectionView() *ConnectionView`

NewConnectionView instantiates a new ConnectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectionViewWithDefaults

`func NewConnectionViewWithDefaults() *ConnectionView`

NewConnectionViewWithDefaults instantiates a new ConnectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ConnectionView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ConnectionView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ConnectionView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ConnectionView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConnectedAt

`func (o *ConnectionView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ConnectionView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ConnectionView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *ConnectionView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetExternalId

`func (o *ConnectionView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ConnectionView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ConnectionView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ConnectionView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetScopes

`func (o *ConnectionView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectionView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectionView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectionView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


