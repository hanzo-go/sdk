# O11yO11yServiceTag

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoolValues** | Pointer to **[]bool** | BoolValues are the boolean operands, when the attribute is a bool. | [optional] 
**Key** | Pointer to **string** | Key is the span attribute to test. | [optional] 
**NumberValues** | Pointer to **[]float32** | NumberValues are the numeric operands, when the attribute is a number. | [optional] 
**Operator** | Pointer to **string** | Operator is how to test it, e.g. in, not_in. | [optional] 
**StringValues** | Pointer to **[]string** | StringValues are the string operands, when the attribute is a string. | [optional] 
**TagType** | Pointer to **string** | TagType says which plane the attribute lives on, e.g. tag or resource. | [optional] 

## Methods

### NewO11yO11yServiceTag

`func NewO11yO11yServiceTag() *O11yO11yServiceTag`

NewO11yO11yServiceTag instantiates a new O11yO11yServiceTag object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yServiceTagWithDefaults

`func NewO11yO11yServiceTagWithDefaults() *O11yO11yServiceTag`

NewO11yO11yServiceTagWithDefaults instantiates a new O11yO11yServiceTag object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolValues

`func (o *O11yO11yServiceTag) GetBoolValues() []bool`

GetBoolValues returns the BoolValues field if non-nil, zero value otherwise.

### GetBoolValuesOk

`func (o *O11yO11yServiceTag) GetBoolValuesOk() (*[]bool, bool)`

GetBoolValuesOk returns a tuple with the BoolValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValues

`func (o *O11yO11yServiceTag) SetBoolValues(v []bool)`

SetBoolValues sets BoolValues field to given value.

### HasBoolValues

`func (o *O11yO11yServiceTag) HasBoolValues() bool`

HasBoolValues returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yServiceTag) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yServiceTag) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yServiceTag) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yServiceTag) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNumberValues

`func (o *O11yO11yServiceTag) GetNumberValues() []float32`

GetNumberValues returns the NumberValues field if non-nil, zero value otherwise.

### GetNumberValuesOk

`func (o *O11yO11yServiceTag) GetNumberValuesOk() (*[]float32, bool)`

GetNumberValuesOk returns a tuple with the NumberValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValues

`func (o *O11yO11yServiceTag) SetNumberValues(v []float32)`

SetNumberValues sets NumberValues field to given value.

### HasNumberValues

`func (o *O11yO11yServiceTag) HasNumberValues() bool`

HasNumberValues returns a boolean if a field has been set.

### GetOperator

`func (o *O11yO11yServiceTag) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *O11yO11yServiceTag) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *O11yO11yServiceTag) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *O11yO11yServiceTag) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetStringValues

`func (o *O11yO11yServiceTag) GetStringValues() []string`

GetStringValues returns the StringValues field if non-nil, zero value otherwise.

### GetStringValuesOk

`func (o *O11yO11yServiceTag) GetStringValuesOk() (*[]string, bool)`

GetStringValuesOk returns a tuple with the StringValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValues

`func (o *O11yO11yServiceTag) SetStringValues(v []string)`

SetStringValues sets StringValues field to given value.

### HasStringValues

`func (o *O11yO11yServiceTag) HasStringValues() bool`

HasStringValues returns a boolean if a field has been set.

### GetTagType

`func (o *O11yO11yServiceTag) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *O11yO11yServiceTag) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *O11yO11yServiceTag) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *O11yO11yServiceTag) HasTagType() bool`

HasTagType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


