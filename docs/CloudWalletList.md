# CloudWalletList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Wallets** | Pointer to [**[]CloudWallet**](CloudWallet.md) | Wallets are the matching wallets, newest first. | [optional] 

## Methods

### NewCloudWalletList

`func NewCloudWalletList() *CloudWalletList`

NewCloudWalletList instantiates a new CloudWalletList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudWalletListWithDefaults

`func NewCloudWalletListWithDefaults() *CloudWalletList`

NewCloudWalletListWithDefaults instantiates a new CloudWalletList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWallets

`func (o *CloudWalletList) GetWallets() []CloudWallet`

GetWallets returns the Wallets field if non-nil, zero value otherwise.

### GetWalletsOk

`func (o *CloudWalletList) GetWalletsOk() (*[]CloudWallet, bool)`

GetWalletsOk returns a tuple with the Wallets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallets

`func (o *CloudWalletList) SetWallets(v []CloudWallet)`

SetWallets sets Wallets field to given value.

### HasWallets

`func (o *CloudWalletList) HasWallets() bool`

HasWallets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


