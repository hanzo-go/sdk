# AccountsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]AccountView**](AccountView.md) | Accounts are the ledger accounts in scope with their balances. | [optional] 
**Scope** | Pointer to **string** | Scope is the scope actually served: \&quot;org\&quot; or \&quot;house\&quot;. | [optional] 
**Tenant** | Pointer to **string** | Tenant is the org whose accounts these are (empty for the house scope&#39;s own rows). | [optional] 

## Methods

### NewAccountsOut

`func NewAccountsOut() *AccountsOut`

NewAccountsOut instantiates a new AccountsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsOutWithDefaults

`func NewAccountsOutWithDefaults() *AccountsOut`

NewAccountsOutWithDefaults instantiates a new AccountsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *AccountsOut) GetAccounts() []AccountView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AccountsOut) GetAccountsOk() (*[]AccountView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AccountsOut) SetAccounts(v []AccountView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AccountsOut) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *AccountsOut) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *AccountsOut) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *AccountsOut) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *AccountsOut) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTenant

`func (o *AccountsOut) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *AccountsOut) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *AccountsOut) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *AccountsOut) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


