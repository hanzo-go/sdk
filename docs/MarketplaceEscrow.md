# MarketplaceEscrow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contract** | Pointer to **string** | Contract is the contract the money moves through: the token the x402 authorization transfers, or the escrow on chain. | [optional] 
**Network** | Pointer to **string** | Network is the CAIP-2 network the payment is signed for. | [optional] 
**PayTo** | Pointer to **string** | PayTo is the seller&#39;s payout address. | [optional] 
**Rail** | Pointer to **string** | Rail is x402 — the buyer&#39;s balance, settled at release — or chain, the Lux escrow contract. | [optional] 
**TxHash** | Pointer to **string** | TxHash is the transaction that paid the seller, once one did: the chain&#39;s hash, or the x402 settlement id on the ledger. | [optional] 

## Methods

### NewMarketplaceEscrow

`func NewMarketplaceEscrow() *MarketplaceEscrow`

NewMarketplaceEscrow instantiates a new MarketplaceEscrow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketplaceEscrowWithDefaults

`func NewMarketplaceEscrowWithDefaults() *MarketplaceEscrow`

NewMarketplaceEscrowWithDefaults instantiates a new MarketplaceEscrow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContract

`func (o *MarketplaceEscrow) GetContract() string`

GetContract returns the Contract field if non-nil, zero value otherwise.

### GetContractOk

`func (o *MarketplaceEscrow) GetContractOk() (*string, bool)`

GetContractOk returns a tuple with the Contract field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContract

`func (o *MarketplaceEscrow) SetContract(v string)`

SetContract sets Contract field to given value.

### HasContract

`func (o *MarketplaceEscrow) HasContract() bool`

HasContract returns a boolean if a field has been set.

### GetNetwork

`func (o *MarketplaceEscrow) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *MarketplaceEscrow) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *MarketplaceEscrow) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *MarketplaceEscrow) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetPayTo

`func (o *MarketplaceEscrow) GetPayTo() string`

GetPayTo returns the PayTo field if non-nil, zero value otherwise.

### GetPayToOk

`func (o *MarketplaceEscrow) GetPayToOk() (*string, bool)`

GetPayToOk returns a tuple with the PayTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayTo

`func (o *MarketplaceEscrow) SetPayTo(v string)`

SetPayTo sets PayTo field to given value.

### HasPayTo

`func (o *MarketplaceEscrow) HasPayTo() bool`

HasPayTo returns a boolean if a field has been set.

### GetRail

`func (o *MarketplaceEscrow) GetRail() string`

GetRail returns the Rail field if non-nil, zero value otherwise.

### GetRailOk

`func (o *MarketplaceEscrow) GetRailOk() (*string, bool)`

GetRailOk returns a tuple with the Rail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRail

`func (o *MarketplaceEscrow) SetRail(v string)`

SetRail sets Rail field to given value.

### HasRail

`func (o *MarketplaceEscrow) HasRail() bool`

HasRail returns a boolean if a field has been set.

### GetTxHash

`func (o *MarketplaceEscrow) GetTxHash() string`

GetTxHash returns the TxHash field if non-nil, zero value otherwise.

### GetTxHashOk

`func (o *MarketplaceEscrow) GetTxHashOk() (*string, bool)`

GetTxHashOk returns a tuple with the TxHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxHash

`func (o *MarketplaceEscrow) SetTxHash(v string)`

SetTxHash sets TxHash field to given value.

### HasTxHash

`func (o *MarketplaceEscrow) HasTxHash() bool`

HasTxHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


