# Tokens

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | Pointer to **string** |  | [optional] 
**Reach** | Pointer to [**Reach**](Reach.md) |  | [optional] 
**Tokens** | Pointer to [**[]Token**](Token.md) | Tokens is &#x60;[]&#x60; where the indexer holds none and &#x60;null&#x60; where the read failed. | [optional] 

## Methods

### NewTokens

`func NewTokens() *Tokens`

NewTokens instantiates a new Tokens object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTokensWithDefaults

`func NewTokensWithDefaults() *Tokens`

NewTokensWithDefaults instantiates a new Tokens object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *Tokens) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *Tokens) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *Tokens) SetChain(v string)`

SetChain sets Chain field to given value.

### HasChain

`func (o *Tokens) HasChain() bool`

HasChain returns a boolean if a field has been set.

### GetReach

`func (o *Tokens) GetReach() Reach`

GetReach returns the Reach field if non-nil, zero value otherwise.

### GetReachOk

`func (o *Tokens) GetReachOk() (*Reach, bool)`

GetReachOk returns a tuple with the Reach field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReach

`func (o *Tokens) SetReach(v Reach)`

SetReach sets Reach field to given value.

### HasReach

`func (o *Tokens) HasReach() bool`

HasReach returns a boolean if a field has been set.

### GetTokens

`func (o *Tokens) GetTokens() []Token`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *Tokens) GetTokensOk() (*[]Token, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *Tokens) SetTokens(v []Token)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *Tokens) HasTokens() bool`

HasTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


