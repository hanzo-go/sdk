# MarketplacePayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** | Address is its address, recovered from the signature that proved it. | [optional] 
**Bound** | Pointer to **bool** | Bound is true once a signature proved it. | [optional] 
**BoundAt** | Pointer to **int64** | BoundAt is when, unix seconds. | [optional] 
**Wallet** | Pointer to **string** | Wallet is the wallet id, in the seller&#39;s org. | [optional] 

## Methods

### NewMarketplacePayout

`func NewMarketplacePayout() *MarketplacePayout`

NewMarketplacePayout instantiates a new MarketplacePayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplacePayoutWithDefaults

`func NewMarketplacePayoutWithDefaults() *MarketplacePayout`

NewMarketplacePayoutWithDefaults instantiates a new MarketplacePayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *MarketplacePayout) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *MarketplacePayout) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *MarketplacePayout) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *MarketplacePayout) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetBound

`func (o *MarketplacePayout) GetBound() bool`

GetBound returns the Bound field if non-nil, zero value otherwise.

### GetBoundOk

`func (o *MarketplacePayout) GetBoundOk() (*bool, bool)`

GetBoundOk returns a tuple with the Bound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBound

`func (o *MarketplacePayout) SetBound(v bool)`

SetBound sets Bound field to given value.

### HasBound

`func (o *MarketplacePayout) HasBound() bool`

HasBound returns a boolean if a field has been set.

### GetBoundAt

`func (o *MarketplacePayout) GetBoundAt() int64`

GetBoundAt returns the BoundAt field if non-nil, zero value otherwise.

### GetBoundAtOk

`func (o *MarketplacePayout) GetBoundAtOk() (*int64, bool)`

GetBoundAtOk returns a tuple with the BoundAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoundAt

`func (o *MarketplacePayout) SetBoundAt(v int64)`

SetBoundAt sets BoundAt field to given value.

### HasBoundAt

`func (o *MarketplacePayout) HasBoundAt() bool`

HasBoundAt returns a boolean if a field has been set.

### GetWallet

`func (o *MarketplacePayout) GetWallet() string`

GetWallet returns the Wallet field if non-nil, zero value otherwise.

### GetWalletOk

`func (o *MarketplacePayout) GetWalletOk() (*string, bool)`

GetWalletOk returns a tuple with the Wallet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWallet

`func (o *MarketplacePayout) SetWallet(v string)`

SetWallet sets Wallet field to given value.

### HasWallet

`func (o *MarketplacePayout) HasWallet() bool`

HasWallet returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


