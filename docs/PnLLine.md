# PnLLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to **string** |  | [optional] 
**Amount** | Pointer to **int32** | cents, display sign (income &amp; expense both positive when normal) | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 

## Methods

### NewPnLLine

`func NewPnLLine() *PnLLine`

NewPnLLine instantiates a new PnLLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPnLLineWithDefaults

`func NewPnLLineWithDefaults() *PnLLine`

NewPnLLineWithDefaults instantiates a new PnLLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *PnLLine) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *PnLLine) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *PnLLine) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *PnLLine) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetAmount

`func (o *PnLLine) GetAmount() int32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PnLLine) GetAmountOk() (*int32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PnLLine) SetAmount(v int32)`

SetAmount sets Amount field to given value.

### HasAmount

`func (o *PnLLine) HasAmount() bool`

HasAmount returns a boolean if a field has been set.

### GetName

`func (o *PnLLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PnLLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PnLLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PnLLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *PnLLine) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PnLLine) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PnLLine) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PnLLine) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


