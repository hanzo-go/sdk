# Accounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]RoutedUsage**](RoutedUsage.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Total** | Pointer to [**AccountsTotal**](AccountsTotal.md) |  | [optional] 

## Methods

### NewAccounts

`func NewAccounts() *Accounts`

NewAccounts instantiates a new Accounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountsWithDefaults

`func NewAccountsWithDefaults() *Accounts`

NewAccountsWithDefaults instantiates a new Accounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *Accounts) GetAccounts() []RoutedUsage`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *Accounts) GetAccountsOk() (*[]RoutedUsage, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *Accounts) SetAccounts(v []RoutedUsage)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *Accounts) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *Accounts) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *Accounts) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *Accounts) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *Accounts) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *Accounts) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Accounts) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Accounts) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *Accounts) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotal

`func (o *Accounts) GetTotal() AccountsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *Accounts) GetTotalOk() (*AccountsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *Accounts) SetTotal(v AccountsTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *Accounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


