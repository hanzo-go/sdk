# IntegrationsConnectionView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Human account label (e.g. Slack team name) | [optional] 
**ExternalId** | Pointer to **string** | Provider account id (e.g. Slack team id) | [optional] 
**Scopes** | Pointer to **[]string** |  | [optional] 
**ConnectedAt** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewIntegrationsConnectionView

`func NewIntegrationsConnectionView() *IntegrationsConnectionView`

NewIntegrationsConnectionView instantiates a new IntegrationsConnectionView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationsConnectionViewWithDefaults

`func NewIntegrationsConnectionViewWithDefaults() *IntegrationsConnectionView`

NewIntegrationsConnectionViewWithDefaults instantiates a new IntegrationsConnectionView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *IntegrationsConnectionView) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *IntegrationsConnectionView) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *IntegrationsConnectionView) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *IntegrationsConnectionView) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetExternalId

`func (o *IntegrationsConnectionView) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *IntegrationsConnectionView) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *IntegrationsConnectionView) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *IntegrationsConnectionView) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetScopes

`func (o *IntegrationsConnectionView) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *IntegrationsConnectionView) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *IntegrationsConnectionView) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *IntegrationsConnectionView) HasScopes() bool`

HasScopes returns a boolean if a field has been set.

### GetConnectedAt

`func (o *IntegrationsConnectionView) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *IntegrationsConnectionView) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *IntegrationsConnectionView) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.

### HasConnectedAt

`func (o *IntegrationsConnectionView) HasConnectedAt() bool`

HasConnectedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


