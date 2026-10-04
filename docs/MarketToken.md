# MarketToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the token&#39;s contract address, lowercase. | [optional] 
**Decimals** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Symbol** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketToken

`func NewMarketToken() *MarketToken`

NewMarketToken instantiates a new MarketToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketTokenWithDefaults

`func NewMarketTokenWithDefaults() *MarketToken`

NewMarketTokenWithDefaults instantiates a new MarketToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketToken) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketToken) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketToken) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketToken) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetDecimals

`func (o *MarketToken) GetDecimals() int64`

GetDecimals returns the Decimals field if non-nil, zero value otherwise.

### GetDecimalsOk

`func (o *MarketToken) GetDecimalsOk() (*int64, bool)`

GetDecimalsOk returns a tuple with the Decimals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecimals

`func (o *MarketToken) SetDecimals(v int64)`

SetDecimals sets Decimals field to given value.

### HasDecimals

`func (o *MarketToken) HasDecimals() bool`

HasDecimals returns a boolean if a field has been set.

### GetName

`func (o *MarketToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketToken) HasName() bool`

HasName returns a boolean if a field has been set.

### GetSymbol

`func (o *MarketToken) GetSymbol() string`

GetSymbol returns the Symbol field if non-nil, zero value otherwise.

### GetSymbolOk

`func (o *MarketToken) GetSymbolOk() (*string, bool)`

GetSymbolOk returns a tuple with the Symbol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSymbol

`func (o *MarketToken) SetSymbol(v string)`

SetSymbol sets Symbol field to given value.

### HasSymbol

`func (o *MarketToken) HasSymbol() bool`

HasSymbol returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


