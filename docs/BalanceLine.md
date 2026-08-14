# BalanceLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** | cents, display sign | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewBalanceLine

`func NewBalanceLine() *BalanceLine`

NewBalanceLine instantiates a new BalanceLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBalanceLineWithDefaults

`func NewBalanceLineWithDefaults() *BalanceLine`

NewBalanceLineWithDefaults instantiates a new BalanceLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *BalanceLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *BalanceLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *BalanceLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *BalanceLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *BalanceLine) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *BalanceLine) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *BalanceLine) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *BalanceLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetName

`func (o *BalanceLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *BalanceLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *BalanceLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *BalanceLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *BalanceLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BalanceLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BalanceLine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BalanceLine) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


