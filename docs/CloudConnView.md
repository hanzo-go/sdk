# CloudConnView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the provider&#39;s label for the connected account. | [optional] 
**ConnectedAt** | Pointer to **string** | ConnectedAt is when the connector was last (re)established, RFC 3339 UTC. | [optional] 
**ExpiresAt** | Pointer to **string** | ExpiresAt is when the access token expires, RFC 3339 UTC; empty for a non-expiring credential. Reading the token auto-rotates inside the window. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s own id for that account. | [optional] 
**Id** | Pointer to **string** | ID is provider + \&quot;:\&quot; + label — what every other connector route addresses. | [optional] 
**Label** | Pointer to **string** | Label is the caller&#39;s name for this connection (\&quot;default\&quot;, \&quot;work\&quot;). | [optional] 
**Provider** | Pointer to **string** | Provider is the user-scoped provider&#39;s registry id. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions the credential carries. Never null; [] when none. | [optional] 

## Methods

### NewCloudConnView

`func NewCloudConnView() *CloudConnView`

NewCloudConnView instantiates a new CloudConnView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudConnViewWithDefaults

`func NewCloudConnViewWithDefaults() *CloudConnView`

NewCloudConnViewWithDefaults instantiates a new CloudConnView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *CloudConnView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CloudConnView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CloudConnView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CloudConnView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetConnectedAt

`func (o *CloudConnView) GetConnectedAt() string`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *CloudConnView) GetConnectedAtOk() (*string, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *CloudConnView) SetConnectedAt(v string)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *CloudConnView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.

### GetExpiresAt

`func (o *CloudConnView) GetExpiresAt() string`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *CloudConnView) GetExpiresAtOk() (*string, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *CloudConnView) SetExpiresAt(v string)`

SetExpiresAt sets ExpiresAt field to given value.

### HasExpiresAt

`func (o *CloudConnView) HasExpiresAt() bool`

HasExpiresAt returns a boolean if a field has been set.

### GetExternalId

`func (o *CloudConnView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *CloudConnView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *CloudConnView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *CloudConnView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetId

`func (o *CloudConnView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudConnView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudConnView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudConnView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *CloudConnView) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CloudConnView) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CloudConnView) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CloudConnView) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetProvider

`func (o *CloudConnView) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *CloudConnView) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *CloudConnView) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *CloudConnView) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScopes

`func (o *CloudConnView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *CloudConnView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *CloudConnView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *CloudConnView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


