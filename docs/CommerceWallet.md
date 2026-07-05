# CommerceWallet

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Accounts** | Pointer to [**[]CommerceWalletAccount**](CommerceWalletAccount.md) |  | [optional] 

## Methods

### NewCommerceWallet

`func NewCommerceWallet() *CommerceWallet`

NewCommerceWallet instantiates a new CommerceWallet object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommerceWalletWithDefaults

`func NewCommerceWalletWithDefaults() *CommerceWallet`

NewCommerceWalletWithDefaults instantiates a new CommerceWallet object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CommerceWallet) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommerceWallet) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommerceWallet) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommerceWallet) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccounts

`func (o *CommerceWallet) GetAccounts() []CommerceWalletAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *CommerceWallet) GetAccountsOk() (*[]CommerceWalletAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *CommerceWallet) SetAccounts(v []CommerceWalletAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *CommerceWallet) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


