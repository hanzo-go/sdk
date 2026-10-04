# BillingAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]BillingRoutedUsage**](BillingRoutedUsage.md) |  | [optional] 
**Scope** | Pointer to **string** |  | [optional] 
**Source** | Pointer to **string** |  | [optional] 
**Total** | Pointer to [**BillingAccountsTotal**](BillingAccountsTotal.md) |  | [optional] 

## Methods

### NewBillingAccounts

`func NewBillingAccounts() *BillingAccounts`

NewBillingAccounts instantiates a new BillingAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingAccountsWithDefaults

`func NewBillingAccountsWithDefaults() *BillingAccounts`

NewBillingAccountsWithDefaults instantiates a new BillingAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *BillingAccounts) GetAccounts() []BillingRoutedUsage`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *BillingAccounts) GetAccountsOk() (*[]BillingRoutedUsage, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *BillingAccounts) SetAccounts(v []BillingRoutedUsage)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *BillingAccounts) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetScope

`func (o *BillingAccounts) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *BillingAccounts) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *BillingAccounts) SetScope(v string)`

SetScope sets Scope field to given value.

### HasScope

`func (o *BillingAccounts) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSource

`func (o *BillingAccounts) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *BillingAccounts) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *BillingAccounts) SetSource(v string)`

SetSource sets Source field to given value.

### HasSource

`func (o *BillingAccounts) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetTotal

`func (o *BillingAccounts) GetTotal() BillingAccountsTotal`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *BillingAccounts) GetTotalOk() (*BillingAccountsTotal, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *BillingAccounts) SetTotal(v BillingAccountsTotal)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *BillingAccounts) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


