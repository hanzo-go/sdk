# MarketPrecompile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**At** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **bool** | Code is whether the address carries any. False is an ANSWER — the node replied and there is nothing deployed there — and is not the same as the read having failed, which the enclosing Reach reports instead. | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Role** | Pointer to **string** |  | [optional] 

## Methods

### NewMarketPrecompile

`func NewMarketPrecompile() *MarketPrecompile`

NewMarketPrecompile instantiates a new MarketPrecompile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMarketPrecompileWithDefaults

`func NewMarketPrecompileWithDefaults() *MarketPrecompile`

NewMarketPrecompileWithDefaults instantiates a new MarketPrecompile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAt

`func (o *MarketPrecompile) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *MarketPrecompile) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *MarketPrecompile) SetAt(v string)`

SetAt sets At field to given value.

### HasAt

`func (o *MarketPrecompile) HasAt() bool`

HasAt returns a boolean if a field has been set.

### GetCode

`func (o *MarketPrecompile) GetCode() bool`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *MarketPrecompile) GetCodeOk() (*bool, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *MarketPrecompile) SetCode(v bool)`

SetCode sets Code field to given value.

### HasCode

`func (o *MarketPrecompile) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *MarketPrecompile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MarketPrecompile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MarketPrecompile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *MarketPrecompile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRole

`func (o *MarketPrecompile) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *MarketPrecompile) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *MarketPrecompile) SetRole(v string)`

SetRole sets Role field to given value.

### HasRole

`func (o *MarketPrecompile) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


