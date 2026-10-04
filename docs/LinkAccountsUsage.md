# LinkAccountsUsage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]LinkRoutedUsage**](LinkRoutedUsage.md) | Accounts is one row per linked account the gateway actually routed through. | [optional] 
**Scope** | Pointer to **string** | Scope is always \&quot;user\&quot;: the caller&#39;s own linked accounts. | [optional] 
**Source** | Pointer to **string** | Source is always \&quot;routed\&quot;: the gateway&#39;s own routed ledger. | [optional] 
**Total** | Pointer to [**LinkAccountsTotal**](LinkAccountsTotal.md) | Total is the honest sum across the rows. | [optional] 

## Methods

### NewLinkAccountsUsage

`func NewLinkAccountsUsage() *LinkAccountsUsage`

NewLinkAccountsUsage instantiates a new LinkAccountsUsage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkAccountsUsageWithDefaults

`func NewLinkAccountsUsageWithDefaults() *LinkAccountsUsage`

NewLinkAccountsUsageWithDefaults instantiates a new LinkAccountsUsage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *LinkAccountsUsage) GetAccounts() []LinkRoutedUsage`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *LinkAccountsUsage) GetAccountsOk() (*[]LinkRoutedUsage, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *LinkAccountsUsage) SetAccounts(v []LinkRoutedUsage)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *LinkAccountsUsage) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *LinkAccountsUsage) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *LinkAccountsUsage) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *LinkAccountsUsage) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *LinkAccountsUsage) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *LinkAccountsUsage) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *LinkAccountsUsage) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *LinkAccountsUsage) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *LinkAccountsUsage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotal

`func (o *LinkAccountsUsage) GetTotal() LinkAccountsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *LinkAccountsUsage) GetTotalOk() (*LinkAccountsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *LinkAccountsUsage) SetTotal(v LinkAccountsTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *LinkAccountsUsage) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


