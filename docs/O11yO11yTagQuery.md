# O11yO11yTagQuery

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoolValues** | Pointer to **[]bool** | BoolValues are the boolean values to test against. | [optional] 
**Key** | Pointer to **string** | Key is the tag to test. | [optional] 
**NumberValues** | Pointer to **[]float64** | NumberValues are the numeric values to test against. | [optional] 
**Operator** | Pointer to **string** | Operator is the comparison, e.g. in, nin, contains, exists. | [optional] 
**StringValues** | Pointer to **[]string** | StringValues are the string values to test against. | [optional] 
**TagType** | Pointer to **string** | TagType is where the tag lives, e.g. ResourceAttribute, SpanAttribute. | [optional] 

## Methods

### NewO11yO11yTagQuery

`func NewO11yO11yTagQuery() *O11yO11yTagQuery`

NewO11yO11yTagQuery instantiates a new O11yO11yTagQuery object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTagQueryWithDefaults

`func NewO11yO11yTagQueryWithDefaults() *O11yO11yTagQuery`

NewO11yO11yTagQueryWithDefaults instantiates a new O11yO11yTagQuery object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolValues

`func (o *O11yO11yTagQuery) GetBoolValues() []bool`

GetBoolValues returns the BoolValues field if non-nil, zero value otherwise.

### GetBoolValuesOk

`func (o *O11yO11yTagQuery) GetBoolValuesOk() (*[]bool, bool)`

GetBoolValuesOk returns a tuple with the BoolValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValues

`func (o *O11yO11yTagQuery) SetBoolValues(v []bool)`

SetBoolValues sets BoolValues field to given value.

### HasBoolValues

`func (o *O11yO11yTagQuery) HasBoolValues() bool`

HasBoolValues returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yTagQuery) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yTagQuery) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yTagQuery) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yTagQuery) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNumberValues

`func (o *O11yO11yTagQuery) GetNumberValues() []float64`

GetNumberValues returns the NumberValues field if non-nil, zero value otherwise.

### GetNumberValuesOk

`func (o *O11yO11yTagQuery) GetNumberValuesOk() (*[]float64, bool)`

GetNumberValuesOk returns a tuple with the NumberValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValues

`func (o *O11yO11yTagQuery) SetNumberValues(v []float64)`

SetNumberValues sets NumberValues field to given value.

### HasNumberValues

`func (o *O11yO11yTagQuery) HasNumberValues() bool`

HasNumberValues returns a boolean if a field has been set.

### GetOperator

`func (o *O11yO11yTagQuery) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *O11yO11yTagQuery) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *O11yO11yTagQuery) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *O11yO11yTagQuery) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetStringValues

`func (o *O11yO11yTagQuery) GetStringValues() []string`

GetStringValues returns the StringValues field if non-nil, zero value otherwise.

### GetStringValuesOk

`func (o *O11yO11yTagQuery) GetStringValuesOk() (*[]string, bool)`

GetStringValuesOk returns a tuple with the StringValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValues

`func (o *O11yO11yTagQuery) SetStringValues(v []string)`

SetStringValues sets StringValues field to given value.

### HasStringValues

`func (o *O11yO11yTagQuery) HasStringValues() bool`

HasStringValues returns a boolean if a field has been set.

### GetTagType

`func (o *O11yO11yTagQuery) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *O11yO11yTagQuery) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *O11yO11yTagQuery) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *O11yO11yTagQuery) HasTagType() bool`

HasTagType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


