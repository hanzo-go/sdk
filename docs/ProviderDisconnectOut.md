# ProviderDisconnectOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to **[]string** | Accounts are the account labels this disconnect removed. Omitted when none was connected. | [optional] 
**Disconnected** | Pointer to **bool** | Disconnected is always true — the org&#39;s secrets and connection row are gone. | [optional] 
**Uninstalled** | Pointer to **[]string** | Uninstalled are the accounts whose provider installation was removed too, because no other org held them. An account another org still holds keeps its installation and is absent here. | [optional] 

## Methods

### NewProviderDisconnectOut

`func NewProviderDisconnectOut() *ProviderDisconnectOut`

NewProviderDisconnectOut instantiates a new ProviderDisconnectOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProviderDisconnectOutWithDefaults

`func NewProviderDisconnectOutWithDefaults() *ProviderDisconnectOut`

NewProviderDisconnectOutWithDefaults instantiates a new ProviderDisconnectOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ProviderDisconnectOut) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ProviderDisconnectOut) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ProviderDisconnectOut) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ProviderDisconnectOut) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDisconnected

`func (o *ProviderDisconnectOut) GetDisconnected() bool`

GetDisconnected returns the Disconnected field if non-nil, zero value otherwise.

### GetDisconnectedOk

`func (o *ProviderDisconnectOut) GetDisconnectedOk() (*bool, bool)`

GetDisconnectedOk returns a tuple with the Disconnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisconnected

`func (o *ProviderDisconnectOut) SetDisconnected(v bool)`

SetDisconnected sets Disconnected field to given value.

### HasDisconnected

`func (o *ProviderDisconnectOut) HasDisconnected() bool`

HasDisconnected returns a boolean if a field has been set.

### GetUninstalled

`func (o *ProviderDisconnectOut) GetUninstalled() []string`

GetUninstalled returns the Uninstalled field if non-nil, zero value otherwise.

### GetUninstalledOk

`func (o *ProviderDisconnectOut) GetUninstalledOk() (*[]string, bool)`

GetUninstalledOk returns a tuple with the Uninstalled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUninstalled

`func (o *ProviderDisconnectOut) SetUninstalled(v []string)`

SetUninstalled sets Uninstalled field to given value.

### HasUninstalled

`func (o *ProviderDisconnectOut) HasUninstalled() bool`

HasUninstalled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


