# MarketPool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the pool contract&#39;s address, lowercase. | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Fee** | Pointer to **int64** | Fee is the pool&#39;s tier in hundredths of a basis point — 3000 is 0.3%. It is the integer the contract stores, unconverted, so nothing here rounds a rate. | [optional] 
**Locked** | Pointer to **string** |  | [optional] 
**Token0** | Pointer to [**MarketToken**](MarketToken.md) |  | [optional] 
**Token0Price** | Pointer to **string** | Token0Price is token1 per token0, and Token1Price its reciprocal, both as the indexer computed them. Neither is a price ON anything: it is the ratio the pool&#39;s reserves stand at. | [optional] 
**Token1** | Pointer to [**MarketToken**](MarketToken.md) |  | [optional] 
**Token1Price** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketPool

`func NewMarketPool() *MarketPool`

NewMarketPool instantiates a new MarketPool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPoolWithDefaults

`func NewMarketPoolWithDefaults() *MarketPool`

NewMarketPoolWithDefaults instantiates a new MarketPool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketPool) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketPool) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketPool) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketPool) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCount

`func (o *MarketPool) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *MarketPool) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *MarketPool) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *MarketPool) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFee

`func (o *MarketPool) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *MarketPool) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *MarketPool) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *MarketPool) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetLocked

`func (o *MarketPool) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *MarketPool) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *MarketPool) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *MarketPool) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetToken0

`func (o *MarketPool) GetToken0() MarketToken`

GetToken0 returns the Token0 field if non-nil, zero value otherwise.

### GetToken0Ok

`func (o *MarketPool) GetToken0Ok() (*MarketToken, bool)`

GetToken0Ok returns a tuple with the Token0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0

`func (o *MarketPool) SetToken0(v MarketToken)`

SetToken0 sets Token0 field to given value.

### HasToken0

`func (o *MarketPool) HasToken0() bool`

HasToken0 returns a boolean if a field has been set.

### GetToken0Price

`func (o *MarketPool) GetToken0Price() string`

GetToken0Price returns the Token0Price field if non-nil, zero value otherwise.

### GetToken0PriceOk

`func (o *MarketPool) GetToken0PriceOk() (*string, bool)`

GetToken0PriceOk returns a tuple with the Token0Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0Price

`func (o *MarketPool) SetToken0Price(v string)`

SetToken0Price sets Token0Price field to given value.

### HasToken0Price

`func (o *MarketPool) HasToken0Price() bool`

HasToken0Price returns a boolean if a field has been set.

### GetToken1

`func (o *MarketPool) GetToken1() MarketToken`

GetToken1 returns the Token1 field if non-nil, zero value otherwise.

### GetToken1Ok

`func (o *MarketPool) GetToken1Ok() (*MarketToken, bool)`

GetToken1Ok returns a tuple with the Token1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1

`func (o *MarketPool) SetToken1(v MarketToken)`

SetToken1 sets Token1 field to given value.

### HasToken1

`func (o *MarketPool) HasToken1() bool`

HasToken1 returns a boolean if a field has been set.

### GetToken1Price

`func (o *MarketPool) GetToken1Price() string`

GetToken1Price returns the Token1Price field if non-nil, zero value otherwise.

### GetToken1PriceOk

`func (o *MarketPool) GetToken1PriceOk() (*string, bool)`

GetToken1PriceOk returns a tuple with the Token1Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1Price

`func (o *MarketPool) SetToken1Price(v string)`

SetToken1Price sets Token1Price field to given value.

### HasToken1Price

`func (o *MarketPool) HasToken1Price() bool`

HasToken1Price returns a boolean if a field has been set.

### GetVolume

`func (o *MarketPool) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *MarketPool) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *MarketPool) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *MarketPool) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


