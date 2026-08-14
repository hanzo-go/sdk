# ConnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** | Account is the account label the provider reported for the credential. apikey path only; a pointer because \&quot;\&quot; is a real answer the provider gave. | [optional] 
**AuthorizeUrl** | Pointer to **string** | AuthorizeURL is the provider consent URL to send the user to. OAuth path only. | [optional] 
**Connected** | Pointer to **bool** | Connected is true on the apikey path once the credential verified and sealed. | [optional] 
**ExternalId** | Pointer to **string** | ExternalID is the provider&#39;s account id for the credential. apikey path only. | [optional] 
**Provider** | Pointer to **string** | Provider is the connector&#39;s registry id. apikey path only. | [optional] 
**Scopes** | Pointer to **[]string** | Scopes are the permissions the credential carries. apikey path only; never null on that path ([] when the provider reported none). | [optional] 

## Methods

### NewConnectOut

`func NewConnectOut() *ConnectOut`

NewConnectOut instantiates a new ConnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectOutWithDefaults

`func NewConnectOutWithDefaults() *ConnectOut`

NewConnectOutWithDefaults instantiates a new ConnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *ConnectOut) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ConnectOut) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ConnectOut) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ConnectOut) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAuthorizeUrl

`func (o *ConnectOut) GetAuthorizeUrl() string`

GetAuthorizeUrl returns the AuthorizeUrl field if non-nil, zero value otherwise.

### GetAuthorizeUrlOk

`func (o *ConnectOut) GetAuthorizeUrlOk() (*string, bool)`

GetAuthorizeUrlOk returns a tuple with the AuthorizeUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizeUrl

`func (o *ConnectOut) SetAuthorizeUrl(v string)`

SetAuthorizeUrl sets AuthorizeUrl field to given value.

### HasAuthorizeUrl

`func (o *ConnectOut) HasAuthorizeUrl() bool`

HasAuthorizeUrl returns a boolean if a field has been set.

### GetConnected

`func (o *ConnectOut) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *ConnectOut) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *ConnectOut) SetConnected(v bool)`

SetConnected sets Connected field to given value.

### HasConnected

`func (o *ConnectOut) HasConnected() bool`

HasConnected returns a boolean if a field has been set.

### GetExternalId

`func (o *ConnectOut) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ConnectOut) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ConnectOut) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ConnectOut) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetProvider

`func (o *ConnectOut) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ConnectOut) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ConnectOut) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ConnectOut) HasProvider() bool`

HasProvider returns a boolean if a field has been set.

### GetScopes

`func (o *ConnectOut) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ConnectOut) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ConnectOut) SetScopes(v []string)`

SetScopes sets Scopes field to given value.

### HasScopes

`func (o *ConnectOut) HasScopes() bool`

HasScopes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


