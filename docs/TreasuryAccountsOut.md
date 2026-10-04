# TreasuryAccountsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]TreasuryAccountView**](TreasuryAccountView.md) | Accounts are the ledger accounts in scope with their balances. | [optional] 
**Scope** | Pointer to **string** | Scope is the scope actually served: \&quot;org\&quot; or \&quot;house\&quot;. | [optional] 
**Tenant** | Pointer to **string** | Tenant is the org whose accounts these are (empty for the house scope&#39;s own rows). | [optional] 

## Methods

### NewTreasuryAccountsOut

`func NewTreasuryAccountsOut() *TreasuryAccountsOut`

NewTreasuryAccountsOut instantiates a new TreasuryAccountsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTreasuryAccountsOutWithDefaults

`func NewTreasuryAccountsOutWithDefaults() *TreasuryAccountsOut`

NewTreasuryAccountsOutWithDefaults instantiates a new TreasuryAccountsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *TreasuryAccountsOut) GetAccounts() []TreasuryAccountView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *TreasuryAccountsOut) GetAccountsOk() (*[]TreasuryAccountView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *TreasuryAccountsOut) SetAccounts(v []TreasuryAccountView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *TreasuryAccountsOut) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *TreasuryAccountsOut) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *TreasuryAccountsOut) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *TreasuryAccountsOut) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *TreasuryAccountsOut) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTenant

`func (o *TreasuryAccountsOut) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TreasuryAccountsOut) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TreasuryAccountsOut) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TreasuryAccountsOut) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


