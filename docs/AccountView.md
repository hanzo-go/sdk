# AccountView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is the ledger account address (\&quot;org:acme:wallet\&quot;, \&quot;fund:reserve\&quot;, …). | [optional] 
**BalanceCents** | Pointer to **int64** | BalanceCents is that account&#39;s signed balance in minor units. | [optional] 

## Methods

### NewAccountView

`func NewAccountView() *AccountView`

NewAccountView instantiates a new AccountView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccountViewWithDefaults

`func NewAccountViewWithDefaults() *AccountView`

NewAccountViewWithDefaults instantiates a new AccountView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *AccountView) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *AccountView) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *AccountView) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *AccountView) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBalanceCents

`func (o *AccountView) GetBalanceCents() int64`

GetBalanceCents returns the BalanceCents field if non-nil, zero value otherwise.

### GetBalanceCentsOk

`func (o *AccountView) GetBalanceCentsOk() (*int64, bool)`

GetBalanceCentsOk returns a tuple with the BalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceCents

`func (o *AccountView) SetBalanceCents(v int64)`

SetBalanceCents sets BalanceCents field to given value.

### HasBalanceCents

`func (o *AccountView) HasBalanceCents() bool`

HasBalanceCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


