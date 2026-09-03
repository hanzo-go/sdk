# Pool

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** | At is the pool contract&#39;s address, lowercase. | [optional] 
**Count** | Pointer to **int64** |  | [optional] 
**Fee** | Pointer to **int64** | Fee is the pool&#39;s tier in hundredths of a basis point — 3000 is 0.3%. It is the integer the contract stores, unconverted, so nothing here rounds a rate. | [optional] 
**Locked** | Pointer to **string** |  | [optional] 
**Token0** | Pointer to [**Token**](Token.md) |  | [optional] 
**Token0Price** | Pointer to **string** | Token0Price is token1 per token0, and Token1Price its reciprocal, both as the indexer computed them. Neither is a price ON anything: it is the ratio the pool&#39;s reserves stand at. | [optional] 
**Token1** | Pointer to [**Token**](Token.md) |  | [optional] 
**Token1Price** | Pointer to **string** |  | [optional] 
**Volume** | Pointer to **string** |  | [optional] 

## Methods

### NewPool

`func NewPool() *Pool`

NewPool instantiates a new Pool object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPoolWithDefaults

`func NewPoolWithDefaults() *Pool`

NewPoolWithDefaults instantiates a new Pool object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *Pool) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *Pool) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *Pool) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *Pool) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCount

`func (o *Pool) GetCount() int64`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *Pool) GetCountOk() (*int64, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *Pool) SetCount(v int64)`

SetCount sets Count field to given value.

### HasCount

`func (o *Pool) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetFee

`func (o *Pool) GetFee() int64`

GetFee returns the Fee field if non-nil, zero value otherwise.

### GetFeeOk

`func (o *Pool) GetFeeOk() (*int64, bool)`

GetFeeOk returns a tuple with the Fee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFee

`func (o *Pool) SetFee(v int64)`

SetFee sets Fee field to given value.

### HasFee

`func (o *Pool) HasFee() bool`

HasFee returns a boolean if a field has been set.

### GetLocked

`func (o *Pool) GetLocked() string`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *Pool) GetLockedOk() (*string, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *Pool) SetLocked(v string)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *Pool) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetToken0

`func (o *Pool) GetToken0() Token`

GetToken0 returns the Token0 field if non-nil, zero value otherwise.

### GetToken0Ok

`func (o *Pool) GetToken0Ok() (*Token, bool)`

GetToken0Ok returns a tuple with the Token0 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0

`func (o *Pool) SetToken0(v Token)`

SetToken0 sets Token0 field to given value.

### HasToken0

`func (o *Pool) HasToken0() bool`

HasToken0 returns a boolean if a field has been set.

### GetToken0Price

`func (o *Pool) GetToken0Price() string`

GetToken0Price returns the Token0Price field if non-nil, zero value otherwise.

### GetToken0PriceOk

`func (o *Pool) GetToken0PriceOk() (*string, bool)`

GetToken0PriceOk returns a tuple with the Token0Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken0Price

`func (o *Pool) SetToken0Price(v string)`

SetToken0Price sets Token0Price field to given value.

### HasToken0Price

`func (o *Pool) HasToken0Price() bool`

HasToken0Price returns a boolean if a field has been set.

### GetToken1

`func (o *Pool) GetToken1() Token`

GetToken1 returns the Token1 field if non-nil, zero value otherwise.

### GetToken1Ok

`func (o *Pool) GetToken1Ok() (*Token, bool)`

GetToken1Ok returns a tuple with the Token1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1

`func (o *Pool) SetToken1(v Token)`

SetToken1 sets Token1 field to given value.

### HasToken1

`func (o *Pool) HasToken1() bool`

HasToken1 returns a boolean if a field has been set.

### GetToken1Price

`func (o *Pool) GetToken1Price() string`

GetToken1Price returns the Token1Price field if non-nil, zero value otherwise.

### GetToken1PriceOk

`func (o *Pool) GetToken1PriceOk() (*string, bool)`

GetToken1PriceOk returns a tuple with the Token1Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken1Price

`func (o *Pool) SetToken1Price(v string)`

SetToken1Price sets Token1Price field to given value.

### HasToken1Price

`func (o *Pool) HasToken1Price() bool`

HasToken1Price returns a boolean if a field has been set.

### GetVolume

`func (o *Pool) GetVolume() string`

GetVolume returns the Volume field if non-nil, zero value otherwise.

### GetVolumeOk

`func (o *Pool) GetVolumeOk() (*string, bool)`

GetVolumeOk returns a tuple with the Volume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolume

`func (o *Pool) SetVolume(v string)`

SetVolume sets Volume field to given value.

### HasVolume

`func (o *Pool) HasVolume() bool`

HasVolume returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


