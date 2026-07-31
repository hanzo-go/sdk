# VectorMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Value** | Pointer to [**VectorMatchValue**](VectorMatchValue.md) |  | [optional] 
**Any** | Pointer to **[]interface{}** |  | [optional] 
**Except** | Pointer to **[]interface{}** |  | [optional] 

## Methods

### NewVectorMatch

`func NewVectorMatch() *VectorMatch`

NewVectorMatch instantiates a new VectorMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVectorMatchWithDefaults

`func NewVectorMatchWithDefaults() *VectorMatch`

NewVectorMatchWithDefaults instantiates a new VectorMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValue

`func (o *VectorMatch) GetValue() VectorMatchValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *VectorMatch) GetValueOk() (*VectorMatchValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *VectorMatch) SetValue(v VectorMatchValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *VectorMatch) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetAny

`func (o *VectorMatch) GetAny() []interface{}`

GetAny returns the Any field if non-nil, zero value otherwise.

### GetAnyOk

`func (o *VectorMatch) GetAnyOk() (*[]interface{}, bool)`

GetAnyOk returns a tuple with the Any field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAny

`func (o *VectorMatch) SetAny(v []interface{})`

SetAny sets Any field to given value.

### HasAny

`func (o *VectorMatch) HasAny() bool`

HasAny returns a boolean if a field has been set.

### GetExcept

`func (o *VectorMatch) GetExcept() []interface{}`

GetExcept returns the Except field if non-nil, zero value otherwise.

### GetExceptOk

`func (o *VectorMatch) GetExceptOk() (*[]interface{}, bool)`

GetExceptOk returns a tuple with the Except field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcept

`func (o *VectorMatch) SetExcept(v []interface{})`

SetExcept sets Except field to given value.

### HasExcept

`func (o *VectorMatch) HasExcept() bool`

HasExcept returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


