# CloudAccountsOut

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]CloudAccountView**](CloudAccountView.md) | Accounts are the ledger accounts in scope with their balances. | [optional] 
**Scope** | Pointer to **string** | Scope is the scope actually served: \&quot;org\&quot; or \&quot;house\&quot;. | [optional] 
**Tenant** | Pointer to **string** | Tenant is the org whose accounts these are (empty for the house scope&#39;s own rows). | [optional] 

## Methods

### NewCloudAccountsOut

`func NewCloudAccountsOut() *CloudAccountsOut`

NewCloudAccountsOut instantiates a new CloudAccountsOut object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudAccountsOutWithDefaults

`func NewCloudAccountsOutWithDefaults() *CloudAccountsOut`

NewCloudAccountsOutWithDefaults instantiates a new CloudAccountsOut object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *CloudAccountsOut) GetAccounts() []CloudAccountView`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CloudAccountsOut) GetAccountsOk() (*[]CloudAccountView, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CloudAccountsOut) SetAccounts(v []CloudAccountView)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CloudAccountsOut) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *CloudAccountsOut) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *CloudAccountsOut) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *CloudAccountsOut) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *CloudAccountsOut) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetTenant

`func (o *CloudAccountsOut) GetTenant() string`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *CloudAccountsOut) GetTenantOk() (*string, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *CloudAccountsOut) SetTenant(v string)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *CloudAccountsOut) HasTenant() bool`

HasTenant returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


