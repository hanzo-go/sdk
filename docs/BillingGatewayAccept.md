# BillingGatewayAccept

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Asset** | Pointer to **string** |  | [optional] 
**Decimals** | Pointer to **int64** |  | [optional] 
**Domain** | Pointer to [**BillingGatewayDomain**](BillingGatewayDomain.md) |  | [optional] 
**MaxAmount** | Pointer to **string** |  | [optional] 
**MinAmount** | Pointer to **string** |  | [optional] 
**MinValiditySeconds** | Pointer to **int64** |  | [optional] 
**Network** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewBillingGatewayAccept

`func NewBillingGatewayAccept() *BillingGatewayAccept`

NewBillingGatewayAccept instantiates a new BillingGatewayAccept object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBillingGatewayAcceptWithDefaults

`func NewBillingGatewayAcceptWithDefaults() *BillingGatewayAccept`

NewBillingGatewayAcceptWithDefaults instantiates a new BillingGatewayAccept object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAsset

`func (o *BillingGatewayAccept) GetAsset() string`

GetAsset returns the Asset field if non-nil, zero value otherwise.

### GetAssetOk

`func (o *BillingGatewayAccept) GetAssetOk() (*string, bool)`

GetAssetOk returns a tuple with the Asset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsset

`func (o *BillingGatewayAccept) SetAsset(v string)`

SetAsset sets Asset field to given value.

### HasAsset

`func (o *BillingGatewayAccept) HasAsset() bool`

HasAsset returns a boolean if a field has been set.

### GetDecimals

`func (o *BillingGatewayAccept) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *BillingGatewayAccept) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *BillingGatewayAccept) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *BillingGatewayAccept) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetDomain

`func (o *BillingGatewayAccept) GetDomain() BillingGatewayDomain`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *BillingGatewayAccept) GetDomainOk() (*BillingGatewayDomain, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *BillingGatewayAccept) SetDomain(v BillingGatewayDomain)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *BillingGatewayAccept) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetMaxAmount

`func (o *BillingGatewayAccept) GetMaxAmount() string`

GetMaxAmount returns the MaxAmount field if non-nil, zero value otherwise.

### GetMaxAmountOk

`func (o *BillingGatewayAccept) GetMaxAmountOk() (*string, bool)`

GetMaxAmountOk returns a tuple with the MaxAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxAmount

`func (o *BillingGatewayAccept) SetMaxAmount(v string)`

SetMaxAmount sets MaxAmount field to given value.

### HasMaxAmount

`func (o *BillingGatewayAccept) HasMaxAmount() bool`

HasMaxAmount returns a boolean if a field has been set.

### GetMinAmount

`func (o *BillingGatewayAccept) GetMinAmount() string`

GetMinAmount returns the MinAmount field if non-nil, zero value otherwise.

### GetMinAmountOk

`func (o *BillingGatewayAccept) GetMinAmountOk() (*string, bool)`

GetMinAmountOk returns a tuple with the MinAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinAmount

`func (o *BillingGatewayAccept) SetMinAmount(v string)`

SetMinAmount sets MinAmount field to given value.

### HasMinAmount

`func (o *BillingGatewayAccept) HasMinAmount() bool`

HasMinAmount returns a boolean if a field has been set.

### GetMinValiditySeconds

`func (o *BillingGatewayAccept) GetMinValiditySeconds() int64`

GetMinValiditySeconds returns the MinValiditySeconds field if non-nil, zero value otherwise.

### GetMinValiditySecondsOk

`func (o *BillingGatewayAccept) GetMinValiditySecondsOk() (*int64, bool)`

GetMinValiditySecondsOk returns a tuple with the MinValiditySeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinValiditySeconds

`func (o *BillingGatewayAccept) SetMinValiditySeconds(v int64)`

SetMinValiditySeconds sets MinValiditySeconds field to given value.

### HasMinValiditySeconds

`func (o *BillingGatewayAccept) HasMinValiditySeconds() bool`

HasMinValiditySeconds returns a boolean if a field has been set.

### GetNetwork

`func (o *BillingGatewayAccept) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *BillingGatewayAccept) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *BillingGatewayAccept) SetNetwork(v string)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *BillingGatewayAccept) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSymbol

`func (o *BillingGatewayAccept) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *BillingGatewayAccept) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *BillingGatewayAccept) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *BillingGatewayAccept) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


