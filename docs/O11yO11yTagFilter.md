# O11yO11yTagFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BoolValues** | Pointer to **[]bool** | BoolValues are the values matched when the tag holds booleans. | [optional] 
**Key** | Pointer to **string** | Key is the tag to test. | [optional] 
**NumberValues** | Pointer to **[]float64** | NumberValues are the values matched when the tag holds numbers. | [optional] 
**Operator** | Pointer to **string** | Operator is the comparison to apply — in, not_in, equals, contains and the other operators the trace filter grammar names. | [optional] 
**StringValues** | Pointer to **[]string** | StringValues are the values matched when the tag holds strings. | [optional] 
**TagType** | Pointer to **string** | TagType says which kind of value the tag holds: string, number or bool. | [optional] 

## Methods

### NewO11yO11yTagFilter

`func NewO11yO11yTagFilter() *O11yO11yTagFilter`

NewO11yO11yTagFilter instantiates a new O11yO11yTagFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewO11yO11yTagFilterWithDefaults

`func NewO11yO11yTagFilterWithDefaults() *O11yO11yTagFilter`

NewO11yO11yTagFilterWithDefaults instantiates a new O11yO11yTagFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBoolValues

`func (o *O11yO11yTagFilter) GetBoolValues() []bool`

GetBoolValues returns the BoolValues field if non-nil, zero value otherwise.

### GetBoolValuesOk

`func (o *O11yO11yTagFilter) GetBoolValuesOk() (*[]bool, bool)`

GetBoolValuesOk returns a tuple with the BoolValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoolValues

`func (o *O11yO11yTagFilter) SetBoolValues(v []bool)`

SetBoolValues sets BoolValues field to given value.

### HasBoolValues

`func (o *O11yO11yTagFilter) HasBoolValues() bool`

HasBoolValues returns a boolean if a field has been set.

### GetKey

`func (o *O11yO11yTagFilter) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *O11yO11yTagFilter) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *O11yO11yTagFilter) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *O11yO11yTagFilter) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetNumberValues

`func (o *O11yO11yTagFilter) GetNumberValues() []float64`

GetNumberValues returns the NumberValues field if non-nil, zero value otherwise.

### GetNumberValuesOk

`func (o *O11yO11yTagFilter) GetNumberValuesOk() (*[]float64, bool)`

GetNumberValuesOk returns a tuple with the NumberValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberValues

`func (o *O11yO11yTagFilter) SetNumberValues(v []float64)`

SetNumberValues sets NumberValues field to given value.

### HasNumberValues

`func (o *O11yO11yTagFilter) HasNumberValues() bool`

HasNumberValues returns a boolean if a field has been set.

### GetOperator

`func (o *O11yO11yTagFilter) GetOperator() string`

GetOperator returns the Operator field if non-nil, zero value otherwise.

### GetOperatorOk

`func (o *O11yO11yTagFilter) GetOperatorOk() (*string, bool)`

GetOperatorOk returns a tuple with the Operator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperator

`func (o *O11yO11yTagFilter) SetOperator(v string)`

SetOperator sets Operator field to given value.

### HasOperator

`func (o *O11yO11yTagFilter) HasOperator() bool`

HasOperator returns a boolean if a field has been set.

### GetStringValues

`func (o *O11yO11yTagFilter) GetStringValues() []string`

GetStringValues returns the StringValues field if non-nil, zero value otherwise.

### GetStringValuesOk

`func (o *O11yO11yTagFilter) GetStringValuesOk() (*[]string, bool)`

GetStringValuesOk returns a tuple with the StringValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStringValues

`func (o *O11yO11yTagFilter) SetStringValues(v []string)`

SetStringValues sets StringValues field to given value.

### HasStringValues

`func (o *O11yO11yTagFilter) HasStringValues() bool`

HasStringValues returns a boolean if a field has been set.

### GetTagType

`func (o *O11yO11yTagFilter) GetTagType() string`

GetTagType returns the TagType field if non-nil, zero value otherwise.

### GetTagTypeOk

`func (o *O11yO11yTagFilter) GetTagTypeOk() (*string, bool)`

GetTagTypeOk returns a tuple with the TagType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagType

`func (o *O11yO11yTagFilter) SetTagType(v string)`

SetTagType sets TagType field to given value.

### HasTagType

`func (o *O11yO11yTagFilter) HasTagType() bool`

HasTagType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


