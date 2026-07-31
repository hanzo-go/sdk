# CloudRailView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** | Chain is the human chain name, e.g. \&quot;Base\&quot;. | [optional] 
**ChainId** | Pointer to **int32** | ChainID is the EIP-155 chain id the wallet must be on. | [optional] 
**Decimals** | Pointer to **int32** | Decimals is the token&#39;s decimal places — 6 for USDC, 18 for an 18-decimal token. Cents are derived per-rail from it. | [optional] 
**Id** | Pointer to **string** | ID is the stable rail id to name when submitting a transfer, e.g. \&quot;base-usdc\&quot;. | [optional] 
**Symbol** | Pointer to **string** | Symbol is the display symbol, e.g. \&quot;USDC\&quot;. | [optional] 
**Token** | Pointer to **string** | Token is the ERC-20 contract address to transfer. | [optional] 
**Treasury** | Pointer to **string** | Treasury is the address on this chain to send funds to. | [optional] 

## Methods

### NewCloudRailView

`func NewCloudRailView() *CloudRailView`

NewCloudRailView instantiates a new CloudRailView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudRailViewWithDefaults

`func NewCloudRailViewWithDefaults() *CloudRailView`

NewCloudRailViewWithDefaults instantiates a new CloudRailView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *CloudRailView) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *CloudRailView) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *CloudRailView) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *CloudRailView) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetChainId

`func (o *CloudRailView) GetChainId() int32`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *CloudRailView) GetChainIdOk() (*int32, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *CloudRailView) SetChainId(v int32)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *CloudRailView) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### GetDecimals

`func (o *CloudRailView) GetDecimals() int32`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *CloudRailView) GetDecimalsOk() (*int32, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *CloudRailView) SetDecimals(v int32)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *CloudRailView) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetId

`func (o *CloudRailView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CloudRailView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CloudRailView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CloudRailView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSymbol

`func (o *CloudRailView) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *CloudRailView) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *CloudRailView) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *CloudRailView) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.

### GetToken

`func (o *CloudRailView) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *CloudRailView) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *CloudRailView) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *CloudRailView) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTreasury

`func (o *CloudRailView) GetTreasury() string`

GetTreasury returns the Treasury field if non-nil, zero value otherwise.

### GetTreasuryOk

`func (o *CloudRailView) GetTreasuryOk() (*string, bool)`

GetTreasuryOk returns a tuple with the Treasury field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTreasury

`func (o *CloudRailView) SetTreasury(v string)`

SetTreasury sets Treasury field to given value.

### HasTreasury

`func (o *CloudRailView) HasTreasury() bool`

HasTreasury returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


